// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_secretsmanager_secret", secretDataSource)
}

// secretDataSource returns the Terraform awscc_secretsmanager_secret data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecretsManager::Secret resource.
func secretDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the secret.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GenerateSecretString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.\n We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.",
		//	  "properties": {
		//	    "ExcludeCharacters": {
		//	      "description": "A string of the characters that you don't want in the password.",
		//	      "type": "string"
		//	    },
		//	    "ExcludeLowercase": {
		//	      "description": "Specifies whether to exclude lowercase letters from the password. If you don't include this switch, the password can contain lowercase letters.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludeNumbers": {
		//	      "description": "Specifies whether to exclude numbers from the password. If you don't include this switch, the password can contain numbers.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludePunctuation": {
		//	      "description": "Specifies whether to exclude the following punctuation characters from the password: ``! \" # $ % \u0026 ' ( ) * + , - . / : ; \u003c = \u003e ? @ [ \\ ] ^ _ ` { | } ~``. If you don't include this switch, the password can contain punctuation.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludeUppercase": {
		//	      "description": "Specifies whether to exclude uppercase letters from the password. If you don't include this switch, the password can contain uppercase letters.",
		//	      "type": "boolean"
		//	    },
		//	    "GenerateStringKey": {
		//	      "description": "The JSON key name for the key/value pair, where the value is the generated password. This pair is added to the JSON structure specified by the ``SecretStringTemplate`` parameter. If you specify this parameter, then you must also specify ``SecretStringTemplate``.",
		//	      "type": "string"
		//	    },
		//	    "IncludeSpace": {
		//	      "description": "Specifies whether to include the space character. If you include this switch, the password can contain space characters.",
		//	      "type": "boolean"
		//	    },
		//	    "PasswordLength": {
		//	      "description": "The length of the password. If you don't include this parameter, the default length is 32 characters.",
		//	      "type": "integer"
		//	    },
		//	    "RequireEachIncludedType": {
		//	      "description": "Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation. If you don't include this switch, the password contains at least one of every character type.",
		//	      "type": "boolean"
		//	    },
		//	    "SecretStringTemplate": {
		//	      "description": "A template that the generated string must match. When you make a change to this property, a new secret version is created.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"generate_secret_string": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExcludeCharacters
				"exclude_characters": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A string of the characters that you don't want in the password.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeLowercase
				"exclude_lowercase": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to exclude lowercase letters from the password. If you don't include this switch, the password can contain lowercase letters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeNumbers
				"exclude_numbers": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to exclude numbers from the password. If you don't include this switch, the password can contain numbers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludePunctuation
				"exclude_punctuation": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to exclude the following punctuation characters from the password: ``! \" # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ` { | } ~``. If you don't include this switch, the password can contain punctuation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExcludeUppercase
				"exclude_uppercase": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to exclude uppercase letters from the password. If you don't include this switch, the password can contain uppercase letters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GenerateStringKey
				"generate_string_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON key name for the key/value pair, where the value is the generated password. This pair is added to the JSON structure specified by the ``SecretStringTemplate`` parameter. If you specify this parameter, then you must also specify ``SecretStringTemplate``.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IncludeSpace
				"include_space": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to include the space character. If you include this switch, the password can contain space characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PasswordLength
				"password_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The length of the password. If you don't include this parameter, the default length is 32 characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RequireEachIncludedType
				"require_each_included_type": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation. If you don't include this switch, the password contains at least one of every character type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretStringTemplate
				"secret_string_template": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A template that the generated string must match. When you make a change to this property, a new secret version is created.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.\n We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).\n To use a KMS key in a different account, use the key ARN or the alias ARN.\n If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.\n If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).\n To use a KMS key in a different account, use the key ARN or the alias ARN.\n If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.\n If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the new secret.\n The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-\n Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the new secret.\n The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-\n Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicaRegions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.",
		//	    "properties": {
		//	      "KmsKeyId": {
		//	        "description": "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.",
		//	        "type": "string"
		//	      },
		//	      "Region": {
		//	        "description": "A string that represents a ``Region``, for example \"us-east-1\".",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Region"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"replica_regions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: KmsKeyId
					"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Region
					"region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string that represents a ``Region``, for example \"us-east-1\".",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.",
		//	  "type": "string"
		//	}
		"secret_string": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:\n  ``[{\"Key\":\"CostCenter\",\"Value\":\"12345\"},{\"Key\":\"environment\",\"Value\":\"production\"}]`` \n Secrets Manager tag key names are case sensitive. A tag with the key \"ABC\" is a different tag from one with key \"abc\".\n Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret. \n If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A structure that contains information about a tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key identifier, or name, of the tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The string value associated with the key of the tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key identifier, or name, of the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The string value associated with the key of the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:\n  ``[{\"Key\":\"CostCenter\",\"Value\":\"12345\"},{\"Key\":\"environment\",\"Value\":\"production\"}]`` \n Secrets Manager tag key names are case sensitive. A tag with the key \"ABC\" is a different tag from one with key \"abc\".\n Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret. \n If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecretsManager::Secret",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecretsManager::Secret").WithTerraformTypeName("awscc_secretsmanager_secret")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                "Description",
		"exclude_characters":         "ExcludeCharacters",
		"exclude_lowercase":          "ExcludeLowercase",
		"exclude_numbers":            "ExcludeNumbers",
		"exclude_punctuation":        "ExcludePunctuation",
		"exclude_uppercase":          "ExcludeUppercase",
		"generate_secret_string":     "GenerateSecretString",
		"generate_string_key":        "GenerateStringKey",
		"id":                         "Id",
		"include_space":              "IncludeSpace",
		"key":                        "Key",
		"kms_key_id":                 "KmsKeyId",
		"name":                       "Name",
		"password_length":            "PasswordLength",
		"region":                     "Region",
		"replica_regions":            "ReplicaRegions",
		"require_each_included_type": "RequireEachIncludedType",
		"secret_string":              "SecretString",
		"secret_string_template":     "SecretStringTemplate",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
