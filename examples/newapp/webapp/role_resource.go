package webapp

import (
    "newapp/logic"
    "net/http"
    "newapp/mdl"
    "newapp/types"
    "newapp/settings"
)

    

func RoleResourceFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.RoleResourceFind(requestDto)

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

    

func RoleResourceCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleResourceCreate(requestDto)

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

    

func RoleResourceRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleResourceRead(requestDto)

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

    

func RoleResourceUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleResourceUpdate(requestDto)

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

    

func RoleResourceDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.RoleResourceDelete(requestDto)

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

    

func RoleResourceFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleResourceFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleResourceFindOrCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleResourceFindOrCreate(requestDto)

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

