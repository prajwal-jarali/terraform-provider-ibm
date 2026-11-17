// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package property holds the schemas of the property resource and data sources.
// A property lives in an environment and serves one value, which targeting rules can override per segment.
package property

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourceName  = "property"
	resourcesName = "properties"
)

var includeValues = []string{utils.Collections, utils.Rules, utils.WorkflowApproval}

// getFields returns the fields of a property, as the user writes them (Input) or as the API returns them (Output).
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, resourceName),
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.PropertyID: mode.RequiredForceNew(utils.NewString(
				"Property id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.Type: mode.RequiredForceNew(utils.NewEnum(
				"Type of the property.", "BOOLEAN", "STRING", "NUMERIC", "SECRETREF",
			)),
			utils.Format: mode.OptionalForceNew(utils.GetFormatSchema(resourceName)),
			utils.Value: mode.Required(utils.NewString(
				"Value of the property. The value must match the `type` and `format` attributes.",
			)),
			utils.SegmentRules: utils.GetSegmentRulesSchema(mode, resourceName, utils.GetRuleFields(mode)),
			utils.SegmentExists: utils.AsComputed(utils.NewBool(
				"Denotes if the targeting rules are specified for the property.",
			)),
			utils.Collections: utils.GetCollectionsSchema(mode, resourceName),
			utils.EvaluationTime: utils.AsComputed(utils.NewString(
				"The last occurrence of the property value evaluation.",
			)),
			utils.WorkflowApproval: utils.GetWorkflowApprovalSchema(),
		},
	)
}

func getEnvironmentID() map[string]*schema.Schema {
	return map[string]*schema.Schema{utils.EnvironmentID: utils.GetParentIDSchema("Environment Id.")}
}

var idKeys = []string{utils.EnvironmentID, utils.PropertyID}

// DataSourceIBMAppConfigProperty returns the `ibm_app_config_property` data source.
func DataSourceIBMAppConfigProperty() *schema.Resource {
	return utils.NewDataSource("Property", idKeys, getFields(utils.Output), readProperty,
		getEnvironmentID(),
		map[string]*schema.Schema{
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
		},
	)
}

// DataSourceIBMAppConfigProperties returns the `ibm_app_config_properties` data source.
func DataSourceIBMAppConfigProperties() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of properties.", getFields(utils.Output), readProperties,
		getEnvironmentID(),
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:      utils.GetExpandSchema(),
			utils.Sort:        utils.GetSortValuesSchema(resourceName),
			utils.Tags:        utils.GetOptionalTagsSchema(),
			utils.Collections: utils.GetFilterSchema("Filter properties by a list of collections."),
			utils.Segments:    utils.GetFilterSchema("Filter properties by a list of segments."),
			utils.Include:     utils.GetIncludeValuesSchema(includeValues...),
			utils.Search:      utils.GetSearchSchema("Name OR Tag"),
		},
	)
}

// ResourceIBMAppConfigProperty returns the `ibm_app_config_property` resource. Its id is `<guid>/<environment_id>/<property_id>`.
func ResourceIBMAppConfigProperty() *schema.Resource {
	return utils.NewResource("Property", idKeys, getFields(utils.Input), utils.Calls{
		Create: createProperty,
		Get:    getProperty,
		Update: updateProperty,
		Delete: deleteProperty,
	}, getEnvironmentID())
}
