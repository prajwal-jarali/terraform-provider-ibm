// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"fmt"
	"maps"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Builders for the fields that every App Configuration resource and data source has.
// The `resource` argument is the singular name of the resource, e.g. "environment".

func computedString(description string) *schema.Schema {
	return &schema.Schema{Type: schema.TypeString, Computed: true, Description: description}
}

func GetGuidSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
		Description: "GUID of the App Configuration service. " +
			"Get it from the service instance credentials section of the dashboard.",
	}
}

func GetOptionalTagsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Description: "Filter the resources to be returned based on the associated tags. " +
			"Specify the parameter as a list of comma separated tags. " +
			"Returns resources associated with any of the specified tags.",
	}
}

func GetExpandSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If set to `true`, returns expanded view of the resource details.",
	}
}

func GetLimitSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Description: "The number of records to retrieve. By default, the list operation return the first 10 records. " +
			"To retrieve different set of records, use `limit` with `offset` to page through the available records.",
	}
}

func GetOffsetSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Description: "The number of records to skip. " +
			"By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. " +
			"Use `offset` with `limit` to page through the available records.",
	}
}

func GetTotalCountSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Total number of records.",
	}
}

// GetUrlNavSchema returns a pagination link. `direction` is one of first, previous, next, last.
func GetUrlNavSchema(direction string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("URL to navigate to the %s list of records.", direction),
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			Href: computedString("URL of the response."),
		}},
	}
}

// AsRequired returns a copy of a schema field that the user must set.
func AsRequired(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.Required, clone.Optional, clone.Computed = true, false, false
	return &clone
}

// AsOptional returns a copy of a schema field that the user may set.
func AsOptional(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.Required, clone.Optional, clone.Computed = false, true, false
	return &clone
}

// AsForceNew returns a copy of a schema field that recreates the resource when it changes.
func AsForceNew(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.ForceNew = true
	return &clone
}

// Merge combines schema maps into a new one. Later maps win on duplicate keys.
func Merge(parts ...map[string]*schema.Schema) map[string]*schema.Schema {
	merged := map[string]*schema.Schema{}
	for _, part := range parts {
		maps.Copy(merged, part)
	}
	return merged
}
