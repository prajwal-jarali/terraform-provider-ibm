// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rule_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigRuleDataSourceBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_rule", getRule),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig(featureID, "Beta users", "true") + `
					data "ibm_app_config_rule" "test" {
						guid           = ibm_app_config_rule.test.guid
						environment_id = ibm_app_config_rule.test.environment_id
						feature_id     = ibm_app_config_rule.test.feature_id
						rule_id        = ibm_app_config_rule.test.rule_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "order"),
					resource.TestCheckResourceAttr(data, "rule_name", "Beta users"),
					resource.TestCheckResourceAttr(data, "value", "true"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigRulesDataSourceBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_rules.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_rule", getRule),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig(featureID, "Beta users", "true") + `
					data "ibm_app_config_rules" "test" {
						guid           = ibm_app_config_rule.test.guid
						environment_id = ibm_app_config_rule.test.environment_id
						feature_id     = ibm_app_config_rule.test.feature_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttr(data, "segment_rules.#", "1"),
					resource.TestCheckResourceAttr(data, "segment_rules.0.rule_id", "tf_rule"),
				),
			},
		},
	})
}
