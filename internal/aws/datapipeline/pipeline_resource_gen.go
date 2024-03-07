// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_datapipeline_pipeline", pipelineResource)
}

// pipelineResource returns the Terraform awscc_datapipeline_pipeline resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataPipeline::Pipeline resource.
func pipelineResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Activate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
		//	  "type": "boolean"
		//	}
		"activate": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the pipeline.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the pipeline.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the pipeline.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParameterObjects
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parameter objects used with the pipeline.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Attributes": {
		//	        "description": "The attributes of the parameter object.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "description": "The field identifier.",
		//	              "type": "string"
		//	            },
		//	            "StringValue": {
		//	              "description": "The field value, expressed as a String.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "StringValue"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "Id": {
		//	        "description": "The ID of the parameter object.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Attributes",
		//	      "Id"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"parameter_objects": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Attributes
					"attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The field identifier.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: StringValue
								"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The field value, expressed as a String.",
									Required:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "The attributes of the parameter object.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the parameter object.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The parameter objects used with the pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParameterValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parameter values used with the pipeline.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Id": {
		//	        "description": "The ID of the parameter value.",
		//	        "type": "string"
		//	      },
		//	      "StringValue": {
		//	        "description": "The field value, expressed as a String.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Id",
		//	      "StringValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"parameter_values": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the parameter value.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: StringValue
					"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The field value, expressed as a String.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The parameter values used with the pipeline.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"pipeline_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineObjects
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Fields": {
		//	        "description": "Key-value pairs that define the properties of the object.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "description": "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
		//	              "type": "string"
		//	            },
		//	            "RefValue": {
		//	              "description": "A field value that you specify as an identifier of another object in the same pipeline definition.",
		//	              "type": "string"
		//	            },
		//	            "StringValue": {
		//	              "description": "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "Id": {
		//	        "description": "The ID of the object.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "The name of the object.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Fields",
		//	      "Id",
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"pipeline_objects": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Fields
					"fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: RefValue
								"ref_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field value that you specify as an identifier of another object in the same pipeline definition.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: StringValue
								"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "Key-value pairs that define the properties of the object.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the object.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the object.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of a tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value to associate with the key name.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"pipeline_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of a tag.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value to associate with the key name.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
			Optional:    true,
			Computed:    true,
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataPipeline::Pipeline").WithTerraformTypeName("awscc_datapipeline_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activate":          "Activate",
		"attributes":        "Attributes",
		"description":       "Description",
		"fields":            "Fields",
		"id":                "Id",
		"key":               "Key",
		"name":              "Name",
		"parameter_objects": "ParameterObjects",
		"parameter_values":  "ParameterValues",
		"pipeline_id":       "PipelineId",
		"pipeline_objects":  "PipelineObjects",
		"pipeline_tags":     "PipelineTags",
		"ref_value":         "RefValue",
		"string_value":      "StringValue",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
