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
	registry.AddDataSourceFactory("awscc_iotwireless_wireless_gateway", wirelessGatewayDataSource)
}

// wirelessGatewayDataSource returns the Terraform awscc_iotwireless_wireless_gateway data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::WirelessGateway resource.
func wirelessGatewayDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn for Wireless Gateway. Returned upon successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn for Wireless Gateway. Returned upon successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of Wireless Gateway.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of Wireless Gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id for Wireless Gateway. Returned upon successful create.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id for Wireless Gateway. Returned upon successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUplinkReceivedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the most recent uplink was received.",
		//	  "type": "string"
		//	}
		"last_uplink_received_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the most recent uplink was received.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoRaWAN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.",
		//	  "properties": {
		//	    "GatewayEui": {
		//	      "pattern": "^(([0-9A-Fa-f]{2}-){7}|([0-9A-Fa-f]{2}:){7}|([0-9A-Fa-f]{2}\\s){7}|([0-9A-Fa-f]{2}){7})([0-9A-Fa-f]{2})$",
		//	      "type": "string"
		//	    },
		//	    "RfRegion": {
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "GatewayEui",
		//	    "RfRegion"
		//	  ],
		//	  "type": "object"
		//	}
		"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GatewayEui
				"gateway_eui": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RfRegion
				"rf_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Wireless Gateway.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Wireless Gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the gateway.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
			Description: "A list of key-value pairs that contain metadata for the gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThingArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
		//	  "type": "string"
		//	}
		"thing_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThingName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Thing Name. If there is a Thing created, this can be returned with a Get call.",
		//	  "type": "string"
		//	}
		"thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Thing Name. If there is a Thing created, this can be returned with a Get call.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::WirelessGateway",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessGateway").WithTerraformTypeName("awscc_iotwireless_wireless_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"gateway_eui":             "GatewayEui",
		"id":                      "Id",
		"key":                     "Key",
		"last_uplink_received_at": "LastUplinkReceivedAt",
		"lo_ra_wan":               "LoRaWAN",
		"name":                    "Name",
		"rf_region":               "RfRegion",
		"tags":                    "Tags",
		"thing_arn":               "ThingArn",
		"thing_name":              "ThingName",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
