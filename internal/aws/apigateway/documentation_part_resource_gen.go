// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_documentation_part", documentationPartResource)
}

// documentationPartResource returns the Terraform awscc_apigateway_documentation_part resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::DocumentationPart resource.
func documentationPartResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DocumentationPartId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"documentation_part_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The location of the targeted API entity of the to-be-created documentation part.",
		//	  "properties": {
		//	    "Method": {
		//	      "description": "The HTTP verb of a method. It is a valid field for the API entity types of ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``*`` for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``method`` attribute must match that of the parent entity exactly.",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "The name of the targeted API entity. It is a valid and required field for the API entity types of ``AUTHORIZER``, ``MODEL``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY`` and ``RESPONSE_HEADER``. It is an invalid field for any other entity type.",
		//	      "type": "string"
		//	    },
		//	    "Path": {
		//	      "description": "The URL path of the target. It is a valid field for the API entity types of ``RESOURCE``, ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``/`` for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``path`` attribute must match that of the parent entity as a prefix.",
		//	      "type": "string"
		//	    },
		//	    "StatusCode": {
		//	      "description": "The HTTP status code of a response. It is a valid field for the API entity types of ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``*`` for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``statusCode`` attribute must match that of the parent entity exactly.",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "The type of API entity to which the documentation content applies. Valid values are ``API``, ``AUTHORIZER``, ``MODEL``, ``RESOURCE``, ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. Content inheritance does not apply to any entity of the ``API``, ``AUTHORIZER``, ``METHOD``, ``MODEL``, ``REQUEST_BODY``, or ``RESOURCE`` type.",
		//	      "enum": [
		//	        "API",
		//	        "AUTHORIZER",
		//	        "MODEL",
		//	        "RESOURCE",
		//	        "METHOD",
		//	        "PATH_PARAMETER",
		//	        "QUERY_PARAMETER",
		//	        "REQUEST_HEADER",
		//	        "REQUEST_BODY",
		//	        "RESPONSE",
		//	        "RESPONSE_HEADER",
		//	        "RESPONSE_BODY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Method
				"method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP verb of a method. It is a valid field for the API entity types of ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``*`` for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``method`` attribute must match that of the parent entity exactly.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the targeted API entity. It is a valid and required field for the API entity types of ``AUTHORIZER``, ``MODEL``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY`` and ``RESPONSE_HEADER``. It is an invalid field for any other entity type.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Path
				"path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URL path of the target. It is a valid field for the API entity types of ``RESOURCE``, ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``/`` for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``path`` attribute must match that of the parent entity as a prefix.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StatusCode
				"status_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP status code of a response. It is a valid field for the API entity types of ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. The default value is ``*`` for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other ``location`` attributes, the child entity's ``statusCode`` attribute must match that of the parent entity exactly.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of API entity to which the documentation content applies. Valid values are ``API``, ``AUTHORIZER``, ``MODEL``, ``RESOURCE``, ``METHOD``, ``PATH_PARAMETER``, ``QUERY_PARAMETER``, ``REQUEST_HEADER``, ``REQUEST_BODY``, ``RESPONSE``, ``RESPONSE_HEADER``, and ``RESPONSE_BODY``. Content inheritance does not apply to any entity of the ``API``, ``AUTHORIZER``, ``METHOD``, ``MODEL``, ``REQUEST_BODY``, or ``RESOURCE`` type.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"API",
							"AUTHORIZER",
							"MODEL",
							"RESOURCE",
							"METHOD",
							"PATH_PARAMETER",
							"QUERY_PARAMETER",
							"REQUEST_HEADER",
							"REQUEST_BODY",
							"RESPONSE",
							"RESPONSE_HEADER",
							"RESPONSE_BODY",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The location of the targeted API entity of the to-be-created documentation part.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Properties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.",
		//	  "type": "string"
		//	}
		"properties": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string identifier of the associated RestApi.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string identifier of the associated RestApi.",
			Required:    true,
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
		Description: "The ``AWS::ApiGateway::DocumentationPart`` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationPart").WithTerraformTypeName("awscc_apigateway_documentation_part")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"documentation_part_id": "DocumentationPartId",
		"location":              "Location",
		"method":                "Method",
		"name":                  "Name",
		"path":                  "Path",
		"properties":            "Properties",
		"rest_api_id":           "RestApiId",
		"status_code":           "StatusCode",
		"type":                  "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
