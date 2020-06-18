package cmd

var usualTemplateWebappEntityType = template{
    Path:    "",
    Content: GetUsualTemplateTypeContent(TypeConfig{true, false}),
}

func GetUsualTemplateTypeContent(cfg TypeConfig) string {

    uuidImport := ""

    if cfg.IsUuid {
        uuidImport = `
    "github.com/google/uuid"
`
    }

    var usualWebappEntityType = `package types

import (
    "net/http"
    "{ms-name}/settings"` + uuidImport + `
    "errors"
)

type {Entity} struct {
    ` + getTypeId(cfg) + `
    ` + getRemoveLine("{Entity}") + `
    
    validator
}

func ({entity} *{Entity}) Validate()  {
    ` + getRemoveLine("Validate") + `
}

type {Entity}Filter struct {
    model {Entity}
    list []{Entity}
    ` + getRemoveLine("{Entity}Filter") + `

    AbstractFilter
}

func Get{Entity}Filter(request *http.Request, functionType string) {Entity}Filter {

    var filter {Entity}Filter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    ` + getRemoveLine("Get{Entity}Filter") + `

    switch functionType {
    case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
        ReadJSON(request, &filter.list)
        break
    default:
        ReadJSON(request, &filter.model)
        break
    }

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *{Entity}Filter) Get{Entity}Model() {Entity} {

    filter.model.Validate()

    return  filter.model
}

func (filter *{Entity}Filter) Get{Entity}ModelList() (data []{Entity}, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *{Entity}Filter) Set{Entity}Model(typeModel {Entity}) {

    filter.model = typeModel
}
`


    return assignMsName(usualWebappEntityType)
}

func getTypeId(config TypeConfig) string {

    if config.IsId {

        if config.IsUuid {
            return `Id uuid.UUID`
        }
        return `Id   int`
    }

    return ""
}
