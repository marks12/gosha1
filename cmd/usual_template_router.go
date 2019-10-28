package cmd

const usualRouter = `package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "encoding/json"
    "{ms-name}/webapp"
    "{ms-name}/settings"
)

// Router - маршрутизатор
func Router() http.Handler {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(settings.HomePageRoute, homePage).Methods("GET")

    //router-generator here dont touch this line

    handler := cors.New(cors.Options{
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"token", "content-type"},
    }).Handler(router)

    return handler
}

func homePage(w http.ResponseWriter, r *http.Request) {
    type Response struct {
        Version string
        Date    string
    }

    json.NewEncoder(w).Encode(Response{
        Version: "0.0.1",
        Date:    "{current-date-time}",
    })
}
`

const usualWssRouter = `package router

import (
    "{ms-name}/wsserver"
)

func HandleWss(msg wsserver.UserMessage, con *wsserver.UserConnection) (answer wsserver.UserResponse) {

    answer = wsserver.UserResponse{
        Type:    msg.Type,
        Result:  "response from home",
        Success: true,
    }

    return
}
`

var usualTemplateRouter = template{
	Path:    "./router/router.go",
	Content: assignCurrentDateTime(assignMsName(usualRouter)),
}

var usualTemplateWssRouter = template{
	Path:    "./router/wss.go",
	Content: assignCurrentDateTime(assignMsName(usualWssRouter)),
}
