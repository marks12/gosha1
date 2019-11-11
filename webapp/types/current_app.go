package types

import (
    "net/http"
)

type CurrentApp struct {
    Id   int
    IsValidStructure bool
    //CurrentApp remove this line for disable generator functionality
}

func (currentApp *CurrentApp) Validate()  {
    //Validate remove this line for disable generator functionality
}

type CurrentAppFilter struct {
    model CurrentApp
    Email string
	Password string
	//CurrentAppFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetCurrentAppFilter(request *http.Request, functionType string) CurrentAppFilter {

    var filter CurrentAppFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    filter.Email = request.FormValue("Email")
	filter.Password = request.FormValue("Password")
	//GetCurrentAppFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *CurrentAppFilter) GetCurrentAppModel() CurrentApp {

    filter.model.Validate()

    return  filter.model
}

func (filter *CurrentAppFilter) SetCurrentAppModel(typeModel CurrentApp) {

    filter.model = typeModel
}
