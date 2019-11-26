package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type Region struct {
    Id   int
    Name string
	Code string
	TypeId *int
	EntityId *int
	//Region remove this line for disable generator functionality
    
    validator
}

func (region *Region) Validate()  {
    //Validate remove this line for disable generator functionality
}

type RegionFilter struct {
    model Region
    list []Region
    //RegionFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetRegionFilter(request *http.Request, functionType string) RegionFilter {

    var filter RegionFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetRegionFilter remove this line for disable generator functionality

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


func (filter *RegionFilter) GetRegionModel() Region {

    filter.model.Validate()

    return  filter.model
}

func (filter *RegionFilter) GetRegionModelList() (data []Region, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *RegionFilter) SetRegionModel(typeModel Region) {

    filter.model = typeModel
}
