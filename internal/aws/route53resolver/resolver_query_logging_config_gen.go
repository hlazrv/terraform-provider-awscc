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

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_route53resolver_resolver_query_logging_config", resolverQueryLoggingConfigResourceType)
}

// resolverQueryLoggingConfigResourceType returns the Terraform awscc_route53resolver_resolver_query_logging_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53Resolver::ResolverQueryLoggingConfig resource type.
func resolverQueryLoggingConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
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
		"association_count": {
			// Property: AssociationCount
			// CloudFormation resource type schema:
			// {
			//   "description": "Count",
			//   "type": "integer"
			// }
			Description: "Count",
			Type:        types.NumberType,
			Computed:    true,
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
		"destination_arn": {
			// Property: DestinationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "destination arn",
			//   "maxLength": 600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "destination arn",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 600),
			},
			Optional: true,
			Computed: true,
			// DestinationArn is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "ResourceId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResourceId",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigName",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigName",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			Optional: true,
			Computed: true,
			// Name is a force-new attribute.
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "AccountId",
			//   "maxLength": 32,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "AccountId",
			Type:        types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(12, 32),
			},
			Computed: true,
		},
		"share_status": {
			// Property: ShareStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			//   "enum": [
			//     "NOT_SHARED",
			//     "SHARED_WITH_ME",
			//     "SHARED_BY_ME"
			//   ],
			//   "type": "string"
			// }
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.",
			//   "enum": [
			//     "CREATING",
			//     "CREATED",
			//     "DELETING",
			//     "FAILED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverQueryLoggingConfig").WithTerraformTypeName("awscc_route53resolver_resolver_query_logging_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"association_count":  "AssociationCount",
		"creation_time":      "CreationTime",
		"creator_request_id": "CreatorRequestId",
		"destination_arn":    "DestinationArn",
		"id":                 "Id",
		"name":               "Name",
		"owner_id":           "OwnerId",
		"share_status":       "ShareStatus",
		"status":             "Status",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53resolver_resolver_query_logging_config", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
