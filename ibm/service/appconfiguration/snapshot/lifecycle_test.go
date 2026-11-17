// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package snapshot

import (
	"net/http"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestCreateKeepsDetailsOfCollectionAndEnvironment(t *testing.T) {
	meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
		body := map[string]any{"git_config_id": "g1", "git_config_name": "G", "git_url": "u", "git_branch": "b", "git_file_path": "f.json"}
		if r.Method == http.MethodGet {
			body["collection"] = map[string]any{"collection_id": "c1", "name": "C"}
			body["environment"] = map[string]any{"environment_id": "e1", "name": "E", "color_code": "#FF0000"}
		}
		return schematest.Reply{Status: http.StatusOK, Body: body}
	})
	res := ResourceIBMAppConfigSnapshot()
	state := schematest.Apply(t, res, meta, nil, map[string]any{
		"guid": "guid", "git_config_id": "g1", "git_config_name": "G", "git_url": "u", "git_branch": "b",
		"git_file_path": "f.json", "git_token": "secret", "collection_id": "c1", "environment_id": "e1",
	})
	if (*calls)[0].Body["git_token"] != "secret" || (*calls)[0].Body["collection_id"] != "c1" {
		t.Fatalf("create must send the token and the ids: %v", (*calls)[0].Body)
	}
	state = schematest.Refresh(t, res, meta, state)
	if state.Attributes["collection.0.name"] != "C" || state.Attributes["git_token"] != "secret" {
		t.Fatalf("refresh must add the details and keep the token: %v", state.Attributes)
	}
}
