// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package schematest helps to test the resources and data sources: it validates their schemas and
// runs them against a fake App Configuration API.
package schematest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ValidateDataSource fails the test if the schema is not valid for a data source.
func ValidateDataSource(t *testing.T, name string, res *schema.Resource) {
	t.Helper()
	if err := res.InternalValidate(nil, false); err != nil {
		t.Errorf("data source %s: %v", name, err)
	}
}

// ValidateResource fails the test if the schema is not valid for a resource.
func ValidateResource(t *testing.T, name string, res *schema.Resource) {
	t.Helper()
	if err := res.InternalValidate(nil, true); err != nil {
		t.Errorf("resource %s: %v", name, err)
	}
}
