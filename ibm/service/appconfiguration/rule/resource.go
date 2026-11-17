// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package rule

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createRule(c *utils.Call) (any, error) {
	options := c.Client.NewCreateFeatureRuleOptions(c.IDs[0], c.IDs[1], nil, nil, c.IDs[2])
	if err := c.Fill(utils.Skip(utils.ForCreate, utils.Value), options); err != nil {
		return nil, err
	}
	var err error
	if options.Value, err = getValue(c); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateFeatureRuleWithContext(c.Ctx, options)
	return result, err
}

func getRule(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetFeatureRuleWithContext(c.Ctx, c.Client.NewGetFeatureRuleOptions(c.IDs[0], c.IDs[1], c.IDs[2]))
}

func updateRule(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateFeatureRuleOptions(c.IDs[0], c.IDs[1], c.IDs[2])
	if err := c.Fill(utils.Skip(utils.ForUpdate, utils.Value), options); err != nil {
		return nil, err
	}
	if c.Data.HasChange(utils.Value) {
		var err error
		if options.Value, err = getValue(c); err != nil {
			return nil, err
		}
	}
	result, _, err := c.Client.UpdateFeatureRuleWithContext(c.Ctx, options)
	return result, err
}

func deleteRule(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeleteFeatureRuleWithContext(c.Ctx, c.Client.NewDeleteFeatureRuleOptions(c.IDs[0], c.IDs[1], c.IDs[2]))
	return response, err
}

// getValue converts the value the user wrote. Its type and format are the ones of the feature.
func getValue(c *utils.Call) (any, error) {
	feature, _, err := c.Client.GetFeatureWithContext(c.Ctx, c.Client.NewGetFeatureOptions(c.IDs[0], c.IDs[1]))
	if err != nil {
		return nil, err
	}
	return parseFeatureValue(feature, c.Text(utils.Value))
}

func parseFeatureValue(feature *appconfigurationv1.Feature, text string) (any, error) {
	format := ""
	if feature.Format != nil {
		format = *feature.Format
	}
	return utils.ParseValue(*feature.Type, format, text)
}
