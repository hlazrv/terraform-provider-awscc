// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotwireless_destination", destinationResourceType)
}

// destinationResourceType returns the Terraform awscc_iotwireless_destination resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::Destination resource type.
func destinationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Destination arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "Destination arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Destination description",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "Destination description",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 2048),
			},
			Optional: true,
		},
		"expression": {
			// Property: Expression
			// CloudFormation resource type schema:
			// {
			//   "description": "Destination expression",
			//   "type": "string"
			// }
			Description: "Destination expression",
			Type:        types.StringType,
			Required:    true,
		},
		"expression_type": {
			// Property: ExpressionType
			// CloudFormation resource type schema:
			// {
			//   "description": "Must be RuleName",
			//   "enum": [
			//     "RuleName",
			//     "MqttTopic"
			//   ],
			//   "type": "string"
			// }
			Description: "Must be RuleName",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique name of destination",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Unique name of destination",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 128),
			},
			Required: true,
			// Name is a force-new attribute.
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS role ARN that grants access",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "AWS role ARN that grants access",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
						Optional: true,
					},
					"value": {
						// Property: Value
						Type: types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
						Optional: true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Destination's resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::Destination").WithTerraformTypeName("awscc_iotwireless_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"description":     "Description",
		"expression":      "Expression",
		"expression_type": "ExpressionType",
		"key":             "Key",
		"name":            "Name",
		"role_arn":        "RoleArn",
		"tags":            "Tags",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotwireless_destination", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
