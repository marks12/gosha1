package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/mode"
)

func setReadOnly(c *ishell.Context) {
	mode.SetReadOnlyMode()
}
