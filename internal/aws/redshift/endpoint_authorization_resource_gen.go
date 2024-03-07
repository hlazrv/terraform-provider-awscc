// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_redshift_endpoint_authorization", endpointAuthorizationResource)
}

// endpointAuthorizationResource returns the Terraform awscc_redshift_endpoint_authorization resource.
// This Terraform resource corresponds to the CloudFormation AWS::Redshift::EndpointAuthorization resource.
func endpointAuthorizationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Account
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target AWS account ID to grant or revoke access for.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"account": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target AWS account ID to grant or revoke access for.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\d{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllowedAllVPCs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
		//	  "type": "boolean"
		//	}
		"allowed_all_vp_cs": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AllowedVPCs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The VPCs allowed access to the cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"allowed_vp_cs": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The VPCs allowed access to the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizeTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time (UTC) when the authorization was created.",
		//	  "type": "string"
		//	}
		"authorize_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time (UTC) when the authorization was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cluster identifier.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cluster identifier.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the cluster.",
		//	  "type": "string"
		//	}
		"cluster_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of Redshift-managed VPC endpoints created for the authorization.",
		//	  "type": "integer"
		//	}
		"endpoint_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of Redshift-managed VPC endpoints created for the authorization.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Force
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
		//	  "type": "boolean"
		//	}
		"force": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Force is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Grantee
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account ID of the grantee of the cluster.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"grantee": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account ID of the grantee of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Grantor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account ID of the cluster owner.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"grantor": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account ID of the cluster owner.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the authorization action.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the authorization action.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"vpc_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^vpc-[A-Za-z0-9]{1,17}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAuthorization").WithTerraformTypeName("awscc_redshift_endpoint_authorization")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":            "Account",
		"allowed_all_vp_cs":  "AllowedAllVPCs",
		"allowed_vp_cs":      "AllowedVPCs",
		"authorize_time":     "AuthorizeTime",
		"cluster_identifier": "ClusterIdentifier",
		"cluster_status":     "ClusterStatus",
		"endpoint_count":     "EndpointCount",
		"force":              "Force",
		"grantee":            "Grantee",
		"grantor":            "Grantor",
		"status":             "Status",
		"vpc_ids":            "VpcIds",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Force",
	})
	opts = opts.WithCreateTimeoutInMinutes(60).WithDeleteTimeoutInMinutes(60)

	opts = opts.WithUpdateTimeoutInMinutes(60)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
