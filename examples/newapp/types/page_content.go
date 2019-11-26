package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type PageContent struct {
    Id   int
    Name string
	PageTemplateId *int
	IsActive bool
	RegionId *int
	LanguageId *int
	Position int
	ComponentTemplateCode string
	//PageContent remove this line for disable generator functionality
    
    validator
}

func (pageContent *PageContent) Validate()  {
    //Validate remove this line for disable generator functionality
}

type PageContentFilter struct {
    model PageContent
    list []PageContent
    //PageContentFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetPageContentFilter(request *http.Request, functionType string) PageContentFilter {

    var filter PageContentFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetPageContentFilter remove this line for disable generator functionality

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


func (filter *PageContentFilter) GetPageContentModel() PageContent {

    filter.model.Validate()

    return  filter.model
}

func (filter *PageContentFilter) GetPageContentModelList() (data []PageContent, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *PageContentFilter) SetPageContentModel(typeModel PageContent) {

    filter.model = typeModel
}
