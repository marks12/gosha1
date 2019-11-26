package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type PageType struct {
    Id   int
    Name string
	Code string
	//PageType remove this line for disable generator functionality
    
    validator
}

func (pageType *PageType) Validate()  {
    //Validate remove this line for disable generator functionality
}

type PageTypeFilter struct {
    model PageType
    list []PageType
    //PageTypeFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetPageTypeFilter(request *http.Request, functionType string) PageTypeFilter {

    var filter PageTypeFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetPageTypeFilter remove this line for disable generator functionality

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


func (filter *PageTypeFilter) GetPageTypeModel() PageType {

    filter.model.Validate()

    return  filter.model
}

func (filter *PageTypeFilter) GetPageTypeModelList() (data []PageType, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *PageTypeFilter) SetPageTypeModel(typeModel PageType) {

    filter.model = typeModel
}
