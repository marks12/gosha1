package cmd

const usualMain = `package main

import (
    "{ms-name}/bootstrap"
    "{ms-name}/router"
    "{ms-name}/settings"
    "fmt"
    "net/http"
)

func main() {

    if len(os.Args) > 1 && os.Args[1] == "shell" {
        cmd.Run()
        os.Exit(1)
    }

    // делаем автомиграцию
    bootstrap.FillDBTestData()

	fmt.Println("API running :" + settings.ServerPort)

    if settings.IsDev() {
        fmt.Println("Running in DEV mode")
    } else {
        fmt.Println("Running in PROD mode")
    }

	http.ListenAndServe("0.0.0.0:" + settings.ServerPort, router.Router())

}
`

var usualTemplateMain = template{
    Path:    "./main.go",
    Content: assignMsName(usualMain),
}
