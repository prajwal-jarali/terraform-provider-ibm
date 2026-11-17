// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package property_test

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

const propertyResource = "ibm_app_config_property.test"

func TestAccIbmAppConfigPropertyBasic(t *testing.T) {
	propertyID := fmt.Sprintf("tf_property_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_property", getProperty),
		Steps: []resource.TestStep{
			{
				Config: testAccPropertyConfig(propertyID, "Timeout", "30", "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(propertyResource, getProperty),
					resource.TestCheckResourceAttrSet(propertyResource, "id"),
					resource.TestCheckResourceAttrSet(propertyResource, "href"),
					resource.TestCheckResourceAttrSet(propertyResource, "created_time"),
					resource.TestCheckResourceAttr(propertyResource, "property_id", propertyID),
					resource.TestCheckResourceAttr(propertyResource, "type", "NUMERIC"),
					resource.TestCheckResourceAttr(propertyResource, "name", "Timeout"),
					resource.TestCheckResourceAttr(propertyResource, "value", "30"),
					resource.TestCheckResourceAttr(propertyResource, "description", "first"),
				),
			},
			{
				Config: testAccPropertyConfig(propertyID, "Request timeout", "60", "second"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(propertyResource, "name", "Request timeout"),
					resource.TestCheckResourceAttr(propertyResource, "value", "60"),
					resource.TestCheckResourceAttr(propertyResource, "description", "second"),
				),
			},
			{
				ResourceName:      propertyResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPropertyConfig(propertyID, name, value, description string) string {
	return acceptance.EnvironmentConfig() + fmt.Sprintf(`
		resource "ibm_app_config_property" "test" {
			guid           = local.guid
			environment_id = ibm_app_config_environment.test.environment_id
			property_id    = "%s"
			name           = "%s"
			type           = "NUMERIC"
			value          = "%s"
			description    = "%s"
			tags           = "test"
		}`, propertyID, name, value, description)
}

// getProperty fetches the resource, to check that it exists.
func getProperty(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetProperty(c.NewGetPropertyOptions(ids[0], ids[1]))
	return response, err
}
