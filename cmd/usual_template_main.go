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

	fmt.Println("API сервер запущен :" + settings.ServerPort)
	http.ListenAndServe("0.0.0.0:" + settings.ServerPort, router.Router())

}

func runWsServer() {

    wsserver.SetMessageHandler("", router.HandleWss)
    wsserver.Run("", settings.GetWssPort())
}

`

var usualTemplateMain = template{
    Path:    "./main.go",
    Content: assignMsName(usualMain),
}
