---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration origin config'
description: |-
  Manages the allowed origins.
---

# ibm_app_config_origin_config

Create, update, or delete the allowed origins by using IBM Cloud™ App Configuration. For more information, about App Configuration origin configuration, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

~> **Note:** The App Configuration service has one origin configuration for each instance, and the API has no call to delete it. Destroying this resource only removes it from the Terraform state and leaves the allowed origins as they are.

## Example usage

```terraform
resource "ibm_app_config_origin_config" "app_config_origin_config" {
  guid = "guid"
  allowed_origins = ["https://example.com"]
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `allowed_origins` - (Required, Array of Strings) List of allowed origins.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the origin config resource. It is the `guid` of the App Configuration instance.
- `created_time` - (Timestamp) Creation time of the origin config.
- `href` - (String) origin config URL.
- `updated_time` - (Timestamp) Last modified time of the origin config.

## Import

The `ibm_app_config_origin_config` resource can be imported by using `guid` of the App Configuration instance. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_origin_config.sample  <guid>

```

**Example**

```
terraform import ibm_app_config_origin_config.sample 272111153-c118-4116-8116-b811fbc31132
```
