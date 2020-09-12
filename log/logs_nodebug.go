// +build !debug

package log

import (
	"log"

	"github.com/superloach/nilog"
)

const flags = log.Lshortfile

// Debug is a *nilog.Logger for debugging purposes, and is nil unless the debug tag is enabled.
var Debug = (*nilog.Logger)(nil)
