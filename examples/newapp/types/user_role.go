package types

import (
    "net/http"
)

type UserRole struct {
    Id   int
    UserId int
	RoleId int
    //UserRole remove this line for disable generator functionality
}

func (userRole *UserRole) Validate()  {
    //Validate remove this line for disable generator functionality
}

type UserRoleFilter struct {
    model UserRole
    //UserRoleFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetUserRoleFilter(request *http.Request, functionType string) UserRoleFilter {

    var filter UserRoleFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetUserRoleFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *UserRoleFilter) GetUserRoleModel() UserRole {

    filter.model.Validate()

    return  filter.model
}

func (filter *UserRoleFilter) SetUserRoleModel(typeModel UserRole) {

    filter.model = typeModel
}
