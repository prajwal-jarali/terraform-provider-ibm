---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration features'
description: |-
  List all the features.
---

# ibm_app_config_features

Retrieve information about an existing IBM Cloud App Configuration features. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration features, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_features" "app_config_features" {
  guid = "guid"
  environment_id = "environment_id"
  collections = ["collection_id"]
  sort = "name"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `collections` - (Optional, Array of Strings) Filter features by a list of collections.
- `expand` - (Optional, Boolean) If set to `true`, returns expanded view of the resource details.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are collections, rules, workflow_approval.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `search` - (Optional, String) Searches for the provided keyword and returns the appropriate row with that value. Here the search happens on the '[Name OR Tag]' of the entity.
- `segments` - (Optional, Array of Strings) Filter features by a list of segments.
- `sort` - (Optional, String) Sort the feature details based on the specified attribute. By default, items are sorted by name. Supported values are created_time, id, name, updated_time.
- `tags` - (Optional, String) Filter the resources to be returned based on the associated tags. Specify the parameter as a list of comma separated tags. Returns resources associated with any of the specified tags.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the feature list.
- `features` - (List) Array of features.

  Nested scheme for `features`:
  - `collections` - (List) List of collections associated with the feature.

    Nested scheme for `collections`:
    - `collection_id` - (String) Collection Id.
    - `name` - (String) Collection name.
  - `created_time` - (Timestamp) Creation time of the feature.
  - `description` - (String) The feature description.
  - `disabled_value` - (String) Value of the feature when it is disabled. The value must match the `type` and `format` attributes.
  - `enabled` - (Boolean) The state of the feature flag.
  - `enabled_value` - (String) Value of the feature when it is enabled. The value must match the `type` and `format` attributes.
  - `evaluation_time` - (Timestamp) The last occurrence of the feature flag value evaluation.
  - `feature_id` - (String) Feature id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `format` - (String) Format of the feature. Required when `type` is `STRING`, not used for the other types. Supported values are TEXT, JSON, YAML.
  - `href` - (String) feature URL.
  - `name` - (String) The feature name.
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
  - `rollout_rule` - (List) The rule of the feature flag that is currently under a progressive rollout. Only one rule can be under rollout at a time.

    Nested scheme for `rollout_rule`:
    - `rule_id` - (String) The identifier of the rule that is under rollout.
    - `type` - (String) The rollout strategy type.
  - `rollout_type` - (String) The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics. Supported values are MANUAL, PROGRESSIVE, GUARDED.
  - `segment_count` - (Integer) The number of targeting rules configured for the feature flag.
  - `segment_exists` - (Boolean) Denotes if the targeting rules are specified for the feature flag.
  - `segment_rules` - (List) Targeting rules that set different feature values for different segments.

    Nested scheme for `segment_rules`:
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
  - `tags` - (String) Tags associated with the feature. Allowed special characters are [_. ,-:].
  - `type` - (String) Type of the feature. Supported values are BOOLEAN, STRING, NUMERIC.
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
