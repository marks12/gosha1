package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type PageTemplate struct {
    Id   int
    PageTypeId *int
	Name string
	IsActive bool
	RegionId *int
	LanguageId *int
	RootPageId *int
	LayoutId *int
	//PageTemplate remove this line for disable generator functionality
    
    validator
}

func (pageTemplate *PageTemplate) Validate()  {
    //Validate remove this line for disable generator functionality
}

type PageTemplateFilter struct {
    model PageTemplate
    list []PageTemplate
    //PageTemplateFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetPageTemplateFilter(request *http.Request, functionType string) PageTemplateFilter {

    var filter PageTemplateFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetPageTemplateFilter remove this line for disable generator functionality

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


func (filter *PageTemplateFilter) GetPageTemplateModel() PageTemplate {

    filter.model.Validate()

    return  filter.model
}

func (filter *PageTemplateFilter) GetPageTemplateModelList() (data []PageTemplate, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *PageTemplateFilter) SetPageTemplateModel(typeModel PageTemplate) {

    filter.model = typeModel
}
