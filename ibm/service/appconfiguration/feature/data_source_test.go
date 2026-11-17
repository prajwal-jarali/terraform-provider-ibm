// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package feature_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigFeatureDataSourceBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_feature.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_feature", getFeature),
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureConfig(featureID, "Checkout", "false", "first") + `
					data "ibm_app_config_feature" "test" {
						guid           = ibm_app_config_feature.test.guid
						environment_id = ibm_app_config_feature.test.environment_id
						feature_id     = ibm_app_config_feature.test.feature_id
						include        = ["rules"]
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "href"),
					resource.TestCheckResourceAttrSet(data, "segment_exists"),
					resource.TestCheckResourceAttr(data, "name", "Checkout"),
					resource.TestCheckResourceAttr(data, "type", "BOOLEAN"),
					resource.TestCheckResourceAttr(data, "enabled_value", "true"),
					resource.TestCheckResourceAttr(data, "disabled_value", "false"),
					resource.TestCheckResourceAttr(data, "enabled", "false"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigFeaturesDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Checkout %d", acctest.RandIntRange(10, 100000))
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_features.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_feature", getFeature),
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureConfig(featureID, name, "false", "first") + `
					data "ibm_app_config_features" "test" {
						guid           = ibm_app_config_feature.test.guid
						environment_id = ibm_app_config_feature.test.environment_id
						search         = ibm_app_config_feature.test.name
						expand         = true
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "features.#", "1"),
					resource.TestCheckResourceAttr(data, "features.0.feature_id", featureID),
					resource.TestCheckResourceAttr(data, "features.0.name", name),
					resource.TestCheckResourceAttr(data, "features.0.type", "BOOLEAN"),
				),
			},
		},
	})
}
