// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package workflow

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

// provider is the `provider_config` the user wrote. The API calls it `provider`, and models its
// metadata as an interface, so it is filled apart from the other fields.
type provider struct {
	Type     *string                              `json:"type"`
	Metadata *appconfigurationv1.WorkflowMetadata `json:"metadata"`
}

func createWorkflowConfig(c *utils.Call) (any, error) {
	options := &appconfigurationv1.CreateWorkflowConfigsOptions{}
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	var err error
	if options.Provider, err = getProvider(c); err != nil {
		return nil, err
	}
	fillIds(options.Scope)
	result, _, err := c.Client.CreateWorkflowConfigsWithContext(c.Ctx, options)
	if err != nil {
		return nil, err
	}
	return renameProvider(result)
}

func getWorkflowConfig(c *utils.Call) (any, *core.DetailedResponse, error) {
	result, response, err := c.Client.GetWorkflowConfigWithContext(c.Ctx, c.Client.NewGetWorkflowConfigOptions(c.IDs[0]))
	if err != nil {
		return nil, response, err
	}
	fields, err := renameProvider(result)
	return fields, response, err
}

// updateWorkflowConfig sends every field: the API replaces the whole workflow config.
func updateWorkflowConfig(c *utils.Call) (any, error) {
	options := &appconfigurationv1.UpdateWorkflowConfigsOptions{WorkflowConfigID: &c.IDs[0]}
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	var err error
	if options.Provider, err = getProvider(c); err != nil {
		return nil, err
	}
	fillIds(options.Scope)
	result, _, err := c.Client.UpdateWorkflowConfigsWithContext(c.Ctx, options)
	if err != nil {
		return nil, err
	}
	return renameProvider(result)
}

func deleteWorkflowConfig(c *utils.Call) (*core.DetailedResponse, error) {
	return c.Client.DeleteWorkflowConfigsWithContext(c.Ctx, c.Client.NewDeleteWorkflowConfigsOptions(c.IDs[0]))
}

// fillIds sets the empty list of ids that the API requires on each scope mode, which the user leaves out for ALL.
func fillIds(scope *appconfigurationv1.WorkflowScope) {
	if scope == nil {
		return
	}
	modes := []*appconfigurationv1.WorkflowScopeMode{scope.Collections, scope.Segments}
	for _, environment := range scope.Environments {
		if environment.Resources != nil {
			modes = append(modes, environment.Resources.Features, environment.Resources.Properties)
		}
	}
	for _, mode := range modes {
		if mode != nil && mode.Ids == nil {
			mode.Ids = []string{}
		}
	}
}

func getProvider(c *utils.Call) (*appconfigurationv1.WorkflowProvider, error) {
	config := provider{}
	if err := utils.DecodeField(c.Data, c.Schema, providerConfig, &config); err != nil {
		return nil, err
	}
	return &appconfigurationv1.WorkflowProvider{Type: config.Type, Metadata: config.Metadata}, nil
}

// renameProvider turns a result into fields whose `provider` is named like the schema field.
func renameProvider(result any) (map[string]any, error) {
	fields, err := utils.FieldMap(result)
	if err != nil {
		return nil, err
	}
	renameKey(fields)
	return fields, nil
}

func renameKey(fields map[string]any) {
	fields[providerConfig] = fields[apiProvider]
	delete(fields, apiProvider)
}
