// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package snapshot

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createSnapshot(c *utils.Call) (any, error) {
	// All the fields of a git config are required, so the options are filled from the schema.
	options := &appconfigurationv1.CreateGitconfigOptions{}
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateGitconfigWithContext(c.Ctx, options)
	return result, err
}

func getSnapshot(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetGitconfigWithContext(c.Ctx, c.Client.NewGetGitconfigOptions(c.IDs[0]))
}

func updateSnapshot(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateGitconfigOptions(c.IDs[0])
	if err := c.Fill(utils.ForUpdate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateGitconfigWithContext(c.Ctx, options)
	return result, err
}

func deleteSnapshot(c *utils.Call) (*core.DetailedResponse, error) {
	return c.Client.DeleteGitconfigWithContext(c.Ctx, c.Client.NewDeleteGitconfigOptions(c.IDs[0]))
}
