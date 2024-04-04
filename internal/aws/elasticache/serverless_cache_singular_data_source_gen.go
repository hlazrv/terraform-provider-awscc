// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticache_serverless_cache", serverlessCacheDataSource)
}

// serverlessCacheDataSource returns the Terraform awscc_elasticache_serverless_cache data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElastiCache::ServerlessCache resource.
func serverlessCacheDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheUsageLimits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The cache capacity limit of the Serverless Cache.",
		//	  "properties": {
		//	    "DataStorage": {
		//	      "additionalProperties": false,
		//	      "description": "The cached data capacity of the Serverless Cache.",
		//	      "properties": {
		//	        "Maximum": {
		//	          "description": "The maximum cached data capacity of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Minimum": {
		//	          "description": "The minimum cached data capacity of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Unit": {
		//	          "description": "The unit of cached data capacity of the Serverless Cache.",
		//	          "enum": [
		//	            "GB"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Unit"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ECPUPerSecond": {
		//	      "additionalProperties": false,
		//	      "description": "The ECPU per second of the Serverless Cache.",
		//	      "properties": {
		//	        "Maximum": {
		//	          "description": "The maximum ECPU per second of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Minimum": {
		//	          "description": "The minimum ECPU per second of the Serverless Cache.",
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cache_usage_limits": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataStorage
				"data_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Maximum
						"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The maximum cached data capacity of the Serverless Cache.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Minimum
						"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum cached data capacity of the Serverless Cache.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Unit
						"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The unit of cached data capacity of the Serverless Cache.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The cached data capacity of the Serverless Cache.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ECPUPerSecond
				"ecpu_per_second": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Maximum
						"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The maximum ECPU per second of the Serverless Cache.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Minimum
						"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum ECPU per second of the Serverless Cache.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The ECPU per second of the Serverless Cache.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The cache capacity limit of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The creation time of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"create_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The creation time of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DailySnapshotTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"daily_snapshot_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The address and the port.",
		//	  "properties": {
		//	    "Address": {
		//	      "description": "Endpoint address.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "Endpoint port.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint address.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint port.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The address and the port.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The engine name of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The engine name of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The final snapshot name which is taken before Serverless Cache is deleted.",
		//	  "type": "string"
		//	}
		"final_snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The final snapshot name which is taken before Serverless Cache is deleted.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FullEngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The full engine version of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"full_engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The full engine version of the Serverless Cache.",
			Computed:    true,
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
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MajorEngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The major engine version of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"major_engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The major engine version of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReaderEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The address and the port.",
		//	  "properties": {
		//	    "Address": {
		//	      "description": "Endpoint address.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "Endpoint port.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"reader_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint address.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint port.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The address and the port.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more Amazon VPC security groups associated with this Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "One or more Amazon VPC security groups associated with this Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerlessCacheName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Serverless Cache. This value must be unique.",
		//	  "type": "string"
		//	}
		"serverless_cache_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Serverless Cache. This value must be unique.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnapshotArnsToRestore
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN's of snapshot to restore Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"snapshot_arns_to_restore": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The ARN's of snapshot to restore Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnapshotRetentionLimit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The snapshot retention limit of the Serverless Cache.",
		//	  "type": "integer"
		//	}
		"snapshot_retention_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The snapshot retention limit of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subnet id's of the Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The subnet id's of the Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with Serverless Cache.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9 _\\.\\/=+:\\-@]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this Serverless Cache.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the user group.",
		//	  "type": "string"
		//	}
		"user_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the user group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ElastiCache::ServerlessCache",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::ServerlessCache").WithTerraformTypeName("awscc_elasticache_serverless_cache")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                  "Address",
		"arn":                      "ARN",
		"cache_usage_limits":       "CacheUsageLimits",
		"create_time":              "CreateTime",
		"daily_snapshot_time":      "DailySnapshotTime",
		"data_storage":             "DataStorage",
		"description":              "Description",
		"ecpu_per_second":          "ECPUPerSecond",
		"endpoint":                 "Endpoint",
		"engine":                   "Engine",
		"final_snapshot_name":      "FinalSnapshotName",
		"full_engine_version":      "FullEngineVersion",
		"key":                      "Key",
		"kms_key_id":               "KmsKeyId",
		"major_engine_version":     "MajorEngineVersion",
		"maximum":                  "Maximum",
		"minimum":                  "Minimum",
		"port":                     "Port",
		"reader_endpoint":          "ReaderEndpoint",
		"security_group_ids":       "SecurityGroupIds",
		"serverless_cache_name":    "ServerlessCacheName",
		"snapshot_arns_to_restore": "SnapshotArnsToRestore",
		"snapshot_retention_limit": "SnapshotRetentionLimit",
		"status":                   "Status",
		"subnet_ids":               "SubnetIds",
		"tags":                     "Tags",
		"unit":                     "Unit",
		"user_group_id":            "UserGroupId",
		"value":                    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
