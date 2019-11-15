package types

import (
    "net/http"
)

type CurrentUser struct {
    Id   int
    //CurrentUser remove this line for disable generator functionality
}

func (currentUser *CurrentUser) Validate()  {
    //Validate remove this line for disable generator functionality
}

type CurrentUserFilter struct {
    model CurrentUser
    //CurrentUserFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetCurrentUserFilter(request *http.Request, functionType string) CurrentUserFilter {

    var filter CurrentUserFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetCurrentUserFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *CurrentUserFilter) GetCurrentUserModel() CurrentUser {

    filter.model.Validate()

    return  filter.model
}

func (filter *CurrentUserFilter) SetCurrentUserModel(typeModel CurrentUser) {

    filter.model = typeModel
}
