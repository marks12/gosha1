package cmd

const usualWebappEntityType = `package types

import (
    "net/http"
)

type {Entity} struct {
    Id   int
    //{Entity} ` + removeLineComment + `
}

func ({entity} *{Entity}) Validate()  {
    //Validate ` + removeLineComment + `
}

type {Entity}Filter struct {
    model {Entity}
    //{Entity}Filter ` + removeLineComment + `

    AbstractFilter
}

func Get{Entity}Filter(request *http.Request, functionType string) {Entity}Filter {

    var filter {Entity}Filter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //Get{Entity}Filter ` + removeLineComment + `

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
