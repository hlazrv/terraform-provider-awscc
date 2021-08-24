// Code generated by generators/resource/main.go; DO NOT EDIT.

package opsworkscm

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_opsworkscm_server", serverResourceType)
}

// serverResourceType returns the Terraform awscc_opsworkscm_server resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::OpsWorksCM::Server resource type.
func serverResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Computed: true,
		},
		"associate_public_ip_address": {
			// Property: AssociatePublicIpAddress
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			// AssociatePublicIpAddress is a force-new attribute.
		},
		"backup_id": {
			// Property: BackupId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 79,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 79),
			},
			Optional: true,
			Computed: true,
			// BackupId is a force-new attribute.
		},
		"backup_retention_count": {
			// Property: BackupRetentionCount
			// CloudFormation resource type schema:
			// {
			//   "minLength": 1,
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
		},
		"custom_certificate": {
			// Property: CustomCertificate
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2097152,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 2097152),
			},
			Optional: true,
			Computed: true,
			// CustomCertificate is a force-new attribute.
		},
		"custom_domain": {
			// Property: CustomDomain
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 253,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 253),
			},
			Optional: true,
			Computed: true,
			// CustomDomain is a force-new attribute.
		},
		"custom_private_key": {
			// Property: CustomPrivateKey
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 4096),
			},
			Optional: true,
			Computed: true,
			// CustomPrivateKey is a force-new attribute.
			// CustomPrivateKey is a write-only attribute.
		},
		"disable_automated_backup": {
			// Property: DisableAutomatedBackup
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Computed: true,
		},
		"engine": {
			// Property: Engine
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
			Computed: true,
			// Engine is a force-new attribute.
			// Engine is a write-only attribute.
		},
		"engine_attributes": {
			// Property: EngineAttributes
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 10000,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 10000,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 10000),
						},
						Optional: true,
					},
					"value": {
						// Property: Value
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 10000),
						},
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			// EngineAttributes is a write-only attribute.
		},
		"engine_model": {
			// Property: EngineModel
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
			Computed: true,
			// EngineModel is a force-new attribute.
		},
		"engine_version": {
			// Property: EngineVersion
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
			Computed: true,
			// EngineVersion is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Computed: true,
		},
		"instance_profile_arn": {
			// Property: InstanceProfileArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Required: true,
			// InstanceProfileArn is a force-new attribute.
		},
		"instance_type": {
			// Property: InstanceType
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Required: true,
			// InstanceType is a force-new attribute.
		},
		"key_pair": {
			// Property: KeyPair
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
			Computed: true,
			// KeyPair is a force-new attribute.
		},
		"preferred_backup_window": {
			// Property: PreferredBackupWindow
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
		},
		"preferred_maintenance_window": {
			// Property: PreferredMaintenanceWindow
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Optional: true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 10000,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// SecurityGroupIds is a force-new attribute.
		},
		"server_name": {
			// Property: ServerName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 40,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 40),
			},
			Optional: true,
			Computed: true,
			// ServerName is a force-new attribute.
		},
		"service_role_arn": {
			// Property: ServiceRoleArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 10000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type: types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 10000),
			},
			Required: true,
			// ServiceRoleArn is a force-new attribute.
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 10000,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// SubnetIds is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						Required: true,
					},
					"value": {
						// Property: Value
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::OpsWorksCM::Server",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpsWorksCM::Server").WithTerraformTypeName("awscc_opsworkscm_server")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"associate_public_ip_address":  "AssociatePublicIpAddress",
		"backup_id":                    "BackupId",
		"backup_retention_count":       "BackupRetentionCount",
		"custom_certificate":           "CustomCertificate",
		"custom_domain":                "CustomDomain",
		"custom_private_key":           "CustomPrivateKey",
		"disable_automated_backup":     "DisableAutomatedBackup",
		"endpoint":                     "Endpoint",
		"engine":                       "Engine",
		"engine_attributes":            "EngineAttributes",
		"engine_model":                 "EngineModel",
		"engine_version":               "EngineVersion",
		"id":                           "Id",
		"instance_profile_arn":         "InstanceProfileArn",
		"instance_type":                "InstanceType",
		"key":                          "Key",
		"key_pair":                     "KeyPair",
		"name":                         "Name",
		"preferred_backup_window":      "PreferredBackupWindow",
		"preferred_maintenance_window": "PreferredMaintenanceWindow",
		"security_group_ids":           "SecurityGroupIds",
		"server_name":                  "ServerName",
		"service_role_arn":             "ServiceRoleArn",
		"subnet_ids":                   "SubnetIds",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/CustomPrivateKey",
		"/properties/EngineAttributes",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_opsworkscm_server", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
