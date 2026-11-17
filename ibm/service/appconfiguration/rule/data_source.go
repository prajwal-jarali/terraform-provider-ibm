// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package rule

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readRule(c *utils.Call) (any, error) {
	result, _, err := c.Client.GetFeatureRuleWithContext(c.Ctx, c.Client.NewGetFeatureRuleOptions(c.IDs[0], c.IDs[1], c.IDs[2]))
	return result, err
}

func readRules(c *utils.Call) (map[string]any, error) {
	return utils.ListAll(c, resourcesName, nil, func() (any, error) {
		page, _, err := c.Client.ListFeatureRulesWithContext(c.Ctx, c.Client.NewListFeatureRulesOptions(
			c.Text(utils.EnvironmentID), c.Text(utils.FeatureID),
		))
		return page, err
	})
}
