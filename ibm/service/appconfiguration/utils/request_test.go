// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testSchema = map[string]*schema.Schema{
	GUID:    {Type: schema.TypeString, Required: true},
	"id":    {Type: schema.TypeString, Required: true},
	Expand:  {Type: schema.TypeBool, Optional: true},
	Include: {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
	Tags:    {Type: schema.TypeString, Optional: true},
	Limit:   {Type: schema.TypeInt, Optional: true},
	Name:    {Type: schema.TypeString, Computed: true},
}

// runRead runs fn inside a data source read, with ResourceData built the way the SDK builds it.
func runRead(t *testing.T, sch map[string]*schema.Schema, config map[string]any, fn func(data *schema.ResourceData)) {
	t.Helper()
	res := &schema.Resource{Schema: sch, ReadContext: func(_ context.Context, data *schema.ResourceData, _ any) diag.Diagnostics {
		fn(data)
		return nil
	}}
	block := res.CoreConfigSchema()
	raw, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	value, err := ctyjson.Unmarshal(raw, block.ImpliedType())
	if err != nil {
		t.Fatal(err)
	}
	diff, err := res.Diff(context.Background(), nil, terraform.NewResourceConfigShimmed(value, block), nil)
	if err != nil {
		t.Fatal(err)
	}
	diff.RawConfig = value
	if _, diags := res.ReadDataApply(context.Background(), diff, nil); diags.HasError() {
		t.Fatal(diags)
	}
}

func TestFillOptions(t *testing.T) {
	t.Run("explicit false and list are sent", func(t *testing.T) {
		runRead(t, testSchema, map[string]any{GUID: "g", "id": "e", Expand: false, Include: []string{"features"}}, func(data *schema.ResourceData) {
			options := &appconfigurationv1.GetEnvironmentOptions{}
			if err := FillOptions(data, testSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			if options.Expand == nil || *options.Expand {
				t.Errorf("Expand = %v, want pointer to false", options.Expand)
			}
			if !reflect.DeepEqual(options.Include, []string{"features"}) {
				t.Errorf("Include = %v", options.Include)
			}
		})
	})
	t.Run("omitted optionals stay nil", func(t *testing.T) {
		runRead(t, testSchema, map[string]any{GUID: "g", "id": "e"}, func(data *schema.ResourceData) {
			options := &appconfigurationv1.GetEnvironmentOptions{}
			if err := FillOptions(data, testSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			if options.Expand != nil || options.Include != nil {
				t.Errorf("want nil optionals, got Expand=%v Include=%v", options.Expand, options.Include)
			}
		})
	})
	t.Run("same fields fill a different options struct", func(t *testing.T) {
		runRead(t, testSchema, map[string]any{GUID: "g", "id": "e", Tags: "a,b", Limit: 5}, func(data *schema.ResourceData) {
			options := &appconfigurationv1.ListEnvironmentsOptions{}
			if err := FillOptions(data, testSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			if options.Tags == nil || *options.Tags != "a,b" || options.Limit == nil || *options.Limit != 5 {
				t.Errorf("Tags=%v Limit=%v", options.Tags, options.Limit)
			}
		})
	})
	t.Run("output-only fields are never copied", func(t *testing.T) {
		type withName struct {
			Name *string `json:"name"`
		}
		runRead(t, testSchema, map[string]any{GUID: "g", "id": "e"}, func(data *schema.ResourceData) {
			if err := data.Set(Name, "from-state"); err != nil {
				t.Fatal(err)
			}
			target := &withName{}
			if err := FillOptions(data, testSchema, func(*schema.ResourceData, string) bool { return true }, target); err != nil {
				t.Fatal(err)
			}
			if target.Name != nil {
				t.Errorf("computed field leaked into request: %v", *target.Name)
			}
		})
	})
}

func TestForReadUnknownField(t *testing.T) {
	runRead(t, testSchema, map[string]any{GUID: "g", "id": "e"}, func(data *schema.ResourceData) {
		if ForRead(data, "not_in_schema") {
			t.Error("unknown field must not pass the guard")
		}
	})
	// ResourceData without any config (e.g. outside a read) must not panic either
	data := (&schema.Resource{Schema: testSchema}).Data(nil)
	if ForRead(data, Expand) {
		t.Error("field must not pass the guard without a config")
	}
}

func TestForUpdate(t *testing.T) {
	sch := map[string]*schema.Schema{
		"id":   {Type: schema.TypeString, Required: true},
		"name": {Type: schema.TypeString, Optional: true},
		"tags": {Type: schema.TypeString, Optional: true},
	}
	prior := map[string]string{"id": "x", "name": "n", "tags": "a"}
	cases := []struct {
		name   string
		prior  map[string]string
		config map[string]any
		want   map[string]bool // nil means Terraform does not call Update at all
	}{
		{"changed value", prior, map[string]any{"id": "x", "name": "n", "tags": "b"}, map[string]bool{"name": false, "tags": true}},
		{"changed to zero value", prior, map[string]any{"id": "x", "name": "n", "tags": ""}, map[string]bool{"name": false, "tags": true}},
		{"removed from config", prior, map[string]any{"id": "x", "name": "n"}, map[string]bool{"name": false, "tags": false}},
		{"unchanged", prior, map[string]any{"id": "x", "name": "n", "tags": "a"}, nil},
		// the SDK compares against the zero value, so "never set" -> "set to zero" is no change
		{"never set to explicit zero value", map[string]string{"id": "x", "name": "n"}, map[string]any{"id": "x", "name": "n2", "tags": ""}, map[string]bool{"name": true, "tags": false}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := map[string]bool{}
			res := &schema.Resource{Schema: sch,
				ReadContext: func(context.Context, *schema.ResourceData, any) diag.Diagnostics { return nil },
				UpdateContext: func(_ context.Context, data *schema.ResourceData, _ any) diag.Diagnostics {
					for _, field := range []string{"name", "tags"} {
						got[field] = ForUpdate(data, field)
					}
					return nil
				}}
			block := res.CoreConfigSchema()
			raw, _ := json.Marshal(tc.config)
			value, err := ctyjson.Unmarshal(raw, block.ImpliedType())
			if err != nil {
				t.Fatal(err)
			}
			state := &terraform.InstanceState{ID: "x", Attributes: tc.prior}
			diff, err := res.Diff(context.Background(), state, terraform.NewResourceConfigShimmed(value, block), nil)
			if err != nil {
				t.Fatal(err)
			}
			if diff == nil {
				if tc.want != nil {
					t.Fatal("expected a diff")
				}
				return
			}
			diff.RawConfig = value
			if _, diags := res.Apply(context.Background(), state, diff, nil); diags.HasError() {
				t.Fatal(diags)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// complexOptions has one field per complex schema type.
type complexOptions struct {
	Settings *struct {
		Retries *int64  `json:"retries"`
		Verbose *bool   `json:"verbose"`
		Label   *string `json:"label"`
	} `json:"settings"`
	Blocks []struct {
		A *string `json:"a"`
	} `json:"blocks"`
	Labels map[string]string `json:"labels"`
	Names  []string          `json:"names"`
	Flags  []string          `json:"flags"`
}

var complexSchema = map[string]*schema.Schema{
	GUID: {Type: schema.TypeString, Required: true},
	"settings": {Type: schema.TypeList, Optional: true, MaxItems: 1, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
		"retries": {Type: schema.TypeInt, Optional: true},
		"verbose": {Type: schema.TypeBool, Optional: true},
		"label":   {Type: schema.TypeString, Optional: true},
	}}},
	"blocks": {Type: schema.TypeList, Optional: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
		"a": {Type: schema.TypeString, Optional: true},
	}}},
	"labels": {Type: schema.TypeMap, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
	"names":  {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
	"flags":  {Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
}

func TestFillOptionsComplexTypes(t *testing.T) {
	t.Run("populated", func(t *testing.T) {
		config := map[string]any{
			GUID:       "g",
			"settings": []any{map[string]any{"retries": 0, "verbose": false}},
			"blocks":   []any{map[string]any{"a": "x"}, map[string]any{"a": "y"}},
			"labels":   map[string]any{"k": "v"},
			"names":    []any{"n1", "n2"},
			"flags":    []any{"f"},
		}
		runRead(t, complexSchema, config, func(data *schema.ResourceData) {
			options := &complexOptions{}
			if err := FillOptions(data, complexSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			// MaxItems 1 block becomes an object, and only what the user wrote is set
			if options.Settings == nil || options.Settings.Retries == nil || *options.Settings.Retries != 0 {
				t.Fatalf("settings.retries must be an explicit 0, got %+v", options.Settings)
			}
			if options.Settings.Verbose == nil || *options.Settings.Verbose {
				t.Errorf("settings.verbose must be an explicit false")
			}
			if options.Settings.Label != nil {
				t.Errorf("omitted nested attribute must stay nil, got %q", *options.Settings.Label)
			}
			if len(options.Blocks) != 2 || *options.Blocks[1].A != "y" {
				t.Errorf("blocks = %+v", options.Blocks)
			}
			if options.Labels["k"] != "v" || !reflect.DeepEqual(options.Names, []string{"n1", "n2"}) || !reflect.DeepEqual(options.Flags, []string{"f"}) {
				t.Errorf("labels=%v names=%v flags=%v", options.Labels, options.Names, options.Flags)
			}
		})
	})
	t.Run("omitted", func(t *testing.T) {
		runRead(t, complexSchema, map[string]any{GUID: "g"}, func(data *schema.ResourceData) {
			options := &complexOptions{}
			if err := FillOptions(data, complexSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			if options.Settings != nil || options.Blocks != nil || options.Labels != nil || options.Names != nil || options.Flags != nil {
				t.Errorf("omitted complex fields must stay nil: %+v", options)
			}
		})
	})
	t.Run("empty list is an explicit value", func(t *testing.T) {
		runRead(t, complexSchema, map[string]any{GUID: "g", "names": []any{}}, func(data *schema.ResourceData) {
			options := &complexOptions{}
			if err := FillOptions(data, complexSchema, ForRead, options); err != nil {
				t.Fatal(err)
			}
			if options.Names == nil || len(options.Names) != 0 {
				t.Errorf("names = %#v, want empty non-nil slice", options.Names)
			}
		})
	})
}

func TestForReadComplexTypes(t *testing.T) {
	// a list that contains an unknown element must not be sent half-built
	res := &schema.Resource{Schema: complexSchema}
	block := res.CoreConfigSchema()
	base, err := ctyjson.Unmarshal([]byte(`{"`+GUID+`":"g"}`), block.ImpliedType())
	if err != nil {
		t.Fatal(err)
	}
	values := base.AsValueMap()
	values["names"] = cty.ListVal([]cty.Value{cty.StringVal("a"), cty.UnknownVal(cty.String)})
	value := cty.ObjectVal(values)

	diff, err := res.Diff(context.Background(), nil, terraform.NewResourceConfigShimmed(value, block), nil)
	if err != nil {
		t.Fatal(err)
	}
	diff.RawConfig = value
	res.ReadContext = func(_ context.Context, data *schema.ResourceData, _ any) diag.Diagnostics {
		if ForRead(data, "names") {
			t.Error("a list with an unknown element must not pass the guard")
		}
		return nil
	}
	if _, diags := res.ReadDataApply(context.Background(), diff, nil); diags.HasError() {
		t.Fatal(diags)
	}
}

func TestFillOptionsTypeMismatch(t *testing.T) {
	// a string schema field that maps to a []string struct field must error, not be dropped
	sch := map[string]*schema.Schema{Include: {Type: schema.TypeString, Optional: true}}
	runRead(t, sch, map[string]any{Include: "features"}, func(data *schema.ResourceData) {
		if err := FillOptions(data, sch, ForRead, &appconfigurationv1.GetEnvironmentOptions{}); err == nil {
			t.Error("expected a type mismatch error")
		}
	})
}
