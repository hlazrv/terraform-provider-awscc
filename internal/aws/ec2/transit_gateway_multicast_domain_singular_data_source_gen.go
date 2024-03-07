// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_transit_gateway_multicast_domain", transitGatewayMulticastDomainDataSource)
}

// transitGatewayMulticastDomainDataSource returns the Terraform awscc_ec2_transit_gateway_multicast_domain data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastDomain resource.
func transitGatewayMulticastDomainDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the transit gateway multicast domain was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The time the transit gateway multicast domain was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Options
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The options for the transit gateway multicast domain.",
		//	  "properties": {
		//	    "AutoAcceptSharedAssociations": {
		//	      "description": "Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain. Valid Values: enable | disable",
		//	      "type": "string"
		//	    },
		//	    "Igmpv2Support": {
		//	      "description": "Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain. Valid Values: enable | disable",
		//	      "type": "string"
		//	    },
		//	    "StaticSourcesSupport": {
		//	      "description": "Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoAcceptSharedAssociations
				"auto_accept_shared_associations": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain. Valid Values: enable | disable",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Igmpv2Support
				"igmpv_2_support": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain. Valid Values: enable | disable",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StaticSourcesSupport
				"static_sources_support": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The options for the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the transit gateway multicast domain.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the transit gateway multicast domain.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
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
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags for the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayMulticastDomainArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the transit gateway multicast domain.",
		//	  "type": "string"
		//	}
		"transit_gateway_multicast_domain_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayMulticastDomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway multicast domain.",
		//	  "type": "string"
		//	}
		"transit_gateway_multicast_domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGatewayMulticastDomain",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastDomain").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_accept_shared_associations":      "AutoAcceptSharedAssociations",
		"creation_time":                        "CreationTime",
		"igmpv_2_support":                      "Igmpv2Support",
		"key":                                  "Key",
		"options":                              "Options",
		"state":                                "State",
		"static_sources_support":               "StaticSourcesSupport",
		"tags":                                 "Tags",
		"transit_gateway_id":                   "TransitGatewayId",
		"transit_gateway_multicast_domain_arn": "TransitGatewayMulticastDomainArn",
		"transit_gateway_multicast_domain_id":  "TransitGatewayMulticastDomainId",
		"value":                                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
