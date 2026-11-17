---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration workflow config'
description: |-
  Manages workflow configs.
---

# ibm_app_config_workflow_config

Create, update, or delete a workflow configuration by using IBM Cloud™ App Configuration. For more information, about App Configuration workflow configuration, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

~> **Note:** The `provider_config` argument is the `provider` of the App Configuration API. Terraform reserves the name `provider`.

## Example usage

```terraform
resource "ibm_app_config_workflow_config" "app_config_workflow_config" {
  guid = "guid"
  workflow_id = "workflow_id"
  name = "name"
  enabled = true
  provider_config {
    type = "SERVICENOW_IBM"
    metadata {
      crn_mask = "crn_mask"
      sm_secret_crn = "sm_secret_crn"
      default_email = "default_email"
      approval_expiration = 10
      service_environment = "production"
    }
  }
  scope {
    collections {
      mode = "ALL"
    }
    environments {
      environment_id = "environment_id"
      resources {
        environment {
          enable = true
        }
        features {
          mode = "ALL"
        }
      }
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `workflow_id` - (Required, String) Unique identifier for the workflow configuration.
- `enabled` - (Required, Boolean) Whether the workflow is enabled.
- `name` - (Required, String) Name of the workflow configuration.
- `provider_config` - (Required, List) Provider configuration for the workflow.
  - `metadata` - (Required, List) Metadata of the ServiceNow provider. Which settings apply depends on the provider type.
    - `approval_expiration` - (Optional, Integer) Approval expiration time in days.
    - `approval_group_name` - (Optional, String) Name of the approval group in ServiceNow.
    - `crn_mask` - (Optional, String) CRN mask for the service.
    - `default_email` - (Optional, String) The email id to which the change request is assigned.
    - `service_environment` - (Optional, String) The ServiceNow change request environment context. If not provided, it defaults to production. Supported values are pre_prod, production.
    - `sm_secret_crn` - (Optional, String) Secret Manager secret CRN for the ServiceNow credentials.
    - `sm_secret_id` - (Optional, String) Secret Manager secret ID for the ServiceNow credentials.
    - `workflow_url` - (Optional, String) ServiceNow instance URL.
  - `type` - (Required, String) Provider type. Supported values are SERVICENOW_IBM, SERVICENOW_EXTERNAL.
- `scope` - (Required, List) Scope of the workflow. At least one of collections, segments or environments must be specified.
  - `collections` - (Optional, List) Scope of the collections. Only the ALL mode is supported.
    - `mode` - (Required, String) Scope mode. Supported values are ALL.
    - `ids` - (Optional, Array of Strings) List of ids. Always empty when the mode is ALL.
  - `environments` - (Optional, List) List of environments where the workflow applies.
    - `environment_id` - (Required, String) Environment identifier.
    - `resources` - (Optional, List) Resources of the environment that the workflow applies to.
      - `environment` - (Optional, List) The environment itself.
        - `enable` - (Optional, Boolean) Whether changes to the environment need approval.
      - `features` - (Optional, List) Scope of the features. Only the ALL mode is supported.
        - `mode` - (Required, String) Scope mode. Supported values are ALL.
        - `ids` - (Optional, Array of Strings) List of ids. Always empty when the mode is ALL.
      - `properties` - (Optional, List) Scope of the properties. Only the ALL mode is supported.
        - `mode` - (Required, String) Scope mode. Supported values are ALL.
        - `ids` - (Optional, Array of Strings) List of ids. Always empty when the mode is ALL.
  - `segments` - (Optional, List) Scope of the segments. Only the ALL mode is supported.
    - `mode` - (Required, String) Scope mode. Supported values are ALL.
    - `ids` - (Optional, Array of Strings) List of ids. Always empty when the mode is ALL.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the workflow config resource. The format is `<guid>/<workflow_id>`.
- `created_time` - (Timestamp) Creation time of the workflow config.
- `href` - (String) workflow config URL.
- `updated_time` - (Timestamp) Last modified time of the workflow config.

## Import

The `ibm_app_config_workflow_config` resource can be imported by using `guid` of the App Configuration instance and `workflowId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_workflow_config.sample  <guid/workflowId>

```

**Example**

```
terraform import ibm_app_config_workflow_config.sample 272111153-c118-4116-8116-b811fbc31132/sample_workflow
```
