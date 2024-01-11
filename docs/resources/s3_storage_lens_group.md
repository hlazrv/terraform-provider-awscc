---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3_storage_lens_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::S3::StorageLensGroup resource is an Amazon S3 resource type that you can use to create Storage Lens Group.
---

# awscc_s3_storage_lens_group (Resource)

The AWS::S3::StorageLensGroup resource is an Amazon S3 resource type that you can use to create Storage Lens Group.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter` (Attributes) Sets the Storage Lens Group filter. (see [below for nested schema](#nestedatt--filter))
- `name` (String) The name that identifies the Amazon S3 Storage Lens Group.

### Optional

- `tags` (Attributes List) A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `storage_lens_group_arn` (String) The ARN for the Amazon S3 Storage Lens Group.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Optional:

- `and` (Attributes) The Storage Lens group will include objects that match all of the specified filter values. (see [below for nested schema](#nestedatt--filter--and))
- `match_any_prefix` (Set of String) Filter to match any of the specified prefixes.
- `match_any_suffix` (Set of String) Filter to match any of the specified suffixes.
- `match_any_tag` (Attributes Set) Filter to match any of the specified object tags. (see [below for nested schema](#nestedatt--filter--match_any_tag))
- `match_object_age` (Attributes) Filter to match all of the specified values for the minimum and maximum object age. (see [below for nested schema](#nestedatt--filter--match_object_age))
- `match_object_size` (Attributes) Filter to match all of the specified values for the minimum and maximum object size. (see [below for nested schema](#nestedatt--filter--match_object_size))
- `or` (Attributes) The Storage Lens group will include objects that match any of the specified filter values. (see [below for nested schema](#nestedatt--filter--or))

<a id="nestedatt--filter--and"></a>
### Nested Schema for `filter.and`

Optional:

- `match_any_prefix` (Set of String) Filter to match any of the specified prefixes.
- `match_any_suffix` (Set of String) Filter to match any of the specified suffixes.
- `match_any_tag` (Attributes Set) Filter to match any of the specified object tags. (see [below for nested schema](#nestedatt--filter--and--match_any_tag))
- `match_object_age` (Attributes) Filter to match all of the specified values for the minimum and maximum object age. (see [below for nested schema](#nestedatt--filter--and--match_object_age))
- `match_object_size` (Attributes) Filter to match all of the specified values for the minimum and maximum object size. (see [below for nested schema](#nestedatt--filter--and--match_object_size))

<a id="nestedatt--filter--and--match_any_tag"></a>
### Nested Schema for `filter.and.match_any_tag`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--filter--and--match_object_age"></a>
### Nested Schema for `filter.and.match_object_age`

Optional:

- `days_greater_than` (Number) Minimum object age to which the rule applies.
- `days_less_than` (Number) Maximum object age to which the rule applies.


<a id="nestedatt--filter--and--match_object_size"></a>
### Nested Schema for `filter.and.match_object_size`

Optional:

- `bytes_greater_than` (Number) Minimum object size to which the rule applies.
- `bytes_less_than` (Number) Maximum object size to which the rule applies.



<a id="nestedatt--filter--match_any_tag"></a>
### Nested Schema for `filter.match_any_tag`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--filter--match_object_age"></a>
### Nested Schema for `filter.match_object_age`

Optional:

- `days_greater_than` (Number) Minimum object age to which the rule applies.
- `days_less_than` (Number) Maximum object age to which the rule applies.


<a id="nestedatt--filter--match_object_size"></a>
### Nested Schema for `filter.match_object_size`

Optional:

- `bytes_greater_than` (Number) Minimum object size to which the rule applies.
- `bytes_less_than` (Number) Maximum object size to which the rule applies.


<a id="nestedatt--filter--or"></a>
### Nested Schema for `filter.or`

Optional:

- `match_any_prefix` (Set of String) Filter to match any of the specified prefixes.
- `match_any_suffix` (Set of String) Filter to match any of the specified suffixes.
- `match_any_tag` (Attributes Set) Filter to match any of the specified object tags. (see [below for nested schema](#nestedatt--filter--or--match_any_tag))
- `match_object_age` (Attributes) Filter to match all of the specified values for the minimum and maximum object age. (see [below for nested schema](#nestedatt--filter--or--match_object_age))
- `match_object_size` (Attributes) Filter to match all of the specified values for the minimum and maximum object size. (see [below for nested schema](#nestedatt--filter--or--match_object_size))

<a id="nestedatt--filter--or--match_any_tag"></a>
### Nested Schema for `filter.or.match_any_tag`

Required:

- `key` (String)
- `value` (String)


<a id="nestedatt--filter--or--match_object_age"></a>
### Nested Schema for `filter.or.match_object_age`

Optional:

- `days_greater_than` (Number) Minimum object age to which the rule applies.
- `days_less_than` (Number) Maximum object age to which the rule applies.


<a id="nestedatt--filter--or--match_object_size"></a>
### Nested Schema for `filter.or.match_object_size`

Optional:

- `bytes_greater_than` (Number) Minimum object size to which the rule applies.
- `bytes_less_than` (Number) Maximum object size to which the rule applies.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_s3_storage_lens_group.example <resource ID>
```