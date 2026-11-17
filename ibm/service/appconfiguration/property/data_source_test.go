// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package property_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigPropertyDataSourceBasic(t *testing.T) {
	propertyID := fmt.Sprintf("tf_property_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_property.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_property", getProperty),
		Steps: []resource.TestStep{
			{
				Config: testAccPropertyConfig(propertyID, "Timeout", "30", "first") + `
					data "ibm_app_config_property" "test" {
						guid           = ibm_app_config_property.test.guid
						environment_id = ibm_app_config_property.test.environment_id
						property_id    = ibm_app_config_property.test.property_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "href"),
					resource.TestCheckResourceAttr(data, "name", "Timeout"),
					resource.TestCheckResourceAttr(data, "type", "NUMERIC"),
					resource.TestCheckResourceAttr(data, "value", "30"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigPropertiesDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Timeout %d", acctest.RandIntRange(10, 100000))
	propertyID := fmt.Sprintf("tf_property_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_properties.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_property", getProperty),
		Steps: []resource.TestStep{
			{
				Config: testAccPropertyConfig(propertyID, name, "30", "first") + `
					data "ibm_app_config_properties" "test" {
						guid           = ibm_app_config_property.test.guid
						environment_id = ibm_app_config_property.test.environment_id
						search         = ibm_app_config_property.test.name
						expand         = true
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "properties.#", "1"),
					resource.TestCheckResourceAttr(data, "properties.0.property_id", propertyID),
					resource.TestCheckResourceAttr(data, "properties.0.type", "NUMERIC"),
				),
			},
		},
	})
}
