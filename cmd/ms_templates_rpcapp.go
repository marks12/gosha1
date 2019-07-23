package cmd

const msRpcappErrors = `package rpcapp

import (
    "msrpc/mdl"
)

func errMakeResponse (err string) (response mdl.Response) {

    response.Error = true
    response.ErrorMessage = err

    return
}`

const msRpcappEntity = `
package rpcapp

import (
    "msrpc/api"
    "msrpc/mdl"
    "{ms-name}/logic"
)

func Find{entity-name}(cmdRequest api.CommandRequest) (mdl.Response) {

    if cmdRequest.IsAuthorized {

        if !cmdRequest.Pagination.IsValid { return errMakeResponse("Invalid pagination")}

        filter := api.Get{entity-name}Filter(cmdRequest)

        // Получаем список
        data, totalRecords, err := logic.Find{entity-name}(cmdRequest.RequestApplication, filter)

        // Создаём структуру ответа
        if err != nil {
            return errMakeResponse(err.Error())
        }

        return mdl.Response{
            Data:mdl.ResponseFind{
            data,
            totalRecords,
        }}
    }

    return errMakeResponse("Invalid authorize in Find{entity-name}")
}

func Create{entity-name}(cmdRequest api.CommandRequest) (mdl.Response) {

    if cmdRequest.IsAuthorized {

        model := api.Get{entity-name}Model(cmdRequest)

        if ! model.IsValid() {

            return errMakeResponse(model.GetValidationErrors())
        }

        data, err := logic.Create{entity-name}(cmdRequest.RequestApplication, model)

        // Создаём структуру ответа
        if err != nil {
            return errMakeResponse(err.Error())
        }

        return mdl.Response{Data:data}
    }

    return errMakeResponse("Invalid authorize in Create{entity-name}")
}

func Read{entity-name}(cmdRequest api.CommandRequest) (mdl.Response) {

    if cmdRequest.IsAuthorized {

        filter := api.Get{entity-name}Filter(cmdRequest)

        data, err := logic.Read{entity-name}(cmdRequest.RequestApplication, filter)

        // Создаём структуру ответа
        if err != nil {
            code := http.StatusBadRequest
            if err.Error() == "Not found" {
                code = http.StatusNotFound
            }
            ErrResponse(w, err.Error(), code)
            return
        }

        return mdl.Response{Data: data}
    }

    return errMakeResponse("Invalid authorize in Read{entity-name}")
}

func Update{entity-name}(cmdRequest api.CommandRequest) (mdl.Response) {

    if cmdRequest.IsAuthorized {

        model := api.Get{entity-name}Model(cmdRequest)

        if ! model.IsValid() {

            return errMakeResponse(model.GetValidationErrors())
        }

        data, err := logic.Update{entity-name}(cmdRequest.RequestApplication, model)

        if err != nil {
            return errMakeResponse(err.Error())
        }

        return mdl.Response{Data: data}
    }

    return errMakeResponse("Invalid authorize in Update{entity-name}")
}


func Delete{entity-name}(cmdRequest api.CommandRequest) (mdl.Response) {

    if cmdRequest.IsAuthorized {

        filter := api.Get{entity-name}Filter(cmdRequest)

        err := logic.Delete{entity-name}(cmdRequest.RequestApplication, filter)

        if err != nil {

            return errMakeResponse(err.Error())
        }

        return mdl.Response{Data: true}
    }

    return errMakeResponse("Invalid authorize in Delete{entity-name}")
}
`

var msTemplateRpcappErrors = template{
    Path:    "./rpcapp/errors.go",
    Content: msRpcappErrors,
}

var msTemplateRpcappEntity = template{
    Path:    "./rpcapp/{entity-name}.go",
    Content: assignMsName(msRpcappEntity),
}
