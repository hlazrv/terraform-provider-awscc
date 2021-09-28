---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_gamelift_alias Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GameLift::Alias
---

# awscc_gamelift_alias (Data Source)

Data Source schema for AWS::GameLift::Alias



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **alias_id** (String) Unique alias ID
- **description** (String) A human-readable description of the alias.
- **name** (String) A descriptive label that is associated with an alias. Alias names do not need to be unique.
- **routing_strategy** (Attributes) (see [below for nested schema](#nestedatt--routing_strategy))

<a id="nestedatt--routing_strategy"></a>
### Nested Schema for `routing_strategy`

Read-Only:

- **fleet_id** (String) A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.
- **message** (String) The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.
- **type** (String) Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.

