// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package guardedrollout

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readGuardedRollout(c *utils.Call) (any, error) {
	options := c.Client.NewGetGuardedRolloutOptions(c.IDs[0], c.IDs[1], c.IDs[2])
	result, _, err := c.Client.GetGuardedRolloutWithContext(c.Ctx, options)
	return result, err
}

func readGuardedRolloutMetric(c *utils.Call) (any, error) {
	options := c.Client.NewGetGuardedRolloutMetricOptions(c.IDs[0], c.IDs[1], c.IDs[2], c.IDs[3])
	result, _, err := c.Client.GetGuardedRolloutMetricWithContext(c.Ctx, options)
	return result, err
}

func readGuardedRollouts(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListGuardedRolloutsOptions(c.Text(utils.EnvironmentID), c.Text(utils.FeatureID))
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListGuardedRolloutsWithContext(c.Ctx, options)
		return page, err
	})
}
