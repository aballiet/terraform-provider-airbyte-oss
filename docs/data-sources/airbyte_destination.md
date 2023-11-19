---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination Data Source - aballiet-terraform-provider-airbyte-oss"
subcategory: ""
description: |-
  Destination DataSource
---

# airbyte_destination (Data Source)

Destination DataSource

## Example Usage

```terraform
data "airbyte_destination" "my_destination" {
  destination_id = "cbc07f65-e5d1-4847-8196-e36f66cd38de"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `connection_configuration` (String) Parsed as JSON.
The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
- `destination_definition_id` (String)
- `destination_name` (String)
- `icon` (String)
- `name` (String)
- `workspace_id` (String)

