// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package integration

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createIntegration(c *utils.Call) (any, error) {
	// The API models the metadata as an interface, so it is filled apart from the other fields.
	metadata := &appconfigurationv1.CreateIntegrationMetadata{}
	if err := utils.DecodeField(c.Data, c.Schema, utils.Metadata, metadata); err != nil {
		return nil, err
	}
	options := c.Client.NewCreateIntegrationOptions(c.IDs[0], c.Text(integrationType), metadata)
	result, _, err := c.Client.CreateIntegrationWithContext(c.Ctx, options)
	if err != nil {
		return nil, err
	}
	return keepSourceFields(c, result)
}

func getIntegration(c *utils.Call) (any, *core.DetailedResponse, error) {
	result, response, err := c.Client.GetIntegrationWithContext(c.Ctx, c.Client.NewGetIntegrationOptions(c.IDs[0]))
	if err != nil {
		return nil, response, err
	}
	fields, err := keepSourceFields(c, result)
	return fields, response, err
}

func deleteIntegration(c *utils.Call) (*core.DetailedResponse, error) {
	return c.Client.DeleteIntegrationWithContext(c.Ctx, c.Client.NewDeleteIntegrationOptions(c.IDs[0]))
}

// keepSourceFields adds the source name and description the user wrote to the metadata of the result.
// The API takes them on creation but does not return them, and a changed value would replace the integration.
func keepSourceFields(c *utils.Call, result any) (map[string]any, error) {
	fields, err := utils.FieldMap(result)
	if err != nil {
		return nil, err
	}
	configuredMetadata, _ := c.Data.Get(utils.Metadata).([]any)
	metadata, _ := fields[utils.Metadata].(map[string]any)
	if len(configuredMetadata) == 0 || metadata == nil {
		return fields, nil
	}
	source, _ := configuredMetadata[0].(map[string]any)
	for _, key := range []string{eventNotificationsSourceName, eventNotificationsSourceDescription} {
		if value, _ := source[key].(string); value != "" {
			metadata[key] = value
		}
	}
	return fields, nil
}
