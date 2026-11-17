// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"gopkg.in/yaml.v3"
)

// ParseValue converts the text a user wrote into the value the API expects for a feature or
// property of the given type and format.
func ParseValue(typ, format, text string) (any, error) {
	switch typ {
	case "BOOLEAN":
		return parse(strconv.ParseBool(text))
	case "NUMERIC":
		return parse(strconv.ParseInt(text, 10, 64))
	case "STRING":
		return parseString(format, text)
	case "SECRETREF":
		return parseString("JSON", text)
	}
	return nil, flex.FmtErrorf("invalid type %q, expected BOOLEAN, STRING, NUMERIC or SECRETREF", typ)
}

func parseString(format, text string) (any, error) {
	var parsed map[string]any
	switch format {
	case "TEXT", "":
		return text, nil
	case "JSON":
		return parse(parsed, json.Unmarshal([]byte(text), &parsed))
	case "YAML":
		return parse(parsed, yaml.Unmarshal([]byte(text), &parsed))
	}
	return nil, flex.FmtErrorf("invalid format %q, expected TEXT, JSON or YAML", format)
}

func parse[T any](value T, err error) (any, error) {
	if err != nil {
		return nil, flex.FmtErrorf("value does not match the type or format: %s", err)
	}
	return value, nil
}

// DecodeRules fills a list of targeting rules from a field, if the guard accepts it. The user writes the
// value of each rule as text, so it is converted like the value of the feature or property it belongs to.
func DecodeRules(c *Call, guard Guard, key, typ, format string, target any) error {
	if !guard(c.Data, key) {
		return nil
	}
	encoded, err := encodeField(c.Data, key, c.Schema[key])
	if err != nil || encoded == nil {
		return err
	}
	var rules []map[string]any
	if err := json.Unmarshal(encoded, &rules); err != nil {
		return err
	}
	for _, rule := range rules {
		text, _ := rule[Value].(string)
		if rule[Value], err = ParseValue(typ, format, text); err != nil {
			return errors.Join(flex.FmtErrorf("Error while building request options, field: %s", key), err)
		}
	}
	converted, err := json.Marshal(rules)
	if err != nil {
		return err
	}
	return json.Unmarshal(converted, target)
}

// CheckExisting fails when the type or the format a user wrote differ from those of a feature or property
// that exists already. The API cannot change them, so the values would be checked against the wrong ones.
func CheckExisting(c *Call, kind string, serverType, serverFormat *string) error {
	for key, server := range map[string]*string{Type: serverType, Format: serverFormat} {
		written, found := "", server != nil
		if found {
			written = *server
		}
		if c.Text(key) != written {
			return flex.FmtErrorf(
				"The %s %s exists with %s %q, but the configuration has %q. It cannot be changed, remove the %s from the configuration or use the same one",
				kind, c.IDs[1], key, written, c.Text(key), key,
			)
		}
	}
	return nil
}
