---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration guarded rollouts'
description: |-
  List all the guarded rollouts of a feature flag.
---

# ibm_app_config_guarded_rollouts

Retrieve information about an existing IBM Cloud App Configuration guarded rollouts. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration guarded rollouts, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_guarded_rollouts" "app_config_guarded_rollouts" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature Id.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the guarded rollout list.
- `first` - (List) URL to navigate to the first list of records.

  Nested scheme for `first`:
  - `href` - (String) URL of the response.
- `guarded_rollouts` - (List) Array of guarded rollouts.

  Nested scheme for `guarded_rollouts`:
  - `end_time` - (Timestamp) The date and time the guarded rollout completed. Not set while it is in progress.
  - `rollout_id` - (String) The unique identifier of the guarded rollout.
  - `rule_id` - (String) The rule associated with this guarded rollout. Not set if the rollout applies to the feature.
  - `start_time` - (Timestamp) The date and time the guarded rollout was started.
  - `status` - (String) The current status of the guarded rollout: COMPLETED, IN_PROGRESS, PAUSED, ROLLED_BACK or STOPPED.
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
