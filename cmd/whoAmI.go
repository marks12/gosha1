package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
)

const MSCORE_ENTITY_ADD = "core:entity:add"

const MSRPC_ENTITY_ADD = "rpc:entity:add"

const MS_INIT = "ms:init"

const MS_ENTITY_ADD = "ms:entity:add"

const USUAL_APP_CREATE = "usual:create"

const USUAL_ENTITY_ADD = "usual:entity:add"

const ENTITY_ADD_FIELD = "entity:field:add"

func whoAmI(c *ishell.Context) {

    red := color.New(color.FgRed).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()
    blue := color.New(color.FgBlue).SprintFunc()

    shell.DeleteCmd(MSCORE_ENTITY_ADD)
    shell.DeleteCmd(MSRPC_ENTITY_ADD)

    choice := c.MultiChoice([]string{
        "MsCore",
        "MsRpcApi",
        "Microservice instance",
        "App using MsRpcApi [not available]",
        "Usual app",
    }, "Please select type of current app")

    switch choice {

        // mscore
        case 0:

            c.Println("Hello " + blue("MsCore") + " you have some command:")
            c.Println(green(MSCORE_ENTITY_ADD), " - Creating new Entity and CRUD")

            setMscoreEntityAdd()

            break

        //msrpc
        case 1:

            c.Println("Hello " + blue("MsRpcApi") + " you have some commands:")
            c.Println(green(MSRPC_ENTITY_ADD), " - Creating new Entity and CRUD")

            setMsrpcEntityAdd()

            break

        //ms instance
        case 2:

            c.Println("Hello " + blue("Microservice instance") + " you have some commands:")
            c.Println(green(MS_INIT), " - Create new microservice")

            setMsInit()

            c.Println(green(MS_ENTITY_ADD), " - Add entity witch CRUD to microservice")

            setMsEntityAdd()

            break

        //uisual app
        case 4:

            c.Println("Hello " + blue("autonomic") + " application:")

            c.Println(green(USUAL_APP_CREATE), " - Create new app")
            setUsualAppCreate()

            c.Println(green(USUAL_ENTITY_ADD), " - Add entity")
            setUsualEntityAdd()

            c.Println(green(ENTITY_ADD_FIELD), " - Add field to model")
            setModelFieldAdd()

            break

        default:

            c.Println(red("You cancel select app. I dont know who are you. Sorry"))

            break
    }
}

func setMscoreEntityAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: MSCORE_ENTITY_ADD,
        Help: "Command create new entity in mscore",
        Func: mscoreEntityAdd,
    })
}

func setMsrpcEntityAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: MSRPC_ENTITY_ADD,
        Help: "Command create new file api/entity.go and create default components for working via msrpc",
        Func: msrpcEntityAdd,
    })
}

func setMsInit() {

    shell.AddCmd(&ishell.Cmd{
        Name: MS_INIT,
        Help: "Command create new empty microservice for using with msrpc, mscore and other services",
        Func: msInit,
    })
}

func setMsEntityAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: MS_ENTITY_ADD,
        Help: "Command add new entity to microservice",
        Func: msEntityAdd,
    })
}

func setUsualAppCreate() {

    shell.AddCmd(&ishell.Cmd{
        Name: USUAL_APP_CREATE,
        Help: "Command create new usual non MS application",
        Func: usualAppInit,
    })
}


func setUsualEntityAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: USUAL_ENTITY_ADD,
        Help: "Command add new entity to usual app",
        Func: usualEntityAdd,
    })
}

func setModelFieldAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: ENTITY_ADD_FIELD,
        Help: "Command add new field to model",
        Func: entityFieldAdd,
    })
}