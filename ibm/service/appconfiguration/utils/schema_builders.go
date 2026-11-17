// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Mode tells whether a schema describes what the user writes (Input, a resource) or what the
// API reports (Output, a data source). One builder function can then serve both: the same field
// is required in a resource and computed in a data source.
type Mode bool

const (
	Input  Mode = true
	Output Mode = false
)

// Required makes the field required for the user in Input mode, and computed in Output mode.
func (m Mode) Required(field *schema.Schema) *schema.Schema {
	if m == Input {
		return AsRequired(field)
	}
	return AsComputed(field)
}

// Optional makes the field optional for the user in Input mode, and computed in Output mode.
func (m Mode) Optional(field *schema.Schema) *schema.Schema {
	if m == Input {
		return AsOptional(field)
	}
	return AsComputed(field)
}

// OptionalComputed makes the field optional for the user in Input mode, keeping the value the
// API picks when the user sets none, and computed in Output mode.
func (m Mode) OptionalComputed(field *schema.Schema) *schema.Schema {
	if m == Input {
		return AsOptionalComputed(field)
	}
	return AsComputed(field)
}

// RequiredForceNew is Required for a field that cannot be changed after creation.
func (m Mode) RequiredForceNew(field *schema.Schema) *schema.Schema {
	if m == Input {
		return AsForceNew(AsRequired(field))
	}
	return AsComputed(field)
}

// OptionalForceNew is Optional for a field that cannot be changed after creation.
func (m Mode) OptionalForceNew(field *schema.Schema) *schema.Schema {
	if m == Input {
		return AsForceNew(AsOptional(field))
	}
	return AsComputed(field)
}

// Plain field builders. They set the type and the description only. The caller decides whether
// the field is an input or an output, normally with a Mode.

func NewString(description string) *schema.Schema {
	return &schema.Schema{Type: schema.TypeString, Description: description}
}

func NewInt(description string) *schema.Schema {
	return &schema.Schema{Type: schema.TypeInt, Description: description}
}

func NewFloat(description string) *schema.Schema {
	return &schema.Schema{Type: schema.TypeFloat, Description: description}
}

func NewBool(description string) *schema.Schema {
	return &schema.Schema{Type: schema.TypeBool, Description: description}
}

func NewStringList(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: description,
	}
}

// NewEnum builds a string field that accepts only the given values.
func NewEnum(description string, values ...string) *schema.Schema {
	field := NewString(withValues(description, values))
	field.ValidateDiagFunc = validation.ToDiagFunc(validation.StringInSlice(values, false))
	return field
}

// NewEnumList builds a list of strings that accepts only the given values.
func NewEnumList(description string, values ...string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: withValues(description, values),
		Elem: &schema.Schema{
			Type:             schema.TypeString,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(values, false)),
		},
	}
}

// NewList builds a list of nested objects.
func NewList(description string, elem map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: description,
		Elem:        &schema.Resource{Schema: elem},
	}
}

// NewObject builds a single nested object. Terraform models it as a list with one item.
func NewObject(description string, elem map[string]*schema.Schema) *schema.Schema {
	field := NewList(description, elem)
	field.MaxItems = 1
	return field
}

func withValues(description string, values []string) string {
	return fmt.Sprintf("%s Supported values are %s.", description, strings.Join(values, ", "))
}

// Modifiers return a modified copy, so a field can be declared once and reused.

// AsOptionalComputed returns a copy of a field that the user may set and
// that the API fills with a default when the user does not.
func AsOptionalComputed(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.Required, clone.Optional, clone.Computed = false, true, true
	return &clone
}

// AsComputed returns a copy of a field that only the API sets.
// Settings that only apply to user input are dropped, because the SDK rejects them on computed fields.
func AsComputed(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.Required, clone.Optional, clone.Computed = false, false, true
	clone.ForceNew, clone.Sensitive = false, false
	clone.MaxItems, clone.MinItems = 0, 0
	clone.ValidateFunc, clone.ValidateDiagFunc = nil, nil
	return &clone
}

// AsSensitive returns a copy of a field whose value is hidden in the plan output.
func AsSensitive(field *schema.Schema) *schema.Schema {
	clone := *field
	clone.Sensitive = true
	return &clone
}
