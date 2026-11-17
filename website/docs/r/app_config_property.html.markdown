---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration property'
description: |-
  Manages property.
---

# ibm_app_config_property

Create, update, or delete a property by using IBM Cloud™ App Configuration. For more information, about App Configuration property, see [Properties](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-properties).

~> **Note:** A property is created in every environment of the instance. When the property already exists, for example because it was created from another environment, the arguments that you set are applied to it as an update. Its `type` and `format` must be the same as the existing ones.

## Example usage

```terraform
# Example 1
resource "ibm_app_config_property" "app_config_property" {
  guid = "guid"
  environment_id = "environment_id"
  property_id = "property_id"
  name = "name"
  type = "STRING"
  format = "JSON"
  value = "{\"key\": \"value\"}"
}

# Example 2
resource "ibm_app_config_property" "app_config_property" {
  guid = "guid"
  environment_id = "environment_id"
  property_id = "property_id"
  name = "name"
  type = "NUMERIC"
  value = "10"
  tags = "tag1,tag2"
  collections {
    collection_id = "collection_id1"
  }
  segment_rules {
    rules {
      segments = ["segment_id1", "segment_id2"]
    }
    value = "20"
    order = 1
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id.
- `property_id` - (Required, String) Property id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `name` - (Required, String) The property name.
- `type` - (Required, String) Type of the property. Supported values are BOOLEAN, STRING, NUMERIC, SECRETREF.
- `value` - (Required, String) Value of the property. The value must match the `type` and `format` attributes.
- `collections` - (Optional, List) List of collections associated with the property.
  - `collection_id` - (Required, String) Collection Id.
  - `deleted` - (Optional, Boolean) Remove the association with the collection when set to `true`.
- `description` - (Optional, String) The property description.
- `format` - (Optional, String) Format of the property. Required when `type` is `STRING`, not used for the other types. Supported values are TEXT, JSON, YAML.
- `segment_rules` - (Optional, List) Targeting rules that set different property values for different segments.
  - `order` - (Required, Integer) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. Index 1 based numbering is used for order.
  - `rules` - (Required, List) The list of targeted segments.
    - `segments` - (Required, Array of Strings) The list of segment ids that are used for targeting using the rule.
  - `value` - (Required, String) Value to be used for evaluation for this rule. The value must match the `type` and `format` attributes.
- `tags` - (Optional, String) Tags associated with the property. Allowed special characters are [_. ,-:].

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the property resource. The format is `<guid>/<environment_id>/<property_id>`.
- `created_time` - (Timestamp) Creation time of the property.
- `evaluation_time` - (Timestamp) The last occurrence of the property value evaluation.
- `href` - (String) property URL.
- `segment_exists` - (Boolean) Denotes if the targeting rules are specified for the property.
- `updated_time` - (Timestamp) Last modified time of the property.
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

The `ibm_app_config_property` resource can be imported by using `guid` of the App Configuration instance, `environmentId`, `propertyId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_property.sample  <guid/environmentId/propertyId>

```

**Example**

```
terraform import ibm_app_config_property.sample 272111153-c118-4116-8116-b811fbc31132/dev/sample_property
```
