// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package packer

import (
	"fmt"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func openTTY() (packersdk.TTY, error) {
	return nil, fmt.Errorf("no TTY available on solaris")
}
