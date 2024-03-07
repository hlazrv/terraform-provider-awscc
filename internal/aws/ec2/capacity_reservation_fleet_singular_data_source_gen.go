// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_capacity_reservation_fleet", capacityReservationFleetDataSource)
}

// capacityReservationFleetDataSource returns the Terraform awscc_ec2_capacity_reservation_fleet data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::CapacityReservationFleet resource.
func capacityReservationFleetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocationStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"allocation_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CapacityReservationFleetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"capacity_reservation_fleet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EndDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"end_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceMatchCriteria
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "open"
		//	  ],
		//	  "type": "string"
		//	}
		"instance_match_criteria": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceTypeSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AvailabilityZone": {
		//	        "type": "string"
		//	      },
		//	      "AvailabilityZoneId": {
		//	        "type": "string"
		//	      },
		//	      "EbsOptimized": {
		//	        "type": "boolean"
		//	      },
		//	      "InstancePlatform": {
		//	        "type": "string"
		//	      },
		//	      "InstanceType": {
		//	        "type": "string"
		//	      },
		//	      "Priority": {
		//	        "maximum": 999,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "Weight": {
		//	        "type": "number"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"instance_type_specifications": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AvailabilityZone
					"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: AvailabilityZoneId
					"availability_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: EbsOptimized
					"ebs_optimized": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: InstancePlatform
					"instance_platform": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: InstanceType
					"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Priority
					"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Weight
					"weight": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NoRemoveEndDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"no_remove_end_date": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RemoveEndDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"remove_end_date": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TagSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ResourceType": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Value",
		//	            "Key"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tag_specifications": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ResourceType
					"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tenancy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "default"
		//	  ],
		//	  "type": "string"
		//	}
		"tenancy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TotalTargetCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 25000,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"total_target_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::CapacityReservationFleet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CapacityReservationFleet").WithTerraformTypeName("awscc_ec2_capacity_reservation_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_strategy":           "AllocationStrategy",
		"availability_zone":             "AvailabilityZone",
		"availability_zone_id":          "AvailabilityZoneId",
		"capacity_reservation_fleet_id": "CapacityReservationFleetId",
		"ebs_optimized":                 "EbsOptimized",
		"end_date":                      "EndDate",
		"instance_match_criteria":       "InstanceMatchCriteria",
		"instance_platform":             "InstancePlatform",
		"instance_type":                 "InstanceType",
		"instance_type_specifications":  "InstanceTypeSpecifications",
		"key":                           "Key",
		"no_remove_end_date":            "NoRemoveEndDate",
		"priority":                      "Priority",
		"remove_end_date":               "RemoveEndDate",
		"resource_type":                 "ResourceType",
		"tag_specifications":            "TagSpecifications",
		"tags":                          "Tags",
		"tenancy":                       "Tenancy",
		"total_target_capacity":         "TotalTargetCapacity",
		"value":                         "Value",
		"weight":                        "Weight",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
