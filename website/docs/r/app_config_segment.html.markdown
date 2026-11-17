---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segment'
description: |-
  Manages segments.
---

# ibm_app_config_segment

Create, update, or delete a segment by using IBM Cloud™ App Configuration. For more information, about App Configuration segment, see [segments](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-segments).

## Example usage

```terraform
resource "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  name = "name"
  segment_id = "segment_id"
  description = "description"
  tags = "tag1,tag2"
  rules {
    attribute_name = "email"
    operator = "endsWith"
    values = ["@in.xyz.com", "@us.abc.com"]
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `segment_id` - (Required, String) Segment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `name` - (Required, String) The segment name.
- `rules` - (Required, List) Rules that determine if an entity belongs to the segment during feature and property evaluation. An entity belongs to the segment when it matches any of the rules.
  - `attribute_name` - (Required, String) Attribute name.
  - `operator` - (Required, String) Operator to be used for the evaluation if the entity belongs to the segment. Supported values are contains, endsWith, greaterThan, greaterThanEquals, is, isNot, lesserThan, lesserThanEquals, notContains, notEndsWith, notStartsWith, startsWith.
  - `values` - (Required, Array of Strings) List of values. Entities matching any of the given values belong to the segment.
- `description` - (Optional, String) The segment description.
- `tags` - (Optional, String) Tags associated with the segment. Allowed special characters are [_. ,-:].

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the segment resource. The format is `<guid>/<segment_id>`.
- `created_time` - (Timestamp) Creation time of the segment.
- `features` - (List) List of features associated with this segment.

  Nested scheme for `features`:
  - `feature_id` - (String) feature Id.
  - `name` - (String) feature name.
- `href` - (String) segment URL.
- `properties` - (List) List of propertys associated with this segment.

  Nested scheme for `properties`:
  - `name` - (String) property name.
  - `property_id` - (String) property Id.
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

## Import

The `ibm_app_config_segment` resource can be imported by using `guid` of the App Configuration instance and `segmentId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_segment.sample  <guid/segmentId>

```

**Example**

```
terraform import ibm_app_config_segment.sample 272111153-c118-4116-8116-b811fbc31132/sample_segment
```
