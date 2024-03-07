// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_request_validator", requestValidatorDataSource)
}

// requestValidatorDataSource returns the Terraform awscc_apigateway_request_validator data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::RequestValidator resource.
func requestValidatorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of this RequestValidator",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of this RequestValidator",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RequestValidatorId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"request_validator_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string identifier of the associated RestApi.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string identifier of the associated RestApi.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ValidateRequestBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean flag to indicate whether to validate a request body according to the configured Model schema.",
		//	  "type": "boolean"
		//	}
		"validate_request_body": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean flag to indicate whether to validate a request body according to the configured Model schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ValidateRequestParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).",
		//	  "type": "boolean"
		//	}
		"validate_request_parameters": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::RequestValidator",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::RequestValidator").WithTerraformTypeName("awscc_apigateway_request_validator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"name":                        "Name",
		"request_validator_id":        "RequestValidatorId",
		"rest_api_id":                 "RestApiId",
		"validate_request_body":       "ValidateRequestBody",
		"validate_request_parameters": "ValidateRequestParameters",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
