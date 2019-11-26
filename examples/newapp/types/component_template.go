package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type ComponentTemplate struct {
    Id   int
    Name string
	Code string
	Path string
	GroupCode string
	GroupId *int
	//ComponentTemplate remove this line for disable generator functionality
    
    validator
}

func (componentTemplate *ComponentTemplate) Validate()  {
    //Validate remove this line for disable generator functionality
}

type ComponentTemplateFilter struct {
    model ComponentTemplate
    list []ComponentTemplate
    //ComponentTemplateFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetComponentTemplateFilter(request *http.Request, functionType string) ComponentTemplateFilter {

    var filter ComponentTemplateFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetComponentTemplateFilter remove this line for disable generator functionality

    switch functionType {
    case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
        ReadJSON(request, &filter.list)
        break
    default:
        ReadJSON(request, &filter.model)
        break
    }

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *ComponentTemplateFilter) GetComponentTemplateModel() ComponentTemplate {

    filter.model.Validate()

    return  filter.model
}

func (filter *ComponentTemplateFilter) GetComponentTemplateModelList() (data []ComponentTemplate, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *ComponentTemplateFilter) SetComponentTemplateModel(typeModel ComponentTemplate) {

    filter.model = typeModel
}
