// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package workflow

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestSchemasAreValid(t *testing.T) {
	schematest.ValidateDataSource(t, "workflow config", DataSourceIBMAppConfigWorkflowConfig())
	schematest.ValidateDataSource(t, "workflow configs", DataSourceIBMAppConfigWorkflowConfigs())
	schematest.ValidateResource(t, "workflow config", ResourceIBMAppConfigWorkflowConfig())
}
