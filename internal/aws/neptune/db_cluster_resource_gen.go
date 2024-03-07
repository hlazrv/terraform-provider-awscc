// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package neptune

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
)

func init() {
	registry.AddResourceFactory("awscc_neptune_db_cluster", dBClusterResource)
}

// dBClusterResource returns the Terraform awscc_neptune_db_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::Neptune::DBCluster resource.
func dBClusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociatedRoles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.",
		//	    "properties": {
		//	      "FeatureName": {
		//	        "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.",
		//	        "type": "string"
		//	      },
		//	      "RoleArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "RoleArn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"associated_roles": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FeatureName
					"feature_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: RoleArn
					"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZones
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"availability_zones": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BackupRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 1,
		//	  "description": "Specifies the number of days for which automatic DB snapshots are retained.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"backup_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Specifies the number of days for which automatic DB snapshots are retained.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.AtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(1),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.",
		//	  "type": "string"
		//	}
		"cluster_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CopyTagsToSnapshot
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them.",
		//	  "type": "boolean"
		//	}
		"copy_tags_to_snapshot": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$",
		//	  "type": "string"
		//	}
		"db_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBClusterParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides the name of the DB cluster parameter group.",
		//	  "type": "string"
		//	}
		"db_cluster_parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Provides the name of the DB cluster parameter group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBInstanceParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the DB parameter group to apply to all instances of the DB cluster. Used only in case of a major EngineVersion upgrade request.",
		//	  "type": "string"
		//	}
		"db_instance_parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the DB parameter group to apply to all instances of the DB cluster. Used only in case of a major EngineVersion upgrade request.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// DBInstanceParameterGroupName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DBPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port number on which the DB instances in the DB cluster accept connections. \n\nIf not specified, the default port used is `8182`. \n\nNote: `Port` property will soon be deprecated from this resource. Please update existing templates to rename it with new property `DBPort` having same functionalities.",
		//	  "type": "integer"
		//	}
		"db_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port number on which the DB instances in the DB cluster accept connections. \n\nIf not specified, the default port used is `8182`. \n\nNote: `Port` property will soon be deprecated from this resource. Please update existing templates to rename it with new property `DBPort` having same functionalities.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBSubnetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.",
		//	  "type": "string"
		//	}
		"db_subnet_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeletionProtection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.",
		//	  "type": "boolean"
		//	}
		"deletion_protection": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableCloudwatchLogsExports
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies a list of log types that are enabled for export to CloudWatch Logs.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"enable_cloudwatch_logs_exports": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Specifies a list of log types that are enabled for export to CloudWatch Logs.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The connection endpoint for the DB cluster. For example: `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the database engine version.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the database engine version.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IamAuthEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
		//	  "type": "boolean"
		//	}
		"iam_auth_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port number on which the DB cluster accepts connections. For example: `8182`.",
		//	  "type": "string"
		//	}
		"port": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The port number on which the DB cluster accepts connections. For example: `8182`.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PreferredBackupWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.",
		//	  "type": "string"
		//	}
		"preferred_backup_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PreferredMaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
		//	  "type": "string"
		//	}
		"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReadEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The reader endpoint for the DB cluster. For example: `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`",
		//	  "type": "string"
		//	}
		"read_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The reader endpoint for the DB cluster. For example: `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestoreToTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
		//	  "type": "string"
		//	}
		"restore_to_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// RestoreToTime is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: RestoreType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "full-copy",
		//	  "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
		//	  "type": "string"
		//	}
		"restore_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("full-copy"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// RestoreType is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServerlessScalingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.",
		//	  "properties": {
		//	    "MaxCapacity": {
		//	      "description": "The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.",
		//	      "maximum": 128,
		//	      "minimum": 2.5,
		//	      "type": "number"
		//	    },
		//	    "MinCapacity": {
		//	      "description": "The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.",
		//	      "maximum": 128,
		//	      "minimum": 1,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "required": [
		//	    "MinCapacity",
		//	    "MaxCapacity"
		//	  ],
		//	  "type": "object"
		//	}
		"serverless_scaling_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxCapacity
				"max_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.",
					Required:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(2.500000, 128.000000),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: MinCapacity
				"min_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.",
					Required:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(1.000000, 128.000000),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnapshotIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.\n\nAfter you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.\n\nHowever, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.",
		//	  "type": "string"
		//	}
		"snapshot_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.\n\nAfter you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.\n\nHowever, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SnapshotIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SourceDBClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
		//	  "type": "string"
		//	}
		"source_db_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SourceDBClusterIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: StorageEncrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the DB cluster is encrypted.\n\nIf you specify the `DBClusterIdentifier`, `DBSnapshotIdentifier`, or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the KmsKeyId property, you must enable encryption.\n\nIf you specify the KmsKeyId, you must enable encryption by setting StorageEncrypted to true.",
		//	  "type": "boolean"
		//	}
		"storage_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the DB cluster is encrypted.\n\nIf you specify the `DBClusterIdentifier`, `DBSnapshotIdentifier`, or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the KmsKeyId property, you must enable encryption.\n\nIf you specify the KmsKeyId, you must enable encryption by setting StorageEncrypted to true.",
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
		//	  "description": "The tags assigned to this cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags assigned to this cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UseLatestRestorableTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
		//	  "type": "boolean"
		//	}
		"use_latest_restorable_time": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// UseLatestRestorableTime is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides a list of VPC security groups that the DB cluster belongs to.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Provides a list of VPC security groups that the DB cluster belongs to.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
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
		Description: "The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Neptune::DBCluster").WithTerraformTypeName("awscc_neptune_db_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"associated_roles":                 "AssociatedRoles",
		"availability_zones":               "AvailabilityZones",
		"backup_retention_period":          "BackupRetentionPeriod",
		"cluster_resource_id":              "ClusterResourceId",
		"copy_tags_to_snapshot":            "CopyTagsToSnapshot",
		"db_cluster_identifier":            "DBClusterIdentifier",
		"db_cluster_parameter_group_name":  "DBClusterParameterGroupName",
		"db_instance_parameter_group_name": "DBInstanceParameterGroupName",
		"db_port":                          "DBPort",
		"db_subnet_group_name":             "DBSubnetGroupName",
		"deletion_protection":              "DeletionProtection",
		"enable_cloudwatch_logs_exports":   "EnableCloudwatchLogsExports",
		"endpoint":                         "Endpoint",
		"engine_version":                   "EngineVersion",
		"feature_name":                     "FeatureName",
		"iam_auth_enabled":                 "IamAuthEnabled",
		"key":                              "Key",
		"kms_key_id":                       "KmsKeyId",
		"max_capacity":                     "MaxCapacity",
		"min_capacity":                     "MinCapacity",
		"port":                             "Port",
		"preferred_backup_window":          "PreferredBackupWindow",
		"preferred_maintenance_window":     "PreferredMaintenanceWindow",
		"read_endpoint":                    "ReadEndpoint",
		"restore_to_time":                  "RestoreToTime",
		"restore_type":                     "RestoreType",
		"role_arn":                         "RoleArn",
		"serverless_scaling_configuration": "ServerlessScalingConfiguration",
		"snapshot_identifier":              "SnapshotIdentifier",
		"source_db_cluster_identifier":     "SourceDBClusterIdentifier",
		"storage_encrypted":                "StorageEncrypted",
		"tags":                             "Tags",
		"use_latest_restorable_time":       "UseLatestRestorableTime",
		"value":                            "Value",
		"vpc_security_group_ids":           "VpcSecurityGroupIds",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DBInstanceParameterGroupName",
		"/properties/RestoreToTime",
		"/properties/RestoreType",
		"/properties/SnapshotIdentifier",
		"/properties/SourceDBClusterIdentifier",
		"/properties/UseLatestRestorableTime",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(2160)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
