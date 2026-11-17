// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package environment

import (
	"reflect"
	"strings"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestSchemasAreValid(t *testing.T) {
	for name, res := range map[string]*schema.Resource{
		"environment":  DataSourceIBMAppConfigEnvironment(),
		"environments": DataSourceIBMAppConfigEnvironments(),
	} {
		if err := res.InternalValidate(nil, false); err != nil {
			t.Errorf("%s: %v", name, err)
		}
	}
	if err := ResourceIBMAppConfigEnvironment().InternalValidate(nil, true); err != nil {
		t.Errorf("environment resource: %v", err)
	}
}

// jsonTags returns the json keys of a struct type.
func jsonTags(typ reflect.Type) map[string]bool {
	tags := map[string]bool{}
	for i := 0; i < typ.NumField(); i++ {
		if tag := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]; tag != "" && tag != "-" {
			tags[tag] = true
		}
	}
	return tags
}

// SetFromStruct matches response keys to schema keys by name, so computed keys must exist in the SDK structs.
func TestComputedKeysMatchSDK(t *testing.T) {
	inputs := map[string]bool{"guid": true, utils.EnvironmentID: true, "expand": true, "include": true, "sort": true, "tags_in": true, "limit": true, "offset": true}
	single := jsonTags(reflect.TypeOf(appconfigurationv1.Environment{}))
	for key, sch := range DataSourceIBMAppConfigEnvironment().Schema {
		if sch.Computed && !sch.Optional && !single[key] {
			t.Errorf("environment: computed key %q is not in the SDK response", key)
		}
	}
	for key := range ResourceIBMAppConfigEnvironment().Schema {
		if key != "guid" && !single[key] {
			t.Errorf("environment resource: key %q is not in the SDK response", key)
		}
	}
	list := jsonTags(reflect.TypeOf(appconfigurationv1.EnvironmentList{}))
	for key, sch := range DataSourceIBMAppConfigEnvironments().Schema {
		if sch.Computed && !sch.Optional && !inputs[key] && key != resourcesName && !list[key] {
			t.Errorf("environments: computed key %q is not in the SDK response", key)
		}
	}
	if !list["environments"] {
		t.Error("EnvironmentList no longer has an `environments` key")
	}
}
