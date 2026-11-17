// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Field names of the API that only this package uses.
const (
	approvalUrl         = "approval_url"
	changeRequestNumber = "change_request_number"
	changeRequestStatus = "change_request_status"
	executionStatus     = "execution_status"
	lastPollTime        = "last_poll_time"
	providerType        = "provider_type"
	resourceId          = "resource_id"
	resourceType        = "resource_type"
	workflowName        = "workflow_name"
)

// Schema parts shared by the resources and data sources of the App Configuration API.
//
// A package describes its fields once, with a Mode, and passes them to the constructors in operations.go:
//
//	NewDataSource(kind, idKeys, getFields(Output), get, inputs...)
//	NewListDataSource(itemsKey, description, getFields(Output), list, inputs...)
//	NewResource(kind, idKeys, getFields(Input), calls, parentIDs...)

// defaultSortValues are the attributes that most list calls can be sorted by.
var defaultSortValues = []string{CreatedTime, "id", Name, UpdatedTime}

// GetAuditFields returns the fields the API reports for most resources: when it was
// created, when it was last modified and where it can be fetched.
func GetAuditFields(resource string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		CreatedTime: AsComputed(NewString(fmt.Sprintf("Creation time of the %s.", resource))),
		UpdatedTime: AsComputed(NewString(fmt.Sprintf("Last modified time of the %s.", resource))),
		Href:        AsComputed(NewString(fmt.Sprintf("%s URL.", resource))),
	}
}

// GetIdentityFields returns the name, the description and the tags, which most resources have.
// The name is required and can be changed. Allowed special characters in names are dot, hyphen and underscore.
func GetIdentityFields(mode Mode, resource string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		Name:        mode.Required(NewString(fmt.Sprintf("The %s name.", resource))),
		Description: mode.Optional(NewString(fmt.Sprintf("The %s description.", resource))),
		Tags: mode.Optional(NewString(fmt.Sprintf(
			"Tags associated with the %s. Allowed special characters are [_. ,-:].", resource,
		))),
	}
}

// GetParentIDSchema returns the id of the resource that another resource lives under, e.g. the
// environment of a feature. The id is part of the API path, so it cannot be changed.
func GetParentIDSchema(description string) *schema.Schema {
	return AsForceNew(AsRequired(NewString(description)))
}

// GetAssociatedSchema returns the list of features, properties or snapshots that the API
// reports for a resource. `idKey` is the name of the id field of the associated item.
func GetAssociatedSchema(kind, idKey, owner string) *schema.Schema {
	return AsComputed(NewList(
		fmt.Sprintf("List of %ss associated with this %s.", kind, owner),
		map[string]*schema.Schema{
			Name:  AsComputed(NewString(fmt.Sprintf("%s name.", kind))),
			idKey: AsComputed(NewString(fmt.Sprintf("%s Id.", kind))),
		},
	))
}

// GetWorkflowApprovalSchema returns the approval details the API attaches to a resource while
// a change request for it is open.
func GetWorkflowApprovalSchema() *schema.Schema {
	text := func(description string) *schema.Schema { return AsComputed(NewString(description)) }
	return AsComputed(NewObject(
		"Workflow approval information. "+
			"Present only when the workflow is enabled and the resource has an active change request.",
		map[string]*schema.Schema{
			workflowName:        text("Name of the workflow configuration."),
			WorkflowID:          text("Unique identifier of the workflow configuration."),
			providerType:        text("Type of workflow provider."),
			changeRequestNumber: text("ServiceNow change request number."),
			changeRequestStatus: text("Status of the change request in ServiceNow."),
			executionStatus:     text("Execution status of the workflow."),
			approvalUrl:         text("URL to view the change request in ServiceNow."),
			resourceType:        text("Type of resource."),
			resourceId:          text("Unique identifier of the resource. Not present for environments."),
			EnvironmentID:       text("Environment identifier. Present only for features, properties and environments."),
			CreatedTime:         text("Timestamp when the workflow approval was created."),
			UpdatedTime:         text("Timestamp when the workflow approval was last updated."),
			lastPollTime:        text("Timestamp when the workflow approval status was last polled from ServiceNow."),
		},
	))
}

// Query parameters of the get and list data sources.

// GetIncludeValuesSchema returns the `include` query parameter with the details that can be included.
func GetIncludeValuesSchema(values ...string) *schema.Schema {
	return AsOptional(NewEnumList("Additional details to include in the response.", values...))
}

// GetSortValuesSchema returns the `sort` query parameter. Without `values` the common sort keys are used.
func GetSortValuesSchema(resource string, values ...string) *schema.Schema {
	if len(values) == 0 {
		values = defaultSortValues
	}
	return AsOptional(NewEnum(fmt.Sprintf(
		"Sort the %s details based on the specified attribute. By default, items are sorted by name.",
		resource,
	), values...))
}

// GetSearchSchema returns the `search` query parameter. `fields` names what is searched, e.g. "Name OR Tag".
func GetSearchSchema(fields string) *schema.Schema {
	return AsOptional(NewString(fmt.Sprintf(
		"Searches for the provided keyword and returns the appropriate row with that value. "+
			"Here the search happens on the '[%s]' of the entity.",
		fields,
	)))
}

// GetFilterSchema returns a query parameter that keeps only the items related to the given ids.
func GetFilterSchema(description string) *schema.Schema {
	return AsOptional(NewStringList(description))
}

// GetPageQuery returns the `limit` and `offset` query parameters of the list data sources.
func GetPageQuery() map[string]*schema.Schema {
	return map[string]*schema.Schema{Limit: GetLimitSchema(), Offset: GetOffsetSchema()}
}

// GetPaginationFields returns the fields of a list response that describe the page:
// the total number of records and the links to the other pages.
func GetPaginationFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		TotalCount: GetTotalCountSchema(),
		First:      GetUrlNavSchema(First),
		Previous:   GetUrlNavSchema(Previous),
		Next:       GetUrlNavSchema(Next),
		Last:       GetUrlNavSchema(Last),
	}
}

// Schema builders used by the constructors. `inputs` are the extra arguments of the call,
// e.g. parent ids and query parameters.

// getDataSourceSchema builds the schema of a data source that fetches one item. The fields of the item
// are outputs, except the `idKeys`, which identify the item and are required.
func getDataSourceSchema(
	fields map[string]*schema.Schema,
	idKeys []string,
	inputs ...map[string]*schema.Schema,
) map[string]*schema.Schema {
	sch := Merge(append([]map[string]*schema.Schema{fields, {GUID: GetGuidSchema()}}, inputs...)...)
	for _, key := range idKeys {
		sch[key] = AsRequired(sch[key])
	}
	return sch
}

// getListSchema builds the schema of a data source that lists items: the GUID, the extra arguments,
// the list of items under `itemsKey` and the pagination fields of the response.
func getListSchema(
	itemsKey, description string,
	item map[string]*schema.Schema,
	inputs ...map[string]*schema.Schema,
) map[string]*schema.Schema {
	list := map[string]*schema.Schema{
		GUID:     GetGuidSchema(),
		itemsKey: AsComputed(NewList(description, item)),
	}
	return Merge(append([]map[string]*schema.Schema{list, GetPaginationFields()}, inputs...)...)
}

// getResourceSchema builds the schema of a resource from its fields. `parents` are the ids of the
// resources it lives under.
func getResourceSchema(
	fields map[string]*schema.Schema,
	parents ...map[string]*schema.Schema,
) map[string]*schema.Schema {
	return Merge(append([]map[string]*schema.Schema{fields, {GUID: GetGuidSchema()}}, parents...)...)
}
