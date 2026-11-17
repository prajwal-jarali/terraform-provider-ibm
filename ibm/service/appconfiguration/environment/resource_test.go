// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package environment_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const environmentResource = "ibm_app_config_environment.test"

func TestAccIbmAppConfigEnvironmentBasic(t *testing.T) {
	environmentID := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_environment", getEnvironment),
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfig(environmentID, "Development", "#e2a222", "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(environmentResource, getEnvironment),
					resource.TestCheckResourceAttrSet(environmentResource, "id"),
					resource.TestCheckResourceAttrSet(environmentResource, "href"),
					resource.TestCheckResourceAttrSet(environmentResource, "created_time"),
					resource.TestCheckResourceAttrSet(environmentResource, "updated_time"),
					resource.TestCheckResourceAttr(environmentResource, "environment_id", environmentID),
					resource.TestCheckResourceAttr(environmentResource, "name", "Development"),
					resource.TestCheckResourceAttr(environmentResource, "color_code", "#e2a222"),
					resource.TestCheckResourceAttr(environmentResource, "description", "first"),
				),
			},
			{
				Config: testAccEnvironmentConfig(environmentID, "Testing", "#431133", "second"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(environmentResource, "name", "Testing"),
					resource.TestCheckResourceAttr(environmentResource, "color_code", "#431133"),
					resource.TestCheckResourceAttr(environmentResource, "description", "second"),
				),
			},
			{
				ResourceName:      environmentResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccEnvironmentConfig(environmentID, name, colorCode, description string) string {
	return acceptance.GUIDConfig() + fmt.Sprintf(`
		resource "ibm_app_config_environment" "test" {
			guid           = local.guid
			environment_id = "%s"
			name           = "%s"
			color_code     = "%s"
			description    = "%s"
			tags           = "version v1"
		}`, environmentID, name, colorCode, description)
}

// getEnvironment fetches the resource, to check that it exists.
func getEnvironment(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetEnvironment(c.NewGetEnvironmentOptions(ids[0]))
	return response, err
}
