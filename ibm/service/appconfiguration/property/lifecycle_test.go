// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package property

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func config() map[string]any {
	return map[string]any{
		"guid": "guid", "environment_id": "dev", "property_id": "p1", "name": "P", "type": "STRING",
		"format": "JSON", "value": `{"a":1}`,
	}
}

func fakePropertyAPI(t *testing.T, exists bool) (any, *[]schematest.Request) {
	return schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Method == http.MethodGet && !exists {
			return schematest.Reply{Status: http.StatusNotFound, Body: map[string]any{}}
		}
		exists = true
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"property_id": "p1", "name": "P", "type": "STRING", "format": "JSON", "value": map[string]any{"a": 1},
		}}
	})
}

func TestCreateNewProperty(t *testing.T) {
	meta, calls := fakePropertyAPI(t, false)
	state := schematest.Apply(t, ResourceIBMAppConfigProperty(), meta, nil, config())

	created := (*calls)[len(*calls)-1]
	if created.Method != http.MethodPost || created.Body["value"].(map[string]any)["a"] != float64(1) {
		t.Fatalf("a JSON value must be sent as an object, got %+v", created)
	}
	if state.ID != "guid/dev/p1" || state.Attributes["value"] != `{"a":1}` {
		t.Fatalf("unexpected state: %v", state)
	}
}

func TestCreateExistingPropertyUpdatesIt(t *testing.T) {
	meta, calls := fakePropertyAPI(t, true)
	schematest.Apply(t, ResourceIBMAppConfigProperty(), meta, nil, config())

	updated := (*calls)[len(*calls)-1]
	if updated.Method != http.MethodPut || updated.Path != "/environments/dev/properties/p1" {
		t.Fatalf("a property that exists must be updated, got %+v", updated)
	}
}

func TestExistingPropertyWithAnotherFormatIsRejected(t *testing.T) {
	meta, calls := fakePropertyAPI(t, true)
	input := config()
	input["format"], input["value"] = "TEXT", "plain"
	if err := schematest.TryApply(t, ResourceIBMAppConfigProperty(), meta, nil, input); err == nil || len(*calls) != 1 {
		t.Fatalf("a format that differs from the server must stop before any update, got %v after %d calls", err, len(*calls))
	}
}
