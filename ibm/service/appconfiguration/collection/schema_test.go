// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package collection

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestSchemasAreValid(t *testing.T) {
	schematest.ValidateDataSource(t, "collection", DataSourceIBMAppConfigCollection())
	schematest.ValidateDataSource(t, "collections", DataSourceIBMAppConfigCollections())
	schematest.ValidateResource(t, "collection", ResourceIBMAppConfigCollection())
}
