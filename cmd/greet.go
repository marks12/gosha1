package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "strings"
)

func great(c *ishell.Context) {
    c.Println("Hello", strings.Join(c.Args, "&"))
}
