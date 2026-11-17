// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package originconfig

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/go-sdk-core/v5/core"
)

func getOriginConfig(c *utils.Call) (any, *core.DetailedResponse, error) {
	return c.Client.ListOriginconfigsWithContext(c.Ctx, c.Client.NewListOriginconfigsOptions())
}

func readOriginConfig(c *utils.Call) (any, error) {
	result, _, err := getOriginConfig(c)
	return result, err
}

// saveOriginConfig sets the allowed origins of the instance. The API has one call for it, for the first
// time and for each change.
func saveOriginConfig(c *utils.Call) (any, error) {
	options := c.Client.NewUpdateOriginconfigsOptions(nil)
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	result, _, err := c.Client.UpdateOriginconfigsWithContext(c.Ctx, options)
	return result, err
}
