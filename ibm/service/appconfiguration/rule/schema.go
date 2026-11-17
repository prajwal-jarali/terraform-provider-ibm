// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package rule holds the schemas of the targeting rule of a feature flag: a resource and data sources.
// A rule serves a value to the entities of the segments it targets. Rules are evaluated in `order`.
package rule

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const resourcesName = utils.SegmentRules

// getFields returns the fields of a feature rule, as the user writes them (Input) or as the API returns them (Output).
// The order of a rule is not set on the rule itself, the API changes it with a separate reorder call.
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetFeatureRuleFields(mode),
		map[string]*schema.Schema{
			utils.RuleID: mode.RequiredForceNew(utils.NewString("Rule id is an identifier used for identifying the rule.")),
			utils.Order: utils.AsComputed(utils.NewInt(
				"Order of the rule, used during evaluation. The value of the first matching rule is used for evaluation.",
			)),
			utils.Href: utils.AsComputed(utils.NewString("Feature flag rule URL.")),
		},
	)
}

// getParents returns the feature the rule belongs to. The ids are part of the API path, so they cannot change.
func getParents() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.EnvironmentID: utils.GetParentIDSchema("Environment Id."),
		utils.FeatureID:     utils.GetParentIDSchema("Feature Id."),
	}
}

var idKeys = []string{utils.EnvironmentID, utils.FeatureID, utils.RuleID}

// DataSourceIBMAppConfigRule returns the data source of one rule of a feature.
func DataSourceIBMAppConfigRule() *schema.Resource {
	return utils.NewDataSource("Rule", idKeys, getFields(utils.Output), readRule, getParents())
}

// DataSourceIBMAppConfigRules returns the data source that lists the rules of a feature.
func DataSourceIBMAppConfigRules() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of rules of the feature.", getFields(utils.Output), readRules,
		getParents(),
	)
}

// ResourceIBMAppConfigRule returns the targeting rule resource.
func ResourceIBMAppConfigRule() *schema.Resource {
	return utils.NewResource("Rule", idKeys, getFields(utils.Input), utils.Calls{
		Create: createRule,
		Get:    getRule,
		Update: updateRule,
		Delete: deleteRule,
	}, getParents())
}
