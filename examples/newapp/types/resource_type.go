package types

import (
    "net/http"
)

type ResourceType struct {
    Id   int
    Name string
	//ResourceType remove this line for disable generator functionality
}

func (resourceType *ResourceType) Validate()  {
    //Validate remove this line for disable generator functionality
}

type ResourceTypeFilter struct {
    model ResourceType
    //ResourceTypeFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetResourceTypeFilter(request *http.Request, functionType string) ResourceTypeFilter {

    var filter ResourceTypeFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetResourceTypeFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *ResourceTypeFilter) GetResourceTypeModel() ResourceType {

    filter.model.Validate()

    return  filter.model
}

func (filter *ResourceTypeFilter) SetResourceTypeModel(typeModel ResourceType) {

    filter.model = typeModel
}
