// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_resource_specific_logging", resourceSpecificLoggingDataSource)
}

// resourceSpecificLoggingDataSource returns the Terraform awscc_iot_resource_specific_logging data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::ResourceSpecificLogging resource.
func resourceSpecificLoggingDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LogLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
		//	  "enum": [
		//	    "ERROR",
		//	    "WARN",
		//	    "INFO",
		//	    "DEBUG",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
		//	  "maxLength": 140,
		//	  "minLength": 13,
		//	  "pattern": "[a-zA-Z0-9.:\\s_\\-]+",
		//	  "type": "string"
		//	}
		"target_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target name.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9.:\\s_\\-]+",
		//	  "type": "string"
		//	}
		"target_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID, or EVENT_TYPE.",
		//	  "enum": [
		//	    "THING_GROUP",
		//	    "CLIENT_ID",
		//	    "SOURCE_IP",
		//	    "PRINCIPAL_ID",
		//	    "EVENT_TYPE"
		//	  ],
		//	  "type": "string"
		//	}
		"target_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID, or EVENT_TYPE.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::ResourceSpecificLogging",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ResourceSpecificLogging").WithTerraformTypeName("awscc_iot_resource_specific_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_level":   "LogLevel",
		"target_id":   "TargetId",
		"target_name": "TargetName",
		"target_type": "TargetType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
