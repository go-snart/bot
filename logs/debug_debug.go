// +build snart_debug

package logs

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

func debug(name string) *nilog.Logger {
	return nilog.New(
		colorable.NewColorableStderr(),
		aurora.Yellow("[debug] "+name+": ").String(),
		log.LstdFlags,
	)
}