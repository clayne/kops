/*
Copyright 2020 The Kubernetes Authors.

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

package azuretasks

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	network "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"k8s.io/klog/v2"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/azure"
)

// VirtualNetwork is an Azure Virtual Network.
// +kops:fitask
type VirtualNetwork struct {
	Name      *string
	Lifecycle fi.Lifecycle

	ResourceGroup *ResourceGroup
	CIDR          *string
	Tags          map[string]*string
	Shared        *bool
}

var (
	_ fi.CloudupTask          = &VirtualNetwork{}
	_ fi.CompareWithID        = &VirtualNetwork{}
	_ fi.CloudupTaskNormalize = &VirtualNetwork{}
)

// CompareWithID returns the Name of the VM Scale Set.
func (n *VirtualNetwork) CompareWithID() *string {
	return n.Name
}

// Find discovers the VirtualNetwork in the cloud provider.
func (n *VirtualNetwork) Find(c *fi.CloudupContext) (*VirtualNetwork, error) {
	cloud := c.T.Cloud.(azure.AzureCloud)
	l, err := cloud.VirtualNetwork().List(context.TODO(), *n.ResourceGroup.Name)
	if err != nil {
		return nil, err
	}
	var found *network.VirtualNetwork
	for _, v := range l {
		if *v.Name == *n.Name {
			found = v
			break
		}
	}
	if found == nil {
		return nil, nil
	}

	if found.ID == nil {
		return nil, fmt.Errorf("found virtual network without ID")
	}
	if found.Properties == nil {
		return nil, fmt.Errorf("found virtual network without properties")
	}
	if found.Properties.AddressSpace == nil {
		return nil, fmt.Errorf("found virtual network without address space")
	}

	addrPrefixes := found.Properties.AddressSpace.AddressPrefixes
	if len(addrPrefixes) != 1 {
		return nil, fmt.Errorf("expecting exactly 1 address prefix for %q, found %d: %+v", *n.Name, len(addrPrefixes), addrPrefixes)
	}
	return &VirtualNetwork{
		Name:      n.Name,
		Lifecycle: n.Lifecycle,
		Shared:    n.Shared,
		ResourceGroup: &ResourceGroup{
			Name: n.ResourceGroup.Name,
		},
		CIDR: addrPrefixes[0],
		Tags: found.Tags,
	}, nil
}

func (n *VirtualNetwork) Normalize(c *fi.CloudupContext) error {
	c.T.Cloud.(azure.AzureCloud).AddClusterTags(n.Tags)
	return nil
}

// Run implements fi.Task.Run.
func (n *VirtualNetwork) Run(c *fi.CloudupContext) error {
	return fi.CloudupDefaultDeltaRunMethod(n, c)
}

// CheckChanges returns an error if a change is not allowed.
func (*VirtualNetwork) CheckChanges(a, e, changes *VirtualNetwork) error {
	if a == nil {
		// Check if required fields are set when a new resource is created.
		if e.Name == nil {
			return fi.RequiredField("Name")
		}
		return nil
	}

	// Check if unchanegable fields won't be changed.
	if changes.Name != nil {
		return fi.CannotChangeField("Name")
	}
	if changes.CIDR != nil {
		return fi.CannotChangeField("CIDR")
	}
	return nil
}

// RenderAzure creates or updates a Virtual Network.
func (*VirtualNetwork) RenderAzure(t *azure.AzureAPITarget, a, e, changes *VirtualNetwork) error {
	if a == nil {
		klog.Infof("Creating a new Virtual Network with name: %s", fi.ValueOf(e.Name))
	} else {
		// Only allow tags to be updated.
		if changes.Tags == nil {
			return nil
		}
		klog.Infof("Updating a Virtual Network with name: %s", fi.ValueOf(e.Name))
	}

	vnet := network.VirtualNetwork{
		Location: to.Ptr(t.Cloud.Region()),
		Properties: &network.VirtualNetworkPropertiesFormat{
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []*string{e.CIDR},
			},
		},
		Tags: e.Tags,
	}
	_, err := t.Cloud.VirtualNetwork().CreateOrUpdate(
		context.TODO(),
		*e.ResourceGroup.Name,
		*e.Name,
		vnet)

	return err
}
