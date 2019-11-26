package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type Layout struct {
    Id   int
    Name string
	RegionId int
	LanguageId int
	//Layout remove this line for disable generator functionality
    
    validator
}

func (layout *Layout) Validate()  {
    //Validate remove this line for disable generator functionality
}

type LayoutFilter struct {
    model Layout
    list []Layout
    //LayoutFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetLayoutFilter(request *http.Request, functionType string) LayoutFilter {

    var filter LayoutFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetLayoutFilter remove this line for disable generator functionality

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


func (filter *LayoutFilter) GetLayoutModel() Layout {

    filter.model.Validate()

    return  filter.model
}

func (filter *LayoutFilter) GetLayoutModelList() (data []Layout, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *LayoutFilter) SetLayoutModel(typeModel Layout) {

    filter.model = typeModel
}
