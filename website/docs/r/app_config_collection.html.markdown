---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration collection'
description: |-
  Manages collections.
---

# ibm_app_config_collection

Create, update, or delete a collection by using IBM Cloud™ App Configuration. For more information, about App Configuration collections, see [Collections](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-collections).

## Example usage

```terraform
resource "ibm_app_config_collection" "app_config_collection" {
  guid = "guid"
  name = "name"
  collection_id = "collection_id"
  description = "description"
  tags = "tag1,tag2"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `collection_id` - (Required, String) Collection Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `name` - (Required, String) The collection name.
- `description` - (Optional, String) The collection description.
- `tags` - (Optional, String) Tags associated with the collection. Allowed special characters are [_. ,-:].

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the collection resource. The format is `<guid>/<collection_id>`.
- `created_time` - (Timestamp) Creation time of the collection.
- `features` - (List) List of features associated with this collection.

  Nested scheme for `features`:
  - `feature_id` - (String) feature Id.
  - `name` - (String) feature name.
- `features_count` - (Integer) Number of features associated with the collection.
- `href` - (String) collection URL.
- `properties` - (List) List of propertys associated with this collection.

  Nested scheme for `properties`:
  - `name` - (String) property name.
  - `property_id` - (String) property Id.
- `properties_count` - (Integer) Number of properties associated with the collection.
- `snapshot_count` - (Integer) Number of snapshots associated with the collection.
- `snapshots` - (List) List of git configs associated with this collection.

  Nested scheme for `snapshots`:
  - `git_config_id` - (String) git config Id.
  - `name` - (String) git config name.
- `updated_time` - (Timestamp) Last modified time of the collection.
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

The `ibm_app_config_collection` resource can be imported by using `guid` of the App Configuration instance and `collectionId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_collection.sample  <guid/collectionId>

```

**Example**

```
terraform import ibm_app_config_collection.sample 272111153-c118-4116-8116-b811fbc31132/sample_collection
```
