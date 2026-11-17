---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration guarded rollout'
description: |-
  Get information about guarded rollout.
---

# ibm_app_config_guarded_rollout

Retrieve information about an existing IBM Cloud App Configuration guarded rollout. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration guarded rollouts, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_guarded_rollout" "app_config_guarded_rollout" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  rollout_id = "rollout_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature Id.
- `rollout_id` - (Required, String) The unique identifier of the guarded rollout.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the guarded rollout.
- `current_phase_index` - (Integer) The index of the current or last active phase.
- `description` - (String) A human-readable description of the rollout outcome.
- `end_time` - (Timestamp) The date and time the guarded rollout completed. Not set while it is in progress.
- `feature_snapshot` - (List) Snapshot of the feature flag values at the time of the guarded rollout.

  Nested scheme for `feature_snapshot`:
  - `collections` - (List) List of collections associated with the feature.

    Nested scheme for `collections`:
    - `collection_id` - (String) Collection Id.
    - `name` - (String) Collection name.
  - `disabled_value` - (String) The value served when the feature flag is disabled.
  - `enabled_value` - (String) The value served when the feature flag is enabled.
- `metrics` - (List) Metrics monitored during the guarded rollout.

  Nested scheme for `metrics`:
  - `auto_rollback` - (Boolean) Whether a regression in this metric triggers an automatic rollback.
  - `description` - (String) A description of what the metric tracks.
  - `direction` - (String) Whether a higher or lower metric value indicates improvement.
  - `event_key` - (String) The event key used to record metric data.
  - `last_event_received_time` - (Timestamp) The time the most recent event was received for this metric.
  - `last_regression_detected_time` - (Timestamp) The time a regression was last detected for this metric.
  - `last_regression_dismissed_time` - (Timestamp) The time a regression was last dismissed for this metric.
  - `metric_id` - (String) The unique identifier of the metric.
  - `metric_name` - (String) The display name of the metric.
  - `metric_type` - (String) The type of metric.
  - `status` - (String) The status of the metric: HEALTHY, INCONCLUSIVE_HEALTHY, INCONCLUSIVE_WORSE, INSUFFICIENT_DATA, NO_EVENTS or REGRESSED.
  - `sub_status` - (String) The state behind the status: PAUSE_AUTO, PAUSE_CONTINUE, PAUSE_FINAL_STAGE, PAUSE_SAMPLE_COLLECTION, PHASE_EXTEND or ROLLBACK.
- `phases` - (List) Ordered list of rollout phases.

  Nested scheme for `phases`:
  - `duration` - (Integer) The duration this phase runs before advancing. Not set for the final phase.
  - `duration_type` - (String) The unit of the duration value: days, hours or minutes.
  - `extended` - (Boolean) Whether this phase has been extended beyond its original duration.
  - `percentage` - (Integer) The rollout percentage for this phase.
  - `phase_index` - (Integer) Zero-based index of the phase.
- `rule` - (List) The targeting rule that is rolled out.

  Nested scheme for `rule`:
  - `order` - (Integer) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. Index 1 based numbering is used for order.
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
- `rule_id` - (String) The rule associated with this guarded rollout. Not set if the rollout applies to the feature.
- `start_time` - (Timestamp) The date and time the guarded rollout was started.
- `status` - (String) The current status of the guarded rollout: COMPLETED, IN_PROGRESS, PAUSED, ROLLED_BACK or STOPPED.
