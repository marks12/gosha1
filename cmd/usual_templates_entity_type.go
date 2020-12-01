package cmd

import "gosha/mode"

func GetUsualTemplateTypeContent(cfg TypeConfig) string {

    uuidImport := ""

    if mode.GetUuidMode() {
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

func Get{Entity}Filter(request *http.Request, functionType string) (filter {Entity}Filter, err error) {

    filter.request = request
	filter.rawRequestBody, err = GetRawBodyContent(request)
    if err != nil {
        return filter, err
    }
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    ` + getRemoveLine("Get{Entity}Filter") + `

    switch functionType {
    case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
        err = ReadJSON(filter.rawRequestBody, &filter.list)
		if err != nil {
			return
		}
        break
    default:
        err = ReadJSON(filter.rawRequestBody, &filter.model)
		if err != nil {
			return
		}
        break
    }

    filter.AbstractFilter, err = GetAbstractFilter(request, filter.rawRequestBody, functionType)

    return  filter, err
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

        if mode.GetUuidMode() {
            return `Id uuid.UUID`
        }
        return `Id   int`
    }

    return ""
}
