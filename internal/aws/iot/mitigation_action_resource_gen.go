// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iot_mitigation_action", mitigationActionResource)
}

// mitigationActionResource returns the Terraform awscc_iot_mitigation_action resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoT::MitigationAction resource.
func mitigationActionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the mitigation action.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9:_-]+",
		//	  "type": "string"
		//	}
		"action_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the mitigation action.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ActionParams
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
		//	  "properties": {
		//	    "AddThingsToThingGroupParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
		//	      "properties": {
		//	        "OverrideDynamicGroups": {
		//	          "description": "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
		//	          "type": "boolean"
		//	        },
		//	        "ThingGroupNames": {
		//	          "description": "The list of groups to which you want to add the things that triggered the mitigation action.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "ThingGroupNames"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "EnableIoTLoggingParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
		//	      "properties": {
		//	        "LogLevel": {
		//	          "description": " Specifies which types of information are logged.",
		//	          "enum": [
		//	            "DEBUG",
		//	            "INFO",
		//	            "ERROR",
		//	            "WARN",
		//	            "UNSET_VALUE"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "RoleArnForLogging": {
		//	          "description": " The ARN of the IAM role used for logging.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogLevel",
		//	        "RoleArnForLogging"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "PublishFindingToSnsParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
		//	      "properties": {
		//	        "TopicArn": {
		//	          "description": "The ARN of the topic to which you want to publish the findings.",
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TopicArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ReplaceDefaultPolicyVersionParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
		//	      "properties": {
		//	        "TemplateName": {
		//	          "enum": [
		//	            "BLANK_POLICY",
		//	            "UNSET_VALUE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TemplateName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "UpdateCACertificateParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
		//	      "properties": {
		//	        "Action": {
		//	          "enum": [
		//	            "DEACTIVATE",
		//	            "UNSET_VALUE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Action"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "UpdateDeviceCertificateParams": {
		//	      "additionalProperties": false,
		//	      "description": "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
		//	      "properties": {
		//	        "Action": {
		//	          "enum": [
		//	            "DEACTIVATE",
		//	            "UNSET_VALUE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Action"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"action_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AddThingsToThingGroupParams
				"add_things_to_thing_group_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OverrideDynamicGroups
						"override_dynamic_groups": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ThingGroupNames
						"thing_group_names": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The list of groups to which you want to add the things that triggered the mitigation action.",
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 10),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 128),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EnableIoTLoggingParams
				"enable_io_t_logging_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogLevel
						"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: " Specifies which types of information are logged.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DEBUG",
									"INFO",
									"ERROR",
									"WARN",
									"UNSET_VALUE",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: RoleArnForLogging
						"role_arn_for_logging": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: " The ARN of the IAM role used for logging.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(20, 2048),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PublishFindingToSnsParams
				"publish_finding_to_sns_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TopicArn
						"topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the topic to which you want to publish the findings.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(20, 2048),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters, to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ReplaceDefaultPolicyVersionParams
				"replace_default_policy_version_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TemplateName
						"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"BLANK_POLICY",
									"UNSET_VALUE",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that adds a blank policy to restrict permissions.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: UpdateCACertificateParams
				"update_ca_certificate_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Action
						"action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DEACTIVATE",
									"UNSET_VALUE",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that changes the state of the CA certificate to inactive.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: UpdateDeviceCertificateParams
				"update_device_certificate_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Action
						"action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DEACTIVATE",
									"UNSET_VALUE",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Parameters to define a mitigation action that changes the state of the device certificate to inactive.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: MitigationActionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"mitigation_action_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MitigationActionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"mitigation_action_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Mitigation actions can be used to take actions to mitigate issues that were found in an Audit finding or Detect violation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::MitigationAction").WithTerraformTypeName("awscc_iot_mitigation_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                                "Action",
		"action_name":                           "ActionName",
		"action_params":                         "ActionParams",
		"add_things_to_thing_group_params":      "AddThingsToThingGroupParams",
		"enable_io_t_logging_params":            "EnableIoTLoggingParams",
		"key":                                   "Key",
		"log_level":                             "LogLevel",
		"mitigation_action_arn":                 "MitigationActionArn",
		"mitigation_action_id":                  "MitigationActionId",
		"override_dynamic_groups":               "OverrideDynamicGroups",
		"publish_finding_to_sns_params":         "PublishFindingToSnsParams",
		"replace_default_policy_version_params": "ReplaceDefaultPolicyVersionParams",
		"role_arn":                              "RoleArn",
		"role_arn_for_logging":                  "RoleArnForLogging",
		"tags":                                  "Tags",
		"template_name":                         "TemplateName",
		"thing_group_names":                     "ThingGroupNames",
		"topic_arn":                             "TopicArn",
		"update_ca_certificate_params":          "UpdateCACertificateParams",
		"update_device_certificate_params":      "UpdateDeviceCertificateParams",
		"value":                                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
