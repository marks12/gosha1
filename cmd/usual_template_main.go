package cmd

const usualMain = `package main

import (
    "{ms-name}/bootstrap"
    "{ms-name}/router"
    "{ms-name}/settings"
    "{ms-name}/wsserver"
    "fmt"
    "os"
    "gosha/cmd"
    "net/http"
)

func main() {

    if len(os.Args) > 1 && os.Args[1] == "shell" {
        cmd.Run()
        os.Exit(1)
    }

    // делаем автомиграцию
    bootstrap.FillDBTestData()

    go runWsServer()

    runHttpServer()
}

func runWsServer() {

	fmt.Println("Websocket сервер запущен :" + settings.GetWssPort())
    wsserver.SetMessageHandler("", router.HandleWss)
    wsserver.Run("", settings.GetWssPort())
}

func runHttpServer() {

	fmt.Println("API сервер запущен :" + settings.ServerPort)
	http.ListenAndServe("0.0.0.0:" + settings.ServerPort, router.Router())
}
`

var usualTemplateMain = template{
    Path:    "./main.go",
    Content: assignMsName(usualMain),
}
