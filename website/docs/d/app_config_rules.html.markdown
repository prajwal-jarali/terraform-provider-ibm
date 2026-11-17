---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag rules'
description: |-
  List all the targeting rules of a feature flag.
---

# ibm_app_config_rules

Retrieve information about an existing IBM Cloud App Configuration feature flag rules. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration feature flag rules, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_rules" "app_config_rules" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature Id.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the feature flag rule list.
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
- `segment_rules` - (List) Array of rules of the feature.

  Nested scheme for `segment_rules`:
  - `href` - (String) Feature flag rule URL.
  - `order` - (Integer) Order of the rule, used during evaluation. The value of the first matching rule is used for evaluation.
  - `rollout_configuration` - (List) Configuration that controls the rollout behaviour for a `PROGRESSIVE` or `GUARDED` rollout type.

    Nested scheme for `rollout_configuration`:
    - `duration_preset` - (String) A predefined duration preset that sets the overall pace of the rollout. Use `CUSTOM` to define phases manually with explicit `duration` and `duration_type` values on each phase.
    - `metrics` - (List) The metrics monitored during the rollout. Applicable only when `rollout_type` is `GUARDED`. Each `metric_id` must be unique.

      Nested scheme for `metrics`:
      - `auto_rollback` - (Boolean) Whether a regression in this metric triggers an automatic rollback.
      - `direction` - (String) Whether a higher or lower metric value indicates improvement.
      - `metric_id` - (String) The unique identifier of the metric.
      - `metric_name` - (String) The display name of the metric.
      - `metric_type` - (String) The type of metric.
    - `phases` - (List) The ordered list of rollout phases. Each phase defines a target percentage and a wait duration before the next phase begins.

      Nested scheme for `phases`:
      - `duration` - (Integer) The length of time to wait before advancing to the next phase.
      - `duration_type` - (String) The unit of time for the `duration` field. Supported values are days, hours, minutes.
      - `percentage` - (Integer) The rollout percentage target for this phase, from 0 to 100.
    - `start_at` - (String) The UTC date and time at which the rollout should start, in ISO 8601 format.
    - `status` - (String) The current execution status of the rollout: `QUEUED`, `RUNNING` or `STOPPED`.
  - `rollout_percentage` - (Integer) Rollout percentage. Supported only for Lite and Enterprise plans.
  - `rollout_type` - (String) The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics. Supported values are MANUAL, PROGRESSIVE, GUARDED.
  - `rule_id` - (String) Rule id is an identifier used for identifying the rule.
  - `rule_name` - (String) Name assigned to a particular rule.
  - `rules` - (List) The list of targeted segments.

    Nested scheme for `rules`:
    - `segments` - (Array of Strings) The list of segment ids that are used for targeting using the rule.
  - `value` - (String) Value to be used for evaluation for this rule. The value must match the `type` and `format` attributes.
- `total_count` - (Integer) Total number of records.
