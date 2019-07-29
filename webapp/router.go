package webapp

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/marks12/gothamel/tag"
	"github.com/rs/cors"
	"fmt"
	"encoding/base64"
	"strings"
)

// Router - маршрутизатор
func Router() http.Handler {

	// [about]
	router := mux.NewRouter().StrictSlash(true)

	//router.NotFoundHandler = http.HandlerFunc(NotFound)
	router.NotFoundHandler = http.HandlerFunc(MainPage)

	router.HandleFunc("/api", homePage).Methods("GET")


	//[ Entity ]
	router.HandleFunc("/api/v1/entity/{id}", EntityRead).Methods("GET")
	router.HandleFunc("/api/v1/entity", EntityFind).Methods("GET")
	router.HandleFunc("/api/v1/entity", EntityCreate).Methods("POST")
	router.HandleFunc("/api/v1/entity/{id}", EntityUpdate).Methods("PUT")
	router.HandleFunc("/api/v1/entity/{id}", EntityDelete).Methods("DELETE")


	//router.HandleFunc("/", MainPage).Methods("GET")
	//router.HandleFunc("/{route}.html", MainPage).Methods("GET")
	//router.HandleFunc("/js/{route}.js", MainPage).Methods("GET")
	//router.HandleFunc("/js/{route}.js.map", MainPage).Methods("GET")
	//router.HandleFunc("/css/{route}.css", MainPage).Methods("GET")
	//router.HandleFunc("/img/{route}.jpg", MainPage).Methods("GET")
	//router.HandleFunc("/img/{route}.jpeg", MainPage).Methods("GET")
	//router.HandleFunc("/img/{route}.png", MainPage).Methods("GET")


	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"token", "content-type"},
		AllowedOrigins: []string{"*.140140.ru", "*.140140.ru:*", "http://localhost:*", "http://192.168.1.58*"},
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
		Date:    "2019.03.13 22:42:25",
	})
}

func MainPage(w http.ResponseWriter, r *http.Request) {

	files := GetFiles()
	var decoded string

	for _, f := range files {
		if f.Path == ("./front/dist" + r.URL.Path) || r.URL.Path == "/" && f.Path == "./front/dist/index.html" {

			htmlPage := strings.TrimSuffix(f.Content, "\n")
			dbyte, err := base64.StdEncoding.DecodeString(htmlPage)
			if err != nil {
				fmt.Println("decode error:", err)
			}

			decoded = string(dbyte)
			w.Write([]byte(decoded))
			return
		}
	}

	NotFound(w, r)
	return

}


func NotFound (w http.ResponseWriter, httpRequest *http.Request) {

	doc := tag.Html{
		Head: tag.Head{Title: "Not Found title"},
	}

	content := tag.Article{Children: &[]tag.Nodes{
		tag.Div{
			Text: "Not Found",
		},
	}}

	doc.SetContent(content).
		AddHeader(tag.Header{}).
		AddFooter(tag.Footer{})

	htmlPage := doc.GetHtmlPage(true)

	w.Write([]byte(htmlPage))
}

