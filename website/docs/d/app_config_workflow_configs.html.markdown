---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration workflow configs'
description: |-
  List all the workflow configs.
---

# ibm_app_config_workflow_configs

Retrieve information about an existing IBM Cloud App Configuration workflow configs. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration workflow configuration, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_workflow_configs" "app_config_workflow_configs" {
  guid = "guid"
  sort = "name"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `search` - (Optional, String) Searches for the provided keyword and returns the appropriate row with that value. Here the search happens on the '[Name]' of the entity.
- `sort` - (Optional, String) Sort the workflow config details based on the specified attribute. By default, items are sorted by name. Supported values are created_time, name, updated_time, workflow_id.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the workflow config list.
- `first` - (List) URL to navigate to the first list of records.

  Nested scheme for `first`:
  - `href` - (String) URL of the response.
- `last` - (List) URL to navigate to the last list of records.

  Nested scheme for `last`:
  - `href` - (String) URL of the response.
- `next` - (List) URL to navigate to the next list of records.

  Nested scheme for `next`:
  - `href` - (String) URL of the response.
- `previous` - (List) URL to navigate to the previous list of records.

  Nested scheme for `previous`:
  - `href` - (String) URL of the response.
- `total_count` - (Integer) Total number of records.
- `workflow_configs` - (List) Array of workflow configs.

  Nested scheme for `workflow_configs`:
  - `created_time` - (Timestamp) Creation time of the workflow config.
  - `enabled` - (Boolean) Whether the workflow is enabled.
  - `href` - (String) workflow config URL.
  - `name` - (String) Name of the workflow configuration.
  - `provider_config` - (List) Provider configuration for the workflow.

    Nested scheme for `provider_config`:
    - `metadata` - (List) Metadata of the ServiceNow provider. Which settings apply depends on the provider type.

      Nested scheme for `metadata`:
      - `approval_expiration` - (Integer) Approval expiration time in days.
      - `approval_group_name` - (String) Name of the approval group in ServiceNow.
      - `crn_mask` - (String) CRN mask for the service.
      - `default_email` - (String) The email id to which the change request is assigned.
      - `service_environment` - (String) The ServiceNow change request environment context. If not provided, it defaults to production. Supported values are pre_prod, production.
      - `sm_secret_crn` - (String) Secret Manager secret CRN for the ServiceNow credentials.
      - `sm_secret_id` - (String) Secret Manager secret ID for the ServiceNow credentials.
      - `workflow_url` - (String) ServiceNow instance URL.
    - `type` - (String) Provider type. Supported values are SERVICENOW_IBM, SERVICENOW_EXTERNAL.
  - `scope` - (List) Scope of the workflow. At least one of collections, segments or environments must be specified.

    Nested scheme for `scope`:
    - `collections` - (List) Scope of the collections. Only the ALL mode is supported.

      Nested scheme for `collections`:
      - `ids` - (Array of Strings) List of ids. Always empty when the mode is ALL.
      - `mode` - (String) Scope mode. Supported values are ALL.
    - `environments` - (List) List of environments where the workflow applies.

      Nested scheme for `environments`:
      - `environment_id` - (String) Environment identifier.
      - `resources` - (List) Resources of the environment that the workflow applies to.

        Nested scheme for `resources`:
        - `environment` - (List) The environment itself.

          Nested scheme for `environment`:
          - `enable` - (Boolean) Whether changes to the environment need approval.
        - `features` - (List) Scope of the features. Only the ALL mode is supported.

          Nested scheme for `features`:
          - `ids` - (Array of Strings) List of ids. Always empty when the mode is ALL.
          - `mode` - (String) Scope mode. Supported values are ALL.
        - `properties` - (List) Scope of the properties. Only the ALL mode is supported.

          Nested scheme for `properties`:
          - `ids` - (Array of Strings) List of ids. Always empty when the mode is ALL.
          - `mode` - (String) Scope mode. Supported values are ALL.
    - `segments` - (List) Scope of the segments. Only the ALL mode is supported.

      Nested scheme for `segments`:
      - `ids` - (Array of Strings) List of ids. Always empty when the mode is ALL.
      - `mode` - (String) Scope mode. Supported values are ALL.
  - `updated_time` - (Timestamp) Last modified time of the workflow config.
  - `workflow_id` - (String) Unique identifier for the workflow configuration.
