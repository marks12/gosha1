package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
)

func usualAppCreate(c *ishell.Context) {

    yellow := color.New(color.FgYellow).SprintFunc()

    c.Println(yellow("Hello we start creating new usual app"))
}