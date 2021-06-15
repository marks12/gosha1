package types

import (
    "net/http"
)

type SelfUpdate struct {
    Id   int
    IsSuccess bool
	//SelfUpdate remove this line for disable generator functionality
    
    validator
}

func (selfUpdate *SelfUpdate) Validate()  {
    //Validate remove this line for disable generator functionality
}

type SelfUpdateFilter struct {
    model SelfUpdate
    list []SelfUpdate
    //SelfUpdateFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetSelfUpdateFilter(request *http.Request, functionType string) SelfUpdateFilter {

    var filter SelfUpdateFilter

    filter.request = request

    //GetSelfUpdateFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return filter
}

func (filter *SelfUpdateFilter) GetSelfUpdateModel() SelfUpdate {

    filter.model.Validate()

    return filter.model
}
