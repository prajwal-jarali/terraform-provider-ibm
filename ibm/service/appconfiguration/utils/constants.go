// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

// Schema field names shared by the App Configuration resources and data sources.
const (
	GUID        = "guid"
	Name        = "name"
	Description = "description"
	Tags        = "tags"
	Href        = "href"
	Type        = "type"
	Format      = "format"
	Enabled     = "enabled"
	Value       = "value"
	Order       = "order"
	Rules       = "rules"

	CreatedTime = "created_time"
	UpdatedTime = "updated_time"

	EnvironmentID = "environment_id"
	CollectionID  = "collection_id"

	// Associated or filtered resources.
	Collections  = "collections"
	Segments     = "segments"
	Features     = "features"
	Properties   = "properties"
	Snapshots    = "snapshots"
	SegmentRules = "segment_rules"

	WorkflowApproval = "workflow_approval"

	// Query parameters of list and get calls.
	Expand  = "expand"
	Include = "include"
	Sort    = "sort"
	Search  = "search"
	Limit   = "limit"
	Offset  = "offset"

	// Pagination fields of list responses.
	TotalCount = "total_count"
	First      = "first"
	Previous   = "previous"
	Next       = "next"
	Last       = "last"
)

// Identifiers of the App Configuration resources.
const (
	FeatureID     = "feature_id"
	PropertyID    = "property_id"
	SegmentID     = "segment_id"
	GitConfigID   = "git_config_id"
	RuleID        = "rule_id"
	RolloutID     = "rollout_id"
	MetricID      = "metric_id"
	IntegrationID = "integration_id"
	WorkflowID    = "workflow_id"
)

// Field names that more than one resource uses.
const (
	AutoRollback          = "auto_rollback"
	Collection            = "collection"
	ColorCode             = "color_code"
	Deleted               = "deleted"
	Direction             = "direction"
	DisabledValue         = "disabled_value"
	Duration              = "duration"
	DurationType          = "duration_type"
	EnabledValue          = "enabled_value"
	Environment           = "environment"
	EvaluationTime        = "evaluation_time"
	EventKey              = "event_key"
	Ids                   = "ids"
	LastEventReceivedTime = "last_event_received_time"
	Metadata              = "metadata"
	MetricName            = "metric_name"
	MetricType            = "metric_type"
	Metrics               = "metrics"
	ModeKey               = "mode"
	Percentage            = "percentage"
	Phases                = "phases"
	RolloutConfiguration  = "rollout_configuration"
	RolloutPercentage     = "rollout_percentage"
	RolloutType           = "rollout_type"
	Rule                  = "rule"
	RuleName              = "rule_name"
	SegmentExists         = "segment_exists"
	StartAt               = "start_at"
	Status                = "status"
	SubStatus             = "sub_status"
)
