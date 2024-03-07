// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_athena_named_query", namedQueryDataSource)
}

// namedQueryDataSource returns the Terraform awscc_athena_named_query data source.
// This Terraform data source corresponds to the CloudFormation AWS::Athena::NamedQuery resource.
func namedQueryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Database
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The database to which the query belongs.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"database": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The database to which the query belongs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The query description.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The query description.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The query name.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The query name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NamedQueryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique ID of the query.",
		//	  "type": "string"
		//	}
		"named_query_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique ID of the query.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: QueryString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The contents of the query with all query statements.",
		//	  "maxLength": 262144,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"query_string": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The contents of the query with all query statements.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the workgroup that contains the named query.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"work_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the workgroup that contains the named query.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Athena::NamedQuery",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::NamedQuery").WithTerraformTypeName("awscc_athena_named_query")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"database":       "Database",
		"description":    "Description",
		"name":           "Name",
		"named_query_id": "NamedQueryId",
		"query_string":   "QueryString",
		"work_group":     "WorkGroup",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
