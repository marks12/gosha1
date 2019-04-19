package types

import (
    "net/http"
)

type Auth struct {

    Email     string
    Password  string
    Token     string
}

func (auth *Auth) Validate()  {

}

type AuthFilter struct {

    model Auth
    AbstractFilter
}

func GetAuthFilter(request *http.Request, functionType string) AuthFilter {

    var filter AuthFilter

    filter.request = request

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}

func (filter *AuthFilter) GetAuthModel() Auth {

    filter.model.Validate()

    return  filter.model
}
