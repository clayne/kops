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

package awstasks

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kops/upup/pkg/fi/cloudup/terraformWriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	eventbridgetypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
	"k8s.io/kops/upup/pkg/fi/cloudup/terraform"
)

// +kops:fitask
type EventBridgeTarget struct {
	ID        *string
	Name      *string
	Lifecycle fi.Lifecycle

	Rule     *EventBridgeRule
	SQSQueue *SQS
}

var _ fi.CompareWithID = &EventBridgeTarget{}

func (eb *EventBridgeTarget) CompareWithID() *string {
	return eb.Name
}

func (eb *EventBridgeTarget) Find(c *fi.CloudupContext) (*EventBridgeTarget, error) {
	cloud := awsup.GetCloud(c)

	if eb.Rule == nil || eb.SQSQueue == nil {
		return nil, nil
	}

	// find the rule the target is attached to
	rule, err := eb.Rule.Find(c)
	if err != nil {
		return nil, err
	}
	if rule == nil {
		return nil, nil
	}

	request := &eventbridge.ListTargetsByRuleInput{
		Rule: eb.Rule.Name,
	}

	response, err := cloud.EventBridge().ListTargetsByRule(c.Context(), request)
	if err != nil {
		return nil, fmt.Errorf("error listing EventBridge targets: %v", err)
	}
	if response == nil || len(response.Targets) == 0 {
		return nil, nil
	}
	for _, target := range response.Targets {
		if fi.ValueOf(target.Arn) == fi.ValueOf(eb.SQSQueue.ARN) {
			actual := &EventBridgeTarget{
				ID:        target.Id,
				Name:      eb.Name,
				Lifecycle: eb.Lifecycle,
				Rule:      eb.Rule,
				SQSQueue:  eb.SQSQueue,
			}
			return actual, nil
		}
	}

	return nil, nil
}

func (eb *EventBridgeTarget) Run(c *fi.CloudupContext) error {
	return fi.CloudupDefaultDeltaRunMethod(eb, c)
}

func (_ *EventBridgeTarget) CheckChanges(a, e, changes *EventBridgeTarget) error {
	if a == nil {
		if e.Rule == nil {
			return field.Required(field.NewPath("Rule"), "")
		}
		if e.SQSQueue == nil {
			return field.Required(field.NewPath("SQSQueue"), "")
		}
	}

	return nil
}

func (eb *EventBridgeTarget) RenderAWS(t *awsup.AWSAPITarget, a, e, changes *EventBridgeTarget) error {
	if a == nil {
		target := eventbridgetypes.Target{
			Arn: eb.SQSQueue.ARN,
			Id:  aws.String("1"),
		}

		request := &eventbridge.PutTargetsInput{
			Rule:    eb.Rule.Name,
			Targets: []eventbridgetypes.Target{target},
		}

		_, err := t.Cloud.EventBridge().PutTargets(context.TODO(), request)
		if err != nil {
			return fmt.Errorf("error creating EventBridge target: %v", err)
		}
	}

	return nil
}

type terraformEventBridgeTarget struct {
	RuleName  *terraformWriter.Literal `cty:"rule"`
	TargetArn *terraformWriter.Literal `cty:"arn"`
}

func (_ *EventBridgeTarget) RenderTerraform(t *terraform.TerraformTarget, a, e, changes *EventBridgeTarget) error {
	tf := &terraformEventBridgeTarget{
		RuleName:  e.Rule.TerraformLink(),
		TargetArn: e.SQSQueue.TerraformLink(),
	}

	return t.RenderResource("aws_cloudwatch_event_target", *e.Name, tf)
}
