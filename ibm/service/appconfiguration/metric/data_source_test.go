// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metric_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigMetricDataSourceBasic(t *testing.T) {
	metricID := fmt.Sprintf("tf_metric_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_metric.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_metric", getMetric),
		Steps: []resource.TestStep{
			{
				Config: testAccMetricConfig(metricID, "Purchases", "first", "false") + `
					data "ibm_app_config_metric" "test" {
						guid      = ibm_app_config_metric.test.guid
						metric_id = ibm_app_config_metric.test.metric_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttr(data, "metric_name", "Purchases"),
					resource.TestCheckResourceAttr(data, "event_key", "purchase"),
					resource.TestCheckResourceAttr(data, "direction", "HIGHER"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigMetricsDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Purchases %d", acctest.RandIntRange(10, 100000))
	metricID := fmt.Sprintf("tf_metric_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_metrics.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_metric", getMetric),
		Steps: []resource.TestStep{
			{
				Config: testAccMetricConfig(metricID, name, "first", "false") + `
					data "ibm_app_config_metrics" "test" {
						guid   = ibm_app_config_metric.test.guid
						search = ibm_app_config_metric.test.metric_name
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "metrics.#", "1"),
					resource.TestCheckResourceAttr(data, "metrics.0.metric_id", metricID),
				),
			},
		},
	})
}
