// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package environment

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readEnvironment(c *utils.Call) (any, error) {
	options := c.Client.NewGetEnvironmentOptions(c.IDs[0])
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.GetEnvironmentWithContext(c.Ctx, options)
	return result, err
}

func readEnvironments(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListEnvironmentsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListEnvironmentsWithContext(c.Ctx, options)
		return page, err
	})
}
