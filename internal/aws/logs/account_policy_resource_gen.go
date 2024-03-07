// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_logs_account_policy", accountPolicyResource)
}

// accountPolicyResource returns the Terraform awscc_logs_account_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::Logs::AccountPolicy resource.
func accountPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "User account id",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "User account id",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per PolicyType.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
		//	  "maxLength": 30720,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per PolicyType.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 30720),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the account policy",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[^:*]{1,256}$",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the account policy",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[^:*]{1,256}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Type of the policy.",
		//	  "enum": [
		//	    "DATA_PROTECTION_POLICY",
		//	    "SUBSCRIPTION_FILTER_POLICY"
		//	  ],
		//	  "type": "string"
		//	}
		"policy_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Type of the policy.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DATA_PROTECTION_POLICY",
					"SUBSCRIPTION_FILTER_POLICY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Scope for policy application",
		//	  "enum": [
		//	    "ALL"
		//	  ],
		//	  "type": "string"
		//	}
		"scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Scope for policy application",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ALL",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SelectionCriteria
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes",
		//	  "type": "string"
		//	}
		"selection_criteria": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::AccountPolicy").WithTerraformTypeName("awscc_logs_account_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":         "AccountId",
		"policy_document":    "PolicyDocument",
		"policy_name":        "PolicyName",
		"policy_type":        "PolicyType",
		"scope":              "Scope",
		"selection_criteria": "SelectionCriteria",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
