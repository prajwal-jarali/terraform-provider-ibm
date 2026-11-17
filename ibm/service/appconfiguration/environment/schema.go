// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package environment holds the environment resource and data sources.
// An environment is a deployment stage, e.g. dev or production, that holds the features and properties of an application.
package environment

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourceName  = "environment"
	resourcesName = "environments"
)

var includeValues = []string{utils.Features, utils.Properties, utils.Snapshots, utils.WorkflowApproval}

// getFields returns the fields of an environment. Every schema of this package is built from them:
// a resource sets what the user writes (Input), a data source reports what the API returns (Output).
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, resourceName),
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.EnvironmentID: mode.RequiredForceNew(utils.NewString(
				"Environment Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.ColorCode: mode.OptionalComputed(utils.NewString(
				"Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`.",
			)),
			utils.Features:         utils.GetAssociatedSchema("feature", utils.FeatureID, resourceName),
			utils.Properties:       utils.GetAssociatedSchema("property", utils.PropertyID, resourceName),
			utils.Snapshots:        utils.GetAssociatedSchema("git config", utils.GitConfigID, resourceName),
			utils.WorkflowApproval: utils.GetWorkflowApprovalSchema(),
		},
	)
}

var idKeys = []string{utils.EnvironmentID}

// DataSourceIBMAppConfigEnvironment returns the `ibm_app_config_environment` data source.
func DataSourceIBMAppConfigEnvironment() *schema.Resource {
	return utils.NewDataSource("Environment", idKeys, getFields(utils.Output), readEnvironment,
		map[string]*schema.Schema{
			utils.Expand:  utils.GetExpandSchema(),
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
		},
	)
}

// DataSourceIBMAppConfigEnvironments returns the `ibm_app_config_environments` data source.
func DataSourceIBMAppConfigEnvironments() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of environments.", getFields(utils.Output), readEnvironments,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:  utils.GetExpandSchema(),
			utils.Sort:    utils.GetSortValuesSchema(resourceName),
			utils.Tags:    utils.GetOptionalTagsSchema(),
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
			utils.Search:  utils.GetSearchSchema("Name OR Tag"),
		},
	)
}

// ResourceIBMAppConfigEnvironment returns the `ibm_app_config_environment` resource.
func ResourceIBMAppConfigEnvironment() *schema.Resource {
	return utils.NewResource("Environment", idKeys, getFields(utils.Input), utils.Calls{
		Create: createEnvironment,
		Get:    getEnvironment,
		Update: updateEnvironment,
		Delete: deleteEnvironment,
	})
}
