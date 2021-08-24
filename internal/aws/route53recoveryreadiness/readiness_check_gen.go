// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoveryreadiness

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
	registry.AddResourceTypeFactory("awscc_route53recoveryreadiness_readiness_check", readinessCheckResourceType)
}

// readinessCheckResourceType returns the Terraform awscc_route53recoveryreadiness_readiness_check resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::ReadinessCheck resource type.
func readinessCheckResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"readiness_check_arn": {
			// Property: ReadinessCheckArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the readiness check.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the readiness check.",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
			},
			Computed: true,
		},
		"readiness_check_name": {
			// Property: ReadinessCheckName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the ReadinessCheck to create.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the ReadinessCheck to create.",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Required: true,
			// ReadinessCheckName is a force-new attribute.
		},
		"resource_set_name": {
			// Property: ResourceSetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the resource set to check.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the resource set to check.",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxItems": 50,
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Aws Route53 Recovery Readiness Check Schema and API specification.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::ReadinessCheck").WithTerraformTypeName("awscc_route53recoveryreadiness_readiness_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                  "Key",
		"readiness_check_arn":  "ReadinessCheckArn",
		"readiness_check_name": "ReadinessCheckName",
		"resource_set_name":    "ResourceSetName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoveryreadiness_readiness_check", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
