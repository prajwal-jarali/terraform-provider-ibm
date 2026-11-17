// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package segment_test

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

const segmentResource = "ibm_app_config_segment.test"

func TestAccIbmAppConfigSegmentBasic(t *testing.T) {
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: acceptance.CheckDestroy("ibm_app_config_segment", getSegment),
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentConfig(segmentID, "Beta users", "endsWith", "@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acceptance.CheckExists(segmentResource, getSegment),
					resource.TestCheckResourceAttrSet(segmentResource, "id"),
					resource.TestCheckResourceAttrSet(segmentResource, "href"),
					resource.TestCheckResourceAttrSet(segmentResource, "created_time"),
					resource.TestCheckResourceAttr(segmentResource, "segment_id", segmentID),
					resource.TestCheckResourceAttr(segmentResource, "name", "Beta users"),
					resource.TestCheckResourceAttr(segmentResource, "rules.#", "1"),
					resource.TestCheckResourceAttr(segmentResource, "rules.0.attribute_name", "email"),
					resource.TestCheckResourceAttr(segmentResource, "rules.0.operator", "endsWith"),
					resource.TestCheckResourceAttr(segmentResource, "rules.0.values.0", "@example.com"),
				),
			},
			{
				Config: testAccSegmentConfig(segmentID, "Internal users", "startsWith", "dev."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(segmentResource, "name", "Internal users"),
					resource.TestCheckResourceAttr(segmentResource, "rules.0.operator", "startsWith"),
					resource.TestCheckResourceAttr(segmentResource, "rules.0.values.0", "dev."),
				),
			},
			{
				ResourceName:      segmentResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSegmentConfig(segmentID, name, operator, value string) string {
	return acceptance.GUIDConfig() + fmt.Sprintf(`
		resource "ibm_app_config_segment" "test" {
			guid        = local.guid
			segment_id  = "%s"
			name        = "%s"
			description = "segment for tests"
			tags        = "test"
			rules {
				attribute_name = "email"
				operator       = "%s"
				values         = ["%s"]
			}
		}`, segmentID, name, operator, value)
}

// getSegment fetches the resource, to check that it exists.
func getSegment(c *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error) {
	_, response, err := c.GetSegment(c.NewGetSegmentOptions(ids[0]))
	return response, err
}
