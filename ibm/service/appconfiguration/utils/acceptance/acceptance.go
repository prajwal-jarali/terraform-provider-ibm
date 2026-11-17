// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package acceptance holds what the acceptance tests of the App Configuration resources share:
// the checks of the credentials, the configuration that points the tests at an existing instance,
// and the checks that read the real API.
//
// The tests run against an App Configuration instance that already exists, so they need:
//
//	IC_API_KEY            the API key of the account
//	IBM_APPCONFIG_GUID    the GUID of the App Configuration instance
//
// Optional: IC_REGION, and IBMCLOUD_APP_CONFIG_API_ENDPOINT for a non default endpoint.
package acceptance

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	apiKeyVar = "IC_API_KEY"
	guidVar   = "IBM_APPCONFIG_GUID"
)

var configureProvider sync.Once

// PreCheck fails the test when the credentials or the instance GUID are missing,
// and configures the provider that the checks of this package use.
func PreCheck(t *testing.T) {
	t.Helper()
	for _, name := range []string{apiKeyVar, guidVar} {
		if os.Getenv(name) == "" {
			t.Fatalf("%s must be set for the App Configuration acceptance tests", name)
		}
	}
	configureProvider.Do(func() {
		if diags := acc.TestAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil)); diags.HasError() {
			t.Fatalf("configuring provider: %s", diags[0].Summary)
		}
	})
}

// GUID returns the GUID of the instance the tests run in.
func GUID() string {
	return os.Getenv(guidVar)
}

// GUIDConfig returns the config that every test starts with: the local value `guid`, the GUID
// of the instance the resources of the test are created in.
func GUIDConfig() string {
	return fmt.Sprintf(`
		locals {
			guid = "%s"
		}`, GUID())
}

// Getter fetches one resource by the ids that follow the GUID in its Terraform id.
// It returns the response so that a 404 can be told from another failure.
type Getter func(client *appconfigurationv1.AppConfigurationV1, ids []string) (*core.DetailedResponse, error)

func fetch(id string, get Getter) (*core.DetailedResponse, error) {
	// The id of a resource that is one per instance, like the origin config, is only the GUID.
	parts := strings.Split(id, "/")
	client, err := utils.NewClient(acc.TestAccProvider.Meta(), parts[0])
	if err != nil {
		return nil, err
	}
	return get(client, parts[1:])
}

// CheckExists fails when the resource in the state cannot be fetched from the API.
func CheckExists(name string, get Getter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return flex.FmtErrorf("Not found: %s", name)
		}
		if _, err := fetch(rs.Primary.ID, get); err != nil {
			return flex.FmtErrorf("%s does not exist: %s", name, err)
		}
		return nil
	}
}

// CheckDestroy fails when a resource of the given type still exists after the destroy.
func CheckDestroy(resourceType string, get Getter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			response, err := fetch(rs.Primary.ID, get)
			if err == nil {
				return flex.FmtErrorf("%s still exists: %s", resourceType, rs.Primary.ID)
			}
			if response == nil || response.StatusCode != 404 {
				return flex.FmtErrorf("Error checking that %s (%s) is destroyed: %s", resourceType, rs.Primary.ID, err)
			}
		}
		return nil
	}
}

// EnvironmentConfig returns the GUID and the environment `ibm_app_config_environment.test` that the
// features and properties of a test live in.
func EnvironmentConfig() string {
	return GUIDConfig() + `
		resource "ibm_app_config_environment" "test" {
			guid           = local.guid
			environment_id = "tf_env"
			name           = "Terraform test"
		}`
}

// FeatureConfig returns the environment `ibm_app_config_environment.test` and a feature flag
// `ibm_app_config_feature.test` in it.
func FeatureConfig(featureID string) string {
	return EnvironmentConfig() + fmt.Sprintf(`
		resource "ibm_app_config_feature" "test" {
			guid           = local.guid
			environment_id = ibm_app_config_environment.test.environment_id
			feature_id     = "%s"
			name           = "Checkout"
			type           = "BOOLEAN"
			enabled_value  = "true"
			disabled_value = "false"
		}`, featureID)
}
