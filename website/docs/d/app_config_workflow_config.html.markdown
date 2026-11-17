---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration workflow config'
description: |-
  Get information about workflow config.
---

# ibm_app_config_workflow_config

Retrieve information about an existing IBM Cloud App Configuration workflow config. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration workflow configuration, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_workflow_config" "app_config_workflow_config" {
  guid = "guid"
  workflow_id = "workflow_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `workflow_id` - (Required, String) Unique identifier for the workflow configuration.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the workflow config resource.
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
