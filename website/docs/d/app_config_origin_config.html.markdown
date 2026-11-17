---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration origin config'
description: |-
  Get information about the allowed origins.
---

# ibm_app_config_origin_config

Retrieve information about an existing IBM Cloud App Configuration origin config. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration origin configuration, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_origin_config" "app_config_origin_config" {
  guid = "guid"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the origin config. It is the `guid` of the App Configuration instance.
- `allowed_origins` - (Array of Strings) List of allowed origins.
- `created_time` - (Timestamp) Creation time of the origin config.
- `href` - (String) origin config URL.
- `updated_time` - (Timestamp) Last modified time of the origin config.
