//go:build !solaris
// +build !solaris

package packer

import (
	"github.com/mattn/go-tty"
)

var openTTY = tty.Open
