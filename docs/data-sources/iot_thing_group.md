---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_thing_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoT::ThingGroup
---

# awscc_iot_thing_group (Data Source)

Data Source schema for AWS::IoT::ThingGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `parent_group_name` (String)
- `query_string` (String)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `thing_group_id` (String)
- `thing_group_name` (String)
- `thing_group_properties` (Attributes) (see [below for nested schema](#nestedatt--thing_group_properties))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--thing_group_properties"></a>
### Nested Schema for `thing_group_properties`

Read-Only:

- `attribute_payload` (Attributes) (see [below for nested schema](#nestedatt--thing_group_properties--attribute_payload))
- `thing_group_description` (String)

<a id="nestedatt--thing_group_properties--attribute_payload"></a>
### Nested Schema for `thing_group_properties.attribute_payload`

Read-Only:

- `attributes` (Map of String)
