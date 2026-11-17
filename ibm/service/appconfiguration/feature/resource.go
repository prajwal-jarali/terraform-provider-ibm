// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package feature

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

// textFields are written as text by the user, and sent as the value their type and format describe.
var textFields = []string{utils.EnabledValue, utils.DisabledValue, utils.SegmentRules}

// createFeature creates a feature. The API creates it in every environment of the instance, so when
// it exists already, its type and format must match the configuration, and the fields the user wrote
// are checked and applied to it as an update.
func createFeature(c *utils.Call) (any, error) {
	existing, response, err := c.Client.GetFeatureWithContext(c.Ctx, c.Client.NewGetFeatureOptions(c.IDs[0], c.IDs[1]))
	if err == nil {
		if err := utils.CheckExisting(c, "feature", existing.Type, existing.Format); err != nil {
			return nil, err
		}
		return saveFeature(c, utils.ForCreate)
	}
	if !utils.IsNotFound(response) {
		return nil, err
	}

	typ, format := c.Text(utils.Type), c.Text(utils.Format)
	enabled, err := utils.ParseValue(typ, format, c.Text(utils.EnabledValue))
	if err != nil {
		return nil, err
	}
	disabled, err := utils.ParseValue(typ, format, c.Text(utils.DisabledValue))
	if err != nil {
		return nil, err
	}
	options := c.Client.NewCreateFeatureOptions(c.IDs[0], c.Text(utils.Name), c.IDs[1], typ, enabled, disabled)
	if err := c.Fill(utils.Skip(utils.ForCreate, textFields...), options); err != nil {
		return nil, err
	}
	if err := utils.DecodeRules(c, utils.ForCreate, utils.SegmentRules, typ, format, &options.SegmentRules); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateFeatureWithContext(c.Ctx, options)
	return result, err
}

func getFeature(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetFeatureWithContext(c.Ctx, c.Client.NewGetFeatureOptions(c.IDs[0], c.IDs[1]))
}

func updateFeature(c *utils.Call) (any, error) {
	return saveFeature(c, utils.ForUpdate)
}

// saveFeature updates the fields that the guard accepts.
func saveFeature(c *utils.Call, guard utils.Guard) (any, error) {
	options := c.Client.NewUpdateFeatureOptions(c.IDs[0], c.IDs[1])
	if err := c.Fill(utils.Skip(guard, textFields...), options); err != nil {
		return nil, err
	}
	typ, format := c.Text(utils.Type), c.Text(utils.Format)
	for key, target := range map[string]*any{utils.EnabledValue: &options.EnabledValue, utils.DisabledValue: &options.DisabledValue} {
		if !guard(c.Data, key) {
			continue
		}
		value, err := utils.ParseValue(typ, format, c.Text(key))
		if err != nil {
			return nil, err
		}
		*target = value
	}
	if err := utils.DecodeRules(c, guard, utils.SegmentRules, typ, format, &options.SegmentRules); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateFeatureWithContext(c.Ctx, options)
	return result, err
}

func deleteFeature(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeleteFeatureWithContext(c.Ctx, c.Client.NewDeleteFeatureOptions(c.IDs[0], c.IDs[1]))
	return response, err
}
