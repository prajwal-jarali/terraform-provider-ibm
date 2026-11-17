---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  List all the environments.
---

# ibm_app_config_environments

Retrieve information about an existing IBM Cloud App Configuration environments. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration environments, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-service-access-level-management).

## Example usage

```terraform
data "ibm_app_config_environments" "app_config_environments" {
  guid = "guid"
  tags = "tag1,tag2"
  sort = "name"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `expand` - (Optional, Boolean) If set to `true`, returns expanded view of the resource details.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are features, properties, snapshots, workflow_approval.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `search` - (Optional, String) Searches for the provided keyword and returns the appropriate row with that value. Here the search happens on the '[Name OR Tag]' of the entity.
- `sort` - (Optional, String) Sort the environment details based on the specified attribute. By default, items are sorted by name. Supported values are created_time, id, name, updated_time.
- `tags` - (Optional, String) Filter the resources to be returned based on the associated tags. Specify the parameter as a list of comma separated tags. Returns resources associated with any of the specified tags.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the environment list.
- `environments` - (List) Array of environments.

  Nested scheme for `environments`:
  - `color_code` - (String) Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`.
  - `created_time` - (Timestamp) Creation time of the environment.
  - `description` - (String) The environment description.
  - `environment_id` - (String) Environment Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `features` - (List) List of features associated with this environment.

    Nested scheme for `features`:
    - `feature_id` - (String) feature Id.
    - `name` - (String) feature name.
  - `href` - (String) environment URL.
  - `name` - (String) The environment name.
  - `properties` - (List) List of propertys associated with this environment.

    Nested scheme for `properties`:
    - `name` - (String) property name.
    - `property_id` - (String) property Id.
  - `snapshots` - (List) List of git configs associated with this environment.

    Nested scheme for `snapshots`:
    - `git_config_id` - (String) git config Id.
    - `name` - (String) git config name.
  - `tags` - (String) Tags associated with the environment. Allowed special characters are [_. ,-:].
  - `updated_time` - (Timestamp) Last modified time of the environment.
  - `workflow_approval` - (List) Workflow approval information. Present only when the workflow is enabled and the resource has an active change request.

    Nested scheme for `workflow_approval`:
    - `approval_url` - (String) URL to view the change request in ServiceNow.
    - `change_request_number` - (String) ServiceNow change request number.
    - `change_request_status` - (String) Status of the change request in ServiceNow.
    - `created_time` - (Timestamp) Timestamp when the workflow approval was created.
    - `environment_id` - (String) Environment identifier. Present only for features, properties and environments.
    - `execution_status` - (String) Execution status of the workflow.
    - `last_poll_time` - (Timestamp) Timestamp when the workflow approval status was last polled from ServiceNow.
    - `provider_type` - (String) Type of workflow provider.
    - `resource_id` - (String) Unique identifier of the resource. Not present for environments.
    - `resource_type` - (String) Type of resource.
    - `updated_time` - (Timestamp) Timestamp when the workflow approval was last updated.
    - `workflow_id` - (String) Unique identifier of the workflow configuration.
    - `workflow_name` - (String) Name of the workflow configuration.
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
