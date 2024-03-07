// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_traffic_distribution_group", trafficDistributionGroupDataSource)
}

// trafficDistributionGroupDataSource returns the Terraform awscc_connect_traffic_distribution_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::TrafficDistributionGroup resource.
func trafficDistributionGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the traffic distribution group.",
		//	  "maxLength": 250,
		//	  "minLength": 1,
		//	  "pattern": "(^[\\S].*[\\S]$)|(^[\\S]$)",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the traffic distribution group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance that has been replicated.",
		//	  "maxLength": 250,
		//	  "minLength": 1,
		//	  "pattern": "^arn:(aws|aws-us-gov):connect:[a-z]{2}-[a-z]+-[0-9]{1}:[0-9]{1,20}:instance/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance that has been replicated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsDefault
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If this is the default traffic distribution group.",
		//	  "type": "boolean"
		//	}
		"is_default": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If this is the default traffic distribution group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the traffic distribution group.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "(^[\\S].*[\\S]$)|(^[\\S]$)",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the traffic distribution group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the traffic distribution group.",
		//	  "enum": [
		//	    "CREATION_IN_PROGRESS",
		//	    "ACTIVE",
		//	    "CREATION_FAILED",
		//	    "PENDING_DELETION",
		//	    "DELETION_FAILED",
		//	    "UPDATE_IN_PROGRESS"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the traffic distribution group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
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
		//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficDistributionGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the traffic distribution group.",
		//	  "pattern": "^arn:(aws|aws-us-gov):connect:[a-z]{2}-[a-z]+-[0-9]{1}:[0-9]{1,20}:traffic-distribution-group/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"traffic_distribution_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the traffic distribution group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::TrafficDistributionGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::TrafficDistributionGroup").WithTerraformTypeName("awscc_connect_traffic_distribution_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                    "Description",
		"instance_arn":                   "InstanceArn",
		"is_default":                     "IsDefault",
		"key":                            "Key",
		"name":                           "Name",
		"status":                         "Status",
		"tags":                           "Tags",
		"traffic_distribution_group_arn": "TrafficDistributionGroupArn",
		"value":                          "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
