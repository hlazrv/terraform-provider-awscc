// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package omics

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_omics_run_group", runGroupResource)
}

// runGroupResource returns the Terraform awscc_omics_run_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::Omics::RunGroup resource.
func runGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^arn:.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 18,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9]+$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxCpus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100000,
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_cpus": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(1.000000, 100000.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100000,
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_duration": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(1.000000, 100000.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxGpus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100000,
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_gpus": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(1.000000, 100000.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxRuns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100000,
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_runs": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(1.000000, 100000.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}||\\p{M}||\\p{Z}||\\p{S}||\\p{N}||\\p{P}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map of resource tags",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Resource tag value",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of resource tags",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Definition of AWS::Omics::RunGroup Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Omics::RunGroup").WithTerraformTypeName("awscc_omics_run_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"creation_time": "CreationTime",
		"id":            "Id",
		"max_cpus":      "MaxCpus",
		"max_duration":  "MaxDuration",
		"max_gpus":      "MaxGpus",
		"max_runs":      "MaxRuns",
		"name":          "Name",
		"tags":          "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
