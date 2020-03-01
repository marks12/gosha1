package types

import (
    "net/http"
)

type RoleResource struct {
    Id   int
    RoleId int
	ResourceId int
	Find bool
	Read bool
	Create bool
	Update bool
	Delete bool
	FindOrCreate bool
	UpdateOrCreate bool
	//RoleResource remove this line for disable generator functionality
}

func (roleResource *RoleResource) Validate()  {
    //Validate remove this line for disable generator functionality
}

type RoleResourceFilter struct {
    model RoleResource
    //RoleResourceFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetRoleResourceFilter(request *http.Request, functionType string) RoleResourceFilter {

    var filter RoleResourceFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetRoleResourceFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *RoleResourceFilter) GetRoleResourceModel() RoleResource {

    filter.model.Validate()

    return  filter.model
}

func (filter *RoleResourceFilter) SetRoleResourceModel(typeModel RoleResource) {

    filter.model = typeModel
}
