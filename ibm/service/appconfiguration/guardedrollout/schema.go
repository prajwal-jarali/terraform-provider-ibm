// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package guardedrollout holds the schemas of the guarded rollout data sources.
// A guarded rollout is started by setting the rollout type of a feature flag or of one of its rules to
// GUARDED, so the API only reads it: there is no resource.
package guardedrollout

import (
	"slices"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	conversions                 = "conversions"
	currentPhaseIndex           = "current_phase_index"
	difference                  = "difference"
	endTime                     = "end_time"
	estimate                    = "estimate"
	exposures                   = "exposures"
	extended                    = "extended"
	featureSnapshot             = "feature_snapshot"
	lastRegressionDetectedTime  = "last_regression_detected_time"
	lastRegressionDismissedTime = "last_regression_dismissed_time"
	lower                       = "lower"
	original                    = "original"
	phaseIndex                  = "phase_index"
	result                      = "result"
	startTime                   = "start_time"
	target                      = "target"
	timestamp                   = "timestamp"
	upper                       = "upper"
)

const resourcesName = "guarded_rollouts"

func text(description string) *schema.Schema { return utils.AsComputed(utils.NewString(description)) }

func number(description string) *schema.Schema { return utils.AsComputed(utils.NewInt(description)) }

// getSummaryFields returns what the list reports for each guarded rollout.
func getSummaryFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.RolloutID: text("The unique identifier of the guarded rollout."),
		utils.RuleID:    text("The rule associated with this guarded rollout. Not set if the rollout applies to the feature."),
		startTime:       text("The date and time the guarded rollout was started."),
		endTime:         text("The date and time the guarded rollout completed. Not set while it is in progress."),
		utils.Status: text(
			"The current status of the guarded rollout: COMPLETED, IN_PROGRESS, PAUSED, ROLLED_BACK or STOPPED.",
		),
	}
}

// getMetricFields returns what a guarded rollout reports about a metric it monitors.
func getMetricFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.MetricID:     text("The unique identifier of the metric."),
		utils.MetricName:   text("The display name of the metric."),
		utils.Description:  text("A description of what the metric tracks."),
		utils.EventKey:     text("The event key used to record metric data."),
		utils.MetricType:   text("The type of metric."),
		utils.Direction:    text("Whether a higher or lower metric value indicates improvement."),
		utils.AutoRollback: utils.AsComputed(utils.NewBool("Whether a regression in this metric triggers an automatic rollback.")),
		utils.Status: text(
			"The status of the metric: HEALTHY, INCONCLUSIVE_HEALTHY, INCONCLUSIVE_WORSE, INSUFFICIENT_DATA, NO_EVENTS or REGRESSED.",
		),
		utils.SubStatus: text(
			"The state behind the status: PAUSE_AUTO, PAUSE_CONTINUE, PAUSE_FINAL_STAGE, " +
				"PAUSE_SAMPLE_COLLECTION, PHASE_EXTEND or ROLLBACK.",
		),
		lastRegressionDetectedTime:  text("The time a regression was last detected for this metric."),
		lastRegressionDismissedTime: text("The time a regression was last dismissed for this metric."),
		utils.LastEventReceivedTime: text("The time the most recent event was received for this metric."),
	}
}

// getDetailFields returns what the API reports for one guarded rollout.
func getDetailFields() map[string]*schema.Schema {
	return utils.Merge(getSummaryFields(), map[string]*schema.Schema{
		utils.Description: text("A human-readable description of the rollout outcome."),
		utils.Rule: utils.AsComputed(utils.NewObject(
			"The targeting rule that is rolled out.",
			utils.GetOrderedRuleFields(utils.Output, utils.GetFeatureRuleFields(utils.Output)),
		)),
		currentPhaseIndex: number("The index of the current or last active phase."),
		utils.Phases: utils.AsComputed(utils.NewList("Ordered list of rollout phases.", map[string]*schema.Schema{
			phaseIndex:         number("Zero-based index of the phase."),
			utils.Percentage:   number("The rollout percentage for this phase."),
			utils.Duration:     number("The duration this phase runs before advancing. Not set for the final phase."),
			utils.DurationType: text("The unit of the duration value: days, hours or minutes."),
			extended:           utils.AsComputed(utils.NewBool("Whether this phase has been extended beyond its original duration.")),
		})),
		featureSnapshot: utils.AsComputed(utils.NewObject(
			"Snapshot of the feature flag values at the time of the guarded rollout.",
			map[string]*schema.Schema{
				utils.Collections:   utils.GetCollectionsSchema(utils.Output, "feature"),
				utils.EnabledValue:  text("The value served when the feature flag is enabled."),
				utils.DisabledValue: text("The value served when the feature flag is disabled."),
			},
		)),
		utils.Metrics: utils.AsComputed(utils.NewList("Metrics monitored during the guarded rollout.", getMetricFields())),
	})
}

// getEstimate returns the statistics of one variant of a metric: the original, the target or their difference.
func getEstimate(description string) *schema.Schema {
	return utils.AsComputed(utils.NewObject(description, map[string]*schema.Schema{
		estimate:    utils.AsComputed(utils.NewFloat("The point estimate.")),
		lower:       utils.AsComputed(utils.NewFloat("The lower bound of the confidence interval.")),
		upper:       utils.AsComputed(utils.NewFloat("The upper bound of the confidence interval.")),
		exposures:   number("The number of exposures recorded. Set on original and target only."),
		conversions: number("The number of conversions recorded. Set on original and target only."),
	}))
}

func getFeature() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.EnvironmentID: utils.GetParentIDSchema("Environment Id."),
		utils.FeatureID:     utils.GetParentIDSchema("Feature Id."),
	}
}

var (
	rolloutKeys = []string{utils.EnvironmentID, utils.FeatureID, utils.RolloutID}
	metricKeys  = append(slices.Clone(rolloutKeys), utils.MetricID)
)

// DataSourceIBMAppConfigGuardedRollout returns the data source of one guarded rollout.
func DataSourceIBMAppConfigGuardedRollout() *schema.Resource {
	return utils.NewDataSource("Guarded rollout", rolloutKeys, getDetailFields(), readGuardedRollout, getFeature())
}

// DataSourceIBMAppConfigGuardedRollouts returns the data source that lists the guarded rollouts of a feature.
func DataSourceIBMAppConfigGuardedRollouts() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of guarded rollouts.", getSummaryFields(), readGuardedRollouts,
		getFeature(),
		utils.GetPageQuery(),
	)
}

// DataSourceIBMAppConfigGuardedRolloutMetric returns the schema of the data source of the statistics that
// one metric collected during a guarded rollout.
func DataSourceIBMAppConfigGuardedRolloutMetric() *schema.Resource {
	return utils.NewDataSource("Guarded rollout metric", metricKeys, getMetricResultFields(), readGuardedRolloutMetric, getFeature())
}

// getMetricResultFields returns the statistics that one metric collected during a guarded rollout.
func getMetricResultFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.RolloutID: text("The unique identifier of the guarded rollout."),
		utils.MetricID:  text("The unique identifier of the metric."),
		utils.Status:    text("The status of the guarded rollout metric."),
		utils.SubStatus: text("The state behind the status."),
		result: utils.AsComputed(utils.NewList("Time-series statistical results for this metric.", map[string]*schema.Schema{
			timestamp:  text("The time at which this result was computed."),
			difference: getEstimate("Statistics of the difference between the target and the original."),
			original:   getEstimate("Statistics of the original variant."),
			target:     getEstimate("Statistics of the target variant."),
		})),
	}
}
