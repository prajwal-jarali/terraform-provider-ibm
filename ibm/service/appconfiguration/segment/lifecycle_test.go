// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package segment

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestCreateSendsRules(t *testing.T) {
	rules := []any{map[string]any{"attribute_name": "email", "operator": "endsWith", "values": []any{"@ibm.com"}}}
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"segment_id": "s1", "name": "S", "rules": rules,
		}}
	})
	state := schematest.Apply(t, ResourceIBMAppConfigSegment(), meta, nil, map[string]any{
		"guid": "guid", "segment_id": "s1", "name": "S", "rules": rules,
	})
	sent := (*calls)[0].Body["rules"].([]any)[0].(map[string]any)
	if sent["attribute_name"] != "email" || sent["values"].([]any)[0] != "@ibm.com" {
		t.Fatalf("rules were not sent as the API expects: %v", (*calls)[0].Body)
	}
	if state.Attributes["rules.0.operator"] != "endsWith" {
		t.Fatalf("rules were not stored: %v", state.Attributes)
	}
}
