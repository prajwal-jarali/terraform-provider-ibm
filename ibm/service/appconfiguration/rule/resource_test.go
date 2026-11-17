// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rule_test

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

const ruleResource = "ibm_app_config_rule.test"

func TestAccIbmAppConfigRuleBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_rule", getRule),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig(featureID, "Beta users", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(ruleResource, getRule),
					resource.TestCheckResourceAttrSet(ruleResource, "id"),
					resource.TestCheckResourceAttrSet(ruleResource, "href"),
					resource.TestCheckResourceAttrSet(ruleResource, "order"),
					resource.TestCheckResourceAttr(ruleResource, "rule_id", "tf_rule"),
					resource.TestCheckResourceAttr(ruleResource, "rule_name", "Beta users"),
					resource.TestCheckResourceAttr(ruleResource, "value", "true"),
					resource.TestCheckResourceAttr(ruleResource, "rules.#", "1"),
				),
			},
			{
				Config: testAccRuleConfig(featureID, "Beta testers", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(ruleResource, "rule_name", "Beta testers"),
					resource.TestCheckResourceAttr(ruleResource, "value", "false"),
				),
			},
			{
				ResourceName:      ruleResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRuleConfig(featureID, name, value string) string {
	return acceptance.FeatureConfig(featureID) + fmt.Sprintf(`
		resource "ibm_app_config_segment" "test" {
			guid       = local.guid
			segment_id = "tf_segment"
			name       = "Beta users"
			rules {
				attribute_name = "email"
				operator       = "endsWith"
				values         = ["@example.com"]
			}
		}
		resource "ibm_app_config_rule" "test" {
			guid           = local.guid
			environment_id = ibm_app_config_feature.test.environment_id
			feature_id     = ibm_app_config_feature.test.feature_id
			rule_id        = "tf_rule"
			rule_name      = "%s"
			value          = "%s"
			rules {
				segments = [ibm_app_config_segment.test.segment_id]
			}
		}`, name, value)
}

// getRule fetches the resource, to check that it exists.
func getRule(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetFeatureRule(c.NewGetFeatureRuleOptions(ids[0], ids[1], ids[2]))
	return response, err
}
