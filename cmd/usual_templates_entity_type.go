package cmd

const usualWebappEntityType = `package types

import (
    "net/http"
)

type {Entity} struct {
    Id   int
    Name string
}

func ({entity} *{Entity}) Validate()  {

}

type {Entity}Filter struct {

    AbstractFilter
}

func Get{Entity}Filter(request *http.Request, functionType string) {Entity}Filter {

    var filter {Entity}Filter

    filter.request = request

    ReadJSON(filter.request, &filter)
    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}

func (filter *{Entity}Filter) Get{Entity}Model() {Entity} {

    var model {Entity}

    ReadJSON(filter.request, &model)

    model.Validate()

    return  model
}
`

var usualTemplateWebappEntityType = template{
    Path:    "",
    Content: assignMsName(usualWebappEntityType),
}
