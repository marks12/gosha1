package cmd

type TypeConfig struct {
    IsId bool
}

const usualTypesAuthenticator = `package types

import (
    "{ms-name}/core"
    "{ms-name}/dbmodels"
    "{ms-name}/flags"
    "{ms-name}/settings"
    "{ms-name}/common"
    "net/http"
    "strings"
)

type Access struct {
	Find bool
	Read bool
	Create bool
	Update bool
	Delete bool
	FindOrCreate bool
}

type Authenticator struct {
    Token        string
    functionType string
    urlPath      string
    userId       int
    roleIds      []int
    validator
}

func (auth *Authenticator) GetCurrentUserId() int {
    return auth.userId
}

func (auth *Authenticator) SetCurrentUserId(id int) {
    auth.userId = id
}

func (auth *Authenticator) GetCurrentUserRoleIds() []int {
    return auth.roleIds
}

func (auth *Authenticator) IsCurrentUserAdmin() bool {
    return common.InArray(settings.AdminRoleId, auth.roleIds)
}

func (auth *Authenticator) IsAuthorized() bool {

    if *flags.Auth {
        return true
    }

    if len(auth.Token) < 1 {
        return false
    }

    dbAuth := dbmodels.Auth{}
    core.Db.Where(dbmodels.Auth{Token: auth.Token}).First(&dbAuth)

    if dbAuth.IsActive {

        if dbAuth.UserId < 1 {
            return false
        }

        auth.SetCurrentUserId(dbAuth.UserId)

        userRoles := []dbmodels.UserRole{}
        core.Db.Where(dbmodels.UserRole{UserId: dbAuth.UserId}).Find(&userRoles)

        for _, ur := range userRoles {
            auth.roleIds = append(auth.roleIds, ur.RoleId)
        }

        usedResources := []dbmodels.Resource{}

        core.Db.Where(dbmodels.Resource{
            Code:   clearPath(auth.urlPath),
            TypeId: settings.HttpRouteResourceType,
        }).Find(&usedResources)

        if len(usedResources) < 1 {
            return false
        }

        ids := []int{}

        for _, r := range usedResources {
            ids = append(ids, r.ID)
        }

        roleResources := []dbmodels.RoleResource{}

        core.Db.Model(dbmodels.RoleResource{}).
            Where("role_id in (select role_id from user_roles where deleted_at IS NULL and user_id = ?) and resource_id in (?)", dbAuth.UserId, ids).Find(&roleResources)

        switch auth.functionType {
        case settings.FunctionTypeFind:
            for _, rr := range roleResources {
                if rr.Find {
                    return true
                }
            }
            return false

        case settings.FunctionTypeRead:
            for _, rr := range roleResources {
                if rr.Read {
                    return true
                }
            }
            return false

        case settings.FunctionTypeCreate, settings.FunctionTypeMultiCreate:
            for _, rr := range roleResources {
                if rr.Create {
                    return true
                }
            }
            return false

        case settings.FunctionTypeUpdate, settings.FunctionTypeMultiUpdate:
            for _, rr := range roleResources {
                if rr.Update {
                    return true
                }
            }
            return false

        case settings.FunctionTypeDelete, settings.FunctionTypeMultiDelete:
            for _, rr := range roleResources {
                if rr.Delete {
                    return true
                }
            }
            return false

        case settings.FunctionTypeFindOrCreate:
            for _, rr := range roleResources {
                if rr.FindOrCreate {
                    return true
                }
            }
            return false
        }
    }

    return false
}

func clearPath(s string) string {
    if strings.Count(s, "/") > 3 {
        return s[0:strings.LastIndex(s, "/")]
    }

    return s
}

func (auth *Authenticator) SetToken(r *http.Request) error {

    auth.Token = r.Header.Get("Token")

    return nil
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
    case settings.FunctionTypeMultiCreate:
        break
    case settings.FunctionTypeMultiUpdate:
        break
    case settings.FunctionTypeMultiDelete:
        break
    default:
        authenticator.validator.validationErrors = append(authenticator.validator.validationErrors, "Unsupported function type: "+functionType)
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
        case settings.FunctionTypeRead:
            break;
        case settings.FunctionTypeUpdate:
            break;
        case settings.FunctionTypeDelete:
            break;
        case settings.FunctionTypeFindOrCreate:
            break;
        case settings.FunctionTypeMultiCreate:
            break;
        case settings.FunctionTypeMultiUpdate:
            break;
        case settings.FunctionTypeMultiDelete:
            break;
        default:
            filter.validator.validationErrors = append(filter.validator.validationErrors, "Usupported method")
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
    filter.functionType = functionType
    filter.urlPath = request.URL.Path

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

