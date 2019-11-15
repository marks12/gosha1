package webapp

import (
    "newapp/logic"
    "net/http"
    "newapp/mdl"
    "newapp/types"
    "newapp/settings"
)

    

func RoleFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.RoleFind(requestDto)

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

    

func RoleCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeCreate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleCreate(requestDto)

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

    

func RoleRead(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeRead)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleRead", http.StatusForbidden)
        return
    }

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleRead(requestDto)

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

    

func RoleUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeUpdate)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleUpdate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleUpdate(requestDto)

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

    

func RoleDelete(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleDelete", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    isOk, err := logic.RoleDelete(requestDto)

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

    

func RoleFindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetRoleFilter(httpRequest, settings.FunctionTypeDelete)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in RoleFindOrCreate", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, err := logic.RoleFindOrCreate(requestDto)

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

