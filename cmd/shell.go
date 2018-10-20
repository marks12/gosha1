package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
)

var shell = ishell.New()

func RunShell() {


    // display welcome info.
    shell.Println("Welcome to interactive shell")
    shell.Println("type help for help")

    // register a function for "greet" command.
    shell.AddCmd(&ishell.Cmd{
        Name: "greet",
        Help: "greet user",
        Func: great,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: "setAppType",
        Help: "Set app type",
        Func: whoAmI,
    })

    // run shell
    shell.Run()
}
