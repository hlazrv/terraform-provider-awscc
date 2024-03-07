// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_enclave_certificate_iam_role_association", enclaveCertificateIamRoleAssociationResource)
}

// enclaveCertificateIamRoleAssociationResource returns the Terraform awscc_ec2_enclave_certificate_iam_role_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::EnclaveCertificateIamRoleAssociation resource.
func enclaveCertificateIamRoleAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
		//	  "maxLength": 1283,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$",
		//	  "type": "string"
		//	}
		"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1283),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateS3BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon S3 bucket to which the certificate was uploaded.",
		//	  "type": "string"
		//	}
		"certificate_s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon S3 bucket to which the certificate was uploaded.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateS3ObjectKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
		//	  "type": "string"
		//	}
		"certificate_s3_object_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EncryptionKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
		//	  "type": "string"
		//	}
		"encryption_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
		//	  "maxLength": 1283,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1283),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate. This association is based on Amazon Resource Names and it enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EnclaveCertificateIamRoleAssociation").WithTerraformTypeName("awscc_ec2_enclave_certificate_iam_role_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_arn":            "CertificateArn",
		"certificate_s3_bucket_name": "CertificateS3BucketName",
		"certificate_s3_object_key":  "CertificateS3ObjectKey",
		"encryption_kms_key_id":      "EncryptionKmsKeyId",
		"role_arn":                   "RoleArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
