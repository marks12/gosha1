package webapp

import (
	"fmt"
	"net/http"
	"gosha/settings"
)

func Run() {

	fmt.Println("API сервер запущен :" + settings.ServerPort)
	http.ListenAndServe("0.0.0.0:" + settings.ServerPort, Router())
}

