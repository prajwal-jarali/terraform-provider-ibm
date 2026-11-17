---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag'
description: |-
  Manages feature flag.
---

# ibm_app_config_feature

Create, update, or delete a feature flag by using IBM Cloud™ App Configuration. For more information, about App Configuration feature flag, see [Feature flags](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-feature-flags).

~> **Note:** A feature flag is created in every environment of the instance. When the feature flag already exists, for example because it was created from another environment, the arguments that you set are applied to it as an update. Its `type` and `format` must be the same as the existing ones.

## Example usage

```terraform
# Example 1
resource "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  name = "name"
  type = "NUMERIC"
  enabled_value = "10"
  disabled_value = "0"
  enabled = false
}

# Example 2
resource "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  environment_id = "environment_id"
  feature_id = "feature_id"
  name = "name"
  type = "STRING"
  format = "TEXT"
  tags = "tag1,tag2"
  enabled_value = "enabled"
  disabled_value = "disabled"
  enabled = true
  collections {
    collection_id = "collection_id1"
  }
  # only use this deleted attribute during
  # update of feature
  collections {
    collection_id = "collection_id2"
    deleted = true
  }
  segment_rules {
    rules {
      segments = ["segment_id1", "segment_id2"]
    }
    value = "value1"
    rollout_percentage = 80
    order = 1
  }
  segment_rules {
    rules {
      segments = ["segment_id3", "segment_id4"]
    }
    value = "value2"
    order = 2
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `feature_id` - (Required, String) Feature id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `disabled_value` - (Required, String) Value of the feature when it is disabled. The value must match the `type` and `format` attributes.
- `enabled_value` - (Required, String) Value of the feature when it is enabled. The value must match the `type` and `format` attributes.
- `name` - (Required, String) The feature name.
- `type` - (Required, String) Type of the feature. Supported values are BOOLEAN, STRING, NUMERIC.
- `collections` - (Optional, List) List of collections associated with the feature.
  - `collection_id` - (Required, String) Collection Id.
  - `deleted` - (Optional, Boolean) Remove the association with the collection when set to `true`.
- `description` - (Optional, String) The feature description.
- `enabled` - (Optional, Boolean) The state of the feature flag.
- `format` - (Optional, String) Format of the feature. Required when `type` is `STRING`, not used for the other types. Supported values are TEXT, JSON, YAML.
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
- `segment_rules` - (Optional, List) Targeting rules that set different feature values for different segments.
  - `order` - (Required, Integer) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. Index 1 based numbering is used for order.
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
  - `rule_id` - (Optional, String) Rule id is an identifier used for identifying the rule.
  - `rule_name` - (Optional, String) Name assigned to a particular rule.
- `tags` - (Optional, String) Tags associated with the feature. Allowed special characters are [_. ,-:].

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the feature flag resource. The format is `<guid>/<environment_id>/<feature_id>`.
- `created_time` - (Timestamp) Creation time of the feature.
- `enabled` - (Boolean) The state of the feature flag.
- `evaluation_time` - (Timestamp) The last occurrence of the feature flag value evaluation.
- `href` - (String) feature URL.
- `rollout_percentage` - (Integer) Rollout percentage. Supported only for Lite and Enterprise plans.
- `rollout_rule` - (List) The rule of the feature flag that is currently under a progressive rollout. Only one rule can be under rollout at a time.

  Nested scheme for `rollout_rule`:
  - `rule_id` - (String) The identifier of the rule that is under rollout.
  - `type` - (String) The rollout strategy type.
- `rollout_type` - (String) The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics. Supported values are MANUAL, PROGRESSIVE, GUARDED.
- `segment_count` - (Integer) The number of targeting rules configured for the feature flag.
- `segment_exists` - (Boolean) Denotes if the targeting rules are specified for the feature flag.
- `updated_time` - (Timestamp) Last modified time of the feature.
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

## Import

The `ibm_app_config_feature` resource can be imported by using `guid` of the App Configuration instance, `environmentId`, `featureId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_feature.sample  <guid/environmentId/featureId>

```

**Example**

```
terraform import ibm_app_config_feature.sample 272111153-c118-4116-8116-b811fbc31132/dev/sample_feature
```
