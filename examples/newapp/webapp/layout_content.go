package webapp

import (
    "newapp/core"
    "newapp/logic"
    "net/http"
    "newapp/mdl"
    "newapp/types"
    "newapp/settings"
)

    

func LayoutContentFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.LayoutContentFind(requestDto)

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

    


func LayoutContentMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeMultiCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentMultiCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentMultiCreate(requestDto)

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

func LayoutContentCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentCreate(requestDto, core.Db)

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

    

func LayoutContentRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentRead(requestDto)

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

    


func LayoutContentMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeMultiUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentMultiUpdate(requestDto)

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

func LayoutContentUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentUpdate(requestDto, core.Db)

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

    

func LayoutContentMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeMultiDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.LayoutContentMultiDelete(requestDto)

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

func LayoutContentDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.LayoutContentDelete(requestDto, core.Db)

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

    

func LayoutContentFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetLayoutContentFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in LayoutContentFindOrCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.LayoutContentFindOrCreate(requestDto)

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

