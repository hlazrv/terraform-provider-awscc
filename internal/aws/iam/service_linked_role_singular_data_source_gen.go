// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_service_linked_role", serviceLinkedRoleDataSource)
}

// serviceLinkedRoleDataSource returns the Terraform awscc_iam_service_linked_role data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::ServiceLinkedRole resource.
func serviceLinkedRoleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AWSServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The service principal for the AWS service to which this role is attached.",
		//	  "type": "string"
		//	}
		"aws_service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The service principal for the AWS service to which this role is attached.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomSuffix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string that you provide, which is combined with the service-provided prefix to form the complete role name.",
		//	  "type": "string"
		//	}
		"custom_suffix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string that you provide, which is combined with the service-provided prefix to form the complete role name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the role.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the role.",
		//	  "type": "string"
		//	}
		"role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::ServiceLinkedRole",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::ServiceLinkedRole").WithTerraformTypeName("awscc_iam_service_linked_role")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_service_name": "AWSServiceName",
		"custom_suffix":    "CustomSuffix",
		"description":      "Description",
		"role_name":        "RoleName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
