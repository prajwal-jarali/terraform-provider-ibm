---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration property'
description: |-
  Get information about property.
---

# ibm_app_config_property

Retrieve information about an existing IBM Cloud App Configuration property. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration property, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_property" "app_config_property" {
  guid = "guid"
  environment_id = "environment_id"
  property_id = "property_id"
  include = ["collections", "rules"]
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `property_id` - (Required, String) Property id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are collections, rules, workflow_approval.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the property resource.
- `collections` - (List) List of collections associated with the property.

  Nested scheme for `collections`:
  - `collection_id` - (String) Collection Id.
  - `name` - (String) Collection name.
- `created_time` - (Timestamp) Creation time of the property.
- `description` - (String) The property description.
- `evaluation_time` - (Timestamp) The last occurrence of the property value evaluation.
- `format` - (String) Format of the property. Required when `type` is `STRING`, not used for the other types. Supported values are TEXT, JSON, YAML.
- `href` - (String) property URL.
- `name` - (String) The property name.
- `segment_exists` - (Boolean) Denotes if the targeting rules are specified for the property.
- `segment_rules` - (List) Targeting rules that set different property values for different segments.

  Nested scheme for `segment_rules`:
  - `order` - (Integer) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. Index 1 based numbering is used for order.
  - `rules` - (List) The list of targeted segments.

    Nested scheme for `rules`:
    - `segments` - (Array of Strings) The list of segment ids that are used for targeting using the rule.
  - `value` - (String) Value to be used for evaluation for this rule. The value must match the `type` and `format` attributes.
- `tags` - (String) Tags associated with the property. Allowed special characters are [_. ,-:].
- `type` - (String) Type of the property. Supported values are BOOLEAN, STRING, NUMERIC, SECRETREF.
- `updated_time` - (Timestamp) Last modified time of the property.
- `value` - (String) Value of the property. The value must match the `type` and `format` attributes.
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
