// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package workflow

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readWorkflowConfig(c *utils.Call) (any, error) {
	result, _, err := c.Client.GetWorkflowConfigWithContext(c.Ctx, c.Client.NewGetWorkflowConfigOptions(c.IDs[0]))
	if err != nil {
		return nil, err
	}
	return renameProvider(result)
}

func readWorkflowConfigs(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListWorkflowConfigsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	fields, err := utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListWorkflowConfigsWithContext(c.Ctx, options)
		return page, err
	})
	if err != nil {
		return nil, err
	}
	items, _ := fields[resourcesName].([]any)
	for _, item := range items {
		renameKey(item.(map[string]any))
	}
	return fields, nil
}
