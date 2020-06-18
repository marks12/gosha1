package webapp

import (
    "gosha/settings"
    "gosha/webapp/logic"
    "gosha/webapp/mdl"
    "gosha/webapp/types"
    "net/http"
)

    

func CurrentAppFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentAppFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentAppFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.CurrentAppFind(requestDto)

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
}

    

func CurrentAppCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentAppFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentAppCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentAppCreate(requestDto)

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

    

func CurrentAppRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentAppFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentAppRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentAppRead(requestDto)

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
}
