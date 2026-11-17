// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	durationPreset = "duration_preset"
)

// Schema parts that features, properties and feature rules have in common:
// value formats, collections, targeting rules and rollouts.

// GetFormatSchema returns the format of a STRING value.
func GetFormatSchema(resource string) *schema.Schema {
	return NewEnum(fmt.Sprintf(
		"Format of the %s. Required when `type` is `STRING`, not used for the other types.", resource,
	), "TEXT", "JSON", "YAML")
}

// GetCollectionsSchema returns the collections a feature or property belongs to.
func GetCollectionsSchema(mode Mode, resource string) *schema.Schema {
	fields := map[string]*schema.Schema{
		CollectionID: mode.Required(NewString("Collection Id.")),
		Name:         AsComputed(NewString("Collection name.")),
	}
	if mode == Input {
		fields[Deleted] = AsOptional(NewBool("Remove the association with the collection when set to `true`."))
	}
	return mode.Optional(NewList(
		fmt.Sprintf("List of collections associated with the %s.", resource), fields,
	))
}

// GetRuleFields returns the fields of a targeting rule: the segments it targets and the value served.
func GetRuleFields(mode Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		Rules: mode.Required(NewList("The list of targeted segments.", map[string]*schema.Schema{
			Segments: mode.Required(NewStringList("The list of segment ids that are used for targeting using the rule.")),
		})),
		Value: mode.Required(NewString(
			"Value to be used for evaluation for this rule. The value must match the `type` and `format` attributes.",
		)),
	}
}

// GetFeatureRuleFields returns the fields of a targeting rule of a feature flag: the common rule
// fields, an id and a name, and how the rule is rolled out.
func GetFeatureRuleFields(mode Mode) map[string]*schema.Schema {
	return Merge(GetRuleFields(mode), GetRolloutFields(mode), map[string]*schema.Schema{
		RuleID:   mode.OptionalComputed(NewString("Rule id is an identifier used for identifying the rule.")),
		RuleName: mode.Optional(NewString("Name assigned to a particular rule.")),
	})
}

// GetSegmentRulesSchema returns the targeting rules of a feature or property. `rule` holds the fields of
// one rule, and the evaluation order is added to them.
func GetSegmentRulesSchema(mode Mode, resource string, rule map[string]*schema.Schema) *schema.Schema {
	return mode.OptionalComputed(NewList(
		fmt.Sprintf("Targeting rules that set different %s values for different segments.", resource),
		GetOrderedRuleFields(mode, rule),
	))
}

// GetOrderedRuleFields adds the evaluation order to the fields of a targeting rule.
func GetOrderedRuleFields(mode Mode, rule map[string]*schema.Schema) map[string]*schema.Schema {
	return Merge(rule, map[string]*schema.Schema{
		Order: mode.Required(NewInt(
			"Order of the rule, used during evaluation. The evaluation is performed in the order defined " +
				"and the value associated with the first matching rule is used for evaluation. " +
				"Index 1 based numbering is used for order.",
		)),
	})
}

// GetRolloutFields returns how a value is rolled out: all at once, in phases, or in phases that
// are watched by metrics.
func GetRolloutFields(mode Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		RolloutPercentage: mode.OptionalComputed(NewInt(
			"Rollout percentage. Supported only for Lite and Enterprise plans.",
		)),
		RolloutType: mode.OptionalComputed(NewEnum(
			"The rollout strategy type. `MANUAL` is the default. `PROGRESSIVE` enables automatic phase-based "+
				"rollout. `GUARDED` enables automatic phase-based rollout that is monitored against the configured metrics.",
			"MANUAL", "PROGRESSIVE", "GUARDED",
		)),
		RolloutConfiguration: mode.Optional(NewObject(
			"Configuration that controls the rollout behaviour for a `PROGRESSIVE` or `GUARDED` rollout type.",
			getRolloutConfigurationFields(mode),
		)),
	}
}

func getRolloutConfigurationFields(mode Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		durationPreset: mode.Required(NewString(
			"A predefined duration preset that sets the overall pace of the rollout. Use `CUSTOM` to define " +
				"phases manually with explicit `duration` and `duration_type` values on each phase.",
		)),
		StartAt: mode.Optional(NewString("The UTC date and time at which the rollout should start, in ISO 8601 format.")),
		Phases: mode.Required(NewList(
			"The ordered list of rollout phases. Each phase defines a target percentage and a wait duration "+
				"before the next phase begins.",
			getPhaseFields(mode),
		)),
		Status: AsComputed(NewString("The current execution status of the rollout: `QUEUED`, `RUNNING` or `STOPPED`.")),
		Metrics: mode.Optional(NewList(
			"The metrics monitored during the rollout. Applicable only when `rollout_type` is `GUARDED`. "+
				"Each `metric_id` must be unique.",
			getRolloutMetricFields(mode),
		)),
	}
}

func getPhaseFields(mode Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		Percentage: mode.Required(NewInt("The rollout percentage target for this phase, from 0 to 100.")),
		Duration:   mode.Optional(NewInt("The length of time to wait before advancing to the next phase.")),
		DurationType: mode.Optional(NewEnum(
			"The unit of time for the `duration` field.", "days", "hours", "minutes",
		)),
	}
}

func getRolloutMetricFields(mode Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		MetricID: mode.Required(NewString("The unique identifier of the metric.")),
		AutoRollback: mode.Optional(NewBool(
			"Whether a regression in this metric triggers an automatic rollback.",
		)),
		MetricName: AsComputed(NewString("The display name of the metric.")),
		Direction:  AsComputed(NewString("Whether a higher or lower metric value indicates improvement.")),
		MetricType: AsComputed(NewString("The type of metric.")),
	}
}
