// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"context"
	"reflect"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Client is the App Configuration API client.
type Client = appconfigurationv1.AppConfigurationV1

// Call is what an API call of a resource or a data source needs: the client, the data of the
// request, the schema that describes it and the ids of the resource (the ones in its Terraform id).
type Call struct {
	Ctx    context.Context
	Client *Client
	Data   *schema.ResourceData
	Schema map[string]*schema.Schema
	IDs    []string
}

// Fill copies the fields of the data that pass the guard into an SDK options struct.
func (c *Call) Fill(guard Guard, options any) error {
	return FillOptions(c.Data, c.Schema, guard, options)
}

// Text returns a string field of the data.
func (c *Call) Text(key string) string {
	return c.Data.Get(key).(string)
}

// Calls are the API calls of a resource. A call the API does not have is left nil:
// a resource without Update has only fields that force a new resource, and without Delete
// it is only removed from the state.
//
// Create and Update return the resource that the API returns. Get returns it too, along with the
// response, so that a resource that is gone (404) can be dropped from the state.
// A result can be an SDK struct, or a field map when it needs changes before it is set.
type Calls struct {
	Create func(c *Call) (any, error)
	Get    func(c *Call) (any, *core.DetailedResponse, error)
	Update func(c *Call) (any, error)
	Delete func(c *Call) (*core.DetailedResponse, error)
}

// NewResource builds a resource. `kind` names it in messages, e.g. "Collection". `idKeys` are the
// fields that make up its Terraform id after the GUID: `<guid>/<idKeys[0]>/<idKeys[1]>/...`.
// `fields` are the fields of the resource as the user writes them, `parents` the ids of the
// resources it lives under, e.g. the environment of a feature.
func NewResource(
	kind string,
	idKeys []string,
	fields map[string]*schema.Schema,
	calls Calls,
	parents ...map[string]*schema.Schema,
) *schema.Resource {
	sch := getResourceSchema(fields, parents...)
	r := resource{kind: kind, idKeys: idKeys, schema: sch, calls: calls}
	res := &schema.Resource{
		CreateContext: r.create,
		ReadContext:   r.read,
		DeleteContext: r.delete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema:        sch,
	}
	if calls.Update != nil {
		res.UpdateContext = r.update
	}
	return res
}

type resource struct {
	kind   string
	idKeys []string
	schema map[string]*schema.Schema
	calls  Calls
}

func (r resource) create(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	client, err := GetAppConfigClient(meta, data)
	if err != nil {
		return diag.FromErr(err)
	}
	c := r.newCall(ctx, client, data, r.idsOf(data))
	result, err := r.calls.Create(c)
	if err == nil {
		data.SetId(BuildID(data.Get(GUID).(string), c.IDs...))
		err = setFields(data, r.schema, result, r.kind)
	}
	if err != nil {
		return Fail(err, "Error while creating %s: %s", r.kind, idLabel(c.IDs))
	}
	return nil
}

// read refreshes the state. A resource that no longer exists is removed from the state.
func (r resource) read(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	c, err := r.existingCall(ctx, data, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	result, response, err := r.calls.Get(c)
	if IsNotFound(response) {
		data.SetId("")
		return nil
	}
	if err == nil {
		err = setFields(data, r.schema, result, r.kind)
	}
	if err != nil {
		return Fail(err, "Error while fetching %s: %s", r.kind, idLabel(c.IDs))
	}
	return nil
}

func (r resource) update(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	c, err := r.existingCall(ctx, data, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	result, err := r.calls.Update(c)
	if err == nil {
		err = setFields(data, r.schema, result, r.kind)
	}
	if err != nil {
		return Fail(err, "Error while updating %s: %s", r.kind, idLabel(c.IDs))
	}
	return nil
}

// delete deletes the resource. A resource that is already gone is not an error.
func (r resource) delete(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	if r.calls.Delete != nil {
		c, err := r.existingCall(ctx, data, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		if response, err := r.calls.Delete(c); err != nil && !IsNotFound(response) {
			return Fail(err, "Error while deleting %s: %s", r.kind, idLabel(c.IDs))
		}
	}
	data.SetId("")
	return nil
}

// existingCall prepares a call on a resource that exists: its ids come from the Terraform id,
// and are stored in the data, so that a resource that is only imported has them too.
func (r resource) existingCall(ctx context.Context, data *schema.ResourceData, meta any) (*Call, error) {
	client, ids, err := GetClientAndIds(data, meta, len(r.idKeys))
	if err != nil {
		return nil, err
	}
	for i, key := range r.idKeys {
		if err := data.Set(key, ids[i]); err != nil {
			return nil, err
		}
	}
	return r.newCall(ctx, client, data, ids), nil
}

func (r resource) newCall(ctx context.Context, client *Client, data *schema.ResourceData, ids []string) *Call {
	return &Call{Ctx: ctx, Client: client, Data: data, Schema: r.schema, IDs: ids}
}

func (r resource) idsOf(data *schema.ResourceData) []string {
	ids := make([]string, len(r.idKeys))
	for i, key := range r.idKeys {
		ids[i] = data.Get(key).(string)
	}
	return ids
}

// NewDataSource builds a data source that reads one item. The item is identified by `idKeys`.
// `fields` are the fields the API returns. `inputs` are the extra arguments of the call, e.g. the
// parent ids and the query parameters.
func NewDataSource(
	kind string,
	idKeys []string,
	fields map[string]*schema.Schema,
	get func(c *Call) (any, error),
	inputs ...map[string]*schema.Schema,
) *schema.Resource {
	sch := getDataSourceSchema(fields, idKeys, inputs...)
	return &schema.Resource{
		Schema: sch,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
			client, err := GetAppConfigClient(meta, data)
			if err != nil {
				return diag.FromErr(err)
			}
			ids := make([]string, len(idKeys))
			for i, key := range idKeys {
				ids[i] = data.Get(key).(string)
			}
			result, err := get(&Call{Ctx: ctx, Client: client, Data: data, Schema: sch, IDs: ids})
			if err == nil {
				data.SetId(BuildID(data.Get(GUID).(string), ids...))
				err = setFields(data, sch, result, kind)
			}
			if err != nil {
				return Fail(err, "Error while fetching %s: %s", kind, idLabel(ids))
			}
			return nil
		},
	}
}

// NewListDataSource builds a data source that lists items. The items are stored under `itemsKey`.
// `fields` are the fields of one item. `list` returns the response as a field map, normally with
// ListAll. `inputs` are the extra arguments of the call, e.g. the parent ids and the query parameters.
func NewListDataSource(
	itemsKey, description string,
	fields map[string]*schema.Schema,
	list func(c *Call) (map[string]any, error),
	inputs ...map[string]*schema.Schema,
) *schema.Resource {
	kind := itemsKey
	sch := getListSchema(itemsKey, description, fields, inputs...)
	return &schema.Resource{
		Schema: sch,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
			client, err := GetAppConfigClient(meta, data)
			if err != nil {
				return diag.FromErr(err)
			}
			guid := data.Get(GUID).(string)
			fields, err := list(&Call{Ctx: ctx, Client: client, Data: data, Schema: sch})
			if err == nil {
				data.SetId(guid)
				err = SetFields(data, sch, fields)
			}
			if err != nil {
				return Fail(err, "Error while listing %s of: %s", kind, guid)
			}
			return nil
		},
	}
}

// setFields stores an API result in the state. The result is an SDK struct or a field map.
func setFields(data *schema.ResourceData, sch map[string]*schema.Schema, result any, kind string) error {
	if isNil(result) {
		return flex.FmtErrorf("SDK failed to provide valid %s", kind)
	}
	if fields, ok := result.(map[string]any); ok {
		return SetFields(data, sch, fields)
	}
	return SetFromStruct(data, sch, result)
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Pointer && v.IsNil()
}

func idLabel(ids []string) string {
	return strings.Join(ids, idSeparator)
}
