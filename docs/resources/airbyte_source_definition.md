---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_definition Resource - aballiet-terraform-provider-airbyte-oss"
subcategory: ""
description: |-
  SourceDefinition Resource
---

# airbyte_source_definition (Resource)

SourceDefinition Resource

## Example Usage

```terraform
resource "airbyte_source_definition" "my_sourcedefinition" {
  scope_id   = "3d9d1434-d468-48b6-8221-d5f1402c7e26"
  scope_type = "workspace"
  source_definition = {
    name              = "Felicia Huels"
    docker_repository = "...my_docker_repository..."
    docker_image_tag  = "...my_docker_image_tag..."
    documentation_url = "http://unnatural-falling-out.biz"
    icon              = "...my_icon..."
    resource_requirements = {
      default = {
        cpu_request    = "...my_cpu_request..."
        cpu_limit      = "...my_cpu_limit..."
        memory_request = "...my_memory_request..."
        memory_limit   = "...my_memory_limit..."
      }
      job_specific = [
        {
          job_type = "replicate"
          resource_requirements = {
            cpu_request    = "...my_cpu_request..."
            cpu_limit      = "...my_cpu_limit..."
            memory_request = "...my_memory_request..."
            memory_limit   = "...my_memory_limit..."
          }
        },
      ]
    }
  }
  workspace_id = "e038120e-c22a-4c97-8dbe-310dc62194f1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_definition` (Attributes) (see [below for nested schema](#nestedatt--source_definition))

### Optional

- `scope_id` (String)
- `scope_type` (String) must be one of ["workspace", "organization"]
- `workspace_id` (String)

### Read-Only

- `custom` (Boolean) Default: false
Whether the connector is custom or not
- `docker_image_tag` (String)
- `docker_repository` (String)
- `documentation_url` (String)
- `icon` (String)
- `max_seconds_between_messages` (Number) Number of seconds allowed between 2 airbyte protocol messages. The source will timeout if this delay is reach
- `name` (String)
- `protocol_version` (String) The Airbyte Protocol version supported by the connector
- `release_date` (String) The date when this connector was first released, in yyyy-mm-dd format.
- `release_stage` (String) must be one of ["alpha", "beta", "generally_available", "custom"]
- `resource_requirements` (Attributes) actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level. (see [below for nested schema](#nestedatt--resource_requirements))
- `source_definition_id` (String)
- `source_type` (String) must be one of ["api", "file", "database", "custom"]
- `support_level` (String) must be one of ["community", "certified", "none"]

<a id="nestedatt--source_definition"></a>
### Nested Schema for `source_definition`

Required:

- `docker_image_tag` (String)
- `docker_repository` (String)
- `documentation_url` (String)
- `name` (String)

Optional:

- `icon` (String)
- `resource_requirements` (Attributes) actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level. (see [below for nested schema](#nestedatt--source_definition--resource_requirements))

<a id="nestedatt--source_definition--resource_requirements"></a>
### Nested Schema for `source_definition.resource_requirements`

Optional:

- `default` (Attributes) optional resource requirements to run workers (blank for unbounded allocations) (see [below for nested schema](#nestedatt--source_definition--resource_requirements--default))
- `job_specific` (Attributes List) (see [below for nested schema](#nestedatt--source_definition--resource_requirements--job_specific))

<a id="nestedatt--source_definition--resource_requirements--default"></a>
### Nested Schema for `source_definition.resource_requirements.default`

Optional:

- `cpu_limit` (String)
- `cpu_request` (String)
- `memory_limit` (String)
- `memory_request` (String)


<a id="nestedatt--source_definition--resource_requirements--job_specific"></a>
### Nested Schema for `source_definition.resource_requirements.job_specific`

Required:

- `job_type` (String) must be one of ["get_spec", "check_connection", "discover_schema", "sync", "reset_connection", "connection_updater", "replicate"]
enum that describes the different types of jobs that the platform runs.
- `resource_requirements` (Attributes) optional resource requirements to run workers (blank for unbounded allocations) (see [below for nested schema](#nestedatt--source_definition--resource_requirements--job_specific--resource_requirements))

<a id="nestedatt--source_definition--resource_requirements--job_specific--resource_requirements"></a>
### Nested Schema for `source_definition.resource_requirements.job_specific.resource_requirements`

Optional:

- `cpu_limit` (String)
- `cpu_request` (String)
- `memory_limit` (String)
- `memory_request` (String)





<a id="nestedatt--resource_requirements"></a>
### Nested Schema for `resource_requirements`

Read-Only:

- `default` (Attributes) optional resource requirements to run workers (blank for unbounded allocations) (see [below for nested schema](#nestedatt--resource_requirements--default))
- `job_specific` (Attributes List) (see [below for nested schema](#nestedatt--resource_requirements--job_specific))

<a id="nestedatt--resource_requirements--default"></a>
### Nested Schema for `resource_requirements.default`

Read-Only:

- `cpu_limit` (String)
- `cpu_request` (String)
- `memory_limit` (String)
- `memory_request` (String)


<a id="nestedatt--resource_requirements--job_specific"></a>
### Nested Schema for `resource_requirements.job_specific`

Read-Only:

- `job_type` (String) must be one of ["get_spec", "check_connection", "discover_schema", "sync", "reset_connection", "connection_updater", "replicate"]
enum that describes the different types of jobs that the platform runs.
- `resource_requirements` (Attributes) optional resource requirements to run workers (blank for unbounded allocations) (see [below for nested schema](#nestedatt--resource_requirements--job_specific--resource_requirements))

<a id="nestedatt--resource_requirements--job_specific--resource_requirements"></a>
### Nested Schema for `resource_requirements.job_specific.resource_requirements`

Read-Only:

- `cpu_limit` (String)
- `cpu_request` (String)
- `memory_limit` (String)
- `memory_request` (String)


