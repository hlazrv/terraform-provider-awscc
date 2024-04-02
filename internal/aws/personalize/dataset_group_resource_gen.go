// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package personalize

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
	registry.AddResourceFactory("awscc_personalize_dataset_group", datasetGroupResource)
}

// datasetGroupResource returns the Terraform awscc_personalize_dataset_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::Personalize::DatasetGroup resource.
func datasetGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the dataset group.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the dataset group.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain of a Domain dataset group.",
		//	  "enum": [
		//	    "ECOMMERCE",
		//	    "VIDEO_ON_DEMAND"
		//	  ],
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain of a Domain dataset group.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ECOMMERCE",
					"VIDEO_ON_DEMAND",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.",
		//	  "maxLength": 2048,
		//	  "pattern": "arn:aws.*:kms:.*:[0-9]{12}:key/.*",
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws.*:kms:.*:[0-9]{12}:key/.*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the new dataset group.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the new dataset group.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Schema for AWS::Personalize::DatasetGroup.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::DatasetGroup").WithTerraformTypeName("awscc_personalize_dataset_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_group_arn": "DatasetGroupArn",
		"domain":            "Domain",
		"kms_key_arn":       "KmsKeyArn",
		"name":              "Name",
		"role_arn":          "RoleArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
