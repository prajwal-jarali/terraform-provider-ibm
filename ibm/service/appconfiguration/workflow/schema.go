// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package workflow holds the schemas of the approval workflow config resource and data sources.
// A workflow config makes changes to the resources in its scope wait for the approval of a ServiceNow change request.
package workflow

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	approvalExpiration = "approval_expiration"
	approvalGroupName  = "approval_group_name"
	crnMask            = "crn_mask"
	defaultEmail       = "default_email"
	enable             = "enable"
	environments       = "environments"
	resources          = "resources"
	scope              = "scope"
	serviceEnvironment = "service_environment"
	smSecretCrn        = "sm_secret_crn"
	smSecretId         = "sm_secret_id"
	workflowUrl        = "workflow_url"
)

const (
	resourceName  = "workflow config"
	resourcesName = "workflow_configs"

	scopeMode      = "ALL"
	providerConfig = "provider_config"
	apiProvider    = "provider"
)

var (
	providerTypes       = []string{"SERVICENOW_IBM", "SERVICENOW_EXTERNAL"}
	serviceEnvironments = []string{"pre_prod", "production"}
	sortValues          = []string{utils.CreatedTime, utils.Name, utils.UpdatedTime, utils.WorkflowID}
)

// getProviderMetadata returns the settings of the ServiceNow provider. `SERVICENOW_IBM` uses `crn_mask` and
// `sm_secret_crn`. `SERVICENOW_EXTERNAL` uses `workflow_url`, `approval_group_name` and `sm_secret_id`.
func getProviderMetadata(mode utils.Mode) map[string]*schema.Schema {
	text := func(description string) *schema.Schema { return mode.Optional(utils.NewString(description)) }
	return map[string]*schema.Schema{
		approvalExpiration: mode.Optional(utils.NewInt("Approval expiration time in days.")),
		crnMask:            text("CRN mask for the service."),
		smSecretCrn:        text("Secret Manager secret CRN for the ServiceNow credentials."),
		defaultEmail:       text("The email id to which the change request is assigned."),
		serviceEnvironment: mode.Optional(utils.NewEnum(
			"The ServiceNow change request environment context. If not provided, it defaults to production.",
			serviceEnvironments...,
		)),
		workflowUrl:       text("ServiceNow instance URL."),
		approvalGroupName: text("Name of the approval group in ServiceNow."),
		smSecretId:        text("Secret Manager secret ID for the ServiceNow credentials."),
	}
}

// getScopeMode returns a scope that covers all resources of a kind. Only the ALL mode is supported.
func getScopeMode(mode utils.Mode, kind string) *schema.Schema {
	return mode.Optional(utils.NewObject(
		"Scope of the "+kind+". Only the ALL mode is supported.",
		map[string]*schema.Schema{
			utils.ModeKey: mode.Required(utils.NewEnum("Scope mode.", scopeMode)),
			utils.Ids:     mode.Optional(utils.NewStringList("List of ids. Always empty when the mode is ALL.")),
		},
	))
}

func getScope(mode utils.Mode) *schema.Schema {
	return mode.Required(utils.NewObject(
		"Scope of the workflow. At least one of collections, segments or environments must be specified.",
		map[string]*schema.Schema{
			utils.Collections: getScopeMode(mode, "collections"),
			utils.Segments:    getScopeMode(mode, "segments"),
			environments: mode.Optional(utils.NewList(
				"List of environments where the workflow applies.", getEnvironmentScopeFields(mode),
			)),
		},
	))
}

func getEnvironmentScopeFields(mode utils.Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.EnvironmentID: mode.Required(utils.NewString("Environment identifier.")),
		resources: mode.Optional(utils.NewObject(
			"Resources of the environment that the workflow applies to.", getResourcesScopeFields(mode),
		)),
	}
}

func getResourcesScopeFields(mode utils.Mode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		utils.Environment: mode.Optional(utils.NewObject("The environment itself.", map[string]*schema.Schema{
			enable: mode.Optional(utils.NewBool("Whether changes to the environment need approval.")),
		})),
		utils.Features:   getScopeMode(mode, "features"),
		utils.Properties: getScopeMode(mode, "properties"),
	}
}

// getProviderConfig returns the provider configuration. The API calls it `provider`, which Terraform
// reserves, so it is named `provider_config`.
func getProviderConfig(mode utils.Mode) *schema.Schema {
	return mode.Required(utils.NewObject("Provider configuration for the workflow.", map[string]*schema.Schema{
		utils.Type: mode.Required(utils.NewEnum("Provider type.", providerTypes...)),
		utils.Metadata: mode.Required(utils.NewObject(
			"Metadata of the ServiceNow provider. Which settings apply depends on the provider type.",
			getProviderMetadata(mode),
		)),
	}))
}

func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.Name:       mode.Required(utils.NewString("Name of the workflow configuration.")),
			utils.WorkflowID: mode.RequiredForceNew(utils.NewString("Unique identifier for the workflow configuration.")),
			utils.Enabled:    mode.Required(utils.NewBool("Whether the workflow is enabled.")),
			providerConfig:   getProviderConfig(mode),
			scope:            getScope(mode),
		},
	)
}

var idKeys = []string{utils.WorkflowID}

// DataSourceIBMAppConfigWorkflowConfig returns the data source of one workflow config.
func DataSourceIBMAppConfigWorkflowConfig() *schema.Resource {
	return utils.NewDataSource("Workflow config", idKeys, getFields(utils.Output), readWorkflowConfig)
}

// DataSourceIBMAppConfigWorkflowConfigs returns the data source that lists the workflow configs.
func DataSourceIBMAppConfigWorkflowConfigs() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of workflow configs.", getFields(utils.Output), readWorkflowConfigs,
		utils.GetPageQuery(),
		map[string]*schema.Schema{
			utils.Sort:   utils.GetSortValuesSchema(resourceName, sortValues...),
			utils.Search: utils.GetSearchSchema("Name"),
		},
	)
}

// ResourceIBMAppConfigWorkflowConfig returns the workflow config resource.
func ResourceIBMAppConfigWorkflowConfig() *schema.Resource {
	return utils.NewResource("Workflow config", idKeys, getFields(utils.Input), utils.Calls{
		Create: createWorkflowConfig,
		Get:    getWorkflowConfig,
		Update: updateWorkflowConfig,
		Delete: deleteWorkflowConfig,
	})
}
