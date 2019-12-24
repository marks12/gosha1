package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "os"
    "github.com/fatih/color"
    "gosha/mode"
)

var shell = ishell.New()

const SET_APP_TYPE = "setAppType"
const GET_APP_INFO = "getAppInfo"
const GENERATE_TYPES_JS = "gen:types:js"
const GENERATE_TESTS_WEBAPP = "gen:tests:webapp"
const GENERATE_GENERATORS = "gen:generators"

func RunShell() {

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    // display welcome info in interactive mode.
    InteractiveEcho([]string{
        "Welcome to interactive shell",
        "If you want execute shell command in non-interacive mode " + red("(NIM)") + " please type command:",
        red("$") + green(" gosha exit someCommand"),
    })

    InteractiveEcho([]string{
        "type help for help",
    })

    // register a function for "greet" command.
    shell.AddCmd(&ishell.Cmd{
        Name: "greet",
        Help: "greet user",
        Func: great,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: SET_APP_TYPE,
        Help: "Set app type. " +
              "\n\t\t\t\tNIM: "+SET_APP_TYPE+" --type=[MsCore | MsRpcApi | Microservice | Usual]",
        Func: setAppType,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: GET_APP_INFO,
        Help: "Get current app info. " +
              "\n\t\t\t\tNIM: "+GET_APP_INFO,
        Func: getAppInfo,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: GENERATE_TYPES_JS,
        Help: "Generate types structs to JS for using in frontend " +
            "\n\t\t\t\tNIM: "+GENERATE_TYPES_JS+" --dst=/some/destination/path --map-namespace=im3 add namespace for modules",
        Func: genTypesJs,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: GENERATE_TESTS_WEBAPP,
        Help: "Generate tests for webapp",
        Func: genTestsWebapp,
    })

    shell.AddCmd(&ishell.Cmd{
        Name: GENERATE_GENERATORS,
        Help: "Generate generators functions",
        Func: genGenerators,
    })


    // run shell
    if len(os.Args) > 1 && os.Args[1] == "exit" {

        shell.Process(os.Args[2:]...)

    } else {
        // start shell
        shell.Run()
    }
}

func setMode() {

    if len(os.Args) > 1 && os.Args[1] == "exit" {
        mode.SetNonInteractiveMode()
    } else {
        mode.SetInteractiveMode()
    }
}

func GetShell() *ishell.Shell {
    return shell
}
