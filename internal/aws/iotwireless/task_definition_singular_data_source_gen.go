// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_task_definition", taskDefinitionDataSource)
}

// taskDefinitionDataSource returns the Terraform awscc_iotwireless_task_definition data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::TaskDefinition resource.
func taskDefinitionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "TaskDefinition arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "TaskDefinition arn. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AutoCreateTasks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
		//	  "type": "boolean"
		//	}
		"auto_create_tasks": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the new wireless gateway task definition",
		//	  "pattern": "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the new wireless gateway task definition",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoRaWANUpdateGatewayTaskEntry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The list of task definitions.",
		//	  "properties": {
		//	    "CurrentVersion": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Model": {
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "PackageVersion": {
		//	          "maxLength": 32,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Station": {
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UpdateVersion": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Model": {
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "PackageVersion": {
		//	          "maxLength": 32,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Station": {
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lo_ra_wan_update_gateway_task_entry": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CurrentVersion
				"current_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Model
						"model": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: PackageVersion
						"package_version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Station
						"station": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateVersion
				"update_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Model
						"model": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: PackageVersion
						"package_version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Station
						"station": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The list of task definitions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the new resource.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the new resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the destination.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TaskDefinitionType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A filter to list only the wireless gateway task definitions that use this task definition type",
		//	  "enum": [
		//	    "UPDATE"
		//	  ],
		//	  "type": "string"
		//	}
		"task_definition_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A filter to list only the wireless gateway task definitions that use this task definition type",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Update
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about the gateways to update.",
		//	  "properties": {
		//	    "LoRaWAN": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CurrentVersion": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Model": {
		//	              "maxLength": 4096,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "PackageVersion": {
		//	              "maxLength": 32,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Station": {
		//	              "maxLength": 4096,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "SigKeyCrc": {
		//	          "format": "int64",
		//	          "type": "integer"
		//	        },
		//	        "UpdateSignature": {
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "UpdateVersion": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Model": {
		//	              "maxLength": 4096,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "PackageVersion": {
		//	              "maxLength": 32,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Station": {
		//	              "maxLength": 4096,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UpdateDataRole": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "UpdateDataSource": {
		//	      "maxLength": 4096,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"update": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LoRaWAN
				"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CurrentVersion
						"current_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Model
								"model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: PackageVersion
								"package_version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Station
								"station": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SigKeyCrc
						"sig_key_crc": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: UpdateSignature
						"update_signature": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: UpdateVersion
						"update_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Model
								"model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: PackageVersion
								"package_version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Station
								"station": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateDataRole
				"update_data_role": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateDataSource
				"update_data_source": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about the gateways to update.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::TaskDefinition",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::TaskDefinition").WithTerraformTypeName("awscc_iotwireless_task_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                 "Arn",
		"auto_create_tasks":                   "AutoCreateTasks",
		"current_version":                     "CurrentVersion",
		"id":                                  "Id",
		"key":                                 "Key",
		"lo_ra_wan":                           "LoRaWAN",
		"lo_ra_wan_update_gateway_task_entry": "LoRaWANUpdateGatewayTaskEntry",
		"model":                               "Model",
		"name":                                "Name",
		"package_version":                     "PackageVersion",
		"sig_key_crc":                         "SigKeyCrc",
		"station":                             "Station",
		"tags":                                "Tags",
		"task_definition_type":                "TaskDefinitionType",
		"update":                              "Update",
		"update_data_role":                    "UpdateDataRole",
		"update_data_source":                  "UpdateDataSource",
		"update_signature":                    "UpdateSignature",
		"update_version":                      "UpdateVersion",
		"value":                               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
