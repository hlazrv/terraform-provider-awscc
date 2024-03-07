// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codegurureviewer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codegurureviewer_repository_association", repositoryAssociationDataSource)
}

// repositoryAssociationDataSource returns the Terraform awscc_codegurureviewer_repository_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeGuruReviewer::RepositoryAssociation resource.
func repositoryAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the repository association.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+",
		//	  "type": "string"
		//	}
		"association_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the repository association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^\\S(.*\\S)?$",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConnectionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+",
		//	  "type": "string"
		//	}
		"connection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the repository to be associated.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^\\S[\\w.-]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the repository to be associated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Owner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^\\S(.*\\S)?$",
		//	  "type": "string"
		//	}
		"owner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags associated with a repository association.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags associated with a repository association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of repository to be associated.",
		//	  "enum": [
		//	    "CodeCommit",
		//	    "Bitbucket",
		//	    "GitHubEnterpriseServer",
		//	    "S3Bucket"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of repository to be associated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeGuruReviewer::RepositoryAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeGuruReviewer::RepositoryAssociation").WithTerraformTypeName("awscc_codegurureviewer_repository_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_arn": "AssociationArn",
		"bucket_name":     "BucketName",
		"connection_arn":  "ConnectionArn",
		"key":             "Key",
		"name":            "Name",
		"owner":           "Owner",
		"tags":            "Tags",
		"type":            "Type",
		"value":           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
