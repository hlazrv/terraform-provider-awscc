// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package deadline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_deadline_queue", queueDataSource)
}

// queueDataSource returns the Terraform awscc_deadline_queue data source.
// This Terraform data source corresponds to the CloudFormation AWS::Deadline::Queue resource.
func queueDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllowedStorageProfileIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "pattern": "^sp-[0-9a-f]{32}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"allowed_storage_profile_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultBudgetAction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "NONE",
		//	  "enum": [
		//	    "NONE",
		//	    "STOP_SCHEDULING_AND_COMPLETE_TASKS",
		//	    "STOP_SCHEDULING_AND_CANCEL_TASKS"
		//	  ],
		//	  "type": "string"
		//	}
		"default_budget_action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "",
		//	  "maxLength": 100,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FarmId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^farm-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"farm_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: JobAttachmentSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "RootPrefix": {
		//	      "maxLength": 63,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "S3BucketName": {
		//	      "maxLength": 63,
		//	      "minLength": 3,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RootPrefix",
		//	    "S3BucketName"
		//	  ],
		//	  "type": "object"
		//	}
		"job_attachment_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RootPrefix
				"root_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: S3BucketName
				"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: JobRunAsUser
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Posix": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Group": {
		//	          "maxLength": 31,
		//	          "minLength": 0,
		//	          "pattern": "^(?:[a-z][a-z0-9-]{0,30})?$",
		//	          "type": "string"
		//	        },
		//	        "User": {
		//	          "maxLength": 31,
		//	          "minLength": 0,
		//	          "pattern": "^(?:[a-z][a-z0-9-]{0,30})?$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Group",
		//	        "User"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "RunAs": {
		//	      "enum": [
		//	        "QUEUE_CONFIGURED_USER",
		//	        "WORKER_AGENT_USER"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Windows": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "PasswordArn": {
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "pattern": "",
		//	          "type": "string"
		//	        },
		//	        "User": {
		//	          "maxLength": 111,
		//	          "minLength": 0,
		//	          "pattern": "^[^\"'/\\[\\]:;|=,+*?\u003c\u003e\\s]*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "PasswordArn",
		//	        "User"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "RunAs"
		//	  ],
		//	  "type": "object"
		//	}
		"job_run_as_user": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Posix
				"posix": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Group
						"group": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: User
						"user": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RunAs
				"run_as": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Windows
				"windows": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PasswordArn
						"password_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: User
						"user": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: QueueId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^queue-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"queue_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RequiredFileSystemLocationNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 1,
		//	    "pattern": "^[0-9A-Za-z ]*$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"required_file_system_location_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:(aws[a-zA-Z-]*):iam::\\d{12}:role(/[!-.0-~]+)*/[\\w+=,.@-]+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Deadline::Queue",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Deadline::Queue").WithTerraformTypeName("awscc_deadline_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_storage_profile_ids":         "AllowedStorageProfileIds",
		"arn":                                 "Arn",
		"default_budget_action":               "DefaultBudgetAction",
		"description":                         "Description",
		"display_name":                        "DisplayName",
		"farm_id":                             "FarmId",
		"group":                               "Group",
		"job_attachment_settings":             "JobAttachmentSettings",
		"job_run_as_user":                     "JobRunAsUser",
		"password_arn":                        "PasswordArn",
		"posix":                               "Posix",
		"queue_id":                            "QueueId",
		"required_file_system_location_names": "RequiredFileSystemLocationNames",
		"role_arn":                            "RoleArn",
		"root_prefix":                         "RootPrefix",
		"run_as":                              "RunAs",
		"s3_bucket_name":                      "S3BucketName",
		"user":                                "User",
		"windows":                             "Windows",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
