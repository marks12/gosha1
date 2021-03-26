package types

import (
    "net/http"
    "gosha/settings"
)

type Authenticator struct {
    Token string
    validator
}

func (auth *Authenticator) IsAuthorized() bool {



    return true
}

func (auth *Authenticator) SetToken(r *http.Request) error {

    auth.Token = r.Header.Get("Token")

    return  nil
}

func (authenticator *Authenticator) Validate(functionType string) {

    switch functionType {

    case settings.FunctionTypeFind:

        break;
    case settings.FunctionTypeCreate:



        break;
    case settings.FunctionTypeRead:



        break;
    case settings.FunctionTypeUpdate:



        break;
    case settings.FunctionTypeDelete:

        break;

    default:
        authenticator.validator.validationErrors = append(authenticator.validator.validationErrors, "Usupported function type: " + functionType)
        break;
    }
}
