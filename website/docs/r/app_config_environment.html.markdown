---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration environments'
description: |-
  Manages environments.
---

# ibm_app_config_environment

Create, update, or delete an environment by using IBM Cloud™ App Configuration. For more information, about App Configuration, see [Getting started with App Configuration](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-getting-started).

## Example usage

```terraform
resource "ibm_app_config_environment" "app_config_environment" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  description = "description"
  tags = "tag1,tag2"
  color_code = "#FF0000"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) Environment Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `name` - (Required, String) The environment name.
- `color_code` - (Optional, String) Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`.
- `description` - (Optional, String) The environment description.
- `tags` - (Optional, String) Tags associated with the environment. Allowed special characters are [_. ,-:].

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of an environment resource. The format is `<guid>/<environment_id>`.
- `created_time` - (Timestamp) Creation time of the environment.
- `features` - (List) List of features associated with this environment.

  Nested scheme for `features`:
  - `feature_id` - (String) feature Id.
  - `name` - (String) feature name.
- `href` - (String) environment URL.
- `properties` - (List) List of propertys associated with this environment.

  Nested scheme for `properties`:
  - `name` - (String) property name.
  - `property_id` - (String) property Id.
- `snapshots` - (List) List of git configs associated with this environment.

  Nested scheme for `snapshots`:
  - `git_config_id` - (String) git config Id.
  - `name` - (String) git config name.
- `updated_time` - (Timestamp) Last modified time of the environment.
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

The `ibm_app_config_environment` resource can be imported by using `guid` of the App Configuration instance and `environmentId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_environment.sample  <guid/environmentId>

```

**Example**

```
terraform import ibm_app_config_environment.sample 272111153-c118-4116-8116-b811fbc31132/dev
```
