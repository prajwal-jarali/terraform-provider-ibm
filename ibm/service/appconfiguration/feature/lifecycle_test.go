// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package feature

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func config() map[string]any {
	return map[string]any{
		"guid": "guid", "environment_id": "dev", "feature_id": "f1", "name": "F", "type": "NUMERIC",
		"enabled_value": "10", "disabled_value": "0",
	}
}

// fakeFeatureAPI answers like the API: the feature exists once it was created, or from the start.
func fakeFeatureAPI(t *testing.T, exists bool) (any, *[]schematest.Request) {
	return schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		if r.Method == http.MethodGet && !exists {
			return schematest.Reply{Status: http.StatusNotFound, Body: map[string]any{}}
		}
		exists = true
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"feature_id": "f1", "name": "F", "type": serverType, "enabled_value": 10, "disabled_value": 0,
		}}
	})
}

var serverType = "NUMERIC"

func TestCreateNewFeature(t *testing.T) {
	meta, calls := fakeFeatureAPI(t, false)
	state := schematest.Apply(t, ResourceIBMAppConfigFeature(), meta, nil, config())

	created := (*calls)[len(*calls)-1]
	if created.Method != http.MethodPost || created.Path != "/environments/dev/features" || created.Body["enabled_value"] != float64(10) {
		t.Fatalf("expected a create with a numeric value, got %+v", created)
	}
	if state.ID != "guid/dev/f1" || state.Attributes["enabled_value"] != "10" {
		t.Fatalf("unexpected state: %v", state)
	}
}

func TestCreateExistingFeatureUpdatesIt(t *testing.T) {
	meta, calls := fakeFeatureAPI(t, true)
	input := config()
	input["description"] = "Updated description"
	schematest.Apply(t, ResourceIBMAppConfigFeature(), meta, nil, input)

	updated := (*calls)[len(*calls)-1]
	if updated.Method != http.MethodPut || updated.Path != "/environments/dev/features/f1" {
		t.Fatalf("a feature that exists must be updated, got %+v", updated)
	}
	if updated.Body["description"] != "Updated description" || updated.Body["tags"] != nil || updated.Body["type"] != nil {
		t.Fatalf("only the fields the user wrote may be sent: %v", updated.Body)
	}
}

func TestSegmentRuleValuesAreConverted(t *testing.T) {
	meta, calls := fakeFeatureAPI(t, false)
	input := config()
	input["segment_rules"] = []any{map[string]any{
		"order": 1, "value": "5", "rules": []any{map[string]any{"segments": []any{"s1"}}},
	}}
	schematest.Apply(t, ResourceIBMAppConfigFeature(), meta, nil, input)

	rule := (*calls)[len(*calls)-1].Body["segment_rules"].([]any)[0].(map[string]any)
	if rule["value"] != float64(5) || rule["order"] != float64(1) {
		t.Fatalf("rule value must follow the type of the feature: %v", rule)
	}
}

func TestExistingFeatureWithAnotherTypeIsRejected(t *testing.T) {
	meta, calls := fakeFeatureAPI(t, true)
	input := config()
	input["type"] = "BOOLEAN"
	input["enabled_value"], input["disabled_value"] = "true", "false"
	err := schematest.TryApply(t, ResourceIBMAppConfigFeature(), meta, nil, input)
	if err == nil || len(*calls) != 1 {
		t.Fatalf("a type that differs from the server must stop before any update, got %v after %d calls", err, len(*calls))
	}
}

func TestExistingFeatureWithInvalidValueIsNotUpdated(t *testing.T) {
	meta, calls := fakeFeatureAPI(t, true)
	input := config()
	input["enabled_value"] = "not a number"
	if err := schematest.TryApply(t, ResourceIBMAppConfigFeature(), meta, nil, input); err == nil || len(*calls) != 1 {
		t.Fatalf("an invalid value must stop before any update, got %v after %d calls", err, len(*calls))
	}
}
