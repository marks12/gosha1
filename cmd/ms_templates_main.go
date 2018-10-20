package cmd

import "regexp"

const msMain = `package main

import (
    "msproduct/core"
    "msproduct/settings"
    "msproduct/bootstrap"
    "fmt"
    "msproduct/rpcapp"
    "msrpc/mdl"
    "msrpc/api"
    "msrpc/cnf"
    "msproduct/ms"
    "os"
    "gosha/cmd"
)

func main() {

    if len(os.Args) > 1 && os.Args[1] == "shell" {
        cmd.Run()
        os.Exit(1)
    }

    // Проверяем соединение с сервером БД
    if core.DbErr != nil {
        fmt.Println("DB error:", core.DbErr)
        return
    }

    // делаем автомиграцию
    bootstrap.FillDBTestData()

    config :=  mdl.Config{
        RabbitLogin: settings.RabbitServerLogin,
        RabbitPassword: settings.RabbitServerPassword,
        RabbitPort: settings.RabbitServerPort,
        RabbitServer: settings.RabbitServer,
        RabbitVirtualhost: settings.GetVirtualHost(),
        AuthHeadersKey: settings.MicroserviceAuthKey,
    }

    RabbitServer := api.GetServer(config)

    // rpc команды

    RabbitServer.Run()
}`

var microserviceNameRegexp = regexp.MustCompile("msproduct")

var msTemplateMain = template{
    Path:    "./main.go",
    Content: microserviceNameRegexp.ReplaceAllString(msMain, getCurrentDirName()),
}

