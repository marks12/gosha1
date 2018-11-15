package cmd

const usualWebappErrors = `package webapp

import (
//    "{ms-name}/services"
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

const usualWebappEntity = `
package webapp

import (
    "{ms-name}/logic"
    "{ms-name}/services"
    "net/http"
    "msrpc/mdl"
    "{ms-name}/types"
)


func {entity-name}Find(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest)

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

    validResponse(w, mdl.Response{
        Data:mdl.ResponseFind{
            data,
            totalRecords,
    }})

    return
}

func {entity-name}Create(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest)

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

    validResponse(w, mdl.Response{
        Data:mdl.ResponseCreate{
            data,
    }})

    return
}

func {entity-name}Read(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest)

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
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.Response{
        Data:mdl.ResponseRead{
            data,
        }})

    return
}


func {entity-name}Update(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest)

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

    validResponse(w, mdl.Response{
        Data:mdl.ResponseUpdate{
            data,
        }})

    return
}

func {entity-name}Delete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.Get{entity-name}Filter(httpRequest)

    if !requestDto.IsAuthorized() {
        errResponse(w, "Invalid authorize in {entity-name}Delete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.{entity-name}Delete(requestDto)

    // Создаём структуру ответа
    if err != nil {
        errResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    validResponse(w, mdl.Response{
        Data:mdl.ResponseUpdate{
            data,
        }})

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
