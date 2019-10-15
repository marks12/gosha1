package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "gosha/mode"
    "fmt"
    "os"
    "strings"
    "gosha/settings"
)

const MSCORE_ENTITY_ADD = "core:entity:add"

const MSCORE_TEST = "test-shell-command"

const MSRPC_ENTITY_ADD = "rpc:entity:add"

const MS_INIT = "ms:init"

const MS_ENTITY_ADD = "ms:entity:add"

const USUAL_APP_CREATE = "usual:create"

const USUAL_ENTITY_ADD  = "usual:entity:add"

const USUAL_AUTH_ADD    = "usual:auth:add"

const ENTITY_ADD_FIELD = "entity:Field:add"

func setAppType(c *ishell.Context) {

    var choice int

    shell.DeleteCmd(MSCORE_ENTITY_ADD)
    shell.DeleteCmd(MSRPC_ENTITY_ADD)

    if mode.IsInteractive() {
        choice = getAppTypeFromChoice(c)
    } else {
        choice = getAppTypeFromArgs()
    }

    setAppCommands(choice, c)

    if IsNextCommand(choice) {
         runSecondLevelProgram(c)
    }
}

func getAppInfo(c *ishell.Context) {

    shell.DeleteCmd(MSCORE_ENTITY_ADD)
    shell.DeleteCmd(MSRPC_ENTITY_ADD)

    app := GetCurrentApp()

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    successText := red("invalid structure. Your app must have follow folders: ", settings.UsualDefaultStructure)

    if app.IsAppExists {
        successText = green("success structure")
    }

    InteractiveEcho([]string{
        "Current app has " + successText,
    })
}

func runSecondLevelProgram(c *ishell.Context) {

    for _, command := range c.Cmds() {
        if command.Name == os.Args[4] {
            command.Func(c)
        }
    }
}

func IsNextCommand(choice int) bool {

    return  choice > -1 &&
            mode.IsNonInteractive() &&
            len(os.Args) > 4
}

func getAppTypeFromArgs() int {

    red := color.New(color.FgRed).SprintFunc()
    supportedTypes :=  []string{"MsCore", "MsRpcApi", "Microservice", "", "Usual"}


    arg, er := GetOsArgument("type")

    if er != nil {
        fmt.Println(red("Not enough arguments. Please add for ex.: --type=MsCore to command"))
        os.Exit(1)
    }

    InArr, choice := InArray(arg.StringResult, supportedTypes)

    if ! InArr {
        fmt.Println(red("You set unsupported type " + arg.StringResult + ". Please use one of: " + strings.Join(supportedTypes,", ")))
        os.Exit(1)
    }

    return choice
}

func getAppTypeFromChoice(c *ishell.Context) int {

    return c.MultiChoice([]string{
        "MsCore",
        "MsRpcApi",
        "Microservice instance",
        "App using MsRpcApi [not available]",
        "Usual app",
    }, "Please select type of current app")
}

func setAppCommands(choice int, c *ishell.Context) {

    red := color.New(color.FgRed).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()
    blue := color.New(color.FgBlue).SprintFunc()

    switch choice {

    // mscore
    case 0:

        InteractiveEcho([]string{
            "Hello " + blue("MsCore") + " you have some command:",
            MSCORE_ENTITY_ADD + " - Creating new Entity and CRUD",
            green(MSCORE_TEST) + " - Test run command",
        })

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

        InteractiveEcho([]string{
            "Hello " + blue("usual") + " application:",
            green(USUAL_APP_CREATE), " - Create new app",
            green(USUAL_ENTITY_ADD), " - Add entity. NIM: --entity=SomeName --crud=fcruda --check-auth=fcruda, where a=findOrCreate",
            green(USUAL_AUTH_ADD), " - Add user, roles, auth, to app",
            green(ENTITY_ADD_FIELD), " - Add Field to model. NIM: --entity=SomeName2 --Field=SomeField --data-type=string",
        })

        setUsualAppCreate()

        setUsualEntityAdd()

		setUsualAuthAdd()

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

    shell.AddCmd(&ishell.Cmd{
        Name: MSCORE_TEST,
        Help: "Command test create in mscore",
        Func: mscoreTestCommand,
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
        Func: UsualAppInit,
    })
}


func setUsualEntityAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: USUAL_ENTITY_ADD,
        Help: "Command add new entity to usual app NIM. Use follow arguments:" +
            "\n\t\t\t\t--entity=SomeName - create entity with name SomeName" +
            "\n\t\t\t\t--crud=fcruda - add follow methods to route" +
            "\n\t\t\t\t--check-auth=fcruda  - add auth, where a=findOrCreate",
        Func: usualEntityAdd,
    })
}

func setUsualAuthAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: USUAL_AUTH_ADD,
        Help: "Command add auth, users, role, userRoles for app",
        Func: usualAuthAdd,
    })
}

func setModelFieldAdd() {

    shell.AddCmd(&ishell.Cmd{
        Name: ENTITY_ADD_FIELD,
        Help: "Command add new Field to model." +
            "NIM:" +
            "--entity=SomeName2" +
            "--Field=SomeField" +
            "--data-type=string",
        Func: entityFieldAdd,
    })
}