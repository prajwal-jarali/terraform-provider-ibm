// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package snapshot

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readSnapshot(c *utils.Call) (any, error) {
	result, _, err := c.Client.GetGitconfigWithContext(c.Ctx, c.Client.NewGetGitconfigOptions(c.IDs[0]))
	return result, err
}

func readSnapshots(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListGitconfigsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListGitconfigsWithContext(c.Ctx, options)
		return page, err
	})
}
