// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !openbsd
// +build !openbsd

package packer

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

func checkProcess(currentPID int) (bool, error) {
	myProc, err := process.NewProcess(int32(currentPID))
	if err != nil {
		return false, fmt.Errorf("Process check error: %s", err)
	}
	bg, err := myProc.Background()
	if err != nil {
		return bg, fmt.Errorf("Process background check error: %s", err)
	}

	return bg, nil
}
