// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_servicecatalogappregistry_application", applicationResource)
}

// applicationResource returns the Terraform awscc_servicecatalogappregistry_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::Application resource.
func applicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the application. ",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "\\w+",
		//	  "type": "string"
		//	}
		"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the application. ",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ApplicationTagKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The key of the AWS application tag, which is awsApplication. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value.",
		//	  "maxLength": 128,
		//	  "pattern": "\\w+",
		//	  "type": "string"
		//	}
		"application_tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The key of the AWS application tag, which is awsApplication. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ApplicationTagValue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The value of the AWS application tag, which is the identifier of an associated resource. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value. ",
		//	  "maxLength": 256,
		//	  "pattern": "\\[a-zA-Z0-9_-:/]+",
		//	  "type": "string"
		//	}
		"application_tag_value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The value of the AWS application tag, which is the identifier of an associated resource. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value. ",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/applications/[a-z0-9]+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the application. ",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the application. ",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "[a-z0-9]{26}",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the application. ",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "\\w+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the application. ",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("\\w+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Schema for AWS::ServiceCatalogAppRegistry::Application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::Application").WithTerraformTypeName("awscc_servicecatalogappregistry_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":      "ApplicationName",
		"application_tag_key":   "ApplicationTagKey",
		"application_tag_value": "ApplicationTagValue",
		"arn":                   "Arn",
		"description":           "Description",
		"id":                    "Id",
		"name":                  "Name",
		"tags":                  "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
