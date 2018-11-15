package cmd

const usualWebappEntityType = `
package types

import (
    "net/http"
)


type {entity-name} struct {
    ID                  string

    CreatedAt           string
    UpdatedAt           string
    DeletedAt           string
}

type {entity-name}Filter struct {

    Pagination
    validator
    Authenticator
}

func Get{entity-name}Filter(request *http.Request) {entity-name}Filter {

    var filter {entity-name}Filter

    return  filter
}


func Get{entity-name}Model(request *http.Request) {entity-name} {

    var model {entity-name}

    model.Validate()

    return  model
}

func (user *{entity-name}) Validate()  {

}
`

var usualTemplateWebappEntityType = template{
    Path:    "",
    Content: assignMsName(usualWebappEntityType),
}
