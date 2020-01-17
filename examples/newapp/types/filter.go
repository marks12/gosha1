package types

import (
    "errors"
    "newapp/settings"
    "net/http"
    "strings"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    "net/url"
)

type FilterIds struct {
    Ids []int
    CurrentId int

    validator
}

func (filter *FilterIds) GetFirstId() (int, error) {
    for _, id := range filter.Ids {
        return id, nil
    }
    return 0, errors.New("Empty array")
}

func (filter *FilterIds) GetIds() []int {
    return filter.Ids
}

func (filter *FilterIds) GetCurrentId() int {
    return filter.CurrentId
}

func (filter *FilterIds) SetCurrentId(id int) int {
    filter.CurrentId = id
	return filter.CurrentId
}

func (filter *FilterIds) AddId(id int) *FilterIds {
    filter.Ids = append(filter.Ids, id)
    return filter
}

func (filter *FilterIds) AddIds(ids []int) *FilterIds {
    for _, id := range ids {
        filter.AddId(id)
    }
    return filter
}

func (filter *FilterIds) ClearIds() *FilterIds {

    filter.Ids = []int{}
    return filter
}

// method find read create update delete
func (filter *FilterIds) Validate(functionType string) {

    switch functionType {
        case settings.FunctionTypeFind:
            break;
        case settings.FunctionTypeCreate:
            break;
        case settings.FunctionTypeMultiCreate:
            break;
        case settings.FunctionTypeRead:
            break;
        case settings.FunctionTypeUpdate:
            break;
        case settings.FunctionTypeMultiUpdate:
            break;
        case settings.FunctionTypeDelete:
            break;
        case settings.FunctionTypeMultiDelete:
            break;
        case settings.FunctionTypeFindOrCreate:
            break;
        case settings.FunctionTypeMultiFindOrCreate:
            break;
        default:
            filter.validator.validationErrors = append(filter.validator.validationErrors, "Unsupported method")
            break;
    }
}


type GoshaSearchFilter struct {

    Search string
    SearchBy []string
}

type GoshaOrderFilter struct {

    Order []string
    OrderDirection []string
}

type AbstractFilter struct {

    request *http.Request
    GoshaSearchFilter
    GoshaOrderFilter
    FilterIds
    Pagination
    validator
    Authenticator
}

func GetAbstractFilter(request *http.Request, functionType string) AbstractFilter {

    var filter AbstractFilter

    filter.request = request
    filter.FunctionType = functionType
    filter.UrlPath = request.URL.Path

    ReadJSON(filter.request, &filter)
    ReadJSON(filter.request, &filter.FilterIds)

    filter.Pagination.CurrentPage, _ = strconv.Atoi(request.FormValue("CurrentPage"))
    filter.Pagination.PerPage, _ = strconv.Atoi(request.FormValue("PerPage"))
    filter.Search = request.FormValue("Search")

    arr, _ := url.ParseQuery(request.URL.RawQuery)

    dirs := []string{}

    for _, dir := range arr["OrderDirection[]"] {

        if strings.ToLower(dir) == "desc" {
            dirs = append(dirs, "desc")
        } else {
            dirs = append(dirs, "asc")
        }
    }

    for index, field := range arr["Order[]"] {

        filter.Order = append(filter.Order, gorm.ToColumnName(field))

        if len(dirs) > index && dirs[index] == "desc" {
            filter.OrderDirection = append(filter.OrderDirection, "desc")
        } else {
            filter.OrderDirection = append(filter.OrderDirection, "asc")
        }
    }

    for _, field := range arr["SearchBy[]"] {
        filter.SearchBy = append(filter.SearchBy, gorm.ToColumnName(field))
    }
    for _, s := range arr["Ids[]"] {
        id, _ := strconv.Atoi(s)
        filter.AddId(id)
    }

    filter.SetToken(request)

    ReadJSON(filter.request, &filter.validator)

    vars := mux.Vars(request)
    id, _ := strconv.Atoi(vars["id"])

    if id > 0 {
        filter.SetCurrentId(id)
    }

    filter.Validate(functionType)

    return filter
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
