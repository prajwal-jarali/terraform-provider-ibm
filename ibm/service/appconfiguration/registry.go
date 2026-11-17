// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/collection"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/environment"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/feature"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/guardedrollout"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/integration"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/metric"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/originconfig"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/property"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/rule"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/segment"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/snapshot"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/workflow"
)

// The constructors of the resources and data sources of the sub-packages, so that the provider
// imports this package only. A new resource gets one line here and one in provider.go.
var (
	DataSourceIBMAppConfigCollection  = collection.DataSourceIBMAppConfigCollection
	DataSourceIBMAppConfigCollections = collection.DataSourceIBMAppConfigCollections
	ResourceIBMAppConfigCollection    = collection.ResourceIBMAppConfigCollection

	DataSourceIBMAppConfigEnvironment  = environment.DataSourceIBMAppConfigEnvironment
	DataSourceIBMAppConfigEnvironments = environment.DataSourceIBMAppConfigEnvironments
	ResourceIBMAppConfigEnvironment    = environment.ResourceIBMAppConfigEnvironment

	DataSourceIBMAppConfigFeature  = feature.DataSourceIBMAppConfigFeature
	DataSourceIBMAppConfigFeatures = feature.DataSourceIBMAppConfigFeatures
	ResourceIBMAppConfigFeature    = feature.ResourceIBMAppConfigFeature

	DataSourceIBMAppConfigGuardedRollout       = guardedrollout.DataSourceIBMAppConfigGuardedRollout
	DataSourceIBMAppConfigGuardedRollouts      = guardedrollout.DataSourceIBMAppConfigGuardedRollouts
	DataSourceIBMAppConfigGuardedRolloutMetric = guardedrollout.DataSourceIBMAppConfigGuardedRolloutMetric

	DataSourceIBMAppConfigIntegration  = integration.DataSourceIBMAppConfigIntegration
	DataSourceIBMAppConfigIntegrations = integration.DataSourceIBMAppConfigIntegrations
	ResourceIBMAppConfigIntegration    = integration.ResourceIBMAppConfigIntegration

	DataSourceIBMAppConfigMetric  = metric.DataSourceIBMAppConfigMetric
	DataSourceIBMAppConfigMetrics = metric.DataSourceIBMAppConfigMetrics
	ResourceIBMAppConfigMetric    = metric.ResourceIBMAppConfigMetric

	DataSourceIBMAppConfigOriginConfig = originconfig.DataSourceIBMAppConfigOriginConfig
	ResourceIBMAppConfigOriginConfig   = originconfig.ResourceIBMAppConfigOriginConfig

	DataSourceIBMAppConfigProperty   = property.DataSourceIBMAppConfigProperty
	DataSourceIBMAppConfigProperties = property.DataSourceIBMAppConfigProperties
	ResourceIBMAppConfigProperty     = property.ResourceIBMAppConfigProperty

	DataSourceIBMAppConfigRule  = rule.DataSourceIBMAppConfigRule
	DataSourceIBMAppConfigRules = rule.DataSourceIBMAppConfigRules
	ResourceIBMAppConfigRule    = rule.ResourceIBMAppConfigRule

	DataSourceIBMAppConfigSegment  = segment.DataSourceIBMAppConfigSegment
	DataSourceIBMAppConfigSegments = segment.DataSourceIBMAppConfigSegments
	ResourceIBMAppConfigSegment    = segment.ResourceIBMAppConfigSegment

	DataSourceIBMAppConfigSnapshot  = snapshot.DataSourceIBMAppConfigSnapshot
	DataSourceIBMAppConfigSnapshots = snapshot.DataSourceIBMAppConfigSnapshots
	ResourceIBMAppConfigSnapshot    = snapshot.ResourceIBMAppConfigSnapshot

	DataSourceIBMAppConfigWorkflowConfig  = workflow.DataSourceIBMAppConfigWorkflowConfig
	DataSourceIBMAppConfigWorkflowConfigs = workflow.DataSourceIBMAppConfigWorkflowConfigs
	ResourceIBMAppConfigWorkflowConfig    = workflow.ResourceIBMAppConfigWorkflowConfig
)
