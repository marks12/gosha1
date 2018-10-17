package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
)

func RunShell() {

    var shell = ishell.New()

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
        Name: "msrpc:entity:add",
        Help: "Command create new file api/entity.go and create default components for working via msrpc",
        Func: addEntity,
    })

    // run shell
    shell.Run()
}
