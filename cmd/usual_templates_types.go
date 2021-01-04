package cmd

import "gosha/mode"

type TypeConfig struct {
    IsId bool
}

var usualTypesAuthenticator = `package types

import (
    "{ms-name}/core"
    "{ms-name}/dbmodels"
    "{ms-name}/flags"
    "{ms-name}/settings"
    "{ms-name}/common"
    "{ms-name}/errors"
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
	UpdateOrCreate bool
}

type Authenticator struct {
    Token        string
    functionType string
    urlPath      string
	ip           string
	maxPerPage   int
	user         dbmodels.User
	auth         dbmodels.Auth
	userId       {ID}
    roleIds      []{ID}
    validator
}

func (auth *Authenticator) GetCurrentUserId() {ID} {
    return auth.userId
}

func (auth *Authenticator) SetIp(r *http.Request) {
	auth.ip = r.Header.Get("X-Forwarded-For")
}

func (auth *Authenticator) GetIp() string {
	return auth.ip
}

func (auth *Authenticator) SetMaxPerPage(i int) {
	auth.maxPerPage = i
}
func (auth *Authenticator) GetMaxPerPage() int {
	return auth.maxPerPage
}

func (auth *Authenticator) SetCurrentUserId(id {ID}) {
    auth.userId = id
}

func (auth *Authenticator) GetCurrentUserRoleIds() []{ID} {
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

        if dbAuth.UserId {GetIdIsNotValidExp} {
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
            TypeId: settings.HttpRouteResourceType` + GetConfigConverter(mode.GetUuidMode()) + `,
        }).Find(&usedResources)

        if len(usedResources) < 1 {
            return false
        }

        ids := []{ID}{}

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

        case settings.FunctionTypeUpdateOrCreate:
            for _, rr := range roleResources {
                if rr.UpdateOrCreate {
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
        authenticator.validator.AddValidationError("Unsupported function type: " + functionType, errors.ErrorCodeUnsupportedFunctionType, "")
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
    "{ms-name}/settings"
    "{ms-name}/errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type GoshaFilterIds struct {
	Ids       []int
	ExceptIds []int
	currentId int

	validator
}

func (filter *GoshaFilterIds) GetFirstId() (int, error) {
	for _, id := range filter.Ids {
		return id, nil
	}
	return 0, errors.New("Empty array")
}

func (filter *GoshaFilterIds) GetCurrentId() int {
	return filter.currentId
}

func (filter *GoshaFilterIds) SetCurrentId(id int) int {
	filter.currentId = id
	return filter.currentId
}

func (filter *GoshaFilterIds) GetIds() []int {
	return filter.Ids
}

func (filter *GoshaFilterIds) GetExceptIds() []int {
	return filter.ExceptIds
}

func (filter *GoshaFilterIds) AddId(id int) *GoshaFilterIds {
	filter.Ids = append(filter.Ids, id)
	return filter
}

func (filter *GoshaFilterIds) AddExceptIds(id int) *GoshaFilterIds {
	filter.ExceptIds = append(filter.ExceptIds, id)
	return filter
}

func (filter *GoshaFilterIds) AddIds(ids []int) *GoshaFilterIds {
	for _, id := range ids {
		filter.AddId(id)
	}
	return filter
}

func (filter *GoshaFilterIds) Clear() *GoshaFilterIds {
	filter.Ids = []int{}
	return filter
}

func (filter *GoshaFilterIds) ClearIds() *GoshaFilterIds {
	filter.Ids = []int{}
	return filter
}

func (filter *GoshaFilterIds) ClearExceptId() *GoshaFilterIds {
	filter.ExceptIds = []int{}
	return filter
}

// method find read create update delete
func (filter *GoshaFilterIds) Validate(functionType string) {

	switch functionType {
	case settings.FunctionTypeFind:

		break
	case settings.FunctionTypeCreate:

		break
	case settings.FunctionTypeRead:
		if len(filter.GetIds()) != 1 || filter.GetIds()[0] < 1 {
			filter.AddValidationError("Error parse Id", errors.ErrorCodeParseId, "Id")
		}
		break
	case settings.FunctionTypeUpdate:
		if len(filter.GetIds()) != 1 || filter.GetIds()[0] < 1 {
			filter.AddValidationError("Error parse Id", errors.ErrorCodeParseId, "Id")
		}
		break
	case settings.FunctionTypeDelete:
		if len(filter.GetIds()) != 1 || filter.GetIds()[0] < 1 {
			filter.AddValidationError("Error parse Id", errors.ErrorCodeParseId, "Id")
		}
		break
	case settings.FunctionTypeMultiDelete:
		break
	case settings.FunctionTypeFindOrCreate:
		break
	case settings.FunctionTypeMultiCreate:
		break
	case settings.FunctionTypeMultiUpdate:
		break
	case settings.FunctionTypeUpdateOrCreate:
		break
	default:
		filter.AddValidationError("Unsupported function type: "+functionType, errors.ErrorCodeUnsupportedFunctionType, "")
		break
	}
}

type GoshaSearchFilter struct {
	Search   string
	SearchBy []string
}

type GoshaOrderFilter struct {
	Order          []string
	OrderDirection []string
}

type GoshaDebugFilter struct {
	isDebug bool
}

func (filter *GoshaDebugFilter) SetDebug(isDebug bool) {
	filter.isDebug = isDebug
}

func (filter GoshaDebugFilter) IsDebug() bool {
	return filter.isDebug
}

type AbstractFilter struct {
	request *http.Request
	rawRequestBody []byte

	GoshaSearchFilter
	GoshaOrderFilter
	GoshaFilterIds
	Pagination
	validator
	Authenticator
	GoshaDebugFilter
}

func GetAbstractFilter(request *http.Request, requestBody []byte, functionType string) (filter AbstractFilter, err error) {

	filter.request = request
    filter.rawRequestBody = requestBody
	filter.functionType = functionType
	filter.urlPath = request.URL.Path

	if !isGroupFunctionType(functionType) {
        err = ReadJSON(filter.rawRequestBody, &filter.GoshaFilterIds)
        if err != nil {
            return
        }
    }

	filter.Pagination.CurrentPage, _ = strconv.Atoi(request.FormValue("CurrentPage"))
	filter.Pagination.PerPage, _ = strconv.Atoi(request.FormValue("PerPage"))
	filter.Search = request.FormValue("Search")

	isDebug, _ := strconv.ParseBool(request.FormValue("IsDebug"))
	filter.GoshaDebugFilter.SetDebug(isDebug)

	arr, _ := url.ParseQuery(request.URL.RawQuery)

	dirs := []string{}

	for _, dir := range arr["OrderDirection[]"] {

		if strings.ToLower(dir) == settings.OrderDirectionDesc {
			dirs = append(dirs, settings.OrderDirectionDesc)
		} else {
			dirs = append(dirs, settings.OrderDirectionAsc)
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

	if len(filter.Order) < 1 && len(filter.OrderDirection) < 1 {
		filter.Order = append(filter.Order, "id")
		filter.OrderDirection = append(filter.OrderDirection, "asc")
	}

	for _, field := range arr["SearchBy[]"] {
		filter.SearchBy = append(filter.SearchBy, gorm.ToColumnName(field))
	}

	filter.SetToken(request)
	filter.SetIp(request)

	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		filter.AddId(id)
		filter.SetCurrentId(id)
	}

	for _, field := range arr["Ids[]"] {
		id, _ := strconv.Atoi(field)
		filter.AddId(id)
	}
	for _, field := range arr["ExceptIds[]"] {
		id, _ := strconv.Atoi(field)
		filter.AddExceptIds(id)
	}

	filter.Validate(functionType)

	return filter, err
}

func (filter *AbstractFilter) IsValid() bool {

	return filter.GoshaFilterIds.IsValid() &&
		filter.Pagination.IsValid() &&
		filter.validator.IsValid() &&
		filter.Authenticator.IsValid()
}

func (filter *AbstractFilter) Validate(functionType string) {

	filter.GoshaFilterIds.Validate(functionType)
	filter.Pagination.Validate(functionType)
	filter.validator.Validate(functionType)
	filter.Authenticator.Validate(functionType)
}

func (filter *AbstractFilter) GetValidationError() error {
	return errors.JoinValidatorError(filter.GoshaFilterIds.GetValidationError(),
		filter.Pagination.GetValidationError(),
		filter.validator.GetValidationError(),
		filter.Authenticator.GetValidationError())
}

func (filter *AbstractFilter) ValidatePerPage() {
	if filter.PerPage > filter.GetMaxPerPage() {
		filter.AddValidationError("PerPage more than maximum", errors.ErrorCodeInvalidPerPage, "PerPage")
	}
}

func (filter *AbstractFilter) GetHost() string {
	return filter.request.Host
}

func (filter *AbstractFilter) GetCurrentIp() string {

	ip := filter.request.Header.Get("X-Forwarded-For")
	return ip
}

func (filter *AbstractFilter) GetCurrentUserAgent() string {
	return filter.request.UserAgent()
}

func isGroupFunctionType(functionType string) bool {
    switch functionType {
    case settings.FunctionTypeMultiCreate, settings.FunctionTypeMultiUpdate, settings.FunctionTypeMultiDelete, settings.FunctionTypeMultiFindOrCreate:
        return true
    default:
        return false
    }
}

`

const usualTypesRequest = `package types

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func GetRawBodyContent(request *http.Request) (data []byte, err error) {
    defer request.Body.Close()
    data, err = ioutil.ReadAll(request.Body)
    if err == http.ErrBodyReadAfterClose {
        err = nil
    }
    return
}
func ReadJSON(body []byte, entity interface{}) (err error) {
    if len(body) > 0 {
        err = json.Unmarshal(body, entity)
    }
    if err != nil && err.Error() == "invalid character '-' in numeric literal" {
        err = nil
    }
    return
}
`

const usualTypesValidator = `package types

import (
    "{ms-name}/settings"
    "{ms-name}/errors"
)

type validator struct {
    validationError errors.ValidatorError
}

func (val *validator) IsValid() bool {

    return val.validationError.IsEmpty()
}

func (val *validator) GetValidationError() errors.ValidatorErrorInterface {
    return &val.validationError
}

func (val *validator) AddValidationError(err string, code errors.ErrorCode, field string) {
    val.validationError.AddError(errors.NewErrorWithCode(err, code, field))
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
        val.AddValidationError("Unsupported function type: "+functionType, errors.ErrorCodeUnsupportedFunctionType, "")
        break
    }
}
`

const usualTypesResponse = `package types

import (
	"{ms-name}/settings"
	"{ms-name}/errors"
)

type APIStatus struct {
	Status string
}

type APIError struct {
	Error  bool
	Errors []Error
}

type Error struct {
	ErrorMessage string
	ErrorCode    int
	Field        string `+"`" + `json:"Field,omitempty"`+"`" + `
	ErrorDebug   string `+"`" + `json:"ErrorDebug,omitempty"`+"`" + `
}

type Pagination struct {
	CurrentPage int
	PerPage     int

	validator
}


func (pagination *Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}

func (pagination *Pagination) Validate(functionType string) {

	switch functionType {

	case settings.FunctionTypeFind:

		if pagination.CurrentPage < 1 {
			pagination.AddValidationError("CurrentPage parameter is not set", errors.ErrorCodeInvalidCurrentPage,"CurrentPage")
		}

		if pagination.PerPage < 1 {
			pagination.AddValidationError("PerPage parameter is not set", errors.ErrorCodeInvalidPerPage,"PerPage")
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
		pagination.validator.AddValidationError("Usupported function type: " + functionType, errors.ErrorCodeUnsupportedFunctionType, "")
		break
	}
}
`

func getUsualTemplateTypesAuthenticator(isUuidAsPk bool) template {

    cont := AssignVar(
        assignMsName(usualTypesAuthenticator),
        "{ID}",
        GetPKType(isUuidAsPk),
    )

    cont = AssignVar(
        cont,
        "{GetIdIsNotValidExp}",
        GetIdIsNotValidExp(isUuidAsPk),
    )

    usualTemplateTypesAuthenticator := template{
        Path:    "./types/authenticator.go",
        Content: cont,
    }
    return usualTemplateTypesAuthenticator
}

var usualTemplateTypesEntity = template{
    Path:    "./types/entity.go",
    Content: usualTypesEntity,
}

func getUsualTemplateTypesFilter(isUuidAsPk bool) template {

    tpl := AssignVar(
        assignMsName(usualTypesFilter),
        "{ID}",
        GetPKType(isUuidAsPk),
    )

    tpl = AssignVar(
        tpl,
        "{STRTOID}",
        GetStrToIdFuncName(isUuidAsPk),
    )

    tpl = AssignVar(
        tpl,
        "{PkNil}",
        GetIdNil(isUuidAsPk),
    )

    tpl = AssignVar(
        tpl,
        "{GetIdIsValidExp}",
        GetIdIsValidExp(isUuidAsPk),
    )

    usualTemplateTypesFilter := template{
        Path:    "./types/filter.go",
        Content: tpl,
    }

    return usualTemplateTypesFilter
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
