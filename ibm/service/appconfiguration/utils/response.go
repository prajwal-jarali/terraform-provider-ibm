// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SetFromStruct sets every schema field that is present in an SDK result.
//
// The result is converted to JSON, so a schema key is filled when it matches a `json` tag of the
// result struct. Nested lists and objects follow the same rule. Fields missing from the result
// (nil or omitted) are left untouched. See SetFields for how values are fitted to the schema.
func SetFromStruct(data *schema.ResourceData, sch map[string]*schema.Schema, result any) error {
	fields, err := FieldMap(result)
	if err != nil {
		return err
	}
	return SetFields(data, sch, fields)
}

// SetFields sets every schema field that has a value in `fields`, which is keyed like the API response.
// Values are fitted to the schema first: keys the schema does not have are dropped, an object
// is wrapped into a one item list, and a value that is not a string is turned into one
// (a JSON value of a feature, for example). A computed list the response lacks is set to an empty list. Use it to rename or adjust response fields before setting them.
func SetFields(data *schema.ResourceData, sch map[string]*schema.Schema, fields map[string]any) error {
	var errs []error
	for key, field := range sch {
		value := fields[key]
		if value == nil {
			// The API leaves out a list that is empty or not asked for with `expand`. A computed list that is
			// not in the state is planned as unknown on every run, so it is stored as an empty list.
			if isComputedList(field) {
				value = []any{}
			} else {
				continue
			}
		}
		if err := data.Set(key, fit(value, field)); err != nil {
			errs = append(errs, flex.FmtErrorf("Error while setting field: %s: %w", key, err))
		}
	}
	return errors.Join(errs...)
}

// isComputedList tells a list or set that only the API sets.
func isComputedList(field *schema.Schema) bool {
	return field.Computed && !field.Optional && (field.Type == schema.TypeList || field.Type == schema.TypeSet)
}

// FieldMap converts a struct to a map keyed by its `json` tags.
func FieldMap(result any) (map[string]any, error) {
	raw, err := json.Marshal(result)
	if err != nil {
		return nil, errors.Join(flex.FmtErrorf("Error while reading response"), err)
	}
	// UseNumber keeps large integers exact and rejects fractions for integer fields.
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	fields := map[string]any{}
	if err := decoder.Decode(&fields); err != nil {
		return nil, errors.Join(flex.FmtErrorf("Error while reading response"), err)
	}
	return fields, nil
}

// fit shapes a response value like its schema field.
func fit(value any, field *schema.Schema) any {
	if field.Type == schema.TypeString {
		return toText(value)
	}
	nested, ok := field.Elem.(*schema.Resource)
	if !ok || (field.Type != schema.TypeList && field.Type != schema.TypeSet) {
		return value
	}
	if object, isObject := value.(map[string]any); isObject {
		value = []any{object}
	}
	items, isList := value.([]any)
	if !isList {
		return value
	}
	fitted := make([]any, len(items))
	for i, item := range items {
		fitted[i] = item
		if object, isObject := item.(map[string]any); isObject {
			fitted[i] = fitObject(object, nested.Schema)
		}
	}
	return fitted
}

func fitObject(object map[string]any, sch map[string]*schema.Schema) map[string]any {
	fitted := map[string]any{}
	for key, field := range sch {
		if value := object[key]; value != nil {
			fitted[key] = fit(value, field)
		}
	}
	return fitted
}

// toText returns a value as a string: JSON for objects and lists, the plain text for the rest.
func toText(value any) any {
	switch typed := value.(type) {
	case string, nil:
		return value
	case map[string]any, []any:
		encoded, err := json.Marshal(typed)
		if err != nil {
			return value
		}
		return string(encoded)
	default:
		return fmt.Sprint(typed)
	}
}
