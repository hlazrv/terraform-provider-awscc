// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package greengrassv2

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	registry.AddResourceFactory("awscc_greengrassv2_component_version", componentVersionResource)
}

// componentVersionResource returns the Terraform awscc_greengrassv2_component_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::GreengrassV2::ComponentVersion resource.
func componentVersionResource(ctx context.Context) (resource.Resource, error) {
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
		// Property: ComponentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"component_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ComponentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"component_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InlineRecipe
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"inline_recipe": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// InlineRecipe is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: LambdaFunction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ComponentDependencies": {
		//	      "additionalProperties": false,
		//	      "patternProperties": {
		//	        "": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "DependencyType": {
		//	              "enum": [
		//	                "SOFT",
		//	                "HARD"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "VersionRequirement": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ComponentLambdaParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EnvironmentVariables": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "EventSources": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Topic": {
		//	                "type": "string"
		//	              },
		//	              "Type": {
		//	                "enum": [
		//	                  "PUB_SUB",
		//	                  "IOT_CORE"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "ExecArgs": {
		//	          "insertionOrder": true,
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "InputPayloadEncodingType": {
		//	          "enum": [
		//	            "json",
		//	            "binary"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "LinuxProcessParams": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ContainerParams": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Devices": {
		//	                  "insertionOrder": false,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "AddGroupOwner": {
		//	                        "type": "boolean"
		//	                      },
		//	                      "Path": {
		//	                        "type": "string"
		//	                      },
		//	                      "Permission": {
		//	                        "enum": [
		//	                          "ro",
		//	                          "rw"
		//	                        ],
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "MemorySizeInKB": {
		//	                  "type": "integer"
		//	                },
		//	                "MountROSysfs": {
		//	                  "type": "boolean"
		//	                },
		//	                "Volumes": {
		//	                  "insertionOrder": false,
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "AddGroupOwner": {
		//	                        "type": "boolean"
		//	                      },
		//	                      "DestinationPath": {
		//	                        "type": "string"
		//	                      },
		//	                      "Permission": {
		//	                        "enum": [
		//	                          "ro",
		//	                          "rw"
		//	                        ],
		//	                        "type": "string"
		//	                      },
		//	                      "SourcePath": {
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "IsolationMode": {
		//	              "enum": [
		//	                "GreengrassContainer",
		//	                "NoContainer"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "MaxIdleTimeInSeconds": {
		//	          "type": "integer"
		//	        },
		//	        "MaxInstancesCount": {
		//	          "type": "integer"
		//	        },
		//	        "MaxQueueSize": {
		//	          "type": "integer"
		//	        },
		//	        "Pinned": {
		//	          "type": "boolean"
		//	        },
		//	        "StatusTimeoutInSeconds": {
		//	          "type": "integer"
		//	        },
		//	        "TimeoutInSeconds": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ComponentName": {
		//	      "type": "string"
		//	    },
		//	    "ComponentPlatforms": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attributes": {
		//	            "additionalProperties": false,
		//	            "patternProperties": {
		//	              "": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "ComponentVersion": {
		//	      "type": "string"
		//	    },
		//	    "LambdaArn": {
		//	      "pattern": "^arn:[^:]*:lambda:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lambda_function": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ComponentDependencies
				"component_dependencies":  // Pattern: ""
				schema.MapNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DependencyType
							"dependency_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"SOFT",
										"HARD",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: VersionRequirement
							"version_requirement": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComponentLambdaParameters
				"component_lambda_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnvironmentVariables
						"environment_variables": // Pattern: ""
						schema.MapAttribute{     /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: EventSources
						"event_sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Topic
									"topic": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"PUB_SUB",
												"IOT_CORE",
											),
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
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ExecArgs
						"exec_args": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: InputPayloadEncodingType
						"input_payload_encoding_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"json",
									"binary",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: LinuxProcessParams
						"linux_process_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ContainerParams
								"container_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Devices
										"devices": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: AddGroupOwner
													"add_group_owner": schema.BoolAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
															boolplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: Path
													"path": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
															stringplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: Permission
													"permission": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														Validators: []validator.String{ /*START VALIDATORS*/
															stringvalidator.OneOf(
																"ro",
																"rw",
															),
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
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: MemorySizeInKB
										"memory_size_in_kb": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: MountROSysfs
										"mount_ro_sysfs": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
												boolplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Volumes
										"volumes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: AddGroupOwner
													"add_group_owner": schema.BoolAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
															boolplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: DestinationPath
													"destination_path": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
															stringplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: Permission
													"permission": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														Validators: []validator.String{ /*START VALIDATORS*/
															stringvalidator.OneOf(
																"ro",
																"rw",
															),
														}, /*END VALIDATORS*/
														PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
															stringplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
													// Property: SourcePath
													"source_path": schema.StringAttribute{ /*START ATTRIBUTE*/
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
															stringplanmodifier.UseStateForUnknown(),
														}, /*END PLAN MODIFIERS*/
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
											Optional:   true,
											Computed:   true,
											PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
												listplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: IsolationMode
								"isolation_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"GreengrassContainer",
											"NoContainer",
										),
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
						// Property: MaxIdleTimeInSeconds
						"max_idle_time_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MaxInstancesCount
						"max_instances_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MaxQueueSize
						"max_queue_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Pinned
						"pinned": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: StatusTimeoutInSeconds
						"status_timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: TimeoutInSeconds
						"timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComponentName
				"component_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComponentPlatforms
				"component_platforms": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attributes
							"attributes":        // Pattern: ""
							schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
									mapplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
					Optional:   true,
					Computed:   true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComponentVersion
				"component_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LambdaArn
				"lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:[^:]*:lambda:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), ""),
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
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// LambdaFunction is a write-only property.
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

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource for Greengrass component version.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GreengrassV2::ComponentVersion").WithTerraformTypeName("awscc_greengrassv2_component_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_group_owner":             "AddGroupOwner",
		"arn":                         "Arn",
		"attributes":                  "Attributes",
		"component_dependencies":      "ComponentDependencies",
		"component_lambda_parameters": "ComponentLambdaParameters",
		"component_name":              "ComponentName",
		"component_platforms":         "ComponentPlatforms",
		"component_version":           "ComponentVersion",
		"container_params":            "ContainerParams",
		"dependency_type":             "DependencyType",
		"destination_path":            "DestinationPath",
		"devices":                     "Devices",
		"environment_variables":       "EnvironmentVariables",
		"event_sources":               "EventSources",
		"exec_args":                   "ExecArgs",
		"inline_recipe":               "InlineRecipe",
		"input_payload_encoding_type": "InputPayloadEncodingType",
		"isolation_mode":              "IsolationMode",
		"lambda_arn":                  "LambdaArn",
		"lambda_function":             "LambdaFunction",
		"linux_process_params":        "LinuxProcessParams",
		"max_idle_time_in_seconds":    "MaxIdleTimeInSeconds",
		"max_instances_count":         "MaxInstancesCount",
		"max_queue_size":              "MaxQueueSize",
		"memory_size_in_kb":           "MemorySizeInKB",
		"mount_ro_sysfs":              "MountROSysfs",
		"name":                        "Name",
		"path":                        "Path",
		"permission":                  "Permission",
		"pinned":                      "Pinned",
		"source_path":                 "SourcePath",
		"status_timeout_in_seconds":   "StatusTimeoutInSeconds",
		"tags":                        "Tags",
		"timeout_in_seconds":          "TimeoutInSeconds",
		"topic":                       "Topic",
		"type":                        "Type",
		"version_requirement":         "VersionRequirement",
		"volumes":                     "Volumes",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/LambdaFunction",
		"/properties/InlineRecipe",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
