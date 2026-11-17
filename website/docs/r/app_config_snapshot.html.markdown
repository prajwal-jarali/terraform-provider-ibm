---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Snapshots'
description: |-
  Manages snapshots.
---

# ibm_app_config_snapshot

Create, update, or delete a snapshot by using IBM Cloud™ App Configuration. For more information, about App Configuration snapshots, see [snapshots](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-snapshots).

## Example usage

```terraform
resource "ibm_app_config_snapshot" "app_config_snapshot" {
  guid = "guid"
  collection_id = "collection_id"
  environment_id = "environment_id"
  git_config_id = "git_config_id"
  git_config_name = "git_config_name"
  git_url = "git_url"
  git_branch = "git_branch"
  git_file_path = "git_file_path"
  git_token = "git_token"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `collection_id` - (Required, String) Collection Id.
- `environment_id` - (Required, String) Environment Id.
- `git_config_id` - (Required, String) Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `git_branch` - (Required, String) Branch to which the configuration is written. Give the branch name only, without slashes at the start or the end. The branch must exist in the repository.
- `git_config_name` - (Required, String) Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
- `git_file_path` - (Required, String) Path of the file to which the configuration is written. It must contain the file name with a `json` extension. Do not use slashes at the start or the end.
- `git_token` - (Required, String) Git token, with enough permission to write and update the file. It is never returned by the API.
- `git_url` - (Required, String) Git url used to connect to the git account, for example `https://api.github.com/repos/{owner}/{repo_name}` for a personal account.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the snapshot resource. The format is `<guid>/<git_config_id>`.
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
- `href` - (String) git config URL.
- `last_sync_time` - (Timestamp) Latest time when the snapshot was synced to git.
- `updated_time` - (Timestamp) Last modified time of the git config.

## Import

The `ibm_app_config_snapshot` resource can be imported by using `guid` of the App Configuration instance and `git_config_id`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_snapshot.sample  <guid/git_config_id>

```

**Example**

```
terraform import ibm_app_config_snapshot.sample 272111153-c118-4116-8116-b811fbc31132/sample
```
