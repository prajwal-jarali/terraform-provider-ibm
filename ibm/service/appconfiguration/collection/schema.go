// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package collection holds the schemas of the collection resource and data sources.
// A collection groups the features and properties of an application.
package collection

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	featuresCount   = "features_count"
	propertiesCount = "properties_count"
	snapshotCount   = "snapshot_count"
)

const (
	resourceName  = "collection"
	resourcesName = "collections"
)

var includeValues = []string{utils.Features, utils.Properties, utils.Snapshots, utils.WorkflowApproval}

// getFields returns the fields of a collection. Every schema of this package is built from them:
// a resource sets what the user writes (Input), a data source reports what the API returns (Output).
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, resourceName),
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.CollectionID: mode.RequiredForceNew(utils.NewString(
				"Collection Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.Features:         utils.GetAssociatedSchema("feature", utils.FeatureID, resourceName),
			utils.Properties:       utils.GetAssociatedSchema("property", utils.PropertyID, resourceName),
			utils.Snapshots:        utils.GetAssociatedSchema("git config", utils.GitConfigID, resourceName),
			featuresCount:          utils.AsComputed(utils.NewInt("Number of features associated with the collection.")),
			propertiesCount:        utils.AsComputed(utils.NewInt("Number of properties associated with the collection.")),
			snapshotCount:          utils.AsComputed(utils.NewInt("Number of snapshots associated with the collection.")),
			utils.WorkflowApproval: utils.GetWorkflowApprovalSchema(),
		},
	)
}

var idKeys = []string{utils.CollectionID}

// DataSourceIBMAppConfigCollection returns the `ibm_app_config_collection` data source.
func DataSourceIBMAppConfigCollection() *schema.Resource {
	return utils.NewDataSource("Collection", idKeys, getFields(utils.Output), readCollection,
		map[string]*schema.Schema{
			utils.Expand:  utils.GetExpandSchema(),
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
		},
	)
}

// DataSourceIBMAppConfigCollections returns the `ibm_app_config_collections` data source.
func DataSourceIBMAppConfigCollections() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of collections.", getFields(utils.Output), readCollections,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:     utils.GetExpandSchema(),
			utils.Sort:       utils.GetSortValuesSchema(resourceName),
			utils.Tags:       utils.GetOptionalTagsSchema(),
			utils.Features:   utils.GetFilterSchema("Filter collections by a list of features."),
			utils.Properties: utils.GetFilterSchema("Filter collections by a list of properties."),
			utils.Include:    utils.GetIncludeValuesSchema(includeValues...),
			utils.Search:     utils.GetSearchSchema("Name OR Tag"),
		},
	)
}

// ResourceIBMAppConfigCollection returns the `ibm_app_config_collection` resource.
func ResourceIBMAppConfigCollection() *schema.Resource {
	return utils.NewResource("Collection", idKeys, getFields(utils.Input), utils.Calls{
		Create: createCollection,
		Get:    getCollection,
		Update: updateCollection,
		Delete: deleteCollection,
	})
}
