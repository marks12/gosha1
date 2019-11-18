package types

import (
    "strings"
    "newapp/settings"
)

type validator struct {
    validationErrors	[]string
}

func (val *validator) IsValid() bool {

    return len(val.validationErrors) < 1
}

func (val *validator) GetValidationErrors() string {

    return strings.Join(val.validationErrors, ". ")
}

func (val *validator) Validate(functionType string) {

    switch functionType {

    case settings.FunctionTypeFind:
        break

    case settings.FunctionTypeCreate:
        break

    case settings.FunctionTypeMultiCreate:
        break

    case settings.FunctionTypeRead:
        break

    case settings.FunctionTypeUpdate:
        break

    case settings.FunctionTypeMultiUpdate:
        break

    case settings.FunctionTypeDelete:
        break

    case settings.FunctionTypeMultiDelete:
        break

    default:
        val.validationErrors = append(val.validationErrors, "Unsupported function type: " + functionType)
        break
    }
}

