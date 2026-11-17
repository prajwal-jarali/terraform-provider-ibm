---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segment'
description: |-
  Get information about segment.
---

# ibm_app_config_segment

Retrieve information about an existing IBM Cloud App Configuration segment. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration segment, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  segment_id = "segment_id"
  include = ["features", "properties"]
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `segment_id` - (Required, String) Segment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `include` - (Optional, Array of Strings) Additional details to include in the response. Supported values are features, properties, workflow_approval.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the segment resource.
- `created_time` - (Timestamp) Creation time of the segment.
- `description` - (String) The segment description.
- `features` - (List) List of features associated with this segment.

  Nested scheme for `features`:
  - `feature_id` - (String) feature Id.
  - `name` - (String) feature name.
- `href` - (String) segment URL.
- `name` - (String) The segment name.
- `properties` - (List) List of propertys associated with this segment.

  Nested scheme for `properties`:
  - `name` - (String) property name.
  - `property_id` - (String) property Id.
- `rules` - (List) Rules that determine if an entity belongs to the segment during feature and property evaluation. An entity belongs to the segment when it matches any of the rules.

  Nested scheme for `rules`:
  - `attribute_name` - (String) Attribute name.
  - `operator` - (String) Operator to be used for the evaluation if the entity belongs to the segment. Supported values are contains, endsWith, greaterThan, greaterThanEquals, is, isNot, lesserThan, lesserThanEquals, notContains, notEndsWith, notStartsWith, startsWith.
  - `values` - (Array of Strings) List of values. Entities matching any of the given values belong to the segment.
- `tags` - (String) Tags associated with the segment. Allowed special characters are [_. ,-:].
- `updated_time` - (Timestamp) Last modified time of the segment.
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
