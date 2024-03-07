// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_kendra_faq", faqDataSource)
}

// faqDataSource returns the Terraform awscc_kendra_faq data source.
// This Terraform data source corresponds to the CloudFormation AWS::Kendra::Faq resource.
func faqDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FAQ description",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FAQ description",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FileFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FAQ file format",
		//	  "enum": [
		//	    "CSV",
		//	    "CSV_WITH_HEADER",
		//	    "JSON"
		//	  ],
		//	  "type": "string"
		//	}
		"file_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FAQ file format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of the FAQ",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of the FAQ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IndexId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Index ID",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "type": "string"
		//	}
		"index_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Index ID",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FAQ name",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FAQ name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FAQ role ARN",
		//	  "maxLength": 1284,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FAQ role ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "FAQ S3 path",
		//	  "properties": {
		//	    "Bucket": {
		//	      "maxLength": 63,
		//	      "minLength": 3,
		//	      "pattern": "[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]",
		//	      "type": "string"
		//	    },
		//	    "Key": {
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Bucket",
		//	    "Key"
		//	  ],
		//	  "type": "object"
		//	}
		"s3_path": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Key
				"key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "FAQ S3 path",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for labeling the FAQ",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging Kendra resources",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for the tag",
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
		//	  "maxItems": 200,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for the tag",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for labeling the FAQ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Kendra::Faq",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Faq").WithTerraformTypeName("awscc_kendra_faq")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"bucket":      "Bucket",
		"description": "Description",
		"file_format": "FileFormat",
		"id":          "Id",
		"index_id":    "IndexId",
		"key":         "Key",
		"name":        "Name",
		"role_arn":    "RoleArn",
		"s3_path":     "S3Path",
		"tags":        "Tags",
		"value":       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
