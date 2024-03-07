// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package internetmonitor

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
	registry.AddResourceFactory("awscc_internetmonitor_monitor", monitorResource)
}

// monitorResource returns the Terraform awscc_internetmonitor_monitor resource.
// This Terraform resource corresponds to the CloudFormation AWS::InternetMonitor::Monitor resource.
func monitorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
		//	  "pattern": "^([0-2]\\d{3})-(0[0-9]|1[0-2])-([0-2]\\d|3[01])T([01]\\d|2[0-4]):([0-5]\\d):([0-6]\\d)((\\.\\d{3})?)Z$",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthEventsConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AvailabilityLocalHealthEventsConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "HealthScoreThreshold": {
		//	          "maximum": 100.0,
		//	          "minimum": 0.0,
		//	          "type": "number"
		//	        },
		//	        "MinTrafficImpact": {
		//	          "maximum": 100.0,
		//	          "minimum": 0.0,
		//	          "type": "number"
		//	        },
		//	        "Status": {
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "AvailabilityScoreThreshold": {
		//	      "maximum": 100.0,
		//	      "minimum": 0.0,
		//	      "type": "number"
		//	    },
		//	    "PerformanceLocalHealthEventsConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "HealthScoreThreshold": {
		//	          "maximum": 100.0,
		//	          "minimum": 0.0,
		//	          "type": "number"
		//	        },
		//	        "MinTrafficImpact": {
		//	          "maximum": 100.0,
		//	          "minimum": 0.0,
		//	          "type": "number"
		//	        },
		//	        "Status": {
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "PerformanceScoreThreshold": {
		//	      "maximum": 100.0,
		//	      "minimum": 0.0,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"health_events_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AvailabilityLocalHealthEventsConfig
				"availability_local_health_events_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HealthScoreThreshold
						"health_score_threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 100.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MinTrafficImpact
						"min_traffic_impact": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 100.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Status
						"status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"ENABLED",
									"DISABLED",
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
				// Property: AvailabilityScoreThreshold
				"availability_score_threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 100.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PerformanceLocalHealthEventsConfig
				"performance_local_health_events_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HealthScoreThreshold
						"health_score_threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 100.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: MinTrafficImpact
						"min_traffic_impact": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 100.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Status
						"status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"ENABLED",
									"DISABLED",
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
				// Property: PerformanceScoreThreshold
				"performance_score_threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 100.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InternetMeasurementsLogDelivery
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "S3Config": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BucketName": {
		//	          "minLength": 3,
		//	          "type": "string"
		//	        },
		//	        "BucketPrefix": {
		//	          "type": "string"
		//	        },
		//	        "LogDeliveryStatus": {
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"internet_measurements_log_delivery": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Config
				"s3_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtLeast(3),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: BucketPrefix
						"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: LogDeliveryStatus
						"log_delivery_status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"ENABLED",
									"DISABLED",
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
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxCityNetworksToMonitor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 500000,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"max_city_networks_to_monitor": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 500000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
		//	  "pattern": "^([0-2]\\d{3})-(0[0-9]|1[0-2])-([0-2]\\d|3[01])T([01]\\d|2[0-4]):([0-5]\\d):([0-6]\\d)((\\.\\d{3})?)Z$",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "minLength": 20,
		//	  "pattern": "^arn:.*",
		//	  "type": "string"
		//	}
		"monitor_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MonitorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_.-]+$",
		//	  "type": "string"
		//	}
		"monitor_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProcessingStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "OK",
		//	    "INACTIVE",
		//	    "COLLECTING_DATA",
		//	    "INSUFFICIENT_DATA",
		//	    "FAULT_SERVICE",
		//	    "FAULT_ACCESS_CLOUDWATCH"
		//	  ],
		//	  "type": "string"
		//	}
		"processing_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProcessingStatusInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"processing_status_info": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Resources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "pattern": "^arn:.*",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(20, 2048),
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:.*"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourcesToAdd
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"resources_to_add": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(20, 2048),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ResourcesToAdd is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ResourcesToRemove
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"resources_to_remove": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Optional:   true,
			Computed:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(20, 2048),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ResourcesToRemove is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "PENDING",
		//	    "ACTIVE",
		//	    "INACTIVE",
		//	    "ERROR"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"PENDING",
					"ACTIVE",
					"INACTIVE",
					"ERROR",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define.",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		// Property: TrafficPercentageToMonitor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"traffic_percentage_to_monitor": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 100),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
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
		Description: "Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::InternetMonitor::Monitor").WithTerraformTypeName("awscc_internetmonitor_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"availability_local_health_events_config": "AvailabilityLocalHealthEventsConfig",
		"availability_score_threshold":            "AvailabilityScoreThreshold",
		"bucket_name":                             "BucketName",
		"bucket_prefix":                           "BucketPrefix",
		"created_at":                              "CreatedAt",
		"health_events_config":                    "HealthEventsConfig",
		"health_score_threshold":                  "HealthScoreThreshold",
		"internet_measurements_log_delivery":      "InternetMeasurementsLogDelivery",
		"key":                                     "Key",
		"log_delivery_status":                     "LogDeliveryStatus",
		"max_city_networks_to_monitor":            "MaxCityNetworksToMonitor",
		"min_traffic_impact":                      "MinTrafficImpact",
		"modified_at":                             "ModifiedAt",
		"monitor_arn":                             "MonitorArn",
		"monitor_name":                            "MonitorName",
		"performance_local_health_events_config":  "PerformanceLocalHealthEventsConfig",
		"performance_score_threshold":             "PerformanceScoreThreshold",
		"processing_status":                       "ProcessingStatus",
		"processing_status_info":                  "ProcessingStatusInfo",
		"resources":                               "Resources",
		"resources_to_add":                        "ResourcesToAdd",
		"resources_to_remove":                     "ResourcesToRemove",
		"s3_config":                               "S3Config",
		"status":                                  "Status",
		"tags":                                    "Tags",
		"traffic_percentage_to_monitor":           "TrafficPercentageToMonitor",
		"value":                                   "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ResourcesToAdd",
		"/properties/ResourcesToRemove",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
