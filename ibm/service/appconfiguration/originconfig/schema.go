// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package originconfig holds the schemas of the origin config resource and data source.
// The origin config lists the origins that may call the client side of the instance. There is one per instance.
package originconfig

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	allowedOrigins = "allowed_origins"
)

const resourceName = "origin config"

func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			allowedOrigins: mode.Required(utils.NewStringList("List of allowed origins.")),
		},
	)
}

// DataSourceIBMAppConfigOriginConfig returns the origin config data source.
func DataSourceIBMAppConfigOriginConfig() *schema.Resource {
	return utils.NewDataSource("Origin config", nil, getFields(utils.Output), readOriginConfig)
}

// ResourceIBMAppConfigOriginConfig returns the origin config resource. The API has no call to delete the
// origin config, so destroying the resource only removes it from the state.
func ResourceIBMAppConfigOriginConfig() *schema.Resource {
	return utils.NewResource("Origin config", nil, getFields(utils.Input), utils.Calls{
		Create: saveOriginConfig,
		Get:    getOriginConfig,
		Update: saveOriginConfig,
	})
}
