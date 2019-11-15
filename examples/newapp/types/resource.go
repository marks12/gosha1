package types

import (
    "net/http"
)

type Resource struct {
    Id   int
    Name string
	Code string
	TypeId int
	//Resource remove this line for disable generator functionality
}

func (resource *Resource) Validate()  {
    //Validate remove this line for disable generator functionality
}

type ResourceFilter struct {
    model Resource
    //ResourceFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetResourceFilter(request *http.Request, functionType string) ResourceFilter {

    var filter ResourceFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetResourceFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *ResourceFilter) GetResourceModel() Resource {

    filter.model.Validate()

    return  filter.model
}

func (filter *ResourceFilter) SetResourceModel(typeModel Resource) {

    filter.model = typeModel
}
