---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration metric'
description: |-
  Manages metrics.
---

# ibm_app_config_metric

Create, update, or delete a metric by using IBM Cloud™ App Configuration. For more information, about App Configuration metrics, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

~> **Note:** The `is_favorite` argument is set with its own API call after the metric is created, and when its value changes.

## Example usage

```terraform
resource "ibm_app_config_metric" "app_config_metric" {
  guid = "guid"
  metric_id = "metric_id"
  metric_name = "metric_name"
  description = "description"
  event_type = "CUSTOM"
  metric_type = "COUNT"
  event_key = "event_key"
  direction = "HIGHER"
  is_favorite = true
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `metric_id` - (Required, String) Metric ID. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `direction` - (Required, String) Whether a higher or lower metric value indicates improvement. Supported values are HIGHER, LOWER.
- `event_key` - (Required, String) The event key used to record metric data.
- `event_type` - (Required, String) The type of event. Supported values are CLICK, CUSTOM, PAGE_VIEW.
- `metric_name` - (Required, String) Metric name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `metric_type` - (Required, String) The type of metric. Supported values are BINARY, COUNT, NUMERIC.
- `description` - (Optional, String) Metric description.
- `is_favorite` - (Optional, Boolean) Whether the metric is marked as favorite.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the metric resource. The format is `<guid>/<metric_id>`.
- `created_time` - (Timestamp) Creation time of the metric.
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
- `updated_time` - (Timestamp) Last modified time of the metric.

## Import

The `ibm_app_config_metric` resource can be imported by using `guid` of the App Configuration instance and `metricId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_metric.sample  <guid/metricId>

```

**Example**

```
terraform import ibm_app_config_metric.sample 272111153-c118-4116-8116-b811fbc31132/sample_metric
```
