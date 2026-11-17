// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package feature

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readFeature(c *utils.Call) (any, error) {
	options := c.Client.NewGetFeatureOptions(c.IDs[0], c.IDs[1])
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.GetFeatureWithContext(c.Ctx, options)
	return result, err
}

func readFeatures(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListFeaturesOptions(c.Text(utils.EnvironmentID))
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListFeaturesWithContext(c.Ctx, options)
		return page, err
	})
}
