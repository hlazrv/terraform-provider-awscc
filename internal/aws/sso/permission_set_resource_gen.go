// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sso

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_sso_permission_set", permissionSetResource)
}

// permissionSetResource returns the Terraform awscc_sso_permission_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSO::PermissionSet resource.
func permissionSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CustomerManagedPolicyReferences
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[\\w+=,.@-]+",
		//	        "type": "string"
		//	      },
		//	      "Path": {
		//	        "maxLength": 512,
		//	        "minLength": 1,
		//	        "pattern": "((/[A-Za-z0-9\\.,\\+@=_-]+)*)/",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 20,
		//	  "type": "array"
		//	}
		"customer_managed_policy_references": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Path
					"path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 512),
							stringvalidator.RegexMatches(regexp.MustCompile("((/[A-Za-z0-9\\.,\\+@=_-]+)*)/"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(20),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.ListOfStringDefaultValue(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The permission set description.",
		//	  "maxLength": 700,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The permission set description.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 700),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InlinePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The inline policy to put in permission set.",
		//	  "type": "string"
		//	}
		"inline_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The inline policy to put in permission set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sso instance arn that the permission set is owned.",
		//	  "maxLength": 1224,
		//	  "minLength": 10,
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The sso instance arn that the permission set is owned.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(10, 1224),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagedPolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The managed policy to attach.",
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 20,
		//	  "type": "array"
		//	}
		"managed_policies": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(20),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(20, 2048),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.ListOfStringDefaultValue(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name you want to assign to this permission set.",
		//	  "maxLength": 32,
		//	  "minLength": 1,
		//	  "pattern": "[\\w+=,.@-]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name you want to assign to this permission set.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 32),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionSetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The permission set that the policy will be attached to",
		//	  "maxLength": 1224,
		//	  "minLength": 10,
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}",
		//	  "type": "string"
		//	}
		"permission_set_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The permission set that the policy will be attached to",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionsBoundary
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CustomerManagedPolicyReference": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Name": {
		//	          "maxLength": 128,
		//	          "minLength": 1,
		//	          "pattern": "[\\w+=,.@-]+",
		//	          "type": "string"
		//	        },
		//	        "Path": {
		//	          "maxLength": 512,
		//	          "minLength": 1,
		//	          "pattern": "((/[A-Za-z0-9\\.,\\+@=_-]+)*)/",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ManagedPolicyArn": {
		//	      "description": "The managed policy to attach.",
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"permissions_boundary": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomerManagedPolicyReference
				"customer_managed_policy_reference": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Path
						"path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 512),
								stringvalidator.RegexMatches(regexp.MustCompile("((/[A-Za-z0-9\\.,\\+@=_-]+)*)/"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ManagedPolicyArn
				"managed_policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The managed policy to attach.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(20, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RelayStateType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The relay state URL that redirect links to any service in the AWS Management Console.",
		//	  "maxLength": 240,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9\u0026amp;$@#\\/%?=~\\-_'\u0026quot;|!:,.;*+\\[\\]\\ \\(\\)\\{\\}]+",
		//	  "type": "string"
		//	}
		"relay_state_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The relay state URL that redirect links to any service in the AWS Management Console.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 240),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9&amp;$@#\\/%?=~\\-_'&quot;|!:,.;*+\\[\\]\\ \\(\\)\\{\\}]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SessionDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The length of time that a user can be signed in to an AWS account.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"session_duration": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The length of time that a user can be signed in to an AWS account.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The metadata that you apply to the permission set to help you categorize and organize them.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[\\w+=,.@-]+",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "[\\w+=,.@-]+",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for SSO PermissionSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::PermissionSet").WithTerraformTypeName("awscc_sso_permission_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_managed_policy_reference":  "CustomerManagedPolicyReference",
		"customer_managed_policy_references": "CustomerManagedPolicyReferences",
		"description":                        "Description",
		"inline_policy":                      "InlinePolicy",
		"instance_arn":                       "InstanceArn",
		"key":                                "Key",
		"managed_policies":                   "ManagedPolicies",
		"managed_policy_arn":                 "ManagedPolicyArn",
		"name":                               "Name",
		"path":                               "Path",
		"permission_set_arn":                 "PermissionSetArn",
		"permissions_boundary":               "PermissionsBoundary",
		"relay_state_type":                   "RelayStateType",
		"session_duration":                   "SessionDuration",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
