package cmd

type TypeConfig struct {
    IsId bool
}

const usualTypesAuthenticator = `package types

import (
    "{ms-name}/settings"
    "{ms-name}/dbmodels"
    "net/http"
    "{ms-name}/core"
)

type Authenticator struct {
    Token string
    validator
}

func (auth *Authenticator) IsAuthorized() bool {

    if len(auth.Token) < 1 {
        return false
    }

    dbUser := dbmodels.User{}
    core.Db.Where(dbmodels.User{Token: auth.Token}).Find(&dbUser)

    if len(dbUser.Token) < 1 {
        return false
    }

    return  true
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
`

var usualTypesEntity = `package types

import (
    "time"
)
// default entity will used when create new entity
type Entity struct {
    ID        int       ` + "`" + `gorm:"primary_key"` + "`" + `
    ` + getRemoveLine("Entity") + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func (entity *Entity) Validate()  {
    ` + getRemoveLine("Validate") + `
}
`

const usualTypesFilter = `package types

import (
    "errors"
    "{ms-name}/settings"
    "net/http"
    "strings"
    "strconv"
    "github.com/gorilla/mux"
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


type SearchFilter struct {

    Search string
}

type OrderFilter struct {

    Order []string
    OrderDirection []string
}

type AbstractFilter struct {

    request *http.Request
    SearchFilter
    OrderFilter
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
    filter.Search  = request.FormValue("Search")


    dirs := []string{}

    for _, dir := range arr["OrderDirection[]"] {
        if strings.ToLower(dir) == "desc" {
            dirs = append(dirs, "desc")
        } else {
            dirs = append(dirs, "asc")
        }
    }

    for index, field := range arr["Order[]"] {
        filter.Order = append(filter.Order, field)
        if len(dirs) > index && dirs[index] == "desc" {
            filter.OrderDirection = append(filter.OrderDirection, "desc")
        } else {
            filter.OrderDirection = append(filter.OrderDirection, "asc")
        }
    }


    filter.SetToken(request)

    ReadJSON(filter.request, &filter.validator)

    vars := mux.Vars(request)
    id, _ := strconv.Atoi(vars["id"])

    if id > 0 {
        filter.AddId(id)
    }

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
`

const usualTypesRequest= `package types

import (
    "encoding/json"
    "net/http"
)

// ReadJSON -
func ReadJSON(r *http.Request, entity interface{}) {

    decoder := json.NewDecoder(r.Body)
    decoder.Decode(entity)

    defer r.Body.Close()
}
`

const usualTypesValidator = `package types

import (
    "strings"
    "{ms-name}/settings"
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

    case settings.FunctionTypeRead:
        break

    case settings.FunctionTypeUpdate:
        break

    case settings.FunctionTypeDelete:
        break

    default:
        val.validationErrors = append(val.validationErrors, "Usupported function type: " + functionType)
        break
    }
}

`

const usualTypesResponse = `package types

import (
	"{ms-name}/settings"
)

type APIStatus struct {
	Status string
}

type APIError struct {
	Error bool
	ErrorMessage string
}

type Pagination struct {

	TotalRecords	int
	TotalPages		int
	CurrentPage		int
	PerPage			int

	validator
}

func (pagination *Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}

func (pagination *Pagination) Validate(functionType string) {

	switch functionType {

	case settings.FunctionTypeFind:

		if pagination.CurrentPage < 1 {
			pagination.validator.validationErrors = append(pagination.validator.validationErrors, "CurrentPage parameter is not set")
		}

		if pagination.PerPage < 1 {
			pagination.validator.validationErrors = append(pagination.validator.validationErrors, "PerPage parameter is not set")
		}

		break
	case settings.FunctionTypeCreate:



		break
	case settings.FunctionTypeRead:



		break
	case settings.FunctionTypeUpdate:



		break
	case settings.FunctionTypeDelete:

		break

	default:
		pagination.validator.validationErrors = append(pagination.validator.validationErrors, "Usupported function type: " + functionType)
		break
	}
}
`

var usualTemplateTypesAuthenticator = template{
    Path:    "./types/authenticator.go",
    Content: assignMsName(usualTypesAuthenticator),
}

var usualTemplateTypesEntity = template{
    Path:    "./types/entity.go",
    Content: usualTypesEntity,
}

var usualTemplateTypesFilter = template{
    Path:    "./types/filter.go",
    Content: assignMsName(usualTypesFilter),
}

var usualTemplateTypesRequest = template{
    Path:    "./types/request.go",
    Content: usualTypesRequest,
}

var usualTemplateTypesResponse = template{
    Path:    "./types/response.go",
    Content: assignMsName(usualTypesResponse),
}

var usualTemplateTypesValidator = template{
    Path:    "./types/validator.go",
    Content: assignMsName(usualTypesValidator),
}

