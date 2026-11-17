---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Integration'
description: |-
  Manages integrations.
---

# ibm_app_config_integration

Create or delete the integration of App Configuration with the Event Notifications or the Key Protect service. For more information, about App Configuration, see [Getting started with App Configuration](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-getting-started).

~> **Note:** An integration cannot be updated. Changing any argument deletes the integration and creates a new one.

## Example usage

```terraform
# Example 1: Event Notifications
resource "ibm_app_config_integration" "app_config_integration_en" {
  guid = "guid"
  integration_id = "integration_id"
  integration_type = "EVENT_NOTIFICATIONS"
  metadata {
    event_notifications_instance_crn = "event_notifications_instance_crn"
    event_notifications_endpoint = "event_notifications_endpoint"
    event_notifications_source_name = "event_notifications_source_name"
    event_notifications_source_description = "event_notifications_source_description"
  }
}

# Example 2: Key Protect
resource "ibm_app_config_integration" "app_config_integration_kms" {
  guid = "guid"
  integration_id = "integration_id"
  integration_type = "KMS"
  metadata {
    kms_instance_crn = "kms_instance_crn"
    kms_endpoint = "kms_endpoint"
    root_key_id = "root_key_id"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `integration_id` - (Required, String) Integration id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `integration_type` - (Required, String) Integration type. Supported values are EVENT_NOTIFICATIONS, KMS.
- `metadata` - (Required, List) Settings of the integration, which depend on the integration type.
  - `event_notifications_endpoint` - (Optional, String) The URL endpoint of the Event Notifications service instance.
  - `event_notifications_instance_crn` - (Optional, String) The CRN of the Event Notifications service instance.
  - `event_notifications_source_description` - (Optional, String) Source description. This description is shown in the sources page of your Event Notifications instance.
  - `event_notifications_source_name` - (Optional, String) Source name. This name is shown in the sources page of your Event Notifications instance.
  - `kms_endpoint` - (Optional, String) The URL endpoint of the Key Protect instance.
  - `kms_instance_crn` - (Optional, String) The CRN of the Key Protect service instance.
  - `root_key_id` - (Optional, String) The root key id of the Key Protect instance.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the integration resource. The format is `<guid>/<integration_id>`.
- `created_time` - (Timestamp) Creation time of the integration.
- `href` - (String) integration URL.
- `updated_time` - (Timestamp) Last modified time of the integration.

## Import

The `ibm_app_config_integration` resource can be imported by using `guid` of the App Configuration instance and `integrationId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_integration.sample  <guid/integrationId>

```

**Example**

```
terraform import ibm_app_config_integration.sample 272111153-c118-4116-8116-b811fbc31132/sample_integration
```
