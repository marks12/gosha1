package webapp

import (
    "newapp/core"
    "newapp/logic"
    "net/http"
    "newapp/mdl"
    "newapp/types"
    "newapp/settings"
)

    

func PageInfoFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeFind)

    //if !requestDto.IsAuthorized() {
    //    ErrResponse(w, "Invalid authorize in PageInfoFind", http.StatusForbidden)
    //    return
    //}

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.PageInfoFind(requestDto)

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

    


func PageInfoMultiCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeMultiCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoMultiCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoMultiCreate(requestDto)

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

func PageInfoCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoCreate(requestDto, core.Db)

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

    

func PageInfoRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoRead(requestDto)

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

    


func PageInfoMultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeMultiUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoMultiUpdate(requestDto)

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

func PageInfoUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoUpdate(requestDto, core.Db)

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

    

func PageInfoMultiDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeMultiDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.PageInfoMultiDelete(requestDto)

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

func PageInfoDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.PageInfoDelete(requestDto, core.Db)

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

    

func PageInfoFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetPageInfoFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in PageInfoFindOrCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.PageInfoFindOrCreate(requestDto)

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

