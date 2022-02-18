// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecr_pull_through_cache_rule", pullThroughCacheRuleDataSourceType)
}

// pullThroughCacheRuleDataSourceType returns the Terraform awscc_ecr_pull_through_cache_rule data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECR::PullThroughCacheRule resource type.
func pullThroughCacheRuleDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"ecr_repository_prefix": {
			// Property: EcrRepositoryPrefix
			// CloudFormation resource type schema:
			// {
			//   "description": "The ECRRepositoryPrefix is a custom alias for upstream registry url.",
			//   "maxLength": 20,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ECRRepositoryPrefix is a custom alias for upstream registry url.",
			Type:        types.StringType,
			Computed:    true,
		},
		"upstream_registry_url": {
			// Property: UpstreamRegistryUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached",
			//   "type": "string"
			// }
			Description: "The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ECR::PullThroughCacheRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::PullThroughCacheRule").WithTerraformTypeName("awscc_ecr_pull_through_cache_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ecr_repository_prefix": "EcrRepositoryPrefix",
		"upstream_registry_url": "UpstreamRegistryUrl",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}