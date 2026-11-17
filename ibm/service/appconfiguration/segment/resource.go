// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package segment

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

func createSegment(c *utils.Call) (any, error) {
	// The rules are filled from the schema, like the other fields.
	options := c.Client.NewCreateSegmentOptions(c.Text(utils.Name), c.IDs[0], nil)
	if err := c.Fill(utils.ForCreate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.CreateSegmentWithContext(c.Ctx, options)
	return result, err
}

func getSegment(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.GetSegmentWithContext(c.Ctx, c.Client.NewGetSegmentOptions(c.IDs[0]))
}

func updateSegment(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateSegmentOptions(c.IDs[0])
	if err := c.Fill(utils.ForUpdate, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateSegmentWithContext(c.Ctx, options)
	return result, err
}

func deleteSegment(c *utils.Call) (*core.DetailedResponse, error) {
	_, response, err := c.Client.DeleteSegmentWithContext(c.Ctx, c.Client.NewDeleteSegmentOptions(c.IDs[0]))
	return response, err
}
