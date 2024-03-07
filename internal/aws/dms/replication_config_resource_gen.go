// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package dms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_dms_replication_config", replicationConfigResource)
}

// replicationConfigResource returns the Terraform awscc_dms_replication_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::DMS::ReplicationConfig resource.
func replicationConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ComputeConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration parameters for provisioning a AWS DMS Serverless replication",
		//	  "properties": {
		//	    "AvailabilityZone": {
		//	      "type": "string"
		//	    },
		//	    "DnsNameServers": {
		//	      "type": "string"
		//	    },
		//	    "KmsKeyId": {
		//	      "type": "string"
		//	    },
		//	    "MaxCapacityUnits": {
		//	      "type": "integer"
		//	    },
		//	    "MinCapacityUnits": {
		//	      "type": "integer"
		//	    },
		//	    "MultiAZ": {
		//	      "type": "boolean"
		//	    },
		//	    "PreferredMaintenanceWindow": {
		//	      "type": "string"
		//	    },
		//	    "ReplicationSubnetGroupId": {
		//	      "type": "string"
		//	    },
		//	    "VpcSecurityGroupIds": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaxCapacityUnits"
		//	  ],
		//	  "type": "object"
		//	}
		"compute_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AvailabilityZone
				"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DnsNameServers
				"dns_name_servers": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaxCapacityUnits
				"max_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: MinCapacityUnits
				"min_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MultiAZ
				"multi_az": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PreferredMaintenanceWindow
				"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ReplicationSubnetGroupId
				"replication_subnet_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VpcSecurityGroupIds
				"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration parameters for provisioning a AWS DMS Serverless replication",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfigArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Replication Config",
		//	  "type": "string"
		//	}
		"replication_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Replication Config",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfigIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier of replication configuration",
		//	  "type": "string"
		//	}
		"replication_config_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier of replication configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "JSON settings for Servereless replications that are provisioned using this replication configuration",
		//	  "type": "object"
		//	}
		"replication_settings": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "JSON settings for Servereless replications that are provisioned using this replication configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of AWS DMS Serverless replication to provision using this replication configuration",
		//	  "enum": [
		//	    "full-load",
		//	    "full-load-and-cdc",
		//	    "cdc"
		//	  ],
		//	  "type": "string"
		//	}
		"replication_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of AWS DMS Serverless replication to provision using this replication configuration",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"full-load",
					"full-load-and-cdc",
					"cdc",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource",
		//	  "type": "string"
		//	}
		"resource_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceEndpointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration",
		//	  "type": "string"
		//	}
		"source_endpoint_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SupplementalSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "JSON settings for specifying supplemental data",
		//	  "type": "object"
		//	}
		"supplemental_settings": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "JSON settings for specifying supplemental data",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration",
		//	  "type": "object"
		//	}
		"table_mappings": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the dataset.\u003c/p\u003e",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "\u003cp\u003eTag key.\u003c/p\u003e",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "\u003cp\u003eTag value.\u003c/p\u003e",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag key.</p>",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag value.</p>",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "<p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetEndpointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration",
		//	  "type": "string"
		//	}
		"target_endpoint_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "A replication configuration that you later provide to configure and start a AWS DMS Serverless replication",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DMS::ReplicationConfig").WithTerraformTypeName("awscc_dms_replication_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"availability_zone":             "AvailabilityZone",
		"compute_config":                "ComputeConfig",
		"dns_name_servers":              "DnsNameServers",
		"key":                           "Key",
		"kms_key_id":                    "KmsKeyId",
		"max_capacity_units":            "MaxCapacityUnits",
		"min_capacity_units":            "MinCapacityUnits",
		"multi_az":                      "MultiAZ",
		"preferred_maintenance_window":  "PreferredMaintenanceWindow",
		"replication_config_arn":        "ReplicationConfigArn",
		"replication_config_identifier": "ReplicationConfigIdentifier",
		"replication_settings":          "ReplicationSettings",
		"replication_subnet_group_id":   "ReplicationSubnetGroupId",
		"replication_type":              "ReplicationType",
		"resource_identifier":           "ResourceIdentifier",
		"source_endpoint_arn":           "SourceEndpointArn",
		"supplemental_settings":         "SupplementalSettings",
		"table_mappings":                "TableMappings",
		"tags":                          "Tags",
		"target_endpoint_arn":           "TargetEndpointArn",
		"value":                         "Value",
		"vpc_security_group_ids":        "VpcSecurityGroupIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
