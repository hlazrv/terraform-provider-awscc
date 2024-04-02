// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lightsail_container", containerResource)
}

// containerResource returns the Terraform awscc_lightsail_container resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lightsail::Container resource.
func containerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContainerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"container_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContainerServiceDeployment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes a container deployment configuration of an Amazon Lightsail container service.",
		//	  "properties": {
		//	    "Containers": {
		//	      "description": "An object that describes the configuration for the containers of the deployment.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.",
		//	        "properties": {
		//	          "Command": {
		//	            "description": "The launch command for the container.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "ContainerName": {
		//	            "description": "The name of the container.",
		//	            "type": "string"
		//	          },
		//	          "Environment": {
		//	            "description": "The environment variables of the container.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Value": {
		//	                  "type": "string"
		//	                },
		//	                "Variable": {
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "Image": {
		//	            "description": "The name of the image used for the container.",
		//	            "type": "string"
		//	          },
		//	          "Ports": {
		//	            "description": "The open firewall ports of the container.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Port": {
		//	                  "type": "string"
		//	                },
		//	                "Protocol": {
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "PublicEndpoint": {
		//	      "additionalProperties": false,
		//	      "description": "An object that describes the endpoint of the deployment.",
		//	      "properties": {
		//	        "ContainerName": {
		//	          "description": "The name of the container for the endpoint.",
		//	          "type": "string"
		//	        },
		//	        "ContainerPort": {
		//	          "description": "The port of the container to which traffic is forwarded to.",
		//	          "type": "integer"
		//	        },
		//	        "HealthCheckConfig": {
		//	          "additionalProperties": false,
		//	          "description": "An object that describes the health check configuration of the container.",
		//	          "properties": {
		//	            "HealthyThreshold": {
		//	              "description": "The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.",
		//	              "type": "integer"
		//	            },
		//	            "IntervalSeconds": {
		//	              "description": "The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.",
		//	              "type": "integer"
		//	            },
		//	            "Path": {
		//	              "description": "The path on the container on which to perform the health check. The default value is /.",
		//	              "type": "string"
		//	            },
		//	            "SuccessCodes": {
		//	              "description": "The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).",
		//	              "type": "string"
		//	            },
		//	            "TimeoutSeconds": {
		//	              "description": "The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.",
		//	              "type": "integer"
		//	            },
		//	            "UnhealthyThreshold": {
		//	              "description": "The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.",
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"container_service_deployment": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Containers
				"containers": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Command
							"command": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The launch command for the container.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ContainerName
							"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the container.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Environment
							"environment": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Value
										"value": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Variable
										"variable": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "The environment variables of the container.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Image
							"image": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the image used for the container.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Ports
							"ports": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Port
										"port": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Protocol
										"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "The open firewall ports of the container.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An object that describes the configuration for the containers of the deployment.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PublicEndpoint
				"public_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContainerName
						"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the container for the endpoint.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContainerPort
						"container_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The port of the container to which traffic is forwarded to.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HealthCheckConfig
						"health_check_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: HealthyThreshold
								"healthy_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: IntervalSeconds
								"interval_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Path
								"path": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The path on the container on which to perform the health check. The default value is /.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SuccessCodes
								"success_codes": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: TimeoutSeconds
								"timeout_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: UnhealthyThreshold
								"unhealthy_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "An object that describes the health check configuration of the container.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that describes the endpoint of the deployment.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes a container deployment configuration of an Amazon Lightsail container service.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsDisabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean value to indicate whether the container service is disabled.",
		//	  "type": "boolean"
		//	}
		"is_disabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean value to indicate whether the container service is disabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Power
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The power specification for the container service.",
		//	  "type": "string"
		//	}
		"power": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The power specification for the container service.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrincipalArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The principal ARN of the container service.",
		//	  "type": "string"
		//	}
		"principal_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The principal ARN of the container service.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateRegistryAccess
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A Boolean value to indicate whether the container service has access to private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.",
		//	  "properties": {
		//	    "EcrImagePullerRole": {
		//	      "additionalProperties": false,
		//	      "description": "An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.",
		//	      "properties": {
		//	        "IsActive": {
		//	          "description": "A Boolean value that indicates whether to activate the role.",
		//	          "type": "boolean"
		//	        },
		//	        "PrincipalArn": {
		//	          "description": "The Amazon Resource Name (ARN) of the role, if it is activated.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"private_registry_access": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EcrImagePullerRole
				"ecr_image_puller_role": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IsActive
						"is_active": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "A Boolean value that indicates whether to activate the role.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: PrincipalArn
						"principal_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the role, if it is activated.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A Boolean value to indicate whether the container service has access to private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PublicDomainNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The public domain names to use with the container service, such as example.com and www.example.com.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The public domain name to use with the container service, such as example.com and www.example.com.",
		//	    "properties": {
		//	      "CertificateName": {
		//	        "type": "string"
		//	      },
		//	      "DomainNames": {
		//	        "description": "An object that describes the configuration for the containers of the deployment.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"public_domain_names": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CertificateName
					"certificate_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DomainNames
					"domain_names": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "An object that describes the configuration for the containers of the deployment.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The public domain names to use with the container service, such as example.com and www.example.com.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scale
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scale specification for the container service.",
		//	  "maximum": 20,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"scale": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The scale specification for the container service.",
			Required:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 20),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the container service.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9]{1,2}|[a-z0-9][a-z0-9-]+[a-z0-9]$",
		//	  "type": "string"
		//	}
		"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the container service.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9]{1,2}|[a-z0-9][a-z0-9-]+[a-z0-9]$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Url
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The publicly accessible URL of the container service.",
		//	  "type": "string"
		//	}
		"url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The publicly accessible URL of the container service.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Lightsail::Container",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Container").WithTerraformTypeName("awscc_lightsail_container")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_name":             "CertificateName",
		"command":                      "Command",
		"container_arn":                "ContainerArn",
		"container_name":               "ContainerName",
		"container_port":               "ContainerPort",
		"container_service_deployment": "ContainerServiceDeployment",
		"containers":                   "Containers",
		"domain_names":                 "DomainNames",
		"ecr_image_puller_role":        "EcrImagePullerRole",
		"environment":                  "Environment",
		"health_check_config":          "HealthCheckConfig",
		"healthy_threshold":            "HealthyThreshold",
		"image":                        "Image",
		"interval_seconds":             "IntervalSeconds",
		"is_active":                    "IsActive",
		"is_disabled":                  "IsDisabled",
		"key":                          "Key",
		"path":                         "Path",
		"port":                         "Port",
		"ports":                        "Ports",
		"power":                        "Power",
		"principal_arn":                "PrincipalArn",
		"private_registry_access":      "PrivateRegistryAccess",
		"protocol":                     "Protocol",
		"public_domain_names":          "PublicDomainNames",
		"public_endpoint":              "PublicEndpoint",
		"scale":                        "Scale",
		"service_name":                 "ServiceName",
		"success_codes":                "SuccessCodes",
		"tags":                         "Tags",
		"timeout_seconds":              "TimeoutSeconds",
		"unhealthy_threshold":          "UnhealthyThreshold",
		"url":                          "Url",
		"value":                        "Value",
		"variable":                     "Variable",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
