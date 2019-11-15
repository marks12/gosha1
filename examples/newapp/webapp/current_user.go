package webapp

import (
    "newapp/logic"
    "net/http"
    "newapp/mdl"
    "newapp/types"
    "newapp/settings"
)

    

func CurrentUserFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.CurrentUserFind(requestDto)

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

    

func CurrentUserCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentUserCreate(requestDto)

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

    

func CurrentUserRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentUserRead(requestDto)

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

    

func CurrentUserUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentUserUpdate(requestDto)

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

    

func CurrentUserDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.CurrentUserDelete(requestDto)

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

    

func CurrentUserFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetCurrentUserFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in CurrentUserFindOrCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.CurrentUserFindOrCreate(requestDto)

    // Создаём структуру ответа
    if err != nil {
        ErrResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    ValidResponse(w, mdl.ResponseFindOrCreate{
        data,
    })

    return
}

