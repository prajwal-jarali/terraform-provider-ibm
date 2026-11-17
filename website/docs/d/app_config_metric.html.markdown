---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration metric'
description: |-
  Get information about metric.
---

# ibm_app_config_metric

Retrieve information about an existing IBM Cloud App Configuration metric. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration metrics, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_metric" "app_config_metric" {
  guid = "guid"
  metric_id = "metric_id"
  include = ["experiments", "guarded_rollouts"]
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `metric_id` - (Required, String) Metric ID. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are experiments, guarded_rollouts.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the metric resource.
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
- `metric_name` - (String) Metric name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `metric_type` - (String) The type of metric. Supported values are BINARY, COUNT, NUMERIC.
- `updated_time` - (Timestamp) Last modified time of the metric.
