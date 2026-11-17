// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metric_test

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

const metricResource = "ibm_app_config_metric.test"

func TestAccIbmAppConfigMetricBasic(t *testing.T) {
	metricID := fmt.Sprintf("tf_metric_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_metric", getMetric),
		Steps: []resource.TestStep{
			{
				Config: testAccMetricConfig(metricID, "Purchases", "first", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(metricResource, getMetric),
					resource.TestCheckResourceAttrSet(metricResource, "id"),
					resource.TestCheckResourceAttrSet(metricResource, "created_time"),
					resource.TestCheckResourceAttr(metricResource, "metric_id", metricID),
					resource.TestCheckResourceAttr(metricResource, "metric_name", "Purchases"),
					resource.TestCheckResourceAttr(metricResource, "event_type", "CUSTOM"),
					resource.TestCheckResourceAttr(metricResource, "metric_type", "COUNT"),
					resource.TestCheckResourceAttr(metricResource, "direction", "HIGHER"),
					resource.TestCheckResourceAttr(metricResource, "is_favorite", "false"),
				),
			},
			{
				Config: testAccMetricConfig(metricID, "Orders", "second", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(metricResource, "metric_name", "Orders"),
					resource.TestCheckResourceAttr(metricResource, "description", "second"),
					resource.TestCheckResourceAttr(metricResource, "is_favorite", "true"),
				),
			},
			{
				ResourceName:      metricResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMetricConfig(metricID, name, description, favorite string) string {
	return acceptance.GUIDConfig() + fmt.Sprintf(`
		resource "ibm_app_config_metric" "test" {
			guid        = local.guid
			metric_id   = "%s"
			metric_name = "%s"
			description = "%s"
			event_type  = "CUSTOM"
			event_key   = "purchase"
			metric_type = "COUNT"
			direction   = "HIGHER"
			is_favorite = %s
		}`, metricID, name, description, favorite)
}

// getMetric fetches the resource, to check that it exists.
func getMetric(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetMetric(c.NewGetMetricOptions(ids[0]))
	return response, err
}
