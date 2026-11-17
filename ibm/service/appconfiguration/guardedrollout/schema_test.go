// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package guardedrollout

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils/schematest"
)

func TestSchemasAreValid(t *testing.T) {
	schematest.ValidateDataSource(t, "guarded rollout", DataSourceIBMAppConfigGuardedRollout())
	schematest.ValidateDataSource(t, "guarded rollouts", DataSourceIBMAppConfigGuardedRollouts())
	schematest.ValidateDataSource(t, "guarded rollout metric", DataSourceIBMAppConfigGuardedRolloutMetric())
}
