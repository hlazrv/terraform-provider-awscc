---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_msk_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::MSK::Configuration
---

# awscc_msk_configuration (Resource)

Resource Type definition for AWS::MSK::Configuration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `server_properties` (String)

### Optional

- `description` (String)
- `kafka_versions_list` (List of String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_msk_configuration.example <resource ID>
```