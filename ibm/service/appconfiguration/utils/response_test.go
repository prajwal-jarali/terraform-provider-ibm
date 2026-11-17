// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func str(s string) *string { return &s }

var resultSchema = map[string]*schema.Schema{
	GUID:        {Type: schema.TypeString, Required: true},
	Name:        {Type: schema.TypeString, Computed: true},
	Tags:        {Type: schema.TypeString, Computed: true},
	CreatedTime: {Type: schema.TypeString, Computed: true},
	TotalCount:  {Type: schema.TypeInt, Computed: true},
	"features": {Type: schema.TypeList, Computed: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
		Name:         {Type: schema.TypeString, Computed: true},
		"feature_id": {Type: schema.TypeString, Computed: true},
	}}},
	Next: {Type: schema.TypeList, Computed: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
		Href: {Type: schema.TypeString, Computed: true},
	}}},
}

func TestSetFromStruct(t *testing.T) {
	data := (&schema.Resource{Schema: resultSchema}).Data(nil)
	created := strfmt.NewDateTime()
	next := "https://x/environments?offset=10"
	result := &appconfigurationv1.EnvironmentList{
		Environments: []appconfigurationv1.Environment{{Name: str("dev")}},
		TotalCount:   func() *int64 { v := int64(3); return &v }(),
		Next:         &appconfigurationv1.PaginatedListNext{Href: &next},
	}
	env := &appconfigurationv1.Environment{
		Name:        str("dev"),
		CreatedTime: &created,
		Features:    []appconfigurationv1.FeatureOutput{{Name: str("f"), FeatureID: str("fid")}},
	}

	if err := data.Set(GUID, "my-guid"); err != nil {
		t.Fatal(err)
	}
	if err := SetFromStruct(data, resultSchema, env); err != nil {
		t.Fatal(err)
	}
	if got := data.Get(Name); got != "dev" {
		t.Errorf("name = %v", got)
	}
	if got := data.Get(CreatedTime); got != created.String() {
		t.Errorf("created_time = %v, want %v", got, created.String())
	}
	if got := data.Get("features").([]any); len(got) != 1 || got[0].(map[string]any)["feature_id"] != "fid" {
		t.Errorf("features = %v", got)
	}
	if got := data.Get(Tags); got != "" {
		t.Errorf("absent field must stay empty, got %v", got)
	}
	if got := data.Get(GUID); got != "my-guid" {
		t.Errorf("input field was overwritten: %v", got)
	}

	if err := SetFromStruct(data, resultSchema, result); err != nil {
		t.Fatal(err)
	}
	if got := data.Get(TotalCount); got != 3 {
		t.Errorf("total_count = %v", got)
	}
	// pagination object is wrapped into a one element list
	nextLinks := data.Get(Next).([]any)
	if len(nextLinks) != 1 || nextLinks[0].(map[string]any)[Href] != next {
		t.Errorf("next = %v", nextLinks)
	}
}

func TestSetFromStructBigInteger(t *testing.T) {
	data := (&schema.Resource{Schema: resultSchema}).Data(nil)
	count := int64(9007199254740993) // 2^53 + 1, not exact as float64
	if err := SetFromStruct(data, resultSchema, &appconfigurationv1.EnvironmentList{TotalCount: &count}); err != nil {
		t.Fatal(err)
	}
	if got := data.Get(TotalCount); got != int(count) {
		t.Errorf("total_count = %v, want %v", got, count)
	}
}

func TestSetFromStructReportsEveryFailure(t *testing.T) {
	data := (&schema.Resource{Schema: resultSchema}).Data(nil)
	// features is a list in the schema but a string here
	err := SetFromStruct(data, resultSchema, map[string]any{"features": "oops", Name: "ok"})
	if err == nil {
		t.Fatal("expected an error for the mismatched field")
	}
	if got := data.Get(Name); got != "ok" {
		t.Errorf("valid fields must still be set, name = %v", got)
	}
}

func TestNextOffset(t *testing.T) {
	cases := []struct {
		name string
		href *string
		want int64
	}{
		{"nil", nil, 0},
		{"offset", str("https://x/e?limit=10&offset=20"), 20},
		{"start", str("https://x/e?start=30"), 30},
		{"no query", str("https://x/e"), 0},
		{"garbage value", str("https://x/e?offset=abc"), 0},
		{"invalid url", str("://bad"), 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NextOffset(tc.href); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

// A computed list that the API leaves out is stored as an empty list. Without it Terraform plans the list as
// "known after apply" after every apply, and a test or a user sees a change that never goes away.
func TestSetFromStructStoresAbsentComputedListAsEmpty(t *testing.T) {
	data := (&schema.Resource{Schema: resultSchema}).Data(nil)
	if err := SetFromStruct(data, resultSchema, &appconfigurationv1.Environment{Name: str("dev")}); err != nil {
		t.Fatal(err)
	}
	raw, exists := data.GetOkExists("features")
	if got, isList := raw.([]any); !exists || !isList || len(got) != 0 {
		t.Errorf("features = %v (exists %v), want an empty list", raw, exists)
	}
	if got := data.Get("tags"); got != "" {
		t.Errorf("a missing string must stay unset, got %q", got)
	}
}
