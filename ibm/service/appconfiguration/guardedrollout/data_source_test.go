// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package guardedrollout_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// A guarded rollout is started by the service, not by Terraform, so a new feature has none.
func TestAccIbmAppConfigGuardedRolloutsDataSourceBasic(t *testing.T) {
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_guarded_rollouts.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: acceptance.FeatureConfig(featureID) + `
					data "ibm_app_config_guarded_rollouts" "test" {
						guid           = ibm_app_config_feature.test.guid
						environment_id = ibm_app_config_feature.test.environment_id
						feature_id     = ibm_app_config_feature.test.feature_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttr(data, "guarded_rollouts.#", "0"),
				),
			},
		},
	})
}
