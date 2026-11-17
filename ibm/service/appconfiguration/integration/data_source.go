// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package integration

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readIntegration(c *utils.Call) (any, error) {
	result, _, err := c.Client.GetIntegrationWithContext(c.Ctx, c.Client.NewGetIntegrationOptions(c.IDs[0]))
	return result, err
}

func readIntegrations(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListIntegrationsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListIntegrationsWithContext(c.Ctx, options)
		return page, err
	})
}
