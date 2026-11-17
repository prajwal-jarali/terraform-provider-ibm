// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package environment

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestResourceLifecycle(t *testing.T) {
	stored := map[string]any{"environment_id": "dev", "name": "Dev", "color_code": "#FF0000"}
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		switch r.Method {
		case http.MethodPost, http.MethodPut:
			for key, value := range r.Body {
				stored[key] = value
			}
		case http.MethodDelete:
			return schematest.Reply{Status: http.StatusNoContent}
		}
		return schematest.Reply{Status: http.StatusOK, Body: stored}
	})
	res := ResourceIBMAppConfigEnvironment()

	state := schematest.Apply(t, res, meta, nil, map[string]any{
		"guid": "guid", "environment_id": "dev", "name": "Dev", "color_code": "#FF0000",
	})
	if state.ID != "guid/dev" || state.Attributes["color_code"] != "#FF0000" {
		t.Fatalf("unexpected state after create: %v", state)
	}

	state = schematest.Apply(t, res, meta, state, map[string]any{
		"guid": "guid", "environment_id": "dev", "name": "Development", "color_code": "#FF0000",
	})
	last := (*calls)[len(*calls)-1]
	if last.Method != http.MethodPut || last.Body["name"] != "Development" || last.Body["color_code"] != nil {
		t.Fatalf("update must send only the changed name: %+v", last)
	}

	if state = schematest.Refresh(t, res, meta, state); state.Attributes["name"] != "Development" {
		t.Fatalf("refresh did not update the state: %v", state)
	}
	schematest.Destroy(t, res, meta, state)
}

func TestResourceGoneIsRemovedFromState(t *testing.T) {
	gone := false
	meta, _ := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if gone {
			return schematest.Reply{Status: http.StatusNotFound, Body: map[string]any{}}
		}
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{"environment_id": "dev", "name": "Dev"}}
	})
	res := ResourceIBMAppConfigEnvironment()
	state := schematest.Apply(t, res, meta, nil, map[string]any{"guid": "guid", "environment_id": "dev", "name": "Dev"})
	gone = true
	if schematest.Refresh(t, res, meta, state) != nil {
		t.Fatal("a resource that is gone must be removed from the state")
	}
}

func TestListFetchesEveryPage(t *testing.T) {
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Query == "" {
			return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
				"environments": []any{map[string]any{"environment_id": "a", "name": "A"}},
				"next":         map[string]any{"href": "https://x/environments?offset=1"},
			}}
		}
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"environments": []any{map[string]any{"environment_id": "b", "name": "B"}},
		}}
	})
	got := schematest.ReadDataSource(t, DataSourceIBMAppConfigEnvironments(), meta, map[string]any{"guid": "guid"})
	if got["environments.#"] != "2" || got["environments.1.environment_id"] != "b" || len(*calls) != 2 {
		t.Fatalf("expected two environments from two calls, got %v after %d calls", got, len(*calls))
	}
}
