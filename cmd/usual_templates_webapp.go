package cmd

const usualWebappErrors = `package webapp

import (
    "{ms-name}/types"
    "{ms-name}/mdl"
    "net/http"
    "encoding/json"
)

func ErrResponse (w http.ResponseWriter, err string, status int) {

    response := types.APIError{}

    response.Error = true
    response.ErrorMessage = err

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)

    return
}

func ValidResponse (w http.ResponseWriter, data interface{}) {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    switch data.(type) {
    case mdl.ResponseCreate:
        w.WriteHeader(http.StatusCreated)
        break
    default:
        w.WriteHeader(http.StatusOK)
        break
    }
    json.NewEncoder(w).Encode(data)

    return
}`

var usualTemplateWebappErrors = template{
    Path:    "./webapp/errors.go",
    Content: assignMsName(usualWebappErrors),
}

var usualTemplateWebappEntity = template{
    Path:    "./webapp/{entity-name}.go",
    Content: assignMsName(GetUsualTemplateWebAppContent(
        Crud{true, true, true, true, true, true},
        Crud{true, true, true, true, true, true},
    )),
}

func GetUsualTemplateWebAppContent(authCrud Crud, methodsCrud Crud) string {

    var usualWebappEntity = `package webapp

import (
    "{ms-name}/core"
    "{ms-name}/logic"
    "net/http"
    "{ms-name}/mdl"
    "{ms-name}/types"
    "{ms-name}/settings"
)

    ` + getWebappFind(methodsCrud, authCrud) + `

    ` + getWebappCreate(methodsCrud, authCrud) + `

    ` + getWebappRead(methodsCrud, authCrud) + `

    ` + getWebappUpdate(methodsCrud, authCrud) + `

    ` + getWebappDelete(methodsCrud, authCrud) + `

    ` + getWebappFindOrCreate(methodsCrud, authCrud) + `

`
    return usualWebappEntity
}

func getWebappFind(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsFind {
        c = `

func {entity-name}Find(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeFind)

    ` + getAuth("Find", authCrud) + `

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.{entity-name}Find(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseFind{
            data,
            totalRecords,
    })

    return
}`
    }

    return
}

func getWebappCreate(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsCreate {
        c = `
func {entity-name}MultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeMultiCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in {entity-name}MultiCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}MultiCreate(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseCreate{
        data,
    })

    return
}

func {entity-name}Create(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeCreate)

    ` + getAuth("Create", authCrud) + `

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Create(requestDto, core.Db)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseCreate{
        data,
    })

    return
}`
    }

    return
}

func getWebappRead(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsRead {
        c = `

func {entity-name}Read(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeRead)

    ` + getAuth("Read", authCrud) + `

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Read(requestDto)

    // Создаём структуру ответа
    if err != nil {
        code := http.StatusBadRequest
        if err.Error() == "Not found" {
            code = http.StatusNotFound
        }
        ErrResponse(w, err.Error(), code)
        return
    }

    ValidResponse(w, mdl.ResponseRead{
        data,
    })

    return
}`
    }

    return
}

func getWebappUpdate(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsUpdate {
        c = `


func {entity-name}MultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeMultiUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in {entity-name}Update", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}MultiUpdate(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseUpdate{
        data,
    })

    return
}

func {entity-name}Update(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeUpdate)

    ` + getAuth("Update", authCrud) + `

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Update(requestDto, core.Db)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseUpdate{
        data,
    })

    return
}`
    }

    return
}

func getWebappDelete(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsDelete {
        c = `

func {entity-name}MultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeMultiDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in {entity-name}Delete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.{entity-name}MultiDelete(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseDelete{
        isOk,
    })

    return
}

func {entity-name}Delete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeDelete)

    ` + getAuth("Delete", authCrud) + `

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.{entity-name}Delete(requestDto, core.Db)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseDelete{
        isOk,
    })

    return
}`
    }

    return
}


func getWebappFindOrCreate(methodCrud Crud, authCrud Crud) (c string) {

    if methodCrud.IsFindOrCreate {
        c = `

func {entity-name}FindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeDelete)

    ` + getAuth("FindOrCreate", authCrud) + `

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}FindOrCreate(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseFindOrCreate{
        data,
    })

    return
}`
    }

    return
}

func getAuth(method string, crud Crud) (auth string) {

    switch method {

        case "Find":
            if ! crud.IsFind {
                return
            }
            break

        case "Create":
            if ! crud.IsCreate {
                return
            }
            break

        case "Read":
            if ! crud.IsRead {
                return
            }
            break

        case "Update":
            if ! crud.IsUpdate {
                return
            }
            break

        case "Delete":
            if ! crud.IsDelete {
                return
            }
            break

        case "FindOrCreate":
            if ! crud.IsFindOrCreate {
                return
            }
            break
    }

    return `if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in {entity-name}` + method +`", http.StatusForbidden)
        return
    }`
}
