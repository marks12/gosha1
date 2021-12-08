package cmd

const usualMain = `package main

import (
    "{ms-name}/bootstrap"
    "{ms-name}/router"
    "{ms-name}/settings"
    "{ms-name}/wsserver"
    "fmt"
    "net/http"
)

func main() {

    // делаем автомиграцию
    bootstrap.FillDBTestData()

    if settings.IsDev() {
        fmt.Println("Running in DEV mode")
    } else {
        fmt.Println("Running in PROD mode")
    }

    go runWsServer()
    runHttpServer()
}

func runWsServer() {

	fmt.Println("Websocket сервер запущен :" + settings.GetWssPort())
    wsserver.SetMessageHandler("", router.HandleWss)
    wsserver.Run("", settings.GetWssPort())
}

func runHttpServer() {

	fmt.Println("API server running :" + settings.ServerPort)
	err := http.ListenAndServe("0.0.0.0:" + settings.ServerPort, router.Router())
	if err != nil {
		fmt.Printf("Cant run server with err: %+v \n", err)
	}
}
`

var usualTemplateMain = template{
    Path:    "./main.go",
    Content: assignMsName(usualMain),
}
