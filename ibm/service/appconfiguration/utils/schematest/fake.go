// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package schematest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Request is an API call that the fake API received.
type Request struct {
	Method, Path, Query string
	Body                map[string]any
}

// Reply is what the fake API answers: a status and a JSON body.
type Reply struct {
	Status int
	Body   any
}

// Handler answers a call of the fake API.
type Handler func(r Request) Reply

type fakeSession struct {
	conns.ClientSession
	client *appconfigurationv1.AppConfigurationV1
}

func (s fakeSession) BluemixSession() (*session.Session, error) {
	return &session.Session{Config: &bluemix.Config{Region: "test"}}, nil
}

func (s fakeSession) AppConfigurationV1() (*appconfigurationv1.AppConfigurationV1, error) {
	return s.client, nil
}

// NewMeta starts a fake App Configuration API and returns the provider meta that points to it, and
// the calls the API received. Paths are relative to the instance: `/environments/dev`.
func NewMeta(t *testing.T, handler Handler) (meta any, calls *[]Request) {
	t.Helper()
	const instance = "/apprapp/feature/v1/instances/guid"
	received := &[]Request{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		call := Request{Method: r.Method, Path: r.URL.Path[len(instance):], Query: r.URL.RawQuery}
		_ = json.Unmarshal(raw, &call.Body)
		*received = append(*received, call)
		reply := handler(call)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(reply.Status)
		_ = json.NewEncoder(w).Encode(reply.Body)
	}))
	t.Cleanup(server.Close)
	t.Setenv("IBMCLOUD_APP_CONFIG_API_ENDPOINT", server.URL)
	client, err := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
		Authenticator: &core.NoAuthAuthenticator{},
	})
	if err != nil {
		t.Fatal(err)
	}
	return fakeSession{client: client}, received
}

// Apply creates or updates a resource, as Terraform does, from the config the user wrote.
// It returns the new state. `state` is nil for a new resource.
func Apply(t *testing.T, res *schema.Resource, meta any, state *terraform.InstanceState, config map[string]any) *terraform.InstanceState {
	t.Helper()
	block := res.CoreConfigSchema()
	raw, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	value, err := ctyjson.Unmarshal(raw, block.ImpliedType())
	if err != nil {
		t.Fatal(err)
	}
	return apply(t, res, meta, state, value, block)
}

func apply(t *testing.T, res *schema.Resource, meta any, state *terraform.InstanceState, value cty.Value, _ any) *terraform.InstanceState {
	t.Helper()
	next, err := tryApply(res, meta, state, value)
	if err != nil {
		t.Fatal(err)
	}
	return next
}

func tryApply(res *schema.Resource, meta any, state *terraform.InstanceState, value cty.Value) (*terraform.InstanceState, error) {
	diff, err := res.Diff(context.Background(), state, terraform.NewResourceConfigShimmed(value, res.CoreConfigSchema()), meta)
	if err != nil {
		return nil, err
	}
	diff.RawConfig = value
	next, diags := res.Apply(context.Background(), state, diff, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("%v", diags)
	}
	return next, nil
}

// TryApply is Apply for a config that is expected to fail. It returns the error.
func TryApply(t *testing.T, res *schema.Resource, meta any, state *terraform.InstanceState, config map[string]any) error {
	t.Helper()
	raw, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	value, err := ctyjson.Unmarshal(raw, res.CoreConfigSchema().ImpliedType())
	if err != nil {
		t.Fatal(err)
	}
	_, err = tryApply(res, meta, state, value)
	return err
}

// Refresh reads a resource again, as `terraform refresh` does. It returns nil when the resource is gone.
func Refresh(t *testing.T, res *schema.Resource, meta any, state *terraform.InstanceState) *terraform.InstanceState {
	t.Helper()
	next, diags := res.RefreshWithoutUpgrade(context.Background(), state, meta)
	if diags.HasError() {
		t.Fatal(diags)
	}
	return next
}

// Destroy deletes a resource.
func Destroy(t *testing.T, res *schema.Resource, meta any, state *terraform.InstanceState) {
	t.Helper()
	_, diags := res.Apply(context.Background(), state, &terraform.InstanceDiff{Destroy: true}, meta)
	if diags.HasError() {
		t.Fatal(diags)
	}
}

// ReadDataSource reads a data source from the config the user wrote and returns its state.
func ReadDataSource(t *testing.T, res *schema.Resource, meta any, config map[string]any) map[string]string {
	t.Helper()
	block := res.CoreConfigSchema()
	raw, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	value, err := ctyjson.Unmarshal(raw, block.ImpliedType())
	if err != nil {
		t.Fatal(err)
	}
	diff, err := res.Diff(context.Background(), nil, terraform.NewResourceConfigShimmed(value, block), meta)
	if err != nil {
		t.Fatal(err)
	}
	diff.RawConfig = value
	state, diags := res.ReadDataApply(context.Background(), diff, meta)
	if diags.HasError() {
		t.Fatal(diags)
	}
	return state.Attributes
}
