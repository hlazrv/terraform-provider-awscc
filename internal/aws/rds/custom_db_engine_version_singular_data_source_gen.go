// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_custom_db_engine_version", customDBEngineVersionDataSource)
}

// customDBEngineVersionDataSource returns the Terraform awscc_rds_custom_db_engine_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::CustomDBEngineVersion resource.
func customDBEngineVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DBEngineVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the custom engine version.",
		//	  "type": "string"
		//	}
		"db_engine_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the custom engine version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatabaseInstallationFilesS3BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is `my-custom-installation-files`.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "type": "string"
		//	}
		"database_installation_files_s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is `my-custom-installation-files`.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatabaseInstallationFilesS3Prefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is `123456789012/cev1`. If this setting isn't specified, no prefix is assumed.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"database_installation_files_s3_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is `123456789012/cev1`. If this setting isn't specified, no prefix is assumed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional description of your CEV.",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional description of your CEV.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The database engine to use for your custom engine version (CEV). The only supported value is `custom-oracle-ee`.",
		//	  "maxLength": 35,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The database engine to use for your custom engine version (CEV). The only supported value is `custom-oracle-ee`.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of your CEV. The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.",
		//	  "maxLength": 60,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of your CEV. The name format is 19.customized_string . For example, a valid name is 19.my_cev1. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion is unique per customer per Region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of Amazon Machine Image (AMI) used for CEV.",
		//	  "type": "string"
		//	}
		"image_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of Amazon Machine Image (AMI) used for CEV.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KMSKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS KMS key identifier for an encrypted CEV. A symmetric KMS key is required for RDS Custom, but optional for Amazon RDS.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS KMS key identifier for an encrypted CEV. A symmetric KMS key is required for RDS Custom, but optional for Amazon RDS.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Manifest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.",
		//	  "maxLength": 51000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"manifest": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceCustomDbEngineVersionIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the source custom engine version.",
		//	  "type": "string"
		//	}
		"source_custom_db_engine_version_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the source custom engine version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "available",
		//	  "description": "The availability status to be assigned to the CEV.",
		//	  "enum": [
		//	    "available",
		//	    "inactive",
		//	    "inactive-except-restore"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The availability status to be assigned to the CEV.",
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
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
		// Property: UseAwsProvidedLatestImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether AWS provided latest image is applied automatically to the Custom Engine Version. By default, AWS provided latest image is applied automatically. This value is only applied on create.",
		//	  "type": "boolean"
		//	}
		"use_aws_provided_latest_image": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether AWS provided latest image is applied automatically to the Custom Engine Version. By default, AWS provided latest image is applied automatically. This value is only applied on create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RDS::CustomDBEngineVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::CustomDBEngineVersion").WithTerraformTypeName("awscc_rds_custom_db_engine_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"database_installation_files_s3_bucket_name": "DatabaseInstallationFilesS3BucketName",
		"database_installation_files_s3_prefix":      "DatabaseInstallationFilesS3Prefix",
		"db_engine_version_arn":                      "DBEngineVersionArn",
		"description":                                "Description",
		"engine":                                     "Engine",
		"engine_version":                             "EngineVersion",
		"image_id":                                   "ImageId",
		"key":                                        "Key",
		"kms_key_id":                                 "KMSKeyId",
		"manifest":                                   "Manifest",
		"source_custom_db_engine_version_identifier": "SourceCustomDbEngineVersionIdentifier",
		"status":                        "Status",
		"tags":                          "Tags",
		"use_aws_provided_latest_image": "UseAwsProvidedLatestImage",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
