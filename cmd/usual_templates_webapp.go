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

const usualWebappEntity = `package webapp

import (
    "{ms-name}/logic"
    "net/http"
    "{ms-name}/mdl"
    "{ms-name}/types"
    "{ms-name}/settings"
)

func {entity-name}Find(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in Find{entity-name}", http.StatusForbidden)
        return
    }

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

    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}Create", http.StatusForbidden)
        return
    }

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

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1


    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}Read", http.StatusForbidden)
        return
    }

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

    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}Update", http.StatusForbidden)
        return
    }

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

    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}Delete", http.StatusForbidden)
        return
    }

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

var usualTemplateWebappErrors = template{
    Path:    "./webapp/errors.go",
    Content: assignMsName(usualWebappErrors),
}

var usualTemplateWebappEntity = template{
    Path:    "./webapp/{entity-name}.go",
    Content: assignMsName(usualWebappEntity),
}
