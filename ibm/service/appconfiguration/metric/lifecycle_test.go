// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package metric

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestFavoriteHasItsOwnCall(t *testing.T) {
	stored := map[string]any{"metric_id": "m1", "metric_name": "M", "is_favorite": false}
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Method != http.MethodGet {
			for key, value := range r.Body {
				stored[key] = value
			}
		}
		return schematest.Reply{Status: http.StatusOK, Body: stored}
	})
	res := ResourceIBMAppConfigMetric()
	config := map[string]any{
		"guid": "guid", "metric_id": "m1", "metric_name": "M", "event_type": "CUSTOM", "metric_type": "COUNT",
		"event_key": "k", "direction": "HIGHER", "is_favorite": true,
	}
	state := schematest.Apply(t, res, meta, nil, config)
	if got := (*calls)[len(*calls)-1]; got.Method != http.MethodPatch || got.Path != "/metrics/m1/favorite" {
		t.Fatalf("expected the favorite call after create, got %+v", got)
	}

	config["is_favorite"] = false
	before := len(*calls)
	schematest.Apply(t, res, meta, state, config)
	for _, call := range (*calls)[before:] {
		if call.Path == "/metrics/m1" && call.Method == http.MethodPut {
			t.Fatalf("a favorite change must not update the metric: %+v", call)
		}
	}
}
