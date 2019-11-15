package types

import (
    "net/http"
)

type User struct {
    Id   int
    Email       string
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string
    Token       string
    //User remove this line for disable generator functionality
}

func (user *User) Validate()  {
    //Validate remove this line for disable generator functionality
}

type UserFilter struct {
    model User
    //UserFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetUserFilter(request *http.Request, functionType string) UserFilter {

    var filter UserFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetUserFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *UserFilter) GetUserModel() User {

    filter.model.Validate()

    return  filter.model
}

func (filter *UserFilter) SetUserModel(typeModel User) {

    filter.model = typeModel
}
