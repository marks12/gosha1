package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "os"
    "github.com/fatih/color"
    "gosha/mode"
)

var shell = ishell.New()

func RunShell() {

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    // display welcome info.
    shell.Println("Welcome to interactive shell")
    shell.Println("If you want execute shell command in non-interacive mode " + red("(NIM)") + " please type command:")
    shell.Println(red("$"), green(" gosha exit someCommand"))

    shell.Println("type help for help")

    // register a function for "greet" command.
    shell.AddCmd(&ishell.Cmd{
        Name: "greet",
        Help: "greet user",
        Func: great,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: "setAppType",
        Help: "Set app type.\n \t\t\t\t NIM: setAppType --type=[MsCore | MsRpcApi | Microservice | Usual]",
        Func: setAppType,
    })

    // run shell
    if len(os.Args) > 1 && os.Args[1] == "exit" {

        mode.SetNonInteractiveMode()
        shell.Process(os.Args[2:]...)

    } else {
        // start shell
        mode.SetInteractiveMode()
        shell.Run()
    }
}
