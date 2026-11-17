// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package metric

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createMetric(c *utils.Call) (any, error) {
	options := c.Client.NewCreateMetricOptions(
		c.IDs[0], c.Text(utils.MetricName), c.Text(utils.MetricType),
		c.Text(eventType), c.Text(utils.EventKey), c.Text(utils.Direction),
	)
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateMetricWithContext(c.Ctx, options)
	if err != nil || !utils.ForCreate(c.Data, isFavorite) {
		return result, err
	}
	// The API creates a metric that is not a favorite, so the flag is set with its own call.
	return setFavorite(c)
}

func getMetric(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetMetricWithContext(c.Ctx, c.Client.NewGetMetricOptions(c.IDs[0]))
}

// updateMetric sends the changed fields, and the favorite flag with its own call.
func updateMetric(c *utils.Call) (any, error) {
	if !c.Data.HasChange(isFavorite) || c.Data.HasChangesExcept(isFavorite) {
		options := c.Client.NewUpdateMetricOptions(c.IDs[0])
		if err := c.Fill(utils.ForUpdate, options); err != nil {
			return nil, err
		}
		result, _, err := c.Client.UpdateMetricWithContext(c.Ctx, options)
		if err != nil || !c.Data.HasChange(isFavorite) {
			return result, err
		}
	}
	return setFavorite(c)
}

func setFavorite(c *utils.Call) (any, error) {
	options := c.Client.NewSetMetricFavoriteOptions(c.IDs[0], c.Data.Get(isFavorite).(bool))
	result, _, err := c.Client.SetMetricFavoriteWithContext(c.Ctx, options)
	return result, err
}

func deleteMetric(c *utils.Call) (*core.DetailedResponse, error) {
	return c.Client.DeleteMetricWithContext(c.Ctx, c.Client.NewDeleteMetricOptions(c.IDs[0]))
}
