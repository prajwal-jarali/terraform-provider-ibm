---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Snapshots'
description: |-
  Get information about Snapshot.
---

# ibm_app_config_snapshot

Retrieve information about an existing IBM Cloud App Configuration snapshot. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration snapshots, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_snapshot" "app_config_snapshot" {
  guid = "guid"
  git_config_id = "git_config_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `git_config_id` - (Required, String) Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the snapshot resource.
- `collection` - (List) Details of the collection.

  Nested scheme for `collection`:
  - `collection_id` - (String) Collection Id.
  - `name` - (String) Collection name.
- `created_time` - (Timestamp) Creation time of the git config.
- `environment` - (List) Details of the environment.

  Nested scheme for `environment`:
  - `color_code` - (String) Color code of the environment.
  - `environment_id` - (String) Environment Id.
  - `name` - (String) Environment name.
- `git_branch` - (String) Branch to which the configuration is written. Give the branch name only, without slashes at the start or the end. The branch must exist in the repository.
- `git_config_name` - (String) Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `git_file_path` - (String) Path of the file to which the configuration is written. It must contain the file name with a `json` extension. Do not use slashes at the start or the end.
- `git_url` - (String) Git url used to connect to the git account, for example `https://api.github.com/repos/{owner}/{repo_name}` for a personal account.
- `href` - (String) git config URL.
- `last_sync_time` - (Timestamp) Latest time when the snapshot was synced to git.
- `updated_time` - (Timestamp) Last modified time of the git config.
