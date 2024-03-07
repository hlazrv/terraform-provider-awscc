// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_role_policy", rolePolicyDataSource)
}

// rolePolicyDataSource returns the Terraform awscc_iam_role_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::RolePolicy resource.
func rolePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy document.",
		//	  "type": "object"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The policy document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The friendly name (not ARN) identifying the policy.",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The friendly name (not ARN) identifying the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the policy document.",
		//	  "type": "string"
		//	}
		"role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the policy document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::RolePolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::RolePolicy").WithTerraformTypeName("awscc_iam_role_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"policy_document": "PolicyDocument",
		"policy_name":     "PolicyName",
		"role_name":       "RoleName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
