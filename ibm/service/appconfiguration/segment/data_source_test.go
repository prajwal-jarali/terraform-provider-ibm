// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package segment_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigSegmentDataSourceBasic(t *testing.T) {
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_segment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_segment", getSegment),
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentConfig(segmentID, "Beta users", "endsWith", "@example.com") + `
					data "ibm_app_config_segment" "test" {
						guid       = ibm_app_config_segment.test.guid
						segment_id = ibm_app_config_segment.test.segment_id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "href"),
					resource.TestCheckResourceAttr(data, "name", "Beta users"),
					resource.TestCheckResourceAttr(data, "description", "segment for tests"),
					resource.TestCheckResourceAttr(data, "rules.0.operator", "endsWith"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigSegmentsDataSourceBasic(t *testing.T) {
	// A unique name keeps the search from matching items that are already in the instance.
	name := fmt.Sprintf("Beta users %d", acctest.RandIntRange(10, 100000))
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 1000))
	const data = "data.ibm_app_config_segments.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_segment", getSegment),
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentConfig(segmentID, name, "endsWith", "@example.com") + `
					data "ibm_app_config_segments" "test" {
						guid   = ibm_app_config_segment.test.guid
						search = ibm_app_config_segment.test.name
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(data, "id"),
					resource.TestCheckResourceAttrSet(data, "total_count"),
					resource.TestCheckResourceAttr(data, "segments.#", "1"),
					resource.TestCheckResourceAttr(data, "segments.0.segment_id", segmentID),
					resource.TestCheckResourceAttr(data, "segments.0.name", name),
				),
			},
		},
	})
}
