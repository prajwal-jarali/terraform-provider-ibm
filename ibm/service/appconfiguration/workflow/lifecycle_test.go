// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package workflow

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestProviderConfigIsSentAsProvider(t *testing.T) {
	stored := map[string]any{}
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		for key, value := range r.Body {
			stored[key] = value
		}
		stored["workflow_id"] = "w1"
		return schematest.Reply{Status: http.StatusOK, Body: stored}
	})
	res := ResourceIBMAppConfigWorkflowConfig()
	state := schematest.Apply(t, res, meta, nil, map[string]any{
		"guid": "guid", "name": "W", "workflow_id": "w1", "enabled": true,
		"provider_config": []any{map[string]any{
			"type":     "SERVICENOW_IBM",
			"metadata": []any{map[string]any{"crn_mask": "mask", "sm_secret_crn": "crn", "service_environment": "production"}},
		}},
		"scope": []any{map[string]any{"collections": []any{map[string]any{"mode": "ALL"}}}},
	})
	body := (*calls)[0].Body
	provider, _ := body["provider"].(map[string]any)
	if provider["type"] != "SERVICENOW_IBM" || provider["metadata"].(map[string]any)["crn_mask"] != "mask" {
		t.Fatalf("provider_config must be sent as `provider`: %v", body)
	}
	if state.Attributes["provider_config.0.metadata.0.crn_mask"] != "mask" ||
		state.Attributes["scope.0.collections.0.mode"] != "ALL" {
		t.Fatalf("response was not stored: %v", state.Attributes)
	}
}
