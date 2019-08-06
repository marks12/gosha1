package types

import (
    "net/http"
)

type ProjectInfo struct {
    Id   int
    Name string
    Value string
    //ProjectInfo remove this line for disable generator functionality
}

func (projectInfo *ProjectInfo) Validate()  {
    //Validate remove this line for disable generator functionality
}

type ProjectInfoFilter struct {
    model ProjectInfo
    //ProjectInfoFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetProjectInfoFilter(request *http.Request, functionType string) ProjectInfoFilter {

    var filter ProjectInfoFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetProjectInfoFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *ProjectInfoFilter) GetProjectInfoModel() ProjectInfo {

    filter.model.Validate()

    return  filter.model
}

func (filter *ProjectInfoFilter) SetProjectInfoModel(typeModel ProjectInfo) {

    filter.model = typeModel
}
