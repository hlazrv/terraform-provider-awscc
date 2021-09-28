---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_chatbot_slack_channel_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Chatbot::SlackChannelConfiguration.
---

# awscc_chatbot_slack_channel_configuration (Resource)

Resource schema for AWS::Chatbot::SlackChannelConfiguration.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **configuration_name** (String) The name of the configuration
- **iam_role_arn** (String) The ARN of the IAM role that defines the permissions for AWS Chatbot
- **slack_channel_id** (String) The id of the Slack channel
- **slack_workspace_id** (String) The id of the Slack workspace

### Optional

- **logging_level** (String) Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
- **sns_topic_arns** (List of String) ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.

### Read-Only

- **arn** (String) Amazon Resource Name (ARN) of the configuration
- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_chatbot_slack_channel_configuration.example <resource ID>
```