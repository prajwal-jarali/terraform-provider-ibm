// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const idSeparator = "/"

// BuildID joins the GUID and the ids of a resource into a Terraform id: `<guid>/<id>/...`.
func BuildID(guid string, ids ...string) string {
	return strings.Join(append([]string{guid}, ids...), idSeparator)
}

// ParseID splits a resource id of the form `<guid>/<id>/...` and stores the GUID in the data,
// so that GetAppConfigClient works for resources that are read by id only (e.g. after an import).
// A resource without ids has the GUID as its id.
// It expects exactly `parts` ids after the GUID and returns them.
func ParseID(data *schema.ResourceData, parts int) ([]string, error) {
	if parts == 0 {
		return nil, data.Set(GUID, data.Id())
	}
	ids, err := flex.IdParts(data.Id())
	if err != nil {
		return nil, err
	}
	if len(ids) != parts+1 || strings.TrimSpace(ids[0]) == "" {
		return nil, flex.FmtErrorf(
			"The given id %s is invalid, expected format: <guid>%s",
			data.Id(), strings.Repeat(idSeparator+"<id>", parts),
		)
	}
	if err := data.Set(GUID, ids[0]); err != nil {
		return nil, err
	}
	return ids[1:], nil
}
