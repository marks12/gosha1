package cmd

const msMain = `package main

import (
    "{ms-name}/core"
    "{ms-name}/settings"
    "{ms-name}/bootstrap"
    "{ms-name}/rpcapp"
    "msrpc/mdl"
    "msrpc/api"
    "msrpc/cnf"
    "{ms-name}/ms"
    "fmt"
    "os"
    "gosha/cmd"
)

func main() {

    if len(os.Args) > 1 && os.Args[1] == "shell" {
        cmd.Run()
        os.Exit(1)
    }

    config :=  mdl.Config{
        RabbitLogin: settings.RabbitServerLogin,
        RabbitPassword: settings.RabbitServerPassword,
        RabbitPort: settings.RabbitServerPort,
        RabbitServer: settings.RabbitServer,
        RabbitVirtualhost: settings.GetVirtualHost(),
        AuthHeadersKey: settings.MicroserviceAuthKey,
    }

    RabbitServer := api.GetServer(config)

    // Проверяем соединение с сервером БД
    if core.DbErr != nil {
        fmt.Println("DB error:", core.DbErr)
        return
    }

    // делаем автомиграцию
    bootstrap.FillDBTestData()

    // rpc команды

    RabbitServer.Run()
}`


const msMainEntity = `

    // rpc команды {entity-name}
    funcConfig := api.GetFunctionConfig(ms.TicketGuid{entity-name}Find, "Find {entity-name}", cnf.Route{entity-name}Find, cnf.FunctionTypeFind, true)
    RabbitServer.BindRoute(funcConfig, rpcapp.Find{entity-name})

    funcConfig = api.GetFunctionConfig(ms.TicketGuid{entity-name}Read, "Read {entity-name}", cnf.Route{entity-name}Read, cnf.FunctionTypeRead, true)
    RabbitServer.BindRoute(funcConfig, rpcapp.Read{entity-name})

    funcConfig = api.GetFunctionConfig(ms.TicketGuid{entity-name}Create, "Create {entity-name}", cnf.Route{entity-name}Create, cnf.FunctionTypeCreate, true)
    RabbitServer.BindRoute(funcConfig, rpcapp.Create{entity-name})

    funcConfig = api.GetFunctionConfig(ms.TicketGuid{entity-name}Update, "Update {entity-name}", cnf.Route{entity-name}Update, cnf.FunctionTypeUpdate, true)
    RabbitServer.BindRoute(funcConfig, rpcapp.Update{entity-name})

    funcConfig = api.GetFunctionConfig(ms.TicketGuid{entity-name}Delete, "Delete {entity-name}", cnf.Route{entity-name}Delete, cnf.FunctionTypeDelete, true)
    RabbitServer.BindRoute(funcConfig, rpcapp.Delete{entity-name})

    RabbitServer.Run`

var msTemplateMain = template{
    Path:    "./main.go",
    Content: assignMsName(msMain),
}

var msTemplateMainEntity = template{
    Path:    "./main.go",
    Content: msMainEntity,
}

