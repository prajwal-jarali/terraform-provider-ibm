// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package rule

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestValueFollowsTheTypeOfTheFeature(t *testing.T) {
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Path == "/environments/dev/features/f1" {
			return schematest.Reply{Status: http.StatusOK, Body: map[string]any{"feature_id": "f1", "type": "NUMERIC"}}
		}
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"rule_id": "r1", "value": 7, "order": 1, "rules": []any{map[string]any{"segments": []any{"s1"}}},
		}}
	})
	state := schematest.Apply(t, ResourceIBMAppConfigRule(), meta, nil, map[string]any{
		"guid": "guid", "environment_id": "dev", "feature_id": "f1", "rule_id": "r1", "value": "7",
		"rules": []any{map[string]any{"segments": []any{"s1"}}},
	})
	created := (*calls)[len(*calls)-1]
	if created.Method != http.MethodPost || created.Body["value"] != float64(7) {
		t.Fatalf("a numeric feature needs a number as value: %+v", created)
	}
	if state.ID != "guid/dev/f1/r1" || state.Attributes["value"] != "7" {
		t.Fatalf("unexpected state: %v", state)
	}
}
