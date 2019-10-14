package webapp

import (
    "gosha/webapp/logic"
    "net/http"
    "gosha/webapp/mdl"
    "gosha/webapp/types"
    "gosha/settings"
)

    

func FieldTypeFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetFieldTypeFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logic.FieldTypeFind(requestDto)

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


