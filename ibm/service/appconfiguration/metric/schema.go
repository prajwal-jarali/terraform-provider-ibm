// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package metric holds the schemas of the metric resource and data sources.
// A metric records events of the application, so that a guarded rollout can tell if a change made things worse.
package metric

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	environmentName      = "environment_name"
	eventType            = "event_type"
	experimentId         = "experiment_id"
	experimentStatus     = "experiment_status"
	experiments          = "experiments"
	experimentsCount     = "experiments_count"
	featureName          = "feature_name"
	guardedRollouts      = "guarded_rollouts"
	guardedRolloutsCount = "guarded_rollouts_count"
	isFavorite           = "is_favorite"
	rolloutStatus        = "rollout_status"
)

const (
	resourceName  = "metric"
	resourcesName = "metrics"
)

var (
	eventTypes    = []string{"CLICK", "CUSTOM", "PAGE_VIEW"}
	metricTypes   = []string{"BINARY", "COUNT", "NUMERIC"}
	directions    = []string{"HIGHER", "LOWER"}
	includeValues = []string{"experiments", "guarded_rollouts"}
	sortValues    = []string{utils.CreatedTime, utils.MetricID, "metric_name", utils.UpdatedTime}
)

// getFields returns the fields of a metric, as the user writes them (Input) or as the API returns them (Output).
// The favorite flag is not part of the metric itself, the API sets it with a separate call.
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.MetricID: mode.RequiredForceNew(utils.NewString(
				"Metric ID. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.MetricName: mode.Required(utils.NewString(
				"Metric name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			utils.Description: mode.Optional(utils.NewString("Metric description.")),
			eventType:         mode.RequiredForceNew(utils.NewEnum("The type of event.", eventTypes...)),
			utils.MetricType:  mode.RequiredForceNew(utils.NewEnum("The type of metric.", metricTypes...)),
			utils.EventKey:    mode.Required(utils.NewString("The event key used to record metric data.")),
			utils.Direction: mode.Required(utils.NewEnum(
				"Whether a higher or lower metric value indicates improvement.", directions...,
			)),
			isFavorite: mode.OptionalComputed(utils.NewBool("Whether the metric is marked as favorite.")),
			utils.LastEventReceivedTime: utils.AsComputed(utils.NewString(
				"The time the most recent event was received for this metric. Set only when an event was recorded.",
			)),
			experiments: utils.AsComputed(utils.NewList(
				"List of experiments that use this metric. Set only when `include` has `experiments`.",
				map[string]*schema.Schema{
					experimentId:     utils.AsComputed(utils.NewString("The unique identifier of the experiment.")),
					utils.Name:       utils.AsComputed(utils.NewString("The name of the experiment.")),
					experimentStatus: utils.AsComputed(utils.NewString("The status of the experiment.")),
				},
			)),
			experimentsCount: utils.AsComputed(utils.NewInt("Total count of experiments that use this metric.")),
			guardedRollouts: utils.AsComputed(utils.NewList(
				"List of guarded rollouts that use this metric. Set only when `include` has `guarded_rollouts`.",
				map[string]*schema.Schema{
					utils.RolloutID:     utils.AsComputed(utils.NewString("The unique identifier of the guarded rollout.")),
					rolloutStatus:       utils.AsComputed(utils.NewString("The current status of the guarded rollout.")),
					utils.FeatureID:     utils.AsComputed(utils.NewString("The unique identifier of the feature flag.")),
					featureName:         utils.AsComputed(utils.NewString("The name of the feature flag.")),
					utils.EnvironmentID: utils.AsComputed(utils.NewString("The unique identifier of the environment.")),
					environmentName:     utils.AsComputed(utils.NewString("The name of the environment.")),
				},
			)),
			guardedRolloutsCount: utils.AsComputed(utils.NewInt("Total count of guarded rollouts that use this metric.")),
		},
	)
}

var idKeys = []string{utils.MetricID}

// DataSourceIBMAppConfigMetric returns the data source of one metric.
func DataSourceIBMAppConfigMetric() *schema.Resource {
	return utils.NewDataSource("Metric", idKeys, getFields(utils.Output), readMetric,
		map[string]*schema.Schema{
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
		},
	)
}

// DataSourceIBMAppConfigMetrics returns the data source that lists the metrics.
func DataSourceIBMAppConfigMetrics() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of metrics.", getFields(utils.Output), readMetrics,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Expand:  utils.GetExpandSchema(),
			utils.Include: utils.GetIncludeValuesSchema(includeValues...),
			utils.Sort:    utils.GetSortValuesSchema(resourceName, sortValues...),
			utils.Search:  utils.GetSearchSchema("Name"),
		},
	)
}

// ResourceIBMAppConfigMetric returns the metric resource.
func ResourceIBMAppConfigMetric() *schema.Resource {
	return utils.NewResource("Metric", idKeys, getFields(utils.Input), utils.Calls{
		Create: createMetric,
		Get:    getMetric,
		Update: updateMetric,
		Delete: deleteMetric,
	})
}
