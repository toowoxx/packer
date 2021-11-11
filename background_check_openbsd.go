// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package packer

import (
	"fmt"
)

func checkProcess(currentPID int) (bool, error) {
	return false, fmt.Errorf("cannot determine if process is backgrounded in " +
		"openbsd")
}
