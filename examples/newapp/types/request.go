package types

import (
    "encoding/json"
    "net/http"
)

// ReadJSON -
func ReadJSON(r *http.Request, entity interface{}) {

    decoder := json.NewDecoder(r.Body)
    decoder.Decode(entity)

    defer r.Body.Close()
}
