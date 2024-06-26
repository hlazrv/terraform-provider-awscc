// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package docdbelastic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_docdbelastic_cluster", clusterDataSource)
}

// clusterDataSource returns the Terraform awscc_docdbelastic_cluster data source.
// This Terraform data source corresponds to the CloudFormation AWS::DocDBElastic::Cluster resource.
func clusterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdminUserName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"admin_user_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AdminUserPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"admin_user_password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BackupRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"backup_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"cluster_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"cluster_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*",
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredBackupWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"preferred_backup_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredMaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ShardCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"shard_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ShardCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"shard_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ShardInstanceCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"shard_instance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DocDBElastic::Cluster",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DocDBElastic::Cluster").WithTerraformTypeName("awscc_docdbelastic_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_user_name":              "AdminUserName",
		"admin_user_password":          "AdminUserPassword",
		"auth_type":                    "AuthType",
		"backup_retention_period":      "BackupRetentionPeriod",
		"cluster_arn":                  "ClusterArn",
		"cluster_endpoint":             "ClusterEndpoint",
		"cluster_name":                 "ClusterName",
		"key":                          "Key",
		"kms_key_id":                   "KmsKeyId",
		"preferred_backup_window":      "PreferredBackupWindow",
		"preferred_maintenance_window": "PreferredMaintenanceWindow",
		"shard_capacity":               "ShardCapacity",
		"shard_count":                  "ShardCount",
		"shard_instance_count":         "ShardInstanceCount",
		"subnet_ids":                   "SubnetIds",
		"tags":                         "Tags",
		"value":                        "Value",
		"vpc_security_group_ids":       "VpcSecurityGroupIds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
