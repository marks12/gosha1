package cmd

const usualWebappErrors = `package webapp

import (
    "{ms-name}/types"
    "net/http"
    "encoding/json"
)

func errResponse (w http.ResponseWriter, err string, status int) {

    response := types.APIError{}

    response.Error = true
    response.ErrorMessage = err

    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)

    return
}

func validResponse (w http.ResponseWriter, data interface{}) {

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)

    return
}`


var usualTemplateWebappErrors = template{
    Path:    "./webapp/errors.go",
    Content: assignMsName(usualWebappErrors),
}

var usualTemplateWebappEntity = template{
    Path:    "./webapp/{entity-name}.go",
    Content: assignMsName(GetUsualTemplateWebAppContent(Crud{true, true, true, true, true})),
}

func GetUsualTemplateWebAppContent(crud Crud) string {

    var usualWebappEntity = `package webapp

import (
    "{ms-name}/logic"
    "net/http"
    "{ms-name}/mdl"
    "{ms-name}/types"
    "{ms-name}/settings"
)

func {entity-name}Find(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeFind)

    ` + getAuth("Find", crud) + `

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.{entity-name}Find(requestDto)

    // Создаём структуру ответа
    if err != nil {
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.ResponseFind{
            data,
            totalRecords,
    })

    return
}

func {entity-name}Create(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeCreate)

    ` + getAuth("Create", crud) + `

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Create(requestDto)

    // Создаём структуру ответа
    if err != nil {
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.ResponseCreate{
        data,
    })

    return
}

func {entity-name}Read(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeRead)

    ` + getAuth("Read", crud) + `

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
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
        errResponse(w, err.Error(), code)
        return
    }

    validResponse(w, mdl.ResponseRead{
        data,
    })

    return
}


func {entity-name}Update(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeUpdate)

    ` + getAuth("Update", crud) + `

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Update(requestDto)

    // Создаём структуру ответа
    if err != nil {
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.ResponseUpdate{
        data,
    })

    return
}

func {entity-name}Delete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeDelete)

    ` + getAuth("Delete", crud) + `

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.{entity-name}Delete(requestDto)

    // Создаём структуру ответа
    if err != nil {
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.ResponseDelete{
        isOk,
    })

    return
}
`
    return usualWebappEntity
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
    }

    return `if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}` + method +`", http.StatusForbidden)
        return
    }`
}