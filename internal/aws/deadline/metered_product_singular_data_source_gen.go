// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package deadline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_deadline_metered_product", meteredProductDataSource)
}

// meteredProductDataSource returns the Terraform awscc_deadline_metered_product data source.
// This Terraform data source corresponds to the CloudFormation AWS::Deadline::MeteredProduct resource.
func meteredProductDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:(aws[a-zA-Z-]*):deadline:[a-z0-9-]+:[0-9]{12}:license-endpoint/le-[0-9a-z]{32}/metered-product/[0-9a-z]{1,32}-[.0-9a-z]{1,32}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Family
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LicenseEndpointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^le-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"license_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 65535,
		//	  "minimum": 1024,
		//	  "type": "integer"
		//	}
		"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProductId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[0-9a-z]{1,32}-[.0-9a-z]{1,32}$",
		//	  "type": "string"
		//	}
		"product_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Vendor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"vendor": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Deadline::MeteredProduct",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Deadline::MeteredProduct").WithTerraformTypeName("awscc_deadline_metered_product")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"family":              "Family",
		"license_endpoint_id": "LicenseEndpointId",
		"port":                "Port",
		"product_id":          "ProductId",
		"vendor":              "Vendor",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
