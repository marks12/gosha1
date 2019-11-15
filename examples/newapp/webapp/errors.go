package webapp

import (
    "newapp/types"
    "net/http"
    "encoding/json"
)

func ErrResponse (w http.ResponseWriter, err string, status int) {

    response := types.APIError{}

    response.Error = true
    response.ErrorMessage = err

    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)

    return
}

func ValidResponse (w http.ResponseWriter, data interface{}) {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)

    return
}