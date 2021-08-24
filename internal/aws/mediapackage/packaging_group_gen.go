// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage

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
	registry.AddResourceTypeFactory("awscc_mediapackage_packaging_group", packagingGroupResourceType)
}

// packagingGroupResourceType returns the Terraform awscc_mediapackage_packaging_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaPackage::PackagingGroup resource type.
func packagingGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the PackagingGroup.",
			//   "type": "string"
			// }
			Description: "The ARN of the PackagingGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authorization": {
			// Property: Authorization
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CdnIdentifierSecret": {
			//       "description": "The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.",
			//       "type": "string"
			//     },
			//     "SecretsRoleArn": {
			//       "description": "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "CdnIdentifierSecret",
			//     "SecretsRoleArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cdn_identifier_secret": {
						// Property: CdnIdentifierSecret
						Description: "The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.",
						Type:        types.StringType,
						Required:    true,
					},
					"secrets_role_arn": {
						// Property: SecretsRoleArn
						Description: "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The fully qualified domain name for Assets in the PackagingGroup.",
			//   "type": "string"
			// }
			Description: "The fully qualified domain name for Assets in the PackagingGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"egress_access_logs": {
			// Property: EgressAccessLogs
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "LogGroupName": {
			//       "description": "Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.",
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						Description: "Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.",
						Type:        types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 512),
						},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the PackagingGroup.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the PackagingGroup.",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			Required: true,
			// Id is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
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
			//   "uniqueItems": true
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::MediaPackage::PackagingGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::PackagingGroup").WithTerraformTypeName("awscc_mediapackage_packaging_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"authorization":         "Authorization",
		"cdn_identifier_secret": "CdnIdentifierSecret",
		"domain_name":           "DomainName",
		"egress_access_logs":    "EgressAccessLogs",
		"id":                    "Id",
		"key":                   "Key",
		"log_group_name":        "LogGroupName",
		"secrets_role_arn":      "SecretsRoleArn",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_mediapackage_packaging_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
