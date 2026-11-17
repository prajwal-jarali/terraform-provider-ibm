---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration metrics'
description: |-
  List all the metrics.
---

# ibm_app_config_metrics

Retrieve information about an existing IBM Cloud App Configuration metrics. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration metrics, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_metrics" "app_config_metrics" {
  guid = "guid"
  sort = "metric_name"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `expand` - (Optional, Boolean) If set to `true`, returns expanded view of the resource details.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are experiments, guarded_rollouts.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `search` - (Optional, String) Searches for the provided keyword and returns the appropriate row with that value. Here the search happens on the '[Name]' of the entity.
- `sort` - (Optional, String) Sort the metric details based on the specified attribute. By default, items are sorted by name. Supported values are created_time, metric_id, metric_name, updated_time.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the metric list.
- `first` - (List) URL to navigate to the first list of records.

  Nested scheme for `first`:
  - `href` - (String) URL of the response.
- `last` - (List) URL to navigate to the last list of records.

  Nested scheme for `last`:
  - `href` - (String) URL of the response.
- `metrics` - (List) Array of metrics.

  Nested scheme for `metrics`:
  - `created_time` - (Timestamp) Creation time of the metric.
  - `description` - (String) Metric description.
  - `direction` - (String) Whether a higher or lower metric value indicates improvement. Supported values are HIGHER, LOWER.
  - `event_key` - (String) The event key used to record metric data.
  - `event_type` - (String) The type of event. Supported values are CLICK, CUSTOM, PAGE_VIEW.
  - `experiments` - (List) List of experiments that use this metric. Set only when `include` has `experiments`.

    Nested scheme for `experiments`:
    - `experiment_id` - (String) The unique identifier of the experiment.
    - `experiment_status` - (String) The status of the experiment.
    - `name` - (String) The name of the experiment.
  - `experiments_count` - (Integer) Total count of experiments that use this metric.
  - `guarded_rollouts` - (List) List of guarded rollouts that use this metric. Set only when `include` has `guarded_rollouts`.

    Nested scheme for `guarded_rollouts`:
    - `environment_id` - (String) The unique identifier of the environment.
    - `environment_name` - (String) The name of the environment.
    - `feature_id` - (String) The unique identifier of the feature flag.
    - `feature_name` - (String) The name of the feature flag.
    - `rollout_id` - (String) The unique identifier of the guarded rollout.
    - `rollout_status` - (String) The current status of the guarded rollout.
  - `guarded_rollouts_count` - (Integer) Total count of guarded rollouts that use this metric.
  - `href` - (String) metric URL.
  - `is_favorite` - (Boolean) Whether the metric is marked as favorite.
  - `last_event_received_time` - (Timestamp) The time the most recent event was received for this metric. Set only when an event was recorded.
  - `metric_id` - (String) Metric ID. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `metric_name` - (String) Metric name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `metric_type` - (String) The type of metric. Supported values are BINARY, COUNT, NUMERIC.
  - `updated_time` - (Timestamp) Last modified time of the metric.
- `next` - (List) URL to navigate to the next list of records.

  Nested scheme for `next`:
  - `href` - (String) URL of the response.
- `previous` - (List) URL to navigate to the previous list of records.

  Nested scheme for `previous`:
  - `href` - (String) URL of the response.
- `total_count` - (Integer) Total number of records.
