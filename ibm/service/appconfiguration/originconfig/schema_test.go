// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package originconfig

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestSchemasAreValid(t *testing.T) {
	schematest.ValidateDataSource(t, "origin config", DataSourceIBMAppConfigOriginConfig())
	schematest.ValidateResource(t, "origin config", ResourceIBMAppConfigOriginConfig())
}
