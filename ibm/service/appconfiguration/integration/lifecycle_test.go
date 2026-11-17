// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package integration

import (
	"net/http"
	"strings"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestCreateKeepsSourceFields(t *testing.T) {
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{
			"integration_id": "en", "integration_type": "EVENT_NOTIFICATIONS",
			"metadata": map[string]any{"event_notifications_instance_crn": "crn", "event_notifications_source_id": "src"},
		}}
	})
	res := ResourceIBMAppConfigIntegration()
	state := schematest.Apply(t, res, meta, nil, map[string]any{
		"guid": "guid", "integration_id": "en", "integration_type": "EVENT_NOTIFICATIONS",
		"metadata": []any{map[string]any{
			"event_notifications_instance_crn": "crn", "event_notifications_endpoint": "https://e",
			"event_notifications_source_name": "name",
		}},
	})
	metadata := (*calls)[0].Body["metadata"].(map[string]any)
	if metadata["event_notifications_source_name"] != "name" {
		t.Fatalf("metadata must be sent as one object: %v", (*calls)[0].Body)
	}
	if state.Attributes["metadata.0.event_notifications_source_name"] != "name" ||
		state.Attributes["metadata.0.event_notifications_source_id"] != "src" {
		t.Fatalf("source name must be kept and the source id read: %v", state.Attributes)
	}
	if again := schematest.Refresh(t, res, meta, state); again.Attributes["metadata.0.event_notifications_source_name"] != "name" {
		t.Fatal("refresh dropped the source name and would replace the integration")
	}
}

func TestMetadataMustMatchIntegrationType(t *testing.T) {
	meta, _ := schematest.NewMeta(t, func(schematest.Request) schematest.Reply {
		return schematest.Reply{Status: http.StatusOK, Body: map[string]any{}}
	})
	err := schematest.TryApply(t, ResourceIBMAppConfigIntegration(), meta, nil, map[string]any{
		"guid": "guid", "integration_id": "kms", "integration_type": "KMS",
		"metadata": []any{map[string]any{"kms_instance_crn": "crn"}},
	})
	if err == nil || !strings.Contains(err.Error(), "kms_endpoint, root_key_id") {
		t.Fatalf("expected the missing KMS settings to be reported, got %v", err)
	}
}
