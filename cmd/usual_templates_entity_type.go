package cmd

const usualWebappEntityType = `package types

import (
    "net/http"
    "strconv"
)

type {Entity} struct {
    Id   int
    Name string
}

func ({entity} *{Entity}) Validate()  {

}

type {Entity}Filter struct {

    TestFilter int
    model {Entity}
    AbstractFilter
}

func Get{Entity}Filter(request *http.Request, functionType string) {Entity}Filter {

    var filter {Entity}Filter

    filter.request = request
    filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *{Entity}Filter) Get{Entity}Model() {Entity} {

    filter.model.Validate()

    return  filter.model
}
`

var usualTemplateWebappEntityType = template{
    Path:    "",
    Content: assignMsName(usualWebappEntityType),
}
