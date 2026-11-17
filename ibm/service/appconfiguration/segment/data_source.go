// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package segment

import "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"

func readSegment(c *utils.Call) (any, error) {
	options := c.Client.NewGetSegmentOptions(c.IDs[0])
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.GetSegmentWithContext(c.Ctx, options)
	return result, err
}

func readSegments(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListSegmentsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, resourcesName, &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListSegmentsWithContext(c.Ctx, options)
		return page, err
	})
}
