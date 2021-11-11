// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !solaris
// +build !solaris

package packer

import (
	"github.com/mattn/go-tty"
)

var openTTY = tty.Open
