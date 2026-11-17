// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package utils holds the helpers shared by the App Configuration resources and data sources:
// reusable schema fields, conversion between schema data and SDK structs, and client setup.
package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const endpointEnvVar = "IBMCLOUD_APP_CONFIG_API_ENDPOINT"

// GetAppConfigClient returns the App Configuration client pointed at the instance
// whose GUID is stored in the data.
func GetAppConfigClient(meta any, data *schema.ResourceData) (*appconfigurationv1.AppConfigurationV1, error) {
	return NewClient(meta, data.Get(GUID).(string))
}

// NewClient returns the App Configuration client pointed at the instance with the given GUID.
// The endpoint can be overridden with IBMCLOUD_APP_CONFIG_API_ENDPOINT.
func NewClient(meta any, guid string) (*appconfigurationv1.AppConfigurationV1, error) {
	guid = strings.TrimSpace(guid)
	if guid == "" {
		return nil, flex.FmtErrorf("Empty|Blank string has been passed as GUID")
	}
	session := meta.(conns.ClientSession)
	bluemixSession, err := session.BluemixSession()
	if err != nil {
		return nil, err
	}
	client, err := session.AppConfigurationV1()
	if err != nil {
		return nil, err
	}
	defaultEndpoint := fmt.Sprintf("https://%s.apprapp.cloud.ibm.com", bluemixSession.Config.Region)
	endpoint := conns.EnvFallBack([]string{endpointEnvVar}, defaultEndpoint)
	// The client is shared by all calls, so the instance URL is set on every request.
	client.Service.Options.URL = fmt.Sprintf("%s/apprapp/feature/v1/instances/%s", endpoint, guid)
	return client, nil
}

// GetClientAndIds parses the resource id into its parts, which also restores the GUID in the data,
// and returns the client for that instance. See ParseID for the id format.
func GetClientAndIds(data *schema.ResourceData, meta any, parts int) (*appconfigurationv1.AppConfigurationV1, []string, error) {
	ids, err := ParseID(data, parts)
	if err != nil {
		return nil, nil, err
	}
	client, err := GetAppConfigClient(meta, data)
	if err != nil {
		return nil, nil, err
	}
	return client, ids, nil
}

// IsNotFound reports whether an API response is a 404.
func IsNotFound(response *core.DetailedResponse) bool {
	return response != nil && response.StatusCode == http.StatusNotFound
}
