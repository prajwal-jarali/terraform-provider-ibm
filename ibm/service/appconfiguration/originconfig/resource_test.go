// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package originconfig_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const originConfigResource = "ibm_app_config_origin_config.test"

// The API has no call to delete the origin config, so the test puts the original origins back when it ends.
func TestAccIbmAppConfigOriginConfigBasic(t *testing.T) {
	restoreOrigins(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginConfig(`"https://one.example.com"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(originConfigResource, getOriginConfig),
					resource.TestCheckResourceAttrSet(originConfigResource, "id"),
					resource.TestCheckResourceAttr(originConfigResource, "allowed_origins.#", "1"),
					resource.TestCheckResourceAttr(originConfigResource, "allowed_origins.0", "https://one.example.com"),
				),
			},
			{
				Config: testAccOriginConfig(`"https://one.example.com", "https://two.example.com"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(originConfigResource, "allowed_origins.#", "2"),
					resource.TestCheckResourceAttr(originConfigResource, "allowed_origins.1", "https://two.example.com"),
				),
			},
		},
	})
}

func testAccOriginConfig(origins string) string {
	return acceptance.GUIDConfig() + fmt.Sprintf(`
		resource "ibm_app_config_origin_config" "test" {
			guid            = local.guid
			allowed_origins = [%s]
		}`, origins)
}

// restoreOrigins remembers the allowed origins of the instance and puts them back when the test ends,
// also when it fails, because destroying the resource leaves the last value in place.
func restoreOrigins(t *testing.T) {
	t.Helper()
	// This runs before resource.Test, which is what skips an acceptance test, so it skips here too.
	if os.Getenv(resource.EnvTfAcc) == "" {
		t.Skipf("set %s to run the acceptance tests", resource.EnvTfAcc)
	}
	acceptance.PreCheck(t)
	client, err := utils.NewClient(acc.TestAccProvider.Meta(), acceptance.GUID())
	if err != nil {
		t.Fatal(err)
	}
	original, _, err := client.ListOriginconfigs(client.NewListOriginconfigsOptions())
	if err != nil {
		t.Fatalf("reading the allowed origins: %s", err)
	}
	t.Cleanup(func() {
		if _, _, err := client.UpdateOriginconfigs(client.NewUpdateOriginconfigsOptions(original.AllowedOrigins)); err != nil {
			t.Errorf("restoring the allowed origins %v: %s", original.AllowedOrigins, err)
		}
	})
}

// getOriginConfig fetches the resource, to check that it exists.
func getOriginConfig(c *appconfigurationv1.AppConfigurationV1, _ []string) (*core.DetailedResponse, error) {
	_, response, err := c.ListOriginconfigs(c.NewListOriginconfigsOptions())
	return response, err
}
