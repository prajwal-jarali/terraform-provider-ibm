// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package collection_test

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

const collectionResource = "ibm_app_config_collection.test"

func TestAccIbmAppConfigCollectionBasic(t *testing.T) {
	collectionID := fmt.Sprintf("tf_collection_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_collection", getCollection),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(collectionID, "Web", "first", "tag1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(collectionResource, getCollection),
					resource.TestCheckResourceAttrSet(collectionResource, "id"),
					resource.TestCheckResourceAttrSet(collectionResource, "href"),
					resource.TestCheckResourceAttrSet(collectionResource, "created_time"),
					resource.TestCheckResourceAttrSet(collectionResource, "updated_time"),
					resource.TestCheckResourceAttr(collectionResource, "collection_id", collectionID),
					resource.TestCheckResourceAttr(collectionResource, "name", "Web"),
					resource.TestCheckResourceAttr(collectionResource, "description", "first"),
					resource.TestCheckResourceAttr(collectionResource, "tags", "tag1"),
				),
			},
			{
				Config: testAccCollectionConfig(collectionID, "Mobile", "second", "tag2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(collectionResource, "name", "Mobile"),
					resource.TestCheckResourceAttr(collectionResource, "description", "second"),
					resource.TestCheckResourceAttr(collectionResource, "tags", "tag2"),
				),
			},
			{
				ResourceName:      collectionResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCollectionConfig(collectionID, name, description, tags string) string {
	return acceptance.GUIDConfig() + fmt.Sprintf(`
		resource "ibm_app_config_collection" "test" {
			guid          = local.guid
			collection_id = "%s"
			name          = "%s"
			description   = "%s"
			tags          = "%s"
		}`, collectionID, name, description, tags)
}

// getCollection fetches the resource, to check that it exists.
func getCollection(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetCollection(c.NewGetCollectionOptions(ids[0]))
	return response, err
}
