package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"gosha/mode"
)

func setReadOnly(c *ishell.Context) {
	mode.SetReadOnlyMode()
}
