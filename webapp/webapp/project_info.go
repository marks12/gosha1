package webapp

import (
    "net/http"
    "gosha/webapp/mdl"
    "gosha/webapp/types"
    "gosha/settings"
    "gosha/webapp/project"
)

    

func ProjectInfoFind(w http.ResponseWriter, httpRequest *http.Request) {

    requestDto := types.GetProjectInfoFilter(httpRequest, settings.FunctionTypeFind)

    if !requestDto.IsAuthorized() {
        ErrResponse(w, "Invalid authorize in ProjectInfoFind", http.StatusForbidden)
        return
    }

    if !requestDto.IsValid() {
        ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
        return
    }

    // Получаем список
    data, totalRecords, err := logicProjectInfoFind(requestDto)

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

func logicProjectInfoFind(filter types.ProjectInfoFilter) (result []types.ProjectInfo, totalRecords int, err error) {

    id := 0

    id++
    result = append(result, types.ProjectInfo{
        Id: id,
        Name: "ProjectName",
        Value: project.GetName(),
    })

    totalRecords = len(result)

    return
}
