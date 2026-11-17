// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package feature holds the schemas of the feature flag resource and data sources.
// A feature flag lives in an environment and serves one value when it is enabled and another when disabled.
package feature

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	rolloutRule  = "rollout_rule"
	segmentCount = "segment_count"
)

const (
	resourceName  = "feature"
	resourcesName = "features"
)

var includeValues = []string{utils.Collections, utils.Rules, utils.WorkflowApproval}

// getFields returns the fields of a feature flag, as the user writes them (Input) or as the API returns them (Output).
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, resourceName),
		utils.GetAuditFields(resourceName),
		utils.GetRolloutFields(mode),
		map[string]*schema.Schema{
			utils.FeatureID: mode.RequiredForceNew(utils.NewString(
				"Feature id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.Type:   mode.RequiredForceNew(utils.NewEnum("Type of the feature.", "BOOLEAN", "STRING", "NUMERIC")),
			utils.Format: mode.OptionalForceNew(utils.GetFormatSchema(resourceName)),
			utils.EnabledValue: mode.Required(utils.NewString(
				"Value of the feature when it is enabled. The value must match the `type` and `format` attributes.",
			)),
			utils.DisabledValue: mode.Required(utils.NewString(
				"Value of the feature when it is disabled. The value must match the `type` and `format` attributes.",
			)),
			utils.Enabled:      mode.OptionalComputed(utils.NewBool("The state of the feature flag.")),
			utils.SegmentRules: utils.GetSegmentRulesSchema(mode, resourceName, utils.GetFeatureRuleFields(mode)),
			utils.SegmentExists: utils.AsComputed(utils.NewBool(
				"Denotes if the targeting rules are specified for the feature flag.",
			)),
			segmentCount: utils.AsComputed(utils.NewInt(
				"The number of targeting rules configured for the feature flag.",
			)),
			rolloutRule:       getRolloutRule(),
			utils.Collections: utils.GetCollectionsSchema(mode, resourceName),
			utils.EvaluationTime: utils.AsComputed(utils.NewString(
				"The last occurrence of the feature flag value evaluation.",
			)),
			utils.WorkflowApproval: utils.GetWorkflowApprovalSchema(),
		},
	)
}

// getRolloutRule returns the rule that is currently under a progressive rollout.
func getRolloutRule() *schema.Schema {
	return utils.AsComputed(utils.NewObject(
		"The rule of the feature flag that is currently under a progressive rollout. Only one rule can be under rollout at a time.",
		map[string]*schema.Schema{
			utils.Type:   utils.AsComputed(utils.NewString("The rollout strategy type.")),
			utils.RuleID: utils.AsComputed(utils.NewString("The identifier of the rule that is under rollout.")),
		},
	))
}

func getEnvironmentID() map[string]*schema.Schema {
	return map[string]*schema.Schema{utils.EnvironmentID: utils.GetParentIDSchema("Environment Id.")}
}

var idKeys = []string{utils.EnvironmentID, utils.FeatureID}

// DataSourceIBMAppConfigFeature returns the `ibm_app_config_feature` data source.
func DataSourceIBMAppConfigFeature() *schema.Resource {
	return utils.NewDataSource("Feature", idKeys, getFields(utils.Output), readFeature,
		getEnvironmentID(),
		map[string]*schema.Schema{
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
		},
	)
}

// DataSourceIBMAppConfigFeatures returns the `ibm_app_config_features` data source.
func DataSourceIBMAppConfigFeatures() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of features.", getFields(utils.Output), readFeatures,
		getEnvironmentID(),
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:      utils.GetExpandSchema(),
			utils.Sort:        utils.GetSortValuesSchema(resourceName),
			utils.Tags:        utils.GetOptionalTagsSchema(),
			utils.Collections: utils.GetFilterSchema("Filter features by a list of collections."),
			utils.Segments:    utils.GetFilterSchema("Filter features by a list of segments."),
			utils.Include:     utils.GetIncludeValuesSchema(includeValues...),
			utils.Search:      utils.GetSearchSchema("Name OR Tag"),
		},
	)
}

// ResourceIBMAppConfigFeature returns the `ibm_app_config_feature` resource. Its id is `<guid>/<environment_id>/<feature_id>`.
func ResourceIBMAppConfigFeature() *schema.Resource {
	return utils.NewResource("Feature", idKeys, getFields(utils.Input), utils.Calls{
		Create: createFeature,
		Get:    getFeature,
		Update: updateFeature,
		Delete: deleteFeature,
	}, getEnvironmentID())
}
