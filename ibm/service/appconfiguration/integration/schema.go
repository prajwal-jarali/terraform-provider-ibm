// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package integration holds the schemas of the integration resource and data sources.
// An integration connects the instance to Event Notifications or to Key Protect.
// The API has no update call, so every argument of the resource forces a new one.
package integration

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	eventNotificationsEndpoint          = "event_notifications_endpoint"
	eventNotificationsInstanceCrn       = "event_notifications_instance_crn"
	eventNotificationsSourceDescription = "event_notifications_source_description"
	eventNotificationsSourceId          = "event_notifications_source_id"
	eventNotificationsSourceName        = "event_notifications_source_name"
	integrationType                     = "integration_type"
	keyStatus                           = "key_status"
	kmsEndpoint                         = "kms_endpoint"
	kmsInstanceCrn                      = "kms_instance_crn"
	kmsSchemeType                       = "kms_scheme_type"
	rootKeyId                           = "root_key_id"
)

const (
	resourceName  = "integration"
	resourcesName = "integrations"
)

const (
	typeEventNotifications = "EVENT_NOTIFICATIONS"
	typeKMS                = "KMS"
)

var types = []string{typeEventNotifications, typeKMS}

// requiredMetadata lists the settings that an integration of each type needs.
var requiredMetadata = map[string][]string{
	typeEventNotifications: {eventNotificationsInstanceCrn, eventNotificationsEndpoint, eventNotificationsSourceName},
	typeKMS:                {kmsInstanceCrn, kmsEndpoint, rootKeyId},
}

// getMetadata returns the settings of an integration. Which ones apply depends on the integration type:
// the `event_notifications_*` ones for EVENT_NOTIFICATIONS and the `kms_*` ones and `root_key_id` for KMS.
func getMetadata(mode utils.Mode) *schema.Schema {
	text := func(description string) *schema.Schema { return mode.Optional(utils.NewString(description)) }
	fields := map[string]*schema.Schema{
		eventNotificationsInstanceCrn: text("The CRN of the Event Notifications service instance."),
		eventNotificationsEndpoint:    text("The URL endpoint of the Event Notifications service instance."),
		kmsInstanceCrn:                text("The CRN of the Key Protect service instance."),
		kmsEndpoint:                   text("The URL endpoint of the Key Protect instance."),
		rootKeyId:                     text("The root key id of the Key Protect instance."),
	}
	// The API takes the source name and description on creation, and returns the source id and the key
	// status instead. The returned ones are outputs, also in a resource.
	fields[eventNotificationsSourceId] = utils.AsComputed(utils.NewString("Event Notifications source id."))
	fields[keyStatus] = utils.AsComputed(utils.NewString("Key Protect instance root key status."))
	fields[kmsSchemeType] = utils.AsComputed(utils.NewString(`For the service name "kms" the schema type is "BYOK".`))
	if mode == utils.Input {
		fields[eventNotificationsSourceName] = text(
			"Source name. This name is shown in the sources page of your Event Notifications instance.",
		)
		fields[eventNotificationsSourceDescription] = text(
			"Source description. This description is shown in the sources page of your Event Notifications instance.",
		)
	}
	field := utils.NewObject("Settings of the integration, which depend on the integration type.", fields)
	return mode.RequiredForceNew(field)
}

func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetAuditFields(resourceName),
		map[string]*schema.Schema{
			utils.IntegrationID: mode.RequiredForceNew(utils.NewString(
				"Integration id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			)),
			integrationType: mode.RequiredForceNew(utils.NewEnum("Integration type.", types...)),
			utils.Metadata:  getMetadata(mode),
		},
	)
}

var idKeys = []string{utils.IntegrationID}

// DataSourceIBMAppConfigIntegration returns the data source of one integration.
func DataSourceIBMAppConfigIntegration() *schema.Resource {
	return utils.NewDataSource("Integration", idKeys, getFields(utils.Output), readIntegration)
}

// DataSourceIBMAppConfigIntegrations returns the `ibm_app_config_integrations` data source.
func DataSourceIBMAppConfigIntegrations() *schema.Resource {
	return utils.NewListDataSource(resourcesName, "Array of integrations.", getFields(utils.Output), readIntegrations,
		utils.GetPageQuery(),
		map[string]*schema.Schema{utils.Expand: utils.GetExpandSchema()},
	)
}

// ResourceIBMAppConfigIntegration returns the integration resource. The API has no call to update an integration,
// so every field of it forces a new one.
func ResourceIBMAppConfigIntegration() *schema.Resource {
	res := utils.NewResource("Integration", idKeys, getFields(utils.Input), utils.Calls{
		Create: createIntegration,
		Get:    getIntegration,
		Delete: deleteIntegration,
	})
	res.CustomizeDiff = validateMetadata
	return res
}

// validateMetadata fails the plan when the metadata lacks a setting that the integration type needs.
// A value that is not known yet, e.g. the output of another resource, is checked when it is known.
func validateMetadata(_ context.Context, diff *schema.ResourceDiff, _ any) error {
	if !diff.NewValueKnown(integrationType) {
		return nil
	}
	var missing []string
	for _, key := range requiredMetadata[diff.Get(integrationType).(string)] {
		path := fmt.Sprintf("%s.0.%s", utils.Metadata, key)
		if diff.NewValueKnown(path) && diff.Get(path).(string) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("metadata of an integration of type %s requires: %s",
			diff.Get(integrationType), strings.Join(missing, ", "))
	}
	return nil
}
