// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_elasticbeanstalk_configuration_template", configurationTemplateResource)
}

// configurationTemplateResource returns the Terraform awscc_elasticbeanstalk_configuration_template resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticBeanstalk::ConfigurationTemplate resource.
func configurationTemplateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Elastic Beanstalk application to associate with this configuration template. ",
		//	  "type": "string"
		//	}
		"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Elastic Beanstalk application to associate with this configuration template. ",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional description for this configuration.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional description for this configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration. ",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// EnvironmentId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: OptionSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "arrayType": "AttributeList",
		//	  "description": "Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. ",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Namespace": {
		//	        "description": "A unique namespace that identifies the option's associated AWS resource.",
		//	        "type": "string"
		//	      },
		//	      "OptionName": {
		//	        "description": "The name of the configuration option.",
		//	        "type": "string"
		//	      },
		//	      "ResourceName": {
		//	        "description": "A unique resource name for the option setting. Use it for a time–based scaling configuration option. ",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The current value for the configuration option.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Namespace",
		//	      "OptionName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"option_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Namespace
					"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique namespace that identifies the option's associated AWS resource.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: OptionName
					"option_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the configuration option.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceName
					"resource_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique resource name for the option setting. Use it for a time–based scaling configuration option. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The current value for the configuration option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlatformArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide. ",
		//	  "type": "string"
		//	}
		"platform_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SolutionStackName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide.\n\n You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration.\n\n Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks. ",
		//	  "type": "string"
		//	}
		"solution_stack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide.\n\n You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration.\n\n Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.\n\nValues specified in OptionSettings override any values obtained from the SourceConfiguration.\n\nYou must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.\n\nConstraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name. ",
		//	  "properties": {
		//	    "ApplicationName": {
		//	      "description": "The name of the application associated with the configuration.",
		//	      "type": "string"
		//	    },
		//	    "TemplateName": {
		//	      "description": "The name of the configuration template.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "TemplateName",
		//	    "ApplicationName"
		//	  ],
		//	  "type": "object"
		//	}
		"source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApplicationName
				"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the application associated with the configuration.",
					Required:    true,
					// ApplicationName is a write-only property.
				}, /*END ATTRIBUTE*/
				// Property: TemplateName
				"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the configuration template.",
					Required:    true,
					// TemplateName is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.\n\nValues specified in OptionSettings override any values obtained from the SourceConfiguration.\n\nYou must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.\n\nConstraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration template",
		//	  "type": "string"
		//	}
		"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration template",
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
		Description: "Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::ConfigurationTemplate").WithTerraformTypeName("awscc_elasticbeanstalk_configuration_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":     "ApplicationName",
		"description":          "Description",
		"environment_id":       "EnvironmentId",
		"namespace":            "Namespace",
		"option_name":          "OptionName",
		"option_settings":      "OptionSettings",
		"platform_arn":         "PlatformArn",
		"resource_name":        "ResourceName",
		"solution_stack_name":  "SolutionStackName",
		"source_configuration": "SourceConfiguration",
		"template_name":        "TemplateName",
		"value":                "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EnvironmentId",
		"/properties/SourceConfiguration/ApplicationName",
		"/properties/SourceConfiguration/TemplateName",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
