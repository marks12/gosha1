package cmd

const usualRouteEntity = `[ {Entity} ]
	router.HandleFunc("/api/v1/{entity}/{id}",  webapp.{Entity}Read).Methods("GET")
	router.HandleFunc("/api/v1/{entity}",       webapp.{Entity}Find).Methods("GET")
	router.HandleFunc("/api/v1/{entity}",       webapp.{Entity}Create).Methods("POST")
	router.HandleFunc("/api/v1/{entity}",       webapp.{Entity}Update).Methods("PUT")
	router.HandleFunc("/api/v1/{entity}/{id}",  webapp.{Entity}Delete).Methods("DELETE")

    //router-generator here dont touch this line`

var usualTemplateRouteEntity = template{
    Path:    "./path_error.txt",
    Content: usualRouteEntity,
}
