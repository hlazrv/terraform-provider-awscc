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
	registry.AddDataSourceFactory("awscc_ec2_vpcdhcp_options_association", vPCDHCPOptionsAssociationDataSource)
}

// vPCDHCPOptionsAssociationDataSource returns the Terraform awscc_ec2_vpcdhcp_options_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VPCDHCPOptionsAssociation resource.
func vPCDHCPOptionsAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DhcpOptionsId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the DHCP options set, or default to associate no DHCP options with the VPC.",
		//	  "type": "string"
		//	}
		"dhcp_options_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the DHCP options set, or default to associate no DHCP options with the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VPCDHCPOptionsAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCDHCPOptionsAssociation").WithTerraformTypeName("awscc_ec2_vpcdhcp_options_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dhcp_options_id": "DhcpOptionsId",
		"vpc_id":          "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
