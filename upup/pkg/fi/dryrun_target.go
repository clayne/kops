/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fi

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"sync"

	"k8s.io/klog/v2"
	"k8s.io/kops/pkg/assets"
	"k8s.io/kops/pkg/diff"
	"k8s.io/kops/util/pkg/reflectutils"
)

// DryRunTarget is a special Target that does not execute anything, but instead tracks all changes.
// By running against a DryRunTarget, a list of changes that would be made can be easily collected,
// without any special support from the Tasks.
type DryRunTarget[T SubContext] struct {
	mutex sync.Mutex

	changes   []*render[T]
	deletions []Deletion[T]

	// The destination to which the final report will be printed on Finish()
	out io.Writer

	// assetBuilder records all assets used
	assetBuilder *assets.AssetBuilder

	// defaultCheckExisting will control whether we look for existing objects in our dry-run.
	// This is normally true except for special-case dry-runs, like `kops get assets`
	defaultCheckExisting bool
}

type NodeupDryRunTarget = DryRunTarget[NodeupSubContext]
type CloudupDryRunTarget = DryRunTarget[CloudupSubContext]

type render[T SubContext] struct {
	a       Task[T]
	aIsNil  bool
	e       Task[T]
	changes Task[T]
}

// ByTaskKey sorts []*render by TaskKey (type/name)
type ByTaskKey[T SubContext] []*render[T]

func (a ByTaskKey[T]) Len() int      { return len(a) }
func (a ByTaskKey[T]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTaskKey[T]) Less(i, j int) bool {
	return buildTaskKey(a[i].e) < buildTaskKey(a[j].e)
}

// DeletionByTaskName sorts []Deletion by TaskName
type DeletionByTaskName[T SubContext] []Deletion[T]

func (a DeletionByTaskName[T]) Len() int      { return len(a) }
func (a DeletionByTaskName[T]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DeletionByTaskName[T]) Less(i, j int) bool {
	return a[i].TaskName() < a[j].TaskName()
}

var _ Target[CloudupSubContext] = &DryRunTarget[CloudupSubContext]{}

func newDryRunTarget[T SubContext](assetBuilder *assets.AssetBuilder, defaultCheckExisting bool, out io.Writer) *DryRunTarget[T] {
	t := &DryRunTarget[T]{}
	t.out = out
	t.assetBuilder = assetBuilder
	t.defaultCheckExisting = defaultCheckExisting
	return t
}

// NewCloudupDryRunTarget builds a dry-run target.
// checkExisting should normally be true, but can be false for special-case dry-run, such as in `kops get assets`
func NewCloudupDryRunTarget(assetBuilder *assets.AssetBuilder, checkExisting bool, out io.Writer) *CloudupDryRunTarget {
	return newDryRunTarget[CloudupSubContext](assetBuilder, checkExisting, out)
}

func NewNodeupDryRunTarget(assetBuilder *assets.AssetBuilder, out io.Writer) *NodeupDryRunTarget {
	return newDryRunTarget[NodeupSubContext](assetBuilder, true, out)
}

func (t *DryRunTarget[T]) DefaultCheckExisting() bool {
	return t.defaultCheckExisting
}

func (t *DryRunTarget[T]) Render(a, e, changes Task[T]) error {
	valA := reflect.ValueOf(a)
	aIsNil := valA.IsNil()

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.changes = append(t.changes, &render[T]{
		a:       a,
		aIsNil:  aIsNil,
		e:       e,
		changes: changes,
	})
	return nil
}

func (t *DryRunTarget[T]) RecordDeletion(deletion Deletion[T]) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.deletions = append(t.deletions, deletion)

	return nil
}

func idForTask[T SubContext](taskMap map[string]Task[T], t Task[T]) string {
	for k, v := range taskMap {
		if v == t {
			// Skip task type, if present (taskType/taskName)
			firstSlash := strings.Index(k, "/")
			if firstSlash != -1 {
				k = k[firstSlash+1:]
			}
			return k
		}
	}
	klog.Fatalf("unknown task: %v", t)
	return "?"
}

func (t *DryRunTarget[T]) PrintReport(taskMap map[string]Task[T], out io.Writer) error {
	b := &bytes.Buffer{}

	if len(t.changes) != 0 {
		var creates []*render[T]
		var updates []*render[T]

		for _, r := range t.changes {
			if r.aIsNil {
				creates = append(creates, r)
			} else {
				updates = append(updates, r)
			}
		}

		// Give everything a consistent ordering
		sort.Sort(ByTaskKey[T](creates))
		sort.Sort(ByTaskKey[T](updates))

		if len(creates) != 0 {
			fmt.Fprintf(b, "Will create resources:\n")
			for _, r := range creates {
				taskName := getTaskName(r.changes)
				fmt.Fprintf(b, "  %s/%s\n", taskName, idForTask(taskMap, r.e))

				changes := reflect.ValueOf(r.changes)
				if changes.Kind() == reflect.Ptr && !changes.IsNil() {
					changes = changes.Elem()
				}

				if changes.Kind() == reflect.Struct {
					for i := 0; i < changes.NumField(); i++ {

						field := changes.Field(i)

						fieldName := changes.Type().Field(i).Name
						if changes.Type().Field(i).PkgPath != "" {
							// Not exported
							continue
						}

						fieldValue := reflectutils.ValueAsString(field)

						shouldPrint := true
						if fieldName == "Name" {
							// The field name is already printed above, no need to repeat it.
							shouldPrint = false
						}
						if fieldName == "Lifecycle" {
							// Lifecycle is a "system" field; no need to show it
							shouldPrint = false
						}
						if fieldValue == "<nil>" || fieldValue == "<resource>" {
							// Uninformative
							shouldPrint = false
						}
						if fieldValue == "id:<nil>" {
							// Uninformative, but we can often print the name instead
							name := ""
							if field.CanInterface() {
								hasName, ok := field.Interface().(HasName)
								if ok {
									name = ValueOf(hasName.GetName())
								}
							}
							if name != "" {
								fieldValue = "name:" + name
							} else {
								shouldPrint = false
							}
						}
						if shouldPrint {
							fmt.Fprintf(b, "  \t%-20s\t%s\n", fieldName, fieldValue)
						}
					}
				}

				fmt.Fprintf(b, "\n")
			}
		}

		if len(updates) != 0 {
			fmt.Fprintf(b, "Will modify resources:\n")
			// We can't use our reflection helpers here - we want corresponding values from a,e,c
			for _, r := range updates {
				changeList, err := buildChangeList(r.a, r.e, r.changes)
				if err != nil {
					return err
				}
				taskName := getTaskName(r.changes)
				fmt.Fprintf(b, "  %s/%s\n", taskName, idForTask(taskMap, r.e))

				if len(changeList) == 0 {
					fmt.Fprintf(b, "   internal consistency error!\n")
					fmt.Fprintf(b, "    actual: %+v\n", r.a)
					fmt.Fprintf(b, "    expect: %+v\n", r.e)
					fmt.Fprintf(b, "    change: %+v\n", r.changes)
					continue
				}

				for _, change := range changeList {
					lines := strings.Split(change.Description, "\n")
					if len(lines) == 1 {
						fmt.Fprintf(b, "  \t%-20s\t%s\n", change.FieldName, change.Description)
					} else {
						fmt.Fprintf(b, "  \t%-20s\n", change.FieldName)
						for _, line := range lines {
							fmt.Fprintf(b, "  \t%-20s\t%s\n", "", line)
						}
					}
				}
				fmt.Fprintf(b, "\n")
			}
		}
	}

	if len(t.deletions) != 0 {
		// Give everything a consistent ordering
		sort.Sort(DeletionByTaskName[T](t.deletions))

		var deferred []Deletion[T]
		var immediate []Deletion[T]

		for _, d := range t.deletions {
			if d.DeferDeletion() {
				deferred = append(deferred, d)
			} else {
				immediate = append(immediate, d)
			}
		}

		if len(deferred) != 0 {
			fmt.Fprintf(b, "Items will be deleted only when the --prune flag is specified:\n")
			for _, d := range deferred {
				fmt.Fprintf(b, "  %-20s %s\n", d.TaskName(), d.Item())
			}
			fmt.Fprintf(b, "\n")
		}
		if len(immediate) != 0 {
			fmt.Fprintf(b, "Items will be deleted during update:\n")
			for _, d := range immediate {
				fmt.Fprintf(b, "  %-20s %s\n", d.TaskName(), d.Item())
			}
			fmt.Fprintf(b, "\n")
		}
	}

	if len(t.assetBuilder.ImageAssets) != 0 {
		klog.V(4).Infof("ImageAssets:")
		for _, a := range t.assetBuilder.ImageAssets {
			klog.V(4).Infof("  %s %s", a.DownloadLocation, a.CanonicalLocation)
		}
	}

	if len(t.assetBuilder.FileAssets) != 0 {
		klog.V(4).Infof("FileAssets:")
		for _, a := range t.assetBuilder.FileAssets {
			if a.DownloadURL != nil {
				klog.V(4).Infof("  %s %s", a.DownloadURL.String(), a.CanonicalURL.String())
			}
		}
	}

	_, err := out.Write(b.Bytes())
	return err
}

type change struct {
	FieldName   string
	Description string
}

func buildChangeList[T SubContext](a, e, changes Task[T]) ([]change, error) {
	var changeList []change

	valC := reflect.ValueOf(changes)
	valA := reflect.ValueOf(a)
	valE := reflect.ValueOf(e)
	if valC.Kind() == reflect.Ptr && !valC.IsNil() {
		valC = valC.Elem()
	}
	if valA.Kind() == reflect.Ptr && !valA.IsNil() {
		valA = valA.Elem()
	}
	if valE.Kind() == reflect.Ptr && !valE.IsNil() {
		valE = valE.Elem()
	}
	if valC.Kind() == reflect.Struct {
		for i := 0; i < valC.NumField(); i++ {
			if valC.Type().Field(i).PkgPath != "" {
				// Not exported
				continue
			}

			fieldValC := valC.Field(i)
			fieldValE := valE.Field(i)
			fieldValA := valA.Field(i)

			changed := true
			switch fieldValC.Kind() {
			case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map:
				changed = !fieldValC.IsNil()
				if fieldValC.IsNil() && !fieldValA.IsNil() && fieldValE.IsNil() {
					changed = true
				}

			case reflect.String:
				changed = fieldValC.Convert(reflect.TypeOf("")).Interface() != ""

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				changed = fieldValA.Int() != fieldValE.Int()

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				changed = fieldValA.Uint() != fieldValE.Uint()

			case reflect.Bool:
				changed = fieldValA.Bool() != fieldValE.Bool()

			default:
				klog.Warningf("unhandled type in diff construction: %v", fieldValC.Kind())
			}

			if !changed {
				continue
			}

			description := ""
			ignored := false
			if fieldValE.CanInterface() {

				switch fieldValE.Interface().(type) {
				// case SimpleUnit:
				//	ignored = true
				case Resource:
					resA, okA := tryResourceAsString(fieldValA)
					resE, okE := tryResourceAsString(fieldValE)
					if okA && okE {
						description = diff.FormatDiff(resA, resE)
					}
				}

				if !ignored && description == "" {
					description = fmt.Sprintf(" %v -> %v", reflectutils.ValueAsString(fieldValA), reflectutils.ValueAsString(fieldValE))
				}
			}
			if ignored {
				continue
			}
			changeList = append(changeList, change{FieldName: valC.Type().Field(i).Name, Description: description})
		}
	} else {
		return nil, fmt.Errorf("unhandled change type: %v", valC.Type())
	}

	return changeList, nil
}

func tryResourceAsString(v reflect.Value) (string, bool) {
	if !v.IsValid() {
		return "", false
	}
	if !v.CanInterface() {
		return "", false
	}

	// Guard against nil interface go-tcha
	if v.Kind() == reflect.Interface && v.IsNil() {
		return "", false
	}
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return "", false
	}

	intf := v.Interface()
	if res, ok := intf.(Resource); ok {
		s, err := ResourceAsString(res)
		if err != nil {
			klog.Warningf("error converting to resource: %v", err)
			return "", false
		}
		return s, true
	}
	return "", false
}

func getTaskName[T SubContext](t Task[T]) string {
	s := fmt.Sprintf("%T", t)
	lastDot := strings.LastIndexByte(s, '.')
	if lastDot != -1 {
		s = s[lastDot+1:]
	}
	return s
}

// Finish is called at the end of a run, and prints a list of changes to the configured Writer
func (t *DryRunTarget[T]) Finish(taskMap map[string]Task[T]) error {
	return t.PrintReport(taskMap, t.out)
}

// Deletions returns all task names which is going to be deleted
func (t *DryRunTarget[T]) Deletions() []string {
	var deletions []string
	for _, d := range t.deletions {
		deletions = append(deletions, d.TaskName())
	}
	return deletions
}

// Changes returns tasks which is going to be created or updated
func (t *DryRunTarget[T]) Changes() (map[string]Task[T], map[string]Task[T]) {
	creates := make(map[string]Task[T])
	updates := make(map[string]Task[T])
	for _, r := range t.changes {
		if r.aIsNil {
			creates[getTaskName(r.changes)] = r.changes
		} else {
			updates[getTaskName(r.changes)] = r.changes
		}
	}
	return creates, updates
}

// HasChanges returns true iff any changes would have been made
func (t *DryRunTarget[T]) HasChanges() bool {
	return len(t.changes)+len(t.deletions) != 0
}
