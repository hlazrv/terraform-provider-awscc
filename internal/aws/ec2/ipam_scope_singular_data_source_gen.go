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
	registry.AddDataSourceFactory("awscc_ec2_ipam_scope", iPAMScopeDataSource)
}

// iPAMScopeDataSource returns the Terraform awscc_ec2_ipam_scope data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAMScope resource.
func iPAMScopeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM scope.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM scope.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IpamArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
		//	  "type": "string"
		//	}
		"ipam_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the IPAM this scope is a part of.",
		//	  "type": "string"
		//	}
		"ipam_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the IPAM this scope is a part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM scope.",
		//	  "type": "string"
		//	}
		"ipam_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM scope.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines whether this scope contains publicly routable space or space for a private network",
		//	  "enum": [
		//	    "public",
		//	    "private"
		//	  ],
		//	  "type": "string"
		//	}
		"ipam_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsDefault
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Is this one of the default scopes created with the IPAM.",
		//	  "type": "boolean"
		//	}
		"is_default": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Is this one of the default scopes created with the IPAM.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PoolCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of pools that currently exist in this scope.",
		//	  "type": "integer"
		//	}
		"pool_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of pools that currently exist in this scope.",
			Computed:    true,
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
		//	      "Key",
		//	      "Value"
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
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMScope",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMScope").WithTerraformTypeName("awscc_ec2_ipam_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"description":     "Description",
		"ipam_arn":        "IpamArn",
		"ipam_id":         "IpamId",
		"ipam_scope_id":   "IpamScopeId",
		"ipam_scope_type": "IpamScopeType",
		"is_default":      "IsDefault",
		"key":             "Key",
		"pool_count":      "PoolCount",
		"tags":            "Tags",
		"value":           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
