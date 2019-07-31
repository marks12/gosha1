package types

import (
    "net/http"
)

type Setting struct {
    Id   int
    //Setting remove this line for disable generator functionality
}

func (setting *Setting) Validate()  {
    //Validate remove this line for disable generator functionality
}

type SettingFilter struct {
    model Setting
    //SettingFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetSettingFilter(request *http.Request, functionType string) SettingFilter {

    var filter SettingFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetSettingFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *SettingFilter) GetSettingModel() Setting {

    filter.model.Validate()

    return  filter.model
}

func (filter *SettingFilter) SetSettingModel(typeModel Setting) {

    filter.model = typeModel
}
