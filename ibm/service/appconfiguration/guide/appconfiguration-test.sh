#!/usr/bin/env bash
# Runs the tests of the App Configuration service through the Makefile.
#
#   ibm/service/appconfiguration/guide/appconfiguration-test.sh unit [resource]        unit tests (fake API), all or one resource
#   ibm/service/appconfiguration/guide/appconfiguration-test.sh acc all                acceptance tests of every resource
#   ibm/service/appconfiguration/guide/appconfiguration-test.sh acc <resource>...      acceptance tests of one or more resources
#   ibm/service/appconfiguration/guide/appconfiguration-test.sh list                   resources that have acceptance tests
#   ibm/service/appconfiguration/guide/appconfiguration-test.sh                        asks what to run
#
# Acceptance tests run against an existing instance. Fill in the settings below, or set the
# environment variables, or keep them in a `.env` file in the repository root. A value that is
# already in the environment wins over `.env`, and `.env` wins over the settings below.
# Do not commit a real API key: keep it in `.env`, which git ignores.

# ---- Settings (acceptance tests) ----------------------------------------------------------------
API_KEY=""        # IC_API_KEY: API key of the account that owns the instance. Required.
INSTANCE_GUID=""  # IBM_APPCONFIG_GUID: GUID of the existing App Configuration instance. Required.
REGION=""         # IC_REGION: region of the instance, e.g. us-south. Optional.
SERVICE_URL=""    # IBMCLOUD_APP_CONFIG_API_ENDPOINT: override of the App Configuration URL. Optional.
IAM_URL=""        # IBMCLOUD_IAM_API_ENDPOINT: override of the IAM URL. Optional.
# -------------------------------------------------------------------------------------------------

set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../../../.." && pwd)"  # the repository root
PKG_DIR="ibm/service/appconfiguration"
cd "$ROOT" || exit 1
export PATH="/usr/local/go/bin:$PATH"

# Resources that have acceptance tests: the packages with a resource or data source test.
acc_resources() {
	for dir in "$PKG_DIR"/*/; do
		ls "$dir" | grep -q '^\(resource\|data_source\)_test.go$' && basename "$dir"
	done
}

usage() {
	sed -n '2,13p' "$0" | sed 's/^# \{0,1\}//'
	echo "Acceptance resources: $(acc_resources | tr '\n' ' ')"
	exit "${1:-0}"
}

# The provider prints a long [INFO]/[WARN] banner when it loads. It is hidden, the exit code is kept.
run_make() {
	make "$@" 2>&1 | grep -v '^\[INFO\]\|^\[WARN\]'
	return "${PIPESTATUS[0]}"
}

unit() {
	local target="./$PKG_DIR/..."
	[[ -n "${1:-}" ]] && target="./$PKG_DIR/$1/"
	[[ "$target" != *"..." && ! -d "${target#./}" ]] && { echo "Unknown resource: $1"; exit 1; }
	# Acceptance tests skip themselves without TF_ACC, so a plain run is only unit tests.
	run_make test TEST="$target"
}

# Sets the variables the provider reads. Order of precedence: environment, .env, settings of this script.
load_settings() {
	local names=(IC_API_KEY IBM_APPCONFIG_GUID IC_REGION IBMCLOUD_APP_CONFIG_API_ENDPOINT IBMCLOUD_IAM_API_ENDPOINT)
	local defaults=("$API_KEY" "$INSTANCE_GUID" "$REGION" "$SERVICE_URL" "$IAM_URL")
	local shell_values=() i
	for name in "${names[@]}"; do shell_values+=("${!name:-}"); done
	if [[ -f .env ]]; then
		set -a
		# shellcheck disable=SC1091
		. ./.env >/dev/null 2>&1
		set +a
	fi
	for i in "${!names[@]}"; do
		local value="${shell_values[$i]:-${!names[$i]:-${defaults[$i]}}}"
		[[ -n "$value" ]] && export "${names[$i]}=$value"
	done
	return 0
}

acc() {
	[[ $# -eq 0 ]] && usage 1
	load_settings
	for var in IC_API_KEY IBM_APPCONFIG_GUID; do
		[[ -z "${!var:-}" ]] && { echo "$var must be set: fill in the settings of the script, or use .env."; exit 1; }
	done
	echo "Instance: $IBM_APPCONFIG_GUID${IC_REGION:+ ($IC_REGION)}${IBMCLOUD_APP_CONFIG_API_ENDPOINT:+ at $IBMCLOUD_APP_CONFIG_API_ENDPOINT}"

	local targets=()
	if [[ "$1" == "all" ]]; then
		targets=("./$PKG_DIR/...")
	else
		local known
		known="$(acc_resources)"
		for name in "$@"; do
			grep -qx "$name" <<<"$known" || { echo "No acceptance tests for: $name (see: list)"; exit 1; }
			targets+=("./$PKG_DIR/$name/")
		done
	fi
	# -p=1: the tests share one instance and use fixed ids, so packages must not run at the same time.
	run_make testacc TEST="${targets[*]}" TESTARGS='-run=TestAcc -p=1' TEST_TIMEOUT=120m
}

menu() {
	PS3="Run: "
	select choice in "Unit tests (all)" "Unit tests (one resource)" "Acceptance tests (all)" "Acceptance tests (one resource)" "Quit"; do
		case "$REPLY" in
		1) unit ;;
		2) read -r -p "Resource: " name && unit "$name" ;;
		3) acc all ;;
		4) read -r -p "Resource(s), separated by spaces: " names && acc $names ;;
		*) exit 0 ;;
		esac
		break
	done
}

case "${1:-}" in
unit) unit "${2:-}" ;;
acc) shift && acc "$@" ;;
list) acc_resources ;;
"") menu ;;
-h | --help | help) usage ;;
*) usage 1 ;;
esac
