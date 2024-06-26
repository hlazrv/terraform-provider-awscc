---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_gamelift_script Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GameLift::Script
---

# awscc_gamelift_script (Data Source)

Data Source schema for AWS::GameLift::Script



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift script resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift script ARN, the resource ID matches the Id value.
- `creation_time` (String) A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057").
- `name` (String) A descriptive label that is associated with a script. Script names do not need to be unique.
- `script_id` (String) A unique identifier for the Realtime script
- `size_on_disk` (Number) The file size of the uploaded Realtime script, expressed in bytes. When files are uploaded from an S3 location, this value remains at "0".
- `storage_location` (Attributes) The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored. The storage location must specify the Amazon S3 bucket name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon GameLift uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the ObjectVersion parameter to specify an earlier version. (see [below for nested schema](#nestedatt--storage_location))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `version` (String) The version that is associated with a script. Version strings do not need to be unique.

<a id="nestedatt--storage_location"></a>
### Nested Schema for `storage_location`

Read-Only:

- `bucket` (String) An Amazon S3 bucket identifier. This is the name of the S3 bucket.
- `key` (String) The name of the zip file that contains the script files.
- `object_version` (String) The version of the file, if object versioning is turned on for the bucket. Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.
- `role_arn` (String) The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.
