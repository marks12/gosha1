package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type LayoutContent struct {
    Id   int
    LayoutId *int
	ComponentTemplateCode string
	Position int
	IsActive bool
	RegionId *int
	LanguageId *int
	Name string
	//LayoutContent remove this line for disable generator functionality
    
    validator
}

func (layoutContent *LayoutContent) Validate()  {
    //Validate remove this line for disable generator functionality
}

type LayoutContentFilter struct {
    model LayoutContent
    list []LayoutContent
    //LayoutContentFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetLayoutContentFilter(request *http.Request, functionType string) LayoutContentFilter {

    var filter LayoutContentFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetLayoutContentFilter remove this line for disable generator functionality

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


func (filter *LayoutContentFilter) GetLayoutContentModel() LayoutContent {

    filter.model.Validate()

    return  filter.model
}

func (filter *LayoutContentFilter) GetLayoutContentModelList() (data []LayoutContent, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *LayoutContentFilter) SetLayoutContentModel(typeModel LayoutContent) {

    filter.model = typeModel
}
