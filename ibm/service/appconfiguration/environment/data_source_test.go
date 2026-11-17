// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package environment_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigEnvironmentDataSourceBasic(t *testing.T) {
	environmentID := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_environment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_environment", getEnvironment),
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfig(environmentID, "Development", "#e2a222", "first") + `
					data "ibm_app_config_environment" "test" {
						guid           = ibm_app_config_environment.test.guid
						environment_id = ibm_app_config_environment.test.environment_id
						expand         = true
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "href"),
					resource.TestCheckResourceAttrSet(data, "created_time"),
					resource.TestCheckResourceAttrSet(data, "updated_time"),
					resource.TestCheckResourceAttr(data, "name", "Development"),
					resource.TestCheckResourceAttr(data, "color_code", "#e2a222"),
					resource.TestCheckResourceAttr(data, "description", "first"),
					resource.TestCheckResourceAttr(data, "tags", "version v1"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigEnvironmentsDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Development %d", acctest.RandIntRange(10, 100000))
	environmentID := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_environments.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_environment", getEnvironment),
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfig(environmentID, name, "#e2a222", "first") + `
					data "ibm_app_config_environments" "test" {
						guid   = ibm_app_config_environment.test.guid
						search = ibm_app_config_environment.test.name
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "environments.#", "1"),
					resource.TestCheckResourceAttr(data, "environments.0.environment_id", environmentID),
					resource.TestCheckResourceAttr(data, "environments.0.name", name),
				),
			},
		},
	})
}
