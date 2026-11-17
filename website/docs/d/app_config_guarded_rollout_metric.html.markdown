---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration guarded rollout metric'
description: |-
  Get the statistics of a metric of a guarded rollout.
---

# ibm_app_config_guarded_rollout_metric

Retrieve information about an existing IBM Cloud App Configuration guarded rollout metric. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration guarded rollouts, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_guarded_rollout_metric" "app_config_guarded_rollout_metric" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  rollout_id = "rollout_id"
  metric_id = "metric_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature Id.
- `metric_id` - (Required, String) The unique identifier of the metric.
- `rollout_id` - (Required, String) The unique identifier of the guarded rollout.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the guarded rollout metric.
- `result` - (List) Time-series statistical results for this metric.

  Nested scheme for `result`:
  - `difference` - (List) Statistics of the difference between the target and the original.

    Nested scheme for `difference`:
    - `conversions` - (Integer) The number of conversions recorded. Set on original and target only.
    - `estimate` - (Float) The point estimate.
    - `exposures` - (Integer) The number of exposures recorded. Set on original and target only.
    - `lower` - (Float) The lower bound of the confidence interval.
    - `upper` - (Float) The upper bound of the confidence interval.
  - `original` - (List) Statistics of the original variant.

    Nested scheme for `original`:
    - `conversions` - (Integer) The number of conversions recorded. Set on original and target only.
    - `estimate` - (Float) The point estimate.
    - `exposures` - (Integer) The number of exposures recorded. Set on original and target only.
    - `lower` - (Float) The lower bound of the confidence interval.
    - `upper` - (Float) The upper bound of the confidence interval.
  - `target` - (List) Statistics of the target variant.

    Nested scheme for `target`:
    - `conversions` - (Integer) The number of conversions recorded. Set on original and target only.
    - `estimate` - (Float) The point estimate.
    - `exposures` - (Integer) The number of exposures recorded. Set on original and target only.
    - `lower` - (Float) The lower bound of the confidence interval.
    - `upper` - (Float) The upper bound of the confidence interval.
  - `timestamp` - (String) The time at which this result was computed.
- `status` - (String) The status of the guarded rollout metric.
- `sub_status` - (String) The state behind the status.
