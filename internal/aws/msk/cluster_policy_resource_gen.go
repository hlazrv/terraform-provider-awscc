// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package msk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_msk_cluster_policy", clusterPolicyResource)
}

// clusterPolicyResource returns the Terraform awscc_msk_cluster_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::MSK::ClusterPolicy resource.
func clusterPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ClusterArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The arn of the cluster for the resource policy.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"cluster_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The arn of the cluster for the resource policy.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CurrentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current version of the policy attached to the specified cluster",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"current_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current version of the policy attached to the specified cluster",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document containing permissions to add to the specified cluster.",
		//	  "type": "object"
		//	}
		"policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "A policy document containing permissions to add to the specified cluster.",
			Required:    true,
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
		Description: "Resource Type definition for AWS::MSK::ClusterPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::ClusterPolicy").WithTerraformTypeName("awscc_msk_cluster_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster_arn":     "ClusterArn",
		"current_version": "CurrentVersion",
		"policy":          "Policy",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
