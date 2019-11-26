package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type ComponentGroup struct {
    Id   int
    Name string
	Code string
	//ComponentGroup remove this line for disable generator functionality
    
    validator
}

func (componentGroup *ComponentGroup) Validate()  {
    //Validate remove this line for disable generator functionality
}

type ComponentGroupFilter struct {
    model ComponentGroup
    list []ComponentGroup
    //ComponentGroupFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetComponentGroupFilter(request *http.Request, functionType string) ComponentGroupFilter {

    var filter ComponentGroupFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetComponentGroupFilter remove this line for disable generator functionality

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


func (filter *ComponentGroupFilter) GetComponentGroupModel() ComponentGroup {

    filter.model.Validate()

    return  filter.model
}

func (filter *ComponentGroupFilter) GetComponentGroupModelList() (data []ComponentGroup, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *ComponentGroupFilter) SetComponentGroupModel(typeModel ComponentGroup) {

    filter.model = typeModel
}
