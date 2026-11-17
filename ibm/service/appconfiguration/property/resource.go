// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package property

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

// textFields are written as text by the user, and sent as the value their type and format describe.
var textFields = []string{utils.Value, utils.SegmentRules}

// createProperty creates a property. The API creates it in every environment of the instance, so when
// it exists already, its type and format must match the configuration, and the fields the user wrote
// are checked and applied to it as an update.
func createProperty(c *utils.Call) (any, error) {
	existing, response, err := c.Client.GetPropertyWithContext(c.Ctx, c.Client.NewGetPropertyOptions(c.IDs[0], c.IDs[1]))
	if err == nil {
		if err := utils.CheckExisting(c, "property", existing.Type, existing.Format); err != nil {
			return nil, err
		}
		return saveProperty(c, utils.ForCreate)
	}
	if !utils.IsNotFound(response) {
		return nil, err
	}

	typ, format := c.Text(utils.Type), c.Text(utils.Format)
	value, err := utils.ParseValue(typ, format, c.Text(utils.Value))
	if err != nil {
		return nil, err
	}
	options := c.Client.NewCreatePropertyOptions(c.IDs[0], c.Text(utils.Name), c.IDs[1], typ, value)
	if err := c.Fill(utils.Skip(utils.ForCreate, textFields...), options); err != nil {
		return nil, err
	}
	if err := utils.DecodeRules(c, utils.ForCreate, utils.SegmentRules, typ, format, &options.SegmentRules); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreatePropertyWithContext(c.Ctx, options)
	return result, err
}

func getProperty(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetPropertyWithContext(c.Ctx, c.Client.NewGetPropertyOptions(c.IDs[0], c.IDs[1]))
}

func updateProperty(c *utils.Call) (any, error) {
	return saveProperty(c, utils.ForUpdate)
}

// saveProperty updates the fields that the guard accepts.
func saveProperty(c *utils.Call, guard utils.Guard) (any, error) {
	options := c.Client.NewUpdatePropertyOptions(c.IDs[0], c.IDs[1])
	if err := c.Fill(utils.Skip(guard, textFields...), options); err != nil {
		return nil, err
	}
	typ, format := c.Text(utils.Type), c.Text(utils.Format)
	if guard(c.Data, utils.Value) {
		value, err := utils.ParseValue(typ, format, c.Text(utils.Value))
		if err != nil {
			return nil, err
		}
		options.Value = value
	}
	if err := utils.DecodeRules(c, guard, utils.SegmentRules, typ, format, &options.SegmentRules); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdatePropertyWithContext(c.Ctx, options)
	return result, err
}

func deleteProperty(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeletePropertyWithContext(c.Ctx, c.Client.NewDeletePropertyOptions(c.IDs[0], c.IDs[1]))
	return response, err
}
