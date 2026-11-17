// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package snapshot holds the schemas of the snapshot (git config) resource and data sources.
// A snapshot writes the configuration of a collection in an environment to a file in a git repository.
package snapshot

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	gitBranch     = "git_branch"
	gitConfigName = "git_config_name"
	gitFilePath   = "git_file_path"
	gitToken      = "git_token"
	gitUrl        = "git_url"
	lastSyncTime  = "last_sync_time"
)

const (
	resourceName  = "git config"
	resourcesName = "git_config"
)

// getFields returns the fields of a snapshot, as the user writes them (Input) or as the API returns them (Output).
// The API takes the ids of the collection and the environment, and returns their details as objects.
func getFields(mode utils.Mode) map[string]*schema.Schema {
	fields := utils.Merge(
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.GitConfigID: mode.RequiredForceNew(utils.NewString(
				"Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			gitConfigName: mode.Required(utils.NewString(
				"Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			gitUrl: mode.Required(utils.NewString(
				"Git url used to connect to the git account, " +
					"for example `https://api.github.com/repos/{owner}/{repo_name}` for a personal account.",
			)),
			gitBranch: mode.Required(utils.NewString(
				"Branch to which the configuration is written. Give the branch name only, without slashes at the start or the end. " +
					"The branch must exist in the repository.",
			)),
			gitFilePath: mode.Required(utils.NewString(
				"Path of the file to which the configuration is written. It must contain the file name with a `json` extension. " +
					"Do not use slashes at the start or the end.",
			)),
			lastSyncTime: utils.AsComputed(utils.NewString("Latest time when the snapshot was synced to git.")),
			utils.Collection: utils.AsComputed(utils.NewObject("Details of the collection.", map[string]*schema.Schema{
				utils.Name:         utils.AsComputed(utils.NewString("Collection name.")),
				utils.CollectionID: utils.AsComputed(utils.NewString("Collection Id.")),
			})),
			utils.Environment: utils.AsComputed(utils.NewObject("Details of the environment.", map[string]*schema.Schema{
				utils.Name:          utils.AsComputed(utils.NewString("Environment name.")),
				utils.EnvironmentID: utils.AsComputed(utils.NewString("Environment Id.")),
				utils.ColorCode:     utils.AsComputed(utils.NewString("Color code of the environment.")),
			})),
		},
	)
	if mode == utils.Input {
		fields[utils.CollectionID] = utils.AsRequired(utils.NewString("Collection Id."))
		fields[utils.EnvironmentID] = utils.AsRequired(utils.NewString("Environment Id."))
		fields[gitToken] = utils.AsSensitive(utils.AsRequired(utils.NewString(
			"Git token, with enough permission to write and update the file. It is never returned by the API.",
		)))
	}
	return fields
}

var idKeys = []string{utils.GitConfigID}

// DataSourceIBMAppConfigSnapshot returns the `ibm_app_config_snapshot` data source.
func DataSourceIBMAppConfigSnapshot() *schema.Resource {
	return utils.NewDataSource("Snapshot", idKeys, getFields(utils.Output), readSnapshot)
}

// DataSourceIBMAppConfigSnapshots returns the `ibm_app_config_snapshots` data source.
func DataSourceIBMAppConfigSnapshots() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of snapshots.", getFields(utils.Output), readSnapshots,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Sort:          utils.GetSortValuesSchema("git config"),
			utils.CollectionID:  utils.AsOptional(utils.NewString("Filters the response based on the specified collection_id.")),
			utils.EnvironmentID: utils.AsOptional(utils.NewString("Filters the response based on the specified environment_id.")),
			utils.Search:        utils.GetSearchSchema("Name"),
		},
	)
}

// ResourceIBMAppConfigSnapshot returns the `ibm_app_config_snapshot` resource.
func ResourceIBMAppConfigSnapshot() *schema.Resource {
	return utils.NewResource("Snapshot", idKeys, getFields(utils.Input), utils.Calls{
		Create: createSnapshot,
		Get:    getSnapshot,
		Update: updateSnapshot,
		Delete: deleteSnapshot,
	})
}
