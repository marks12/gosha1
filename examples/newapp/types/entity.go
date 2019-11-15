package types

import (
    "errors"
    "net/http"
    "newapp/settings"
)

type Entity struct {
    Id   int
    Name string
	Description string

    validator
	//Entity remove this line for disable generator functionality
}

func (entity *Entity) Validate()  {
    //Validate remove this line for disable generator functionality
}

type EntityFilter struct {
    model Entity
    list []Entity
    //EntityFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetEntityFilter(request *http.Request, functionType string) EntityFilter {

    var filter EntityFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetEntityFilter remove this line for disable generator functionality

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


func (filter *EntityFilter) GetEntityModel() Entity {

    filter.model.Validate()

    return  filter.model
}

func (filter *EntityFilter) GetEntityModelList() (data []Entity, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *EntityFilter) SetEntityModel(typeModel Entity) {

    filter.model = typeModel
}
