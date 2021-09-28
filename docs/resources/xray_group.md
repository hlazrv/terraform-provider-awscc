---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_xray_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This schema provides construct and validation rules for AWS-XRay Group resource parameters.
---

# awscc_xray_group (Resource)

This schema provides construct and validation rules for AWS-XRay Group resource parameters.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filter_expression** (String) The filter expression defining criteria by which to group traces.
- **group_name** (String) The case-sensitive name of the new group. Names must be unique.
- **insights_configuration** (Attributes) (see [below for nested schema](#nestedatt--insights_configuration))
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **group_arn** (String) The ARN of the group that was generated on creation.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--insights_configuration"></a>
### Nested Schema for `insights_configuration`

Optional:

- **insights_enabled** (Boolean) Set the InsightsEnabled value to true to enable insights or false to disable insights.
- **notifications_enabled** (Boolean) Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_xray_group.example <resource ID>
```