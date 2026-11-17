// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package originconfig

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestSingletonWithoutDelete(t *testing.T) {
	stored := map[string]any{"allowed_origins": []any{"*"}}
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Method == http.MethodPut {
			stored["allowed_origins"] = r.Body["allowed_origins"]
		}
		return schematest.Reply{Status: http.StatusOK, Body: stored}
	})
	res := ResourceIBMAppConfigOriginConfig()
	state := schematest.Apply(t, res, meta, nil, map[string]any{"guid": "guid", "allowed_origins": []any{"https://a.com"}})
	if state.ID != "guid" || state.Attributes["allowed_origins.0"] != "https://a.com" {
		t.Fatalf("unexpected state: %v", state)
	}
	state = schematest.Refresh(t, res, meta, state)
	before := len(*calls)
	schematest.Destroy(t, res, meta, state)
	if len(*calls) != before {
		t.Fatal("the API has no delete call, destroying must only touch the state")
	}
}
