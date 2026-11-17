---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Integration'
description: |-
  Get information about Integration.
---

# ibm_app_config_integration

Retrieve information about an existing IBM Cloud App Configuration integration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration integrations, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_integration" "app_config_integration" {
  guid = "guid"
  integration_id = "integration_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `integration_id` - (Required, String) Integration id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the integration resource.
- `created_time` - (Timestamp) Creation time of the integration.
- `href` - (String) integration URL.
- `integration_type` - (String) Integration type. Supported values are EVENT_NOTIFICATIONS, KMS.
- `metadata` - (List) Settings of the integration, which depend on the integration type.

  Nested scheme for `metadata`:
  - `event_notifications_endpoint` - (String) The URL endpoint of the Event Notifications service instance.
  - `event_notifications_instance_crn` - (String) The CRN of the Event Notifications service instance.
  - `event_notifications_source_id` - (String) Event Notifications source id.
  - `key_status` - (String) Key Protect instance root key status.
  - `kms_endpoint` - (String) The URL endpoint of the Key Protect instance.
  - `kms_instance_crn` - (String) The CRN of the Key Protect service instance.
  - `kms_scheme_type` - (String) For the service name "kms" the schema type is "BYOK".
  - `root_key_id` - (String) The root key id of the Key Protect instance.
- `updated_time` - (Timestamp) Last modified time of the integration.
