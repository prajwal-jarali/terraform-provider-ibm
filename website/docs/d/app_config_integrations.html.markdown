---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Integrations'
description: |-
  Get information about Integrations.
---

# ibm_app_config_integrations

Retrieve information about an existing IBM Cloud App Configuration Integrations. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration integrations, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_integrations" "app_config_integrations" {
  guid = "guid"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `expand` - (Optional, Boolean) If set to `true`, returns expanded view of the resource details.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the integration list.
- `first` - (List) URL to navigate to the first list of records.

  Nested scheme for `first`:
  - `href` - (String) URL of the response.
- `integrations` - (List) Array of integrations.

  Nested scheme for `integrations`:
  - `created_time` - (Timestamp) Creation time of the integration.
  - `href` - (String) integration URL.
  - `integration_id` - (String) Integration id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
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
