// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"errors"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Fail turns an error into diagnostics, with a message that says which operation failed.
func Fail(err error, format string, args ...any) diag.Diagnostics {
	return diag.FromErr(errors.Join(flex.FmtErrorf(format, args...), err))
}
