// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_verified_access_trust_provider", verifiedAccessTrustProviderResource)
}

// verifiedAccessTrustProviderResource returns the Terraform awscc_ec2_verified_access_trust_provider resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::VerifiedAccessTrustProvider resource.
func verifiedAccessTrustProviderResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The creation time.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The creation time.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the Amazon Web Services Verified Access trust provider.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the Amazon Web Services Verified Access trust provider.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeviceOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The options for device identity based trust providers.",
		//	  "properties": {
		//	    "PublicSigningKeyUrl": {
		//	      "description": "URL Verified Access will use to verify authenticity of the device tokens.",
		//	      "type": "string"
		//	    },
		//	    "TenantId": {
		//	      "description": "The ID of the tenant application with the device-identity provider.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"device_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PublicSigningKeyUrl
				"public_signing_key_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "URL Verified Access will use to verify authenticity of the device tokens.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TenantId
				"tenant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the tenant application with the device-identity provider.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The options for device identity based trust providers.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeviceTrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of device-based trust provider. Possible values: jamf|crowdstrike",
		//	  "type": "string"
		//	}
		"device_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of device-based trust provider. Possible values: jamf|crowdstrike",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The last updated time.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The last updated time.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OidcOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The OpenID Connect details for an oidc -type, user-identity based trust provider.",
		//	  "properties": {
		//	    "AuthorizationEndpoint": {
		//	      "description": "The OIDC authorization endpoint.",
		//	      "type": "string"
		//	    },
		//	    "ClientId": {
		//	      "description": "The client identifier.",
		//	      "type": "string"
		//	    },
		//	    "ClientSecret": {
		//	      "description": "The client secret.",
		//	      "type": "string"
		//	    },
		//	    "Issuer": {
		//	      "description": "The OIDC issuer.",
		//	      "type": "string"
		//	    },
		//	    "Scope": {
		//	      "description": "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
		//	      "type": "string"
		//	    },
		//	    "TokenEndpoint": {
		//	      "description": "The OIDC token endpoint.",
		//	      "type": "string"
		//	    },
		//	    "UserInfoEndpoint": {
		//	      "description": "The OIDC user info endpoint.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"oidc_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthorizationEndpoint
				"authorization_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC authorization endpoint.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ClientId
				"client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The client identifier.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ClientSecret
				"client_secret": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The client secret.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Issuer
				"issuer": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC issuer.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Scope
				"scope": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TokenEndpoint
				"token_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC token endpoint.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: UserInfoEndpoint
				"user_info_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC user info endpoint.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The OpenID Connect details for an oidc -type, user-identity based trust provider.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyReferenceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier to be used when working with policy rules.",
		//	  "type": "string"
		//	}
		"policy_reference_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier to be used when working with policy rules.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SseSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options for customer provided KMS encryption.",
		//	  "properties": {
		//	    "CustomerManagedKeyEnabled": {
		//	      "description": "Whether to encrypt the policy with the provided key or disable encryption",
		//	      "type": "boolean"
		//	    },
		//	    "KmsKeyArn": {
		//	      "description": "KMS Key Arn used to encrypt the group policy",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sse_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomerManagedKeyEnabled
				"customer_managed_key_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether to encrypt the policy with the provided key or disable encryption",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "KMS Key Arn used to encrypt the group policy",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options for customer provided KMS encryption.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Type of trust provider. Possible values: user|device",
		//	  "type": "string"
		//	}
		"trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Type of trust provider. Possible values: user|device",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserTrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of device-based trust provider. Possible values: oidc|iam-identity-center",
		//	  "type": "string"
		//	}
		"user_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of device-based trust provider. Possible values: oidc|iam-identity-center",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessTrustProviderId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon Web Services Verified Access trust provider.",
		//	  "type": "string"
		//	}
		"verified_access_trust_provider_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon Web Services Verified Access trust provider.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::EC2::VerifiedAccessTrustProvider type describes a verified access trust provider",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VerifiedAccessTrustProvider").WithTerraformTypeName("awscc_ec2_verified_access_trust_provider")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"authorization_endpoint":            "AuthorizationEndpoint",
		"client_id":                         "ClientId",
		"client_secret":                     "ClientSecret",
		"creation_time":                     "CreationTime",
		"customer_managed_key_enabled":      "CustomerManagedKeyEnabled",
		"description":                       "Description",
		"device_options":                    "DeviceOptions",
		"device_trust_provider_type":        "DeviceTrustProviderType",
		"issuer":                            "Issuer",
		"key":                               "Key",
		"kms_key_arn":                       "KmsKeyArn",
		"last_updated_time":                 "LastUpdatedTime",
		"oidc_options":                      "OidcOptions",
		"policy_reference_name":             "PolicyReferenceName",
		"public_signing_key_url":            "PublicSigningKeyUrl",
		"scope":                             "Scope",
		"sse_specification":                 "SseSpecification",
		"tags":                              "Tags",
		"tenant_id":                         "TenantId",
		"token_endpoint":                    "TokenEndpoint",
		"trust_provider_type":               "TrustProviderType",
		"user_info_endpoint":                "UserInfoEndpoint",
		"user_trust_provider_type":          "UserTrustProviderType",
		"value":                             "Value",
		"verified_access_trust_provider_id": "VerifiedAccessTrustProviderId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
