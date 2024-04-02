// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_image_version", imageVersionResource)
}

// imageVersionResource returns the Terraform awscc_sagemaker_image_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::ImageVersion resource.
func imageVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The alias of the image version.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The alias of the image version.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Alias is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Aliases
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of aliases for the image version.",
		//	  "items": {
		//	    "description": "The alias of the image version.",
		//	    "maxLength": 128,
		//	    "minLength": 1,
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"aliases": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of aliases for the image version.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 128),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Aliases is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: BaseImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The registry path of the container image on which this image version is based.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": ".+",
		//	  "type": "string"
		//	}
		"base_image": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The registry path of the container image on which this image version is based.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile(".+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContainerImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The registry path of the container image that contains this image version.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": ".+",
		//	  "type": "string"
		//	}
		"container_image": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The registry path of the container image that contains this image version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Horovod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates Horovod compatibility.",
		//	  "type": "boolean"
		//	}
		"horovod": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates Horovod compatibility.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the parent image.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image\\/[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"image_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the parent image.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the image this version belongs to.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9]([-.]?[A-Za-z0-9])*$",
		//	  "type": "string"
		//	}
		"image_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the image this version belongs to.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9]([-.]?[A-Za-z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image version.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image-version\\/[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*\\/[0-9]+$",
		//	  "type": "string"
		//	}
		"image_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: JobType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates SageMaker job type compatibility.",
		//	  "enum": [
		//	    "TRAINING",
		//	    "INFERENCE",
		//	    "NOTEBOOK_KERNEL"
		//	  ],
		//	  "type": "string"
		//	}
		"job_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates SageMaker job type compatibility.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"TRAINING",
					"INFERENCE",
					"NOTEBOOK_KERNEL",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MLFramework
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The machine learning framework vended in the image version.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z]+ ?\\d+\\.\\d+(\\.\\d+)?$",
		//	  "type": "string"
		//	}
		"ml_framework": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The machine learning framework vended in the image version.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]+ ?\\d+\\.\\d+(\\.\\d+)?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Processor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates CPU or GPU compatibility.",
		//	  "enum": [
		//	    "CPU",
		//	    "GPU"
		//	  ],
		//	  "type": "string"
		//	}
		"processor": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates CPU or GPU compatibility.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CPU",
					"GPU",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProgrammingLang
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The supported programming language and its version.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z]+ ?\\d+\\.\\d+(\\.\\d+)?$",
		//	  "type": "string"
		//	}
		"programming_lang": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The supported programming language and its version.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]+ ?\\d+\\.\\d+(\\.\\d+)?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReleaseNotes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maintainer description of the image version.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"release_notes": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The maintainer description of the image version.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VendorGuidance
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The availability of the image version specified by the maintainer.",
		//	  "enum": [
		//	    "NOT_PROVIDED",
		//	    "STABLE",
		//	    "TO_BE_ARCHIVED",
		//	    "ARCHIVED"
		//	  ],
		//	  "type": "string"
		//	}
		"vendor_guidance": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The availability of the image version specified by the maintainer.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"NOT_PROVIDED",
					"STABLE",
					"TO_BE_ARCHIVED",
					"ARCHIVED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of the image version.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The version number of the image version.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::SageMaker::ImageVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ImageVersion").WithTerraformTypeName("awscc_sagemaker_image_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":             "Alias",
		"aliases":           "Aliases",
		"base_image":        "BaseImage",
		"container_image":   "ContainerImage",
		"horovod":           "Horovod",
		"image_arn":         "ImageArn",
		"image_name":        "ImageName",
		"image_version_arn": "ImageVersionArn",
		"job_type":          "JobType",
		"ml_framework":      "MLFramework",
		"processor":         "Processor",
		"programming_lang":  "ProgrammingLang",
		"release_notes":     "ReleaseNotes",
		"vendor_guidance":   "VendorGuidance",
		"version":           "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Aliases",
		"/properties/Alias",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
