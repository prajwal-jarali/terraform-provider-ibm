// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package collection

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readCollection(c *utils.Call) (any, error) {
	options := c.Client.NewGetCollectionOptions(c.IDs[0])
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.GetCollectionWithContext(c.Ctx, options)
	return result, err
}

func readCollections(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListCollectionsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListCollectionsWithContext(c.Ctx, options)
		return page, err
	})
}
