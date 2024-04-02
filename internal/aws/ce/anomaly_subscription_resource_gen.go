// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ce_anomaly_subscription", anomalySubscriptionResource)
}

// anomalySubscriptionResource returns the Terraform awscc_ce_anomaly_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::CE::AnomalySubscription resource.
func anomalySubscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The accountId",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The accountId",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Frequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The frequency at which anomaly reports are sent over email. ",
		//	  "enum": [
		//	    "DAILY",
		//	    "IMMEDIATE",
		//	    "WEEKLY"
		//	  ],
		//	  "type": "string"
		//	}
		"frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The frequency at which anomaly reports are sent over email. ",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DAILY",
					"IMMEDIATE",
					"WEEKLY",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorArnList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of cost anomaly monitors.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Subscription ARN",
		//	    "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"monitor_arn_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of cost anomaly monitors.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags to assign to subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name for the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name for the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags to assign to subscription.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ResourceTags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Subscribers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of subscriber",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Address": {
		//	        "pattern": "(^[a-zA-Z0-9.!#$%\u0026'*+=?^_‘{|}~-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$)|(^arn:(aws[a-zA-Z-]*):sns:[a-zA-Z0-9-]+:[0-9]{12}:[a-zA-Z0-9_-]+(\\.fifo)?$)",
		//	        "type": "string"
		//	      },
		//	      "Status": {
		//	        "enum": [
		//	          "CONFIRMED",
		//	          "DECLINED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "EMAIL",
		//	          "SNS"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Address",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"subscribers": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Address
					"address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("(^[a-zA-Z0-9.!#$%&'*+=?^_‘{|}~-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$)|(^arn:(aws[a-zA-Z-]*):sns:[a-zA-Z0-9-]+:[0-9]{12}:[a-zA-Z0-9_-]+(\\.fifo)?$)"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"CONFIRMED",
								"DECLINED",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"EMAIL",
								"SNS",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of subscriber",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Subscription ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"subscription_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Subscription ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subscription.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": "[\\S\\s]*",
		//	  "type": "string"
		//	}
		"subscription_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subscription.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\S\\s]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Threshold
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dollar value that triggers a notification if the threshold is exceeded. ",
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The dollar value that triggers a notification if the threshold is exceeded. ",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.AtLeast(0.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ThresholdExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
		//	  "type": "string"
		//	}
		"threshold_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalySubscription").WithTerraformTypeName("awscc_ce_anomaly_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":           "AccountId",
		"address":              "Address",
		"frequency":            "Frequency",
		"key":                  "Key",
		"monitor_arn_list":     "MonitorArnList",
		"resource_tags":        "ResourceTags",
		"status":               "Status",
		"subscribers":          "Subscribers",
		"subscription_arn":     "SubscriptionArn",
		"subscription_name":    "SubscriptionName",
		"threshold":            "Threshold",
		"threshold_expression": "ThresholdExpression",
		"type":                 "Type",
		"value":                "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ResourceTags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
