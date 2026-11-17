// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package feature_test

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

const featureResource = "ibm_app_config_feature.test"

func TestAccIbmAppConfigFeatureBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_feature", getFeature),
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureConfig(featureID, "Checkout", "false", "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(featureResource, getFeature),
					resource.TestCheckResourceAttrSet(featureResource, "id"),
					resource.TestCheckResourceAttrSet(featureResource, "href"),
					resource.TestCheckResourceAttrSet(featureResource, "created_time"),
					resource.TestCheckResourceAttr(featureResource, "feature_id", featureID),
					resource.TestCheckResourceAttr(featureResource, "type", "BOOLEAN"),
					resource.TestCheckResourceAttr(featureResource, "name", "Checkout"),
					resource.TestCheckResourceAttr(featureResource, "enabled", "false"),
					resource.TestCheckResourceAttr(featureResource, "enabled_value", "true"),
					resource.TestCheckResourceAttr(featureResource, "disabled_value", "false"),
					resource.TestCheckResourceAttr(featureResource, "description", "first"),
				),
			},
			{
				// The API rejects enabling a feature flag that is in no collection, so `enabled` stays false.
				Config: testAccFeatureConfig(featureID, "New checkout", "false", "second"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(featureResource, "name", "New checkout"),
					resource.TestCheckResourceAttr(featureResource, "enabled", "false"),
					resource.TestCheckResourceAttr(featureResource, "description", "second"),
				),
			},
			{
				ResourceName:      featureResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// A feature flag that is created in the second environment already exists there, because the API creates
// it in every environment. The resource then updates it instead of failing.
func TestAccIbmAppConfigFeatureExistingInOtherEnvironment(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	// The name and the description belong to the feature in every environment, and names are unique in an
	// instance, so a name left by an earlier run would fail the update.
	name := fmt.Sprintf("Checkout %d", acctest.RandIntRange(10, 100000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureConfig(featureID, name, "false", "first") + fmt.Sprintf(`
					resource "ibm_app_config_environment" "other" {
						guid           = local.guid
						environment_id = "tf_other"
						name           = "Other"
					}
					resource "ibm_app_config_feature" "other" {
						guid           = local.guid
						environment_id = ibm_app_config_environment.other.environment_id
						feature_id     = "%s"
						name           = "%s"
						description    = "first"
						type           = "BOOLEAN"
						enabled_value  = "true"
						disabled_value = "false"
						depends_on     = [ibm_app_config_feature.test]
					}`, featureID, name),
				Check: resource.TestCheckResourceAttr("ibm_app_config_feature.other", "name", name),
			},
		},
	})
}

func testAccFeatureConfig(featureID, name, enabled, description string) string {
	return acceptance.EnvironmentConfig() + fmt.Sprintf(`
		resource "ibm_app_config_feature" "test" {
			guid           = local.guid
			environment_id = ibm_app_config_environment.test.environment_id
			feature_id     = "%s"
			name           = "%s"
			type           = "BOOLEAN"
			enabled_value  = "true"
			disabled_value = "false"
			enabled        = %s
			description    = "%s"
			tags           = "test"
		}`, featureID, name, enabled, description)
}

// getFeature fetches the resource, to check that it exists.
func getFeature(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetFeature(c.NewGetFeatureOptions(ids[0], ids[1]))
	return response, err
}
