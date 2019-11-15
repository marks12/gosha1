package types

import (
    "net/http"
)

type Role struct {
    Id   int
    Name string
	Description string
	//Role remove this line for disable generator functionality
}

func (role *Role) Validate()  {
    //Validate remove this line for disable generator functionality
}

type RoleFilter struct {
    model Role
    //RoleFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetRoleFilter(request *http.Request, functionType string) RoleFilter {

    var filter RoleFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetRoleFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *RoleFilter) GetRoleModel() Role {

    filter.model.Validate()

    return  filter.model
}

func (filter *RoleFilter) SetRoleModel(typeModel Role) {

    filter.model = typeModel
}
