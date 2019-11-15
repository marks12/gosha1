package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "encoding/json"
    "newapp/webapp"
    "newapp/settings"
)

// Router - маршрутизатор
func Router() http.Handler {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(settings.HomePageRoute, homePage).Methods("GET")

    //[ User ]
    router.HandleFunc(settings.UserRoute, webapp.UserFind).Methods("GET")
    router.HandleFunc(settings.UserRoute, webapp.UserCreate).Methods("POST")
    router.HandleFunc(settings.UserRoute+"/{id}", webapp.UserRead).Methods("GET")
    router.HandleFunc(settings.UserRoute+"/{id}", webapp.UserUpdate).Methods("PUT")
    router.HandleFunc(settings.UserRoute+"/{id}", webapp.UserDelete).Methods("DELETE")
    router.HandleFunc(settings.UserRoute, webapp.UserFindOrCreate).Methods("PUT")

    //[ Role ]
    router.HandleFunc(settings.RoleRoute, webapp.RoleFind).Methods("GET")
    router.HandleFunc(settings.RoleRoute, webapp.RoleCreate).Methods("POST")
    router.HandleFunc(settings.RoleRoute+"/{id}", webapp.RoleRead).Methods("GET")
    router.HandleFunc(settings.RoleRoute+"/{id}", webapp.RoleUpdate).Methods("PUT")
    router.HandleFunc(settings.RoleRoute+"/{id}", webapp.RoleDelete).Methods("DELETE")
    router.HandleFunc(settings.RoleRoute, webapp.RoleFindOrCreate).Methods("PUT")

    //[ RoleResource ]
    router.HandleFunc(settings.RoleResourceRoute, webapp.RoleResourceFind).Methods("GET")
    router.HandleFunc(settings.RoleResourceRoute, webapp.RoleResourceCreate).Methods("POST")
    router.HandleFunc(settings.RoleResourceRoute+"/{id}", webapp.RoleResourceRead).Methods("GET")
    router.HandleFunc(settings.RoleResourceRoute+"/{id}", webapp.RoleResourceUpdate).Methods("PUT")
    router.HandleFunc(settings.RoleResourceRoute+"/{id}", webapp.RoleResourceDelete).Methods("DELETE")
    router.HandleFunc(settings.RoleResourceRoute, webapp.RoleResourceFindOrCreate).Methods("PUT")

    //[ Resource ]
    router.HandleFunc(settings.ResourceRoute, webapp.ResourceFind).Methods("GET")
    router.HandleFunc(settings.ResourceRoute, webapp.ResourceCreate).Methods("POST")
    router.HandleFunc(settings.ResourceRoute+"/{id}", webapp.ResourceRead).Methods("GET")
    router.HandleFunc(settings.ResourceRoute+"/{id}", webapp.ResourceUpdate).Methods("PUT")
    router.HandleFunc(settings.ResourceRoute+"/{id}", webapp.ResourceDelete).Methods("DELETE")
    router.HandleFunc(settings.ResourceRoute, webapp.ResourceFindOrCreate).Methods("PUT")

    //[ ResourceType ]
    router.HandleFunc(settings.ResourceTypeRoute, webapp.ResourceTypeFind).Methods("GET")
    router.HandleFunc(settings.ResourceTypeRoute, webapp.ResourceTypeCreate).Methods("POST")
    router.HandleFunc(settings.ResourceTypeRoute+"/{id}", webapp.ResourceTypeRead).Methods("GET")
    router.HandleFunc(settings.ResourceTypeRoute+"/{id}", webapp.ResourceTypeUpdate).Methods("PUT")
    router.HandleFunc(settings.ResourceTypeRoute+"/{id}", webapp.ResourceTypeDelete).Methods("DELETE")
    router.HandleFunc(settings.ResourceTypeRoute, webapp.ResourceTypeFindOrCreate).Methods("PUT")

    //[ UserRole ]
    router.HandleFunc(settings.UserRoleRoute, webapp.UserRoleFind).Methods("GET")
    router.HandleFunc(settings.UserRoleRoute, webapp.UserRoleCreate).Methods("POST")
    router.HandleFunc(settings.UserRoleRoute+"/{id}", webapp.UserRoleRead).Methods("GET")
    router.HandleFunc(settings.UserRoleRoute+"/{id}", webapp.UserRoleUpdate).Methods("PUT")
    router.HandleFunc(settings.UserRoleRoute+"/{id}", webapp.UserRoleDelete).Methods("DELETE")
    router.HandleFunc(settings.UserRoleRoute, webapp.UserRoleFindOrCreate).Methods("PUT")

    //[ Auth ]
    //router.HandleFunc(settings.AuthRoute, webapp.AuthFind).Methods("GET")
    router.HandleFunc(settings.AuthRoute, webapp.AuthCreate).Methods("POST")
    //router.HandleFunc(settings.AuthRoute+"/{id}", webapp.AuthRead).Methods("GET")
    //router.HandleFunc(settings.AuthRoute+"/{id}", webapp.AuthUpdate).Methods("PUT")
    router.HandleFunc(settings.AuthRoute+"/{id}", webapp.AuthDelete).Methods("DELETE")
    //router.HandleFunc(settings.AuthRoute, webapp.AuthFindOrCreate).Methods("PUT")

    //[ CurrentUser ]
    router.HandleFunc(settings.CurrentUserRoute, webapp.CurrentUserFind).Methods("GET")
    //router.HandleFunc(settings.CurrentUserRoute, webapp.CurrentUserCreate).Methods("POST")
    //router.HandleFunc(settings.CurrentUserRoute+"/{id}", webapp.CurrentUserRead).Methods("GET")
    //router.HandleFunc(settings.CurrentUserRoute+"/{id}", webapp.CurrentUserUpdate).Methods("PUT")
    //router.HandleFunc(settings.CurrentUserRoute+"/{id}", webapp.CurrentUserDelete).Methods("DELETE")
    //router.HandleFunc(settings.CurrentUserRoute, webapp.CurrentUserFindOrCreate).Methods("PUT")

    //[ Entity ]
    router.HandleFunc(settings.EntityRoute, webapp.EntityFind).Methods("GET")
    router.HandleFunc(settings.EntityRoute, webapp.EntityCreate).Methods("POST")
    router.HandleFunc(settings.EntityRoute+"/list", webapp.EntityMultiCreate).Methods("POST")
    router.HandleFunc(settings.EntityRoute+"/list", webapp.EntityMultiUpdate).Methods("PUT")
    router.HandleFunc(settings.EntityRoute+"/list", webapp.EntityMultiDelete).Methods("DELETE")
    router.HandleFunc(settings.EntityRoute+"/{id}", webapp.EntityRead).Methods("GET")
    router.HandleFunc(settings.EntityRoute+"/{id}", webapp.EntityUpdate).Methods("PUT")
    router.HandleFunc(settings.EntityRoute+"/{id}", webapp.EntityDelete).Methods("DELETE")
    router.HandleFunc(settings.EntityRoute, webapp.EntityFindOrCreate).Methods("PUT")

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
        Date:    "2019.11.15 15:44:48",
    })
}
