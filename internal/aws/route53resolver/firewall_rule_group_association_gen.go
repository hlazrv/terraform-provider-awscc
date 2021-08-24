// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver

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
	registry.AddResourceTypeFactory("awscc_route53resolver_firewall_rule_group_association", firewallRuleGroupAssociationResourceType)
}

// firewallRuleGroupAssociationResourceType returns the Terraform awscc_route53resolver_firewall_rule_group_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53Resolver::FirewallRuleGroupAssociation resource type.
func firewallRuleGroupAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn",
			//   "maxLength": 600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Arn",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 600),
			},
			Computed: true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 40),
			},
			Computed: true,
		},
		"creator_request_id": {
			// Property: CreatorRequestId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the creator request.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The id of the creator request.",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			Computed: true,
		},
		"firewall_rule_group_id": {
			// Property: FirewallRuleGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupId",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Required: true,
			// FirewallRuleGroupId is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Id",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Computed: true,
		},
		"managed_owner_name": {
			// Property: ManagedOwnerName
			// CloudFormation resource type schema:
			// {
			//   "description": "ServicePrincipal",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ServicePrincipal",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
			},
			Computed: true,
		},
		"modification_time": {
			// Property: ModificationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 40),
			},
			Computed: true,
		},
		"mutation_protection": {
			// Property: MutationProtection
			// CloudFormation resource type schema:
			// {
			//   "description": "MutationProtectionStatus",
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "MutationProtectionStatus",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupAssociationName",
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupAssociationName",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 64),
			},
			Optional: true,
		},
		"priority": {
			// Property: Priority
			// CloudFormation resource type schema:
			// {
			//   "description": "Priority",
			//   "type": "integer"
			// }
			Description: "Priority",
			Type:        types.NumberType,
			Required:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			//   "enum": [
			//     "COMPLETE",
			//     "DELETING",
			//     "UPDATING",
			//     "INACTIVE_OWNER_ACCOUNT_CLOSED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status_message": {
			// Property: StatusMessage
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallDomainListAssociationStatus",
			//   "type": "string"
			// }
			Description: "FirewallDomainListAssociationStatus",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Tags",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
						Required: true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
						Required: true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "VpcId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "VpcId",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Required: true,
			// VpcId is a force-new attribute.
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallRuleGroupAssociation").WithTerraformTypeName("awscc_route53resolver_firewall_rule_group_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"creation_time":          "CreationTime",
		"creator_request_id":     "CreatorRequestId",
		"firewall_rule_group_id": "FirewallRuleGroupId",
		"id":                     "Id",
		"key":                    "Key",
		"managed_owner_name":     "ManagedOwnerName",
		"modification_time":      "ModificationTime",
		"mutation_protection":    "MutationProtection",
		"name":                   "Name",
		"priority":               "Priority",
		"status":                 "Status",
		"status_message":         "StatusMessage",
		"tags":                   "Tags",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53resolver_firewall_rule_group_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
