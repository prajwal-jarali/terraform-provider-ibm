---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag rule'
description: |-
  Manages the targeting rule of a feature flag.
---

# ibm_app_config_rule

Create, update, or delete a targeting rule of a feature flag by using IBM Cloud™ App Configuration. For more information, about App Configuration feature flag rules, see [Feature flags](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-feature-flags).

~> **Note:** The `order` of a rule is set by the API and cannot be changed with this resource.

## Example usage

```terraform
resource "ibm_app_config_rule" "app_config_rule" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  rule_id = "rule_id"
  rule_name = "rule_name"
  rules {
    segments = ["segment_id1", "segment_id2"]
  }
  value = "value"
  rollout_percentage = 80
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature Id.
- `rule_id` - (Required, String) Rule id is an identifier used for identifying the rule.
- `rules` - (Required, List) The list of targeted segments.
  - `segments` - (Required, Array of Strings) The list of segment ids that are used for targeting using the rule.
- `value` - (Required, String) Value to be used for evaluation for this rule. The value must match the `type` and `format` attributes.
- `rollout_configuration` - (Optional, List) Configuration that controls the rollout behaviour for a `PROGRESSIVE` or `GUARDED` rollout type.
  - `duration_preset` - (Required, String) A predefined duration preset that sets the overall pace of the rollout. Use `CUSTOM` to define phases manually with explicit `duration` and `duration_type` values on each phase.
  - `phases` - (Required, List) The ordered list of rollout phases. Each phase defines a target percentage and a wait duration before the next phase begins.
    - `percentage` - (Required, Integer) The rollout percentage target for this phase, from 0 to 100.
    - `duration` - (Optional, Integer) The length of time to wait before advancing to the next phase.
    - `duration_type` - (Optional, String) The unit of time for the `duration` field. Supported values are days, hours, minutes.
  - `metrics` - (Optional, List) The metrics monitored during the rollout. Applicable only when `rollout_type` is `GUARDED`. Each `metric_id` must be unique.
    - `metric_id` - (Required, String) The unique identifier of the metric.
    - `auto_rollback` - (Optional, Boolean) Whether a regression in this metric triggers an automatic rollback.
  - `start_at` - (Optional, String) The UTC date and time at which the rollout should start, in ISO 8601 format.
- `rollout_percentage` - (Optional, Integer) Rollout percentage. Supported only for Lite and Enterprise plans.
- `rollout_type` - (Optional, String) The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics. Supported values are MANUAL, PROGRESSIVE, GUARDED.
- `rule_name` - (Optional, String) Name assigned to a particular rule.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the feature flag rule resource. The format is `<guid>/<environment_id>/<feature_id>/<rule_id>`.
- `href` - (String) Feature flag rule URL.
- `order` - (Integer) Order of the rule, used during evaluation. The value of the first matching rule is used for evaluation.
- `rollout_percentage` - (Integer) Rollout percentage. Supported only for Lite and Enterprise plans.
- `rollout_type` - (String) The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics. Supported values are MANUAL, PROGRESSIVE, GUARDED.

## Import

The `ibm_app_config_rule` resource can be imported by using `guid` of the App Configuration instance, `environmentId`, `featureId`, `ruleId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_rule.sample  <guid/environmentId/featureId/ruleId>

```

**Example**

```
terraform import ibm_app_config_rule.sample 272111153-c118-4116-8116-b811fbc31132/dev/sample_feature/sample_rule
```
