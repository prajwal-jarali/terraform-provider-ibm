---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration Snapshots'
description: |-
  Get information about Snapshots.
---

# ibm_app_config_snapshots

Retrieve information about an existing IBM Cloud App Configuration snapshots. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration snapshots, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_snapshots" "app_config_snapshots" {
  guid = "guid"
  sort = "name"
  collection_id = "collection_id"
  environment_id = "environment_id"
  limit = 10
  offset = 0
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `collection_id` - (Optional, String) Filters the response based on the specified collection_id.
- `environment_id` - (Optional, String) Filters the response based on the specified environment_id.
- `limit` - (Optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (Optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `search` - (Optional, String) Searches for the provided keyword and returns the appropriate row with that value. Here the search happens on the '[Name]' of the entity.
- `sort` - (Optional, String) Sort the git config details based on the specified attribute. By default, items are sorted by name. Supported values are created_time, id, name, updated_time.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the snapshot list.
- `first` - (List) URL to navigate to the first list of records.

  Nested scheme for `first`:
  - `href` - (String) URL of the response.
- `git_config` - (List) Array of snapshots.

  Nested scheme for `git_config`:
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
  - `git_config_id` - (String) Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `git_config_name` - (String) Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.
  - `git_file_path` - (String) Path of the file to which the configuration is written. It must contain the file name with a `json` extension. Do not use slashes at the start or the end.
  - `git_url` - (String) Git url used to connect to the git account, for example `https://api.github.com/repos/{owner}/{repo_name}` for a personal account.
  - `href` - (String) git config URL.
  - `last_sync_time` - (Timestamp) Latest time when the snapshot was synced to git.
  - `updated_time` - (Timestamp) Last modified time of the git config.
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
