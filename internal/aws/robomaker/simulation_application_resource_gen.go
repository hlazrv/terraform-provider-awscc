// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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
	registry.AddResourceFactory("awscc_robomaker_simulation_application", simulationApplicationResource)
}

// simulationApplicationResource returns the Terraform awscc_robomaker_simulation_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::RoboMaker::SimulationApplication resource.
func simulationApplicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CurrentRevisionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current revision id.",
		//	  "type": "string"
		//	}
		"current_revision_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current revision id.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Environment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URI of the Docker image for the robot application.",
		//	  "type": "string"
		//	}
		"environment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URI of the Docker image for the robot application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the simulation application.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_\\-]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the simulation application.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_\\-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RenderingEngine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The rendering engine for the simulation application.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of the rendering engine.",
		//	      "enum": [
		//	        "OGRE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the rendering engine.",
		//	      "pattern": "1.x",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "Version"
		//	  ],
		//	  "type": "object"
		//	}
		"rendering_engine": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the rendering engine.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"OGRE",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the rendering engine.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("1.x"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The rendering engine for the simulation application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// RenderingEngine is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: RobotSoftwareSuite
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The robot software suite used by the simulation application.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of the robot software suite.",
		//	      "enum": [
		//	        "ROS",
		//	        "ROS2",
		//	        "General"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the robot software suite.",
		//	      "enum": [
		//	        "Kinetic",
		//	        "Melodic",
		//	        "Dashing",
		//	        "Foxy"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name"
		//	  ],
		//	  "type": "object"
		//	}
		"robot_software_suite": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the robot software suite.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ROS",
							"ROS2",
							"General",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the robot software suite.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Kinetic",
							"Melodic",
							"Dashing",
							"Foxy",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// Version is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The robot software suite used by the simulation application.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SimulationSoftwareSuite
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The simulation software suite used by the simulation application.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of the simulation software suite.",
		//	      "enum": [
		//	        "Gazebo",
		//	        "RosbagPlay",
		//	        "SimulationRuntime"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the simulation software suite.",
		//	      "enum": [
		//	        "7",
		//	        "9",
		//	        "11",
		//	        "Kinetic",
		//	        "Melodic",
		//	        "Dashing",
		//	        "Foxy"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name"
		//	  ],
		//	  "type": "object"
		//	}
		"simulation_software_suite": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the simulation software suite.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Gazebo",
							"RosbagPlay",
							"SimulationRuntime",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the simulation software suite.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"7",
							"9",
							"11",
							"Kinetic",
							"Melodic",
							"Dashing",
							"Foxy",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// Version is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The simulation software suite used by the simulation application.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Sources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sources of the simulation application.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Information about a source configuration.",
		//	    "properties": {
		//	      "Architecture": {
		//	        "description": "The target processor architecture for the application.",
		//	        "enum": [
		//	          "X86_64",
		//	          "ARM64",
		//	          "ARMHF"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "S3Bucket": {
		//	        "description": "The Amazon S3 bucket name.",
		//	        "pattern": "[a-z0-9][a-z0-9.\\-]*[a-z0-9]",
		//	        "type": "string"
		//	      },
		//	      "S3Key": {
		//	        "description": "The s3 object key.",
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "S3Bucket",
		//	      "S3Key",
		//	      "Architecture"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Architecture
					"architecture": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The target processor architecture for the application.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"X86_64",
								"ARM64",
								"ARMHF",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: S3Bucket
					"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon S3 bucket name.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("[a-z0-9][a-z0-9.\\-]*[a-z0-9]"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: S3Key
					"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The s3 object key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 1024),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The sources of the simulation application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Sources is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
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
		Description: "This schema is for testing purpose only.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::SimulationApplication").WithTerraformTypeName("awscc_robomaker_simulation_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":              "Architecture",
		"arn":                       "Arn",
		"current_revision_id":       "CurrentRevisionId",
		"environment":               "Environment",
		"name":                      "Name",
		"rendering_engine":          "RenderingEngine",
		"robot_software_suite":      "RobotSoftwareSuite",
		"s3_bucket":                 "S3Bucket",
		"s3_key":                    "S3Key",
		"simulation_software_suite": "SimulationSoftwareSuite",
		"sources":                   "Sources",
		"tags":                      "Tags",
		"version":                   "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RenderingEngine",
		"/properties/RobotSoftwareSuite/Version",
		"/properties/Sources",
		"/properties/SimulationSoftwareSuite/Version",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
