package types

import (
    "errors"
    "net/http"
    "strings"
    "strconv"
    "github.com/gorilla/mux"
    "gosha/settings"
)

type FilterIds struct {
    ids []int

    validator
}

func (filter *FilterIds) GetFirstId() (int, error) {
    for _, id := range filter.ids {
        return id, nil
    }
    return 0, errors.New("Empty array")
}

func (filter *FilterIds) GetIds() []int {
    return filter.ids
}

func (filter *FilterIds) AddId(id int) *FilterIds {
    filter.ids = append(filter.ids, id)
    return filter
}

func (filter *FilterIds) AddIds(ids []int) *FilterIds {
    for _, id := range ids {
        filter.AddId(id)
    }
    return filter
}

func (filter *FilterIds) Clear() *FilterIds {

    filter.ids = []int{}
    return filter
}

// method find read create update delete
func (filter *FilterIds) Validate(functionType string) {

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
            filter.validator.validationErrors = append(filter.validator.validationErrors, "Usupported method")
            break;
    }
}


type AbstractFilter struct {

    request *http.Request
    functionType string

    FilterIds
    Pagination
    validator
    Authenticator
}

func GetAbstractFilter(request *http.Request, functionType string) AbstractFilter {

    var filter AbstractFilter

    filter.request = request

    ReadJSON(filter.request, &filter)
    ReadJSON(filter.request, &filter.FilterIds)

    filter.Pagination.CurrentPage,_  = strconv.Atoi(request.FormValue("CurrentPage"))
    filter.Pagination.PerPage,_  = strconv.Atoi(request.FormValue("PerPage"))

    filter.SetToken(request)

    ReadJSON(filter.request, &filter.validator)

    vars := mux.Vars(request)
    id, _ := strconv.Atoi(vars["id"])

    if id > 0 {
        filter.AddId(id)
    }

    filter.SetFunctionType(functionType)

    filter.Validate(functionType)

    return  filter
}

func (filter *AbstractFilter) IsValid() bool  {

    return  filter.FilterIds.IsValid() &&
        filter.Pagination.IsValid() &&
        filter.validator.IsValid() &&
        filter.Authenticator.IsValid()
}

func (filter *AbstractFilter) Validate(functionType string)  {

    filter.FilterIds.Validate(functionType)
    filter.Pagination.Validate(functionType)
    filter.validator.Validate(functionType)
    filter.Authenticator.Validate(functionType)
}

func (filter *AbstractFilter) GetValidationErrors() string  {

    return strings.Join([]string{
        filter.FilterIds.GetValidationErrors(),
        filter.Pagination.GetValidationErrors(),
        filter.validator.GetValidationErrors(),
        filter.Authenticator.GetValidationErrors(),
    }, ". ")
}

func (filter *AbstractFilter) SetFunctionType(t string) *AbstractFilter {

    switch t {
    case settings.FunctionTypeFind:
        filter.functionType = t
        break;
    case settings.FunctionTypeCreate:
        filter.functionType = t
        break;
    case settings.FunctionTypeRead:
        filter.functionType = t
        break;
    case settings.FunctionTypeUpdate:
        filter.functionType = t
        break;
    case settings.FunctionTypeDelete:
        filter.functionType = t
        break;
    default:
        break;
    }

    return filter
}

func (filter *AbstractFilter) GetFunctionType() string  {
    return filter.functionType
}

func (filter *AbstractFilter) GetMethod() string  {
    return filter.GetFunctionType()
}

func (filter *AbstractFilter) SetMethod(t string) *AbstractFilter {
    return filter.SetFunctionType(t)
}
