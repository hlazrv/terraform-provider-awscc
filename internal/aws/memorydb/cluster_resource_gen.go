// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package memorydb

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_memorydb_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_memorydb_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::MemoryDB::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ACLName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Access Control List to associate with the cluster.",
		//	  "pattern": "[a-zA-Z][a-zA-Z0-9\\-]*",
		//	  "type": "string"
		//	}
		"acl_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Access Control List to associate with the cluster.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z][a-zA-Z0-9\\-]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the cluster.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoMinorVersionUpgrade
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag that enables automatic minor version upgrade when set to true.\n\nYou cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.",
		//	  "type": "boolean"
		//	}
		"auto_minor_version_upgrade": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag that enables automatic minor version upgrade when set to true.\n\nYou cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The cluster endpoint.",
		//	  "properties": {
		//	    "Address": {
		//	      "description": "The DNS address of the primary read-write node.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "The port number that the engine is listening on. ",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cluster_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The DNS address of the primary read-write node.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The port number that the engine is listening on. ",
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The cluster endpoint.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the cluster. This value must be unique as it also serves as the cluster identifier.",
		//	  "pattern": "[a-z][a-z0-9\\-]*",
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the cluster. This value must be unique as it also serves as the cluster identifier.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[a-z][a-z0-9\\-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataTiering
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.",
		//	  "enum": [
		//	    "true",
		//	    "false"
		//	  ],
		//	  "type": "string"
		//	}
		"data_tiering": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"true",
					"false",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional description of the cluster.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional description of the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Redis engine version used by the cluster.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Redis engine version used by the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.",
		//	  "type": "string"
		//	}
		"final_snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// FinalSnapshotName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the KMS key used to encrypt the cluster.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the KMS key used to encrypt the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
		//	  "type": "string"
		//	}
		"maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NodeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The compute and memory capacity of the nodes in the cluster.",
		//	  "type": "string"
		//	}
		"node_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The compute and memory capacity of the nodes in the cluster.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: NumReplicasPerShard
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of replicas to apply to each shard. The limit is 5.",
		//	  "type": "integer"
		//	}
		"num_replicas_per_shard": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of replicas to apply to each shard. The limit is 5.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NumShards
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of shards the cluster will contain.",
		//	  "type": "integer"
		//	}
		"num_shards": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of shards the cluster will contain.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the parameter group associated with the cluster.",
		//	  "type": "string"
		//	}
		"parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the parameter group associated with the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParameterGroupStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the parameter group used by the cluster.",
		//	  "type": "string"
		//	}
		"parameter_group_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the parameter group used by the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port number on which each member of the cluster accepts connections.",
		//	  "type": "integer"
		//	}
		"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port number on which each member of the cluster accepts connections.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more Amazon VPC security groups associated with this cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "One or more Amazon VPC security groups associated with this cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnapshotArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"snapshot_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SnapshotArns is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.",
		//	  "type": "string"
		//	}
		"snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SnapshotName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SnapshotRetentionLimit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
		//	  "type": "integer"
		//	}
		"snapshot_retention_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnapshotWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.",
		//	  "type": "string"
		//	}
		"snapshot_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
		//	  "type": "string"
		//	}
		"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.",
		//	  "type": "string"
		//	}
		"sns_topic_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the cluster. For example, Available, Updating, Creating.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the cluster. For example, Available, Updating, Creating.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subnet group to be used for the cluster.",
		//	  "type": "string"
		//	}
		"subnet_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subnet group to be used for the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TLSEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
		//	  "type": "boolean"
		//	}
		"tls_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for the tag. May not be null.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value. May be null.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for the tag. May not be null.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value. May be null.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::Cluster").WithTerraformTypeName("awscc_memorydb_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acl_name":                   "ACLName",
		"address":                    "Address",
		"arn":                        "ARN",
		"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
		"cluster_endpoint":           "ClusterEndpoint",
		"cluster_name":               "ClusterName",
		"data_tiering":               "DataTiering",
		"description":                "Description",
		"engine_version":             "EngineVersion",
		"final_snapshot_name":        "FinalSnapshotName",
		"key":                        "Key",
		"kms_key_id":                 "KmsKeyId",
		"maintenance_window":         "MaintenanceWindow",
		"node_type":                  "NodeType",
		"num_replicas_per_shard":     "NumReplicasPerShard",
		"num_shards":                 "NumShards",
		"parameter_group_name":       "ParameterGroupName",
		"parameter_group_status":     "ParameterGroupStatus",
		"port":                       "Port",
		"security_group_ids":         "SecurityGroupIds",
		"snapshot_arns":              "SnapshotArns",
		"snapshot_name":              "SnapshotName",
		"snapshot_retention_limit":   "SnapshotRetentionLimit",
		"snapshot_window":            "SnapshotWindow",
		"sns_topic_arn":              "SnsTopicArn",
		"sns_topic_status":           "SnsTopicStatus",
		"status":                     "Status",
		"subnet_group_name":          "SubnetGroupName",
		"tags":                       "Tags",
		"tls_enabled":                "TLSEnabled",
		"value":                      "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SnapshotArns",
		"/properties/SnapshotName",
		"/properties/FinalSnapshotName",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
