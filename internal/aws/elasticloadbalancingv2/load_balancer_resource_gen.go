// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_elasticloadbalancingv2_load_balancer", loadBalancerResource)
}

// loadBalancerResource returns the Terraform awscc_elasticloadbalancingv2_load_balancer resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::LoadBalancer resource.
func loadBalancerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CanonicalHostedZoneID
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"canonical_hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DNSName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.",
		//	  "type": "string"
		//	}
		"enforce_security_group_inbound_rules_on_private_link_traffic": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address type. The possible values are ``ipv4`` (for IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can?t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.",
		//	  "type": "string"
		//	}
		"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address type. The possible values are ``ipv4`` (for IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can?t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "arrayType": "AttributeList",
		//	  "description": "The load balancer attributes.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies an attribute for an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ``deletion_protection.enabled`` - Indicates whether deletion protection is enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``load_balancing.cross_zone.enabled`` - Indicates whether cross-zone load balancing is enabled. The possible values are ``true`` and ``false``. The default for Network Load Balancers and Gateway Load Balancers is ``false``. The default for Application Load Balancers is ``true``, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ``access_logs.s3.enabled`` - Indicates whether access logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``access_logs.s3.bucket`` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and h",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the attribute.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"load_balancer_attributes": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ``deletion_protection.enabled`` - Indicates whether deletion protection is enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``load_balancing.cross_zone.enabled`` - Indicates whether cross-zone load balancing is enabled. The possible values are ``true`` and ``false``. The default for Network Load Balancers and Gateway Load Balancers is ``false``. The default for Application Load Balancers is ``true``, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ``access_logs.s3.enabled`` - Indicates whether access logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``access_logs.s3.bucket`` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and h",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the attribute.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The load balancer attributes.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerFullName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_full_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".\n If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".\n If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scheme
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.\n The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.\n The default is an Internet-facing load balancer.\n You cannot specify a scheme for a Gateway Load Balancer.",
		//	  "type": "string"
		//	}
		"scheme": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.\n The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.\n The default is an Internet-facing load balancer.\n You cannot specify a scheme for a Gateway Load Balancer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_groups": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a subnet for a load balancer.",
		//	    "properties": {
		//	      "AllocationId": {
		//	        "description": "[Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.",
		//	        "type": "string"
		//	      },
		//	      "IPv6Address": {
		//	        "description": "[Network Load Balancers] The IPv6 address.",
		//	        "type": "string"
		//	      },
		//	      "PrivateIPv4Address": {
		//	        "description": "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
		//	        "type": "string"
		//	      },
		//	      "SubnetId": {
		//	        "description": "The ID of the subnet.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "SubnetId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_mappings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AllocationId
					"allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: IPv6Address
					"i_pv_6_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The IPv6 address.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: PrivateIPv4Address
					"private_i_pv_4_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SubnetId
					"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the subnet.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subnets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to assign to the load balancer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Information about a tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
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
						Description: "The key of the tag.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to assign to the load balancer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of load balancer. The default is ``application``.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of load balancer. The default is ``application``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::LoadBalancer").WithTerraformTypeName("awscc_elasticloadbalancingv2_load_balancer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_id":            "AllocationId",
		"canonical_hosted_zone_id": "CanonicalHostedZoneID",
		"dns_name":                 "DNSName",
		"enforce_security_group_inbound_rules_on_private_link_traffic": "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		"i_pv_6_address":           "IPv6Address",
		"ip_address_type":          "IpAddressType",
		"key":                      "Key",
		"load_balancer_arn":        "LoadBalancerArn",
		"load_balancer_attributes": "LoadBalancerAttributes",
		"load_balancer_full_name":  "LoadBalancerFullName",
		"load_balancer_name":       "LoadBalancerName",
		"name":                     "Name",
		"private_i_pv_4_address":   "PrivateIPv4Address",
		"scheme":                   "Scheme",
		"security_groups":          "SecurityGroups",
		"subnet_id":                "SubnetId",
		"subnet_mappings":          "SubnetMappings",
		"subnets":                  "Subnets",
		"tags":                     "Tags",
		"type":                     "Type",
		"value":                    "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
