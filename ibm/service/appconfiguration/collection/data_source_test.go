// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package collection_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigCollectionDataSourceBasic(t *testing.T) {
	collectionID := fmt.Sprintf("tf_collection_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_collection.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_collection", getCollection),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(collectionID, "Web", "first", "tag1") + `
					data "ibm_app_config_collection" "test" {
						guid          = ibm_app_config_collection.test.guid
						collection_id = ibm_app_config_collection.test.collection_id
						expand        = true
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "href"),
					resource.TestCheckResourceAttrSet(data, "created_time"),
					resource.TestCheckResourceAttr(data, "name", "Web"),
					resource.TestCheckResourceAttr(data, "description", "first"),
					resource.TestCheckResourceAttr(data, "tags", "tag1"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigCollectionsDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Web %d", acctest.RandIntRange(10, 100000))
	collectionID := fmt.Sprintf("tf_collection_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_collections.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_collection", getCollection),
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(collectionID, name, "first", "tag1") + `
					data "ibm_app_config_collections" "test" {
						guid   = ibm_app_config_collection.test.guid
						search = ibm_app_config_collection.test.name
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "collections.#", "1"),
					resource.TestCheckResourceAttr(data, "collections.0.collection_id", collectionID),
					resource.TestCheckResourceAttr(data, "collections.0.name", name),
				),
			},
		},
	})
}
