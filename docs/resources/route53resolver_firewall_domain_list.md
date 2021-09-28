---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53resolver_firewall_domain_list Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Route53Resolver::FirewallDomainList.
---

# awscc_route53resolver_firewall_domain_list (Resource)

Resource schema for AWS::Route53Resolver::FirewallDomainList.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **domain_file_url** (String) S3 URL to import domains from.
- **domains** (List of String) An inline list of domains to use for this domain list.
- **name** (String) FirewallDomainListName
- **tags** (Attributes Set) Tags (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) Arn
- **creation_time** (String) Rfc3339TimeString
- **creator_request_id** (String) The id of the creator request.
- **domain_count** (Number) Count
- **id** (String) ResourceId
- **managed_owner_name** (String) ServicePrincipal
- **modification_time** (String) Rfc3339TimeString
- **status** (String) ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
- **status_message** (String) FirewallDomainListAssociationStatus

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53resolver_firewall_domain_list.example <resource ID>
```