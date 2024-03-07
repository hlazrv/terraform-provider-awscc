// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package emrcontainers

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_emrcontainers_virtual_cluster", virtualClusterResource)
}

// virtualClusterResource returns the Terraform awscc_emrcontainers_virtual_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::EMRContainers::VirtualCluster resource.
func virtualClusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContainerProvider
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Container provider of the virtual cluster.",
		//	  "properties": {
		//	    "Id": {
		//	      "description": "The ID of the container cluster",
		//	      "maxLength": 100,
		//	      "minLength": 1,
		//	      "pattern": "^[0-9A-Za-z][A-Za-z0-9\\-_]*",
		//	      "type": "string"
		//	    },
		//	    "Info": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EksInfo": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Namespace": {
		//	              "maxLength": 63,
		//	              "minLength": 1,
		//	              "pattern": "[a-z0-9]([-a-z0-9]*[a-z0-9])?",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Namespace"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "EksInfo"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "description": "The type of the container provider",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "Id",
		//	    "Info"
		//	  ],
		//	  "type": "object"
		//	}
		"container_provider": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Id
				"id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the container cluster",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 100),
						stringvalidator.RegexMatches(regexp.MustCompile("^[0-9A-Za-z][A-Za-z0-9\\-_]*"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Info
				"info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EksInfo
						"eks_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Namespace
								"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 63),
										stringvalidator.RegexMatches(regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of the container provider",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Container provider of the virtual cluster.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the virtual cluster.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the virtual cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the virtual cluster.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[\\.\\-_/#A-Za-z0-9]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the virtual cluster.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\.\\-_/#A-Za-z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this virtual cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) for this virtual cluster.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this virtual cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Schema of AWS::EMRContainers::VirtualCluster Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMRContainers::VirtualCluster").WithTerraformTypeName("awscc_emrcontainers_virtual_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"container_provider": "ContainerProvider",
		"eks_info":           "EksInfo",
		"id":                 "Id",
		"info":               "Info",
		"key":                "Key",
		"name":               "Name",
		"namespace":          "Namespace",
		"tags":               "Tags",
		"type":               "Type",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
