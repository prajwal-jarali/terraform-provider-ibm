// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package originconfig_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigOriginConfigDataSourceBasic(t *testing.T) {
	const data = "data.ibm_app_config_origin_config.test"
	restoreOrigins(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginConfig(`"https://one.example.com"`) + `
					data "ibm_app_config_origin_config" "test" {
						guid = ibm_app_config_origin_config.test.guid
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttr(data, "allowed_origins.0", "https://one.example.com"),
				),
			},
		},
	})
}
