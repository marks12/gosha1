package main

import (
    "newapp/bootstrap"
    "newapp/router"
    "newapp/settings"
    "newapp/wsserver"
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

	fmt.Println("API сервер запущен :" + settings.ServerPort)
	http.ListenAndServe("0.0.0.0:" + settings.ServerPort, router.Router())
}
