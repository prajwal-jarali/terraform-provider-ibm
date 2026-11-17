# Developer guide: App Configuration resources

How a resource is built, and how to change it. All names below use a made-up resource `widget`.

## Layout

```
appconfiguration/
  registry.go            re-exports the constructors, so provider.go imports one package
  <resource>/            one folder per resource: environment, feature, rule, ...
    schema.go            fields + the three constructors      <- change fields here
    resource.go          SDK calls: create, get, update, delete
    data_source.go       SDK calls: read one, read list
    *_test.go            schema_test, lifecycle_test (fake API), resource_test / data_source_test (real API)
  utils/                 the shared kit (below)
    schematest/          fake API for unit tests
    acceptance/          helpers shared by the real-API tests
```

A change to a field is a change to `schema.go` only. A change to behavior is `resource.go` / `data_source.go`.

## 1. Fields: `getFields(mode)`

One function describes a resource. It is called twice:

- `utils.Input`: the resource, what the user writes.
- `utils.Output`: the data source, what the API returns.

`Mode` turns the same field into the right flags:

| Helper | Input (resource) | Output (data source) |
|---|---|---|
| `mode.Required(f)` | required | computed |
| `mode.Optional(f)` | optional | computed |
| `mode.OptionalComputed(f)` | optional, API default kept | computed |
| `mode.RequiredForceNew(f)` | required, change = recreate | computed |
| `mode.OptionalForceNew(f)` | optional, change = recreate | computed |

```go
func getFields(mode utils.Mode) map[string]*schema.Schema {
	return utils.Merge(
		utils.GetIdentityFields(mode, "widget"), // name, description, tags
		utils.GetAuditFields("widget"),          // created_time, updated_time, href (computed)
		map[string]*schema.Schema{
			utils.WidgetID: mode.RequiredForceNew(utils.NewString("Widget Id.")),
			"size":         mode.Optional(utils.NewInt("Size in cm.")),
			"kind":         mode.OptionalComputed(utils.NewEnum("Kind.", "small", "large")),
		},
	)
}
```

### Field builders (`utils/schema_builders.go`)

They set the type and description. The `Mode` helper sets the flags.

`NewString`, `NewInt`, `NewFloat`, `NewBool`, `NewStringList`, `NewEnum(desc, values...)`,
`NewEnumList`, `NewList(desc, elem)` (list of objects), `NewObject(desc, elem)` (one object, `MaxItems: 1`).

Modifiers return a copy: `AsRequired`, `AsOptional`, `AsComputed`, `AsOptionalComputed`, `AsForceNew`, `AsSensitive`.

```go
"secret": mode.Required(utils.AsSensitive(utils.NewString("Secret."))),
"limits": mode.Optional(utils.NewObject("Limits.", map[string]*schema.Schema{
	"max": utils.NewInt("Maximum."),
})),
```

Keep nested schemas in small named functions (`getLimitFields(mode)`) instead of inline.

### Shared field groups (`utils/schema_common.go`, `schema_rules.go`)

| Helper | Gives |
|---|---|
| `GetIdentityFields(mode, name)` | `name`, `description`, `tags` |
| `GetAuditFields(name)` | `created_time`, `updated_time`, `href` |
| `GetParentIDSchema(desc)` | required, force-new id of the parent (e.g. `environment_id` of a feature) |
| `GetAssociatedSchema(kind, idKey, owner)` | computed list of related items |
| `GetWorkflowApprovalSchema()` | computed workflow approval object |
| `GetIncludeValuesSchema`, `GetSortValuesSchema`, `GetSearchSchema`, `GetFilterSchema`, `GetPageQuery` | query arguments of data sources |
| `GetRuleFields`, `GetRolloutFields` | rule and rollout blocks |
| `Merge(maps...)` | combines maps, a later map wins |

Key names are constants in `utils/constants.go` (or `const` in the package). Do not repeat a string key.

## 2. Constructors (`schema.go`)

```go
var idKeys = []string{utils.WidgetID} // the Terraform id is <guid>/<widget_id>

func ResourceIBMAppConfigWidget() *schema.Resource {
	return utils.NewResource("Widget", idKeys, getFields(utils.Input), utils.Calls{
		Create: createWidget, Get: getWidget, Update: updateWidget, Delete: deleteWidget,
	}) // extra args = parent ids, e.g. utils.NewResource(..., calls, map[string]*schema.Schema{utils.EnvironmentID: utils.GetParentIDSchema("Environment Id.")})
}

func DataSourceIBMAppConfigWidget() *schema.Resource {
	return utils.NewDataSource("Widget", idKeys, getFields(utils.Output), readWidget)
}

func DataSourceIBMAppConfigWidgets() *schema.Resource {
	return utils.NewListDataSource("widgets", "Array of widgets.", getFields(utils.Output), readWidgets,
		utils.GetPageQuery())
}
```

What the constructors add for you: the `guid` argument, the parent ids, the importer, the id
(`<guid>/<id1>/<id2>`), reading the response into state, 404 handling (resource dropped from state)
and, for lists, the pagination fields.

- A resource without an update API: leave `Update` nil. Every field must then be force-new.
- A resource without a delete API: leave `Delete` nil. It is only removed from state.
- Extra query arguments of a data source (`expand`, `include`, `sort`): pass them as the last argument(s).

## 3. SDK calls (`resource.go`, `data_source.go`)

A call gets `*utils.Call`: `c.Ctx`, `c.Client` (the SDK), `c.Data`, `c.IDs` (ids after the guid), `c.Text(key)`.

```go
func createWidget(c *utils.Call) (any, error) {
	options := c.Client.NewCreateWidgetOptions(c.Text(utils.Name), c.IDs[0]) // required SDK arguments
	if err := c.Fill(utils.ForCreate, options); err != nil {                // everything else
		return nil, err
	}
	result, _, err := c.Client.CreateWidgetWithContext(c.Ctx, options)
	return result, err // the constructor writes it to the state
}
```

### How `Fill` works

`c.Fill(guard, options)` copies fields from the user's config into the SDK options struct.
A schema key goes to the struct field whose `json` tag has the same name:

```go
// SDK                                       // schema
Size *int64 `json:"size,omitempty"`          "size": mode.Optional(utils.NewInt(...))
```

So adding the field `size` to the schema is enough, if the SDK struct has `size`. If the names differ,
rename the schema key or set it by hand after `Fill`.

Guards decide what is sent:

| Guard | Sends a field when |
|---|---|
| `ForCreate`, `ForRead` | the user wrote it (an explicit `false`/`0` counts) |
| `ForUpdate` | it changed and is still in the config |
| `utils.Skip(guard, "a", "b")` | same as `guard`, but never `a` and `b` (you set them yourself) |

Computed-only fields are never sent. Nested objects are sent as one object.

### Response side

The result of a call is written to the state by matching the schema key to the response `json` key.
A computed key that is not in the SDK response fails `TestComputedKeysMatchSDK`.
If a response needs a change first (rename, format a value), return a `map[string]any` instead of the struct:
`utils.FieldMap(result)` makes one from a struct.

### Lists

```go
func readWidgets(c *utils.Call) (map[string]any, error) {
	options := c.Client.NewListWidgetsOptions()
	if err := c.Fill(utils.ForRead, options); err != nil {
		return nil, err
	}
	return utils.ListAll(c, "widgets", &options.Offset, func() (any, error) {
		page, _, err := c.Client.ListWidgetsWithContext(c.Ctx, options)
		return page, err
	}) // fetches every page, unless the user set offset
}
```

### Other helpers (`utils/`)

| Helper | Use |
|---|---|
| `ParseValue(type, format, text)` | turns a text value into bool/number/JSON by the feature type |
| `DecodeRules`, `DecodeField` | read a nested schema block into an SDK struct |
| `CheckExisting` | feature/property create finds an existing item first, then updates it |
| `BuildID`, `ParseID` | `<guid>/<a>/<b>` ids |
| `IsNotFound` | 404 test |

## 4. Add or change things

**Add a field**: add it to `getFields`, add a line to the docs in `website/docs/r|d/app_config_<name>.html.markdown`.
If the SDK struct has the same `json` key, nothing else is needed. Run the unit tests.

**Make a field immutable**: `mode.Optional(...)` becomes `mode.OptionalForceNew(...)`.

**Add a resource**:
1. Copy `environment/` (the smallest example), rename the package and the SDK calls.
2. Add the constructors to `registry.go` (`ResourceIBMAppConfigWidget = widget.ResourceIBMAppConfigWidget`).
3. Add one line each to `DataSourcesMap` / `ResourcesMap` in `ibm/provider/provider.go`.
4. Add docs under `website/docs`.

## 5. Tests

| File | Runs | What |
|---|---|---|
| `schema_test.go` | always | schema is valid; computed keys exist in the SDK structs |
| `lifecycle_test.go` | always | create / update / refresh / delete / pagination against a fake API (`utils/schematest`) |
| `resource_test.go`, `data_source_test.go` | `TF_ACC=1` + credentials | real API |

Fake API example:

```go
meta, calls := schematest.NewMeta(t, func(r schematest.Request) schematest.Reply {
	return schematest.Reply{Status: http.StatusOK, Body: map[string]any{"widget_id": "w1", "name": "W"}}
})
state := schematest.Apply(t, ResourceIBMAppConfigWidget(), meta, nil, map[string]any{"guid": "g", "widget_id": "w1", "name": "W"})
// (*calls) holds every request, so you can check method, query and body.
```

Also: `Refresh`, `Destroy`, `TryApply`, `ReadDataSource`.

`utils/acceptance` holds only what several real-API tests share: `PreCheck`, `GUIDConfig`, `EnvironmentConfig`,
`FeatureConfig` (Terraform config to build on), `CheckExists`, `CheckDestroy`. The fetch of one resource
(`getWidget`) lives in that resource's own `resource_test.go`.
A resource that needs another service (Git, Event Notifications, Key Protect, ServiceNow) has no acceptance test,
only the fake-API one. The API checks the Git repository of a snapshot and the Secrets Manager secret of a workflow,
so dummy values are rejected.

### Run the tests

Use `ibm/service/appconfiguration/guide/appconfiguration-test.sh`. It calls `make test` and `make testacc`, and hides the `[INFO]`/`[WARN]`
lines the provider prints when it loads.

```
ibm/service/appconfiguration/guide/appconfiguration-test.sh unit                  # unit tests of the whole service
ibm/service/appconfiguration/guide/appconfiguration-test.sh unit feature          # unit tests of one resource
ibm/service/appconfiguration/guide/appconfiguration-test.sh acc all               # acceptance tests of every resource
ibm/service/appconfiguration/guide/appconfiguration-test.sh acc collection rule   # acceptance tests of some resources
ibm/service/appconfiguration/guide/appconfiguration-test.sh list                  # resources that have acceptance tests
ibm/service/appconfiguration/guide/appconfiguration-test.sh                       # menu
```

Unit tests need nothing. They use the fake API and finish in seconds.

Acceptance tests run in an App Configuration instance that already exists, and they create and delete their own
resources in it. Give the script the instance and the credentials in one of three ways. A value set higher in this list
wins over the ones below it:

1. Environment variables in your shell.
2. A `.env` file in the repository root (git ignores it), with `export NAME=value` lines.
3. The settings at the top of the script. Do not commit a real API key there.

| Setting in the script | Environment variable | |
|---|---|---|
| `API_KEY` | `IC_API_KEY` | Required. API key of the account that owns the instance. |
| `INSTANCE_GUID` | `IBM_APPCONFIG_GUID` | Required. GUID of the instance. |
| `REGION` | `IC_REGION` | Optional. Region of the instance. |
| `SERVICE_URL` | `IBMCLOUD_APP_CONFIG_API_ENDPOINT` | Optional. Replaces the App Configuration URL, e.g. for a test environment. |
| `IAM_URL` | `IBMCLOUD_IAM_API_ENDPOINT` | Optional. Replaces the IAM URL. |

The script prints the instance it runs against before it starts.

The resources run one after another (`-p=1`) because they share one instance and use fixed ids such as `tf_env`.
A run that is stopped can leave `tf_env` or a test resource behind. Delete it before the next run, or the create
step fails with "already exists".

Without the script:

```
make test TEST=./ibm/service/appconfiguration/...
make testacc TEST=./ibm/service/appconfiguration/environment/ TESTARGS='-run=TestAcc -p=1' TEST_TIMEOUT=120m
```

## 6. Behavior to know

- Feature and property create: get first, `CheckExisting`, then update.
- Workflow `provider_config` is sent to the API as `provider`.
- Metric `is_favorite` is a separate PATCH call.
- Integration keeps the source name and description the user wrote.
- A rule value is parsed by the type of its feature.
- Integration, origin config and guarded rollout lack update/delete where the API has none.
- `ibm_app_config_integration_en` and `_kms` are older resources with their own schema. They stay in the package root (`resource_ibm_app_config_integration_*.go`) for existing users, and their docs point to `ibm_app_config_integration`. Do not extend them.
- Integration: `CustomizeDiff` (`validateMetadata`) checks that the `metadata` has the settings its `integration_type` needs.
