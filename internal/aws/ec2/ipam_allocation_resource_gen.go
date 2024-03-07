// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_ipam_allocation", iPAMAllocationResource)
}

// iPAMAllocationResource returns the Terraform awscc_ec2_ipam_allocation resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::IPAMAllocation resource.
func iPAMAllocationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Cidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents an IPAM custom allocation of a single IPv4 or IPv6 CIDR",
		//	  "type": "string"
		//	}
		"cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents an IPAM custom allocation of a single IPv4 or IPv6 CIDR",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolAllocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the allocation.",
		//	  "type": "string"
		//	}
		"ipam_pool_allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the allocation.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM Pool.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
		//	  "type": "integer"
		//	}
		"netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// NetmaskLength is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Schema of AWS::EC2::IPAMAllocation Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMAllocation").WithTerraformTypeName("awscc_ec2_ipam_allocation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr":                    "Cidr",
		"description":             "Description",
		"ipam_pool_allocation_id": "IpamPoolAllocationId",
		"ipam_pool_id":            "IpamPoolId",
		"netmask_length":          "NetmaskLength",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/NetmaskLength",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
