package types

import (
    "net/http"
    "newapp/settings"
    "errors"
)

type RegionType struct {
    Id   int
    Name string
	//RegionType remove this line for disable generator functionality
    
    validator
}

func (regionType *RegionType) Validate()  {
    //Validate remove this line for disable generator functionality
}

type RegionTypeFilter struct {
    model RegionType
    list []RegionType
    //RegionTypeFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetRegionTypeFilter(request *http.Request, functionType string) RegionTypeFilter {

    var filter RegionTypeFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetRegionTypeFilter remove this line for disable generator functionality

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


func (filter *RegionTypeFilter) GetRegionTypeModel() RegionType {

    filter.model.Validate()

    return  filter.model
}

func (filter *RegionTypeFilter) GetRegionTypeModelList() (data []RegionType, err error) {

    for k, _ := range filter.list {
        filter.list[k].Validate()

        if ! filter.list[k].IsValid() {
            err = errors.New(filter.list[k].GetValidationErrors())
            break
        }
    }

    return  filter.list, nil
}

func (filter *RegionTypeFilter) SetRegionTypeModel(typeModel RegionType) {

    filter.model = typeModel
}
