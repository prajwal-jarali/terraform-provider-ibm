// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package environment

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createEnvironment(c *utils.Call) (any, error) {
	options := c.Client.NewCreateEnvironmentOptions(c.Text(utils.Name), c.IDs[0])
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateEnvironmentWithContext(c.Ctx, options)
	return result, err
}

// getEnvironment asks for the expanded view: without it the API leaves out most fields.
func getEnvironment(c *utils.Call) (any, *core.DetailedResponse, error) {
	options := c.Client.NewGetEnvironmentOptions(c.IDs[0]).SetExpand(true)
	return c.Client.GetEnvironmentWithContext(c.Ctx, options)
}

func updateEnvironment(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateEnvironmentOptions(c.IDs[0])
	if err := c.Fill(utils.ForUpdate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateEnvironmentWithContext(c.Ctx, options)
	return result, err
}

func deleteEnvironment(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeleteEnvironmentWithContext(c.Ctx, c.Client.NewDeleteEnvironmentOptions(c.IDs[0]))
	return response, err
}
