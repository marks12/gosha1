package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type PageInfo struct {
    Id   int
    SeoMeta *int
	//PageInfo remove this line for disable generator functionality
    
    validator
}

func (pageInfo *PageInfo) Validate()  {
    //Validate remove this line for disable generator functionality
}

type PageInfoFilter struct {
    model PageInfo
    list []PageInfo
    //PageInfoFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetPageInfoFilter(request *http.Request, functionType string) PageInfoFilter {

    var filter PageInfoFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetPageInfoFilter remove this line for disable generator functionality

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


func (filter *PageInfoFilter) GetPageInfoModel() PageInfo {

    filter.model.Validate()

    return  filter.model
}

func (filter *PageInfoFilter) GetPageInfoModelList() (data []PageInfo, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *PageInfoFilter) SetPageInfoModel(typeModel PageInfo) {

    filter.model = typeModel
}
