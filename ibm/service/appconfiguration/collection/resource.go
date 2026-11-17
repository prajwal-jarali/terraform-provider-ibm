// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package collection

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createCollection(c *utils.Call) (any, error) {
	options := c.Client.NewCreateCollectionOptions(c.Text(utils.Name), c.IDs[0])
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateCollectionWithContext(c.Ctx, options)
	return result, err
}

// getCollection asks for the expanded view: without it the API leaves out the description, tags, times and href.
func getCollection(c *utils.Call) (any, *core.DetailedResponse, error) {
	options := c.Client.NewGetCollectionOptions(c.IDs[0]).SetExpand(true)
	return c.Client.GetCollectionWithContext(c.Ctx, options)
}

func updateCollection(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateCollectionOptions(c.IDs[0])
	if err := c.Fill(utils.ForUpdate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateCollectionWithContext(c.Ctx, options)
	return result, err
}

func deleteCollection(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeleteCollectionWithContext(c.Ctx, c.Client.NewDeleteCollectionOptions(c.IDs[0]))
	return response, err
}
