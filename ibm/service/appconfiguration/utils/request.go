// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"slices"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Guard decides whether a schema field should be included in a request.
type Guard func(data *schema.ResourceData, field string) bool

// configValue returns the value the user wrote for a field.
// It returns false when there is no raw config, the field is not in the schema,
// or the value (or anything nested in it) is not known yet. It never panics.
func configValue(data *schema.ResourceData, field string) (cty.Value, bool) {
	raw := data.GetRawConfig()
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute(field) {
		return cty.NilVal, false
	}
	value := raw.GetAttr(field)
	if !value.IsWhollyKnown() {
		return cty.NilVal, false
	}
	return value, true
}

// ForRead includes a field only if the user wrote it in the config and its value is fully known.
// This tells `expand = false` (sent) apart from an omitted `expand` (not sent).
// Lists, sets, maps and nested blocks count as written when present, even if empty.
func ForRead(data *schema.ResourceData, field string) bool {
	value, ok := configValue(data, field)
	return ok && !value.IsNull()
}

// ForCreate includes a field only if the user wrote it in the config, same as ForRead.
func ForCreate(data *schema.ResourceData, field string) bool {
	return ForRead(data, field)
}

// ForUpdate includes a field only if it changed and is still set in the config.
// Changing a value to its zero value is a change and is sent.
// Removing a field from the config is skipped, i.e. the API keeps its current value.
//
// A field that has no value in state and is set to its zero value in the config is not a change
// for the SDK (HasChange is false), so it is not sent either.
func ForUpdate(data *schema.ResourceData, field string) bool {
	return data.HasChange(field) && ForRead(data, field)
}

// FillOptions copies the schema fields that pass the guard into an SDK options struct.
//
// Values are taken from the raw config, so only what the user wrote is sent: omitted nested
// attributes stay nil and explicit zero values (false, 0, "") are kept.
// A field reaches the struct only if its schema key matches the `json` tag of a struct field.
// Use pointer fields (as the SDK does) so that an explicit zero value is distinguishable from unset.
// Output-only (computed, not optional) schema fields are never copied.
//
// A list block with MaxItems 1 is sent as a single object, which is how the SDK models it.
func FillOptions(data *schema.ResourceData, sch map[string]*schema.Schema, guard Guard, options any) error {
	fields := map[string]json.RawMessage{}
	for key, field := range sch {
		if isOutputOnly(field) || !guard(data, key) {
			continue
		}
		encoded, err := encodeField(data, key, field)
		if err != nil {
			return errors.Join(flex.FmtErrorf("Error while building request options, field: %s", key), err)
		}
		if encoded != nil {
			fields[key] = encoded
		}
	}
	return decodeInto(fields, options)
}

// encodeField returns the JSON of a field as written in the config, or nil when there is nothing to send.
// Nested objects, which Terraform models as lists of one item, are sent as single objects.
func encodeField(data *schema.ResourceData, key string, field *schema.Schema) (json.RawMessage, error) {
	value, ok := configValue(data, key)
	if !ok || value.IsNull() {
		return nil, nil
	}
	encoded, err := ctyjson.Marshal(value, value.Type())
	if err != nil {
		return nil, err
	}
	var decoded any
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()
	if err := decoder.Decode(&decoded); err != nil {
		return nil, err
	}
	decoded = unwrapObjects(decoded, field)
	if decoded == nil {
		return nil, nil
	}
	return json.Marshal(decoded)
}

// unwrapObjects replaces every list with MaxItems 1 by its only item, at any depth.
func unwrapObjects(value any, field *schema.Schema) any {
	nested, ok := field.Elem.(*schema.Resource)
	items, isList := value.([]any)
	if !ok || !isList {
		return value
	}
	for _, item := range items {
		if object, isObject := item.(map[string]any); isObject {
			for key, child := range nested.Schema {
				if _, found := object[key]; found {
					object[key] = unwrapObjects(object[key], child)
				}
			}
		}
	}
	if field.MaxItems != 1 {
		return items
	}
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

// DecodeField fills a struct from one schema field. It serves the request fields that the SDK
// models as interfaces, which FillOptions cannot fill; use Skip to leave them out of FillOptions.
func DecodeField(data *schema.ResourceData, sch map[string]*schema.Schema, key string, target any) error {
	encoded, err := encodeField(data, key, sch[key])
	if err == nil && encoded != nil {
		err = json.Unmarshal(encoded, target)
	}
	if err != nil {
		return errors.Join(flex.FmtErrorf("Error while building request options, field: %s", key), err)
	}
	return nil
}

// Skip returns a guard that leaves out the given fields.
func Skip(guard Guard, fields ...string) Guard {
	return func(data *schema.ResourceData, field string) bool {
		return !slices.Contains(fields, field) && guard(data, field)
	}
}

// decodeInto fills a struct from JSON fields, matching them by `json` tag.
func decodeInto(fields map[string]json.RawMessage, options any) error {
	raw, err := json.Marshal(fields)
	if err == nil {
		err = json.Unmarshal(raw, options)
	}
	if err != nil {
		return errors.Join(flex.FmtErrorf("Error while building request options"), err)
	}
	return nil
}

func isOutputOnly(field *schema.Schema) bool {
	return field.Computed && !field.Optional && !field.Required
}
