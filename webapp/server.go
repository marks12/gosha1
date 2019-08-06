package webapp

import (
	"fmt"
	"gosha/settings"
	"net/http"
)

func Run() {

	fmt.Println("API server running :" + settings.ServerPort)
	fmt.Println("Try to open: http://localhost:" + settings.ServerPort)
	e := http.ListenAndServe("0.0.0.0:" + settings.ServerPort, Router())

	if e != nil {
		fmt.Println("error running server: "  + e.Error())
	}
}

