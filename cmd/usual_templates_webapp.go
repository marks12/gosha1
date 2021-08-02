package cmd

const usualWebappErrors = `package webapp

import (
    "{ms-name}/common"
    "{ms-name}/errors"
    "{ms-name}/logic"
    "{ms-name}/mdl"
    "{ms-name}/types"
    "net/http"
    "encoding/json"
    "fmt"
    "strings"
	dynamicStruct "github.com/ompluscator/dynamic-struct"
    "reflect"
)

type FilterInterface interface {
    IsDebug() bool
    GetLanguageId() int
}


func ApplyFieldsFilterToData(fields []string, data interface{}) (result interface{}) {
	if len(fields) < 1 {
		return data
	}
	b, _ := json.Marshal(&data)

	resMap := fieldPreprocessing(fields)

	result = buildStruct(resMap, data)
	err := json.Unmarshal(b, &result)

	if err != nil {
		fmt.Printf("ApplyFieldsFilterToData error: %+v", err.Error())
		return data
	}

	return
}


// Формирование карты полей
func fieldPreprocessing(fields []string) map[string]interface{} {
	res := map[string]interface{}{}

	nextFields := map[string][]string{}

	for _, field := range fields {
		if len(field) < 1 {
			continue
		}
		validFieldName := strings.Title(field)
		if !strings.Contains(validFieldName, ".") {
			res[validFieldName] = true
		} else {
			arr := strings.Split(validFieldName, ".")
			nextFields[arr[0]] = append(nextFields[arr[0]], strings.Join(arr[1:], "."))
		}
	}
	for k, v := range nextFields {
		res[k] = fieldPreprocessing(v)
	}
	return res
}


// Построение стуктуры на основе карты полей
func buildStruct(structMap map[string]interface{}, data interface{}) interface{} {
	var i *interface{}
	isArr := isArray(data)
	builder := dynamicStruct.NewStruct()
	for k, v := range structMap {
		if nextMap, isOk := v.(map[string]interface{}); isOk {
			builder.AddField(k, buildStruct(nextMap, getDataFromStructByFieldName(data, k)), "")
		} else {
			builder.AddField(k, i, "")
		}
	}
	if isArr {
		return builder.Build().NewSliceOfStructs()
	}
	return builder.Build().New()
}

//Получаем вложенные данные из структуры
func getDataFromStructByFieldName(data interface{}, field string) interface{} {
	var i *interface{}

	r := reflect.Indirect(reflect.ValueOf(data))
	if r.Kind() == reflect.Slice {
		el := reflect.Indirect(reflect.New(r.Type().Elem()))
		f := el.FieldByName(field)
		if !f.IsValid() {
			return i
		}
		return reflect.New(f.Type()).Interface()
	} else {
		if !r.IsValid() || r.IsZero() || !(r.Kind() == reflect.Struct) && r.IsNil() {
			return i
		}
		f := reflect.Indirect(r).FieldByName(field)
		if !(r.Kind() == reflect.Struct) && f.IsNil() || f.IsZero() || !f.IsValid() {
			return i
		}
		return f.Interface()
	}
}

func isArray(data interface{}) (res bool) {
	rt := reflect.Indirect(reflect.ValueOf(data))
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

func Bad(w http.ResponseWriter, requestDto FilterInterface, err error) {
    ErrResponse(w, err, http.StatusBadRequest, requestDto)
}

func AuthErr(w http.ResponseWriter, requestDto FilterInterface) {

    ErrResponse(w, GetAuthErrTpl(common.MyCaller()), http.StatusForbidden, requestDto)
}

func GetAuthErrTpl(operation string) errors.ErrorWithCode {

    return errors.NewErrorWithCode(
        fmt.Sprintf("Invalid authorize in %s", operation),
        errors.ErrorCodeInvalidAuthorize,
        "Token")
}

func ErrResponse(w http.ResponseWriter, err error, status int, filter FilterInterface) {

    response := types.APIError{}
    response.Error = true

    switch e := err.(type) {
    case errors.ValidatorError:
        for _, errWithCode := range e.Errors() {
            newError := types.Error{
                Field:     errWithCode.GetField(),
                ErrorCode: errWithCode.ErrorCode(),
            }
            if filter.IsDebug() {
                newError.ErrorDebug = errWithCode.Error()
            }
            response.Errors = append(response.Errors, newError)
        }
        break

    case errors.ValidatorErrorInterface:
        for _, errWithCode := range e.Errors() {
            newError := types.Error{
                Field:     errWithCode.GetField(),
                ErrorCode: errWithCode.ErrorCode(),
            }
            if filter.IsDebug() {
                newError.ErrorDebug = errWithCode.Error()
            }
            response.Errors = append(response.Errors, newError)
        }
        break

    case errors.ErrorWithCode:
        newError := types.Error{
            Field:     e.GetField(),
            ErrorCode: e.ErrorCode(),
        }
        if filter.IsDebug() {
            newError.ErrorDebug = e.Error()
        }
        response.Errors = append(response.Errors, newError)
        break

    default:
        newError := types.Error{
        }
        if filter.IsDebug() {
            newError.ErrorDebug = e.Error()
        }
        response.Errors = append(response.Errors, newError)
        break
    }

    var errCodes []int
    for _, e := range response.Errors {
        errCodes = append(errCodes, e.ErrorCode)
    }
    errCodes = common.UniqueIntArray(errCodes)
    f := types.TranslateErrorFilter{}
    f.LanguageId = filter.GetLanguageId()
    if f.LanguageId < 1 {
        f.LanguageId = errors.DefaultErrorLanguageId
    }
    f.ErrorCodes = errCodes
    f.CurrentPage = 1
    f.PerPage = len(errCodes)
    translates, _, err := logic.TranslateErrorFind(f)
    if err != nil {
        fmt.Println("TranslateErrorFind err = ", err)
    }

    for i, err := range response.Errors {
        for _, translate := range translates {
            if err.ErrorCode == translate.Code {
                response.Errors[i].ErrorMessage = translate.Translate
                break
            }
        }
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(response)

    return
}


func ValidResponse (w http.ResponseWriter, data interface{}) {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    switch data.(type) {
    case mdl.ResponseCreate:
        w.WriteHeader(http.StatusCreated)
        break
    default:
        w.WriteHeader(http.StatusOK)
        break
    }
    json.NewEncoder(w).Encode(data)

    return
}`

var usualTemplateWebappErrors = template{
	Path:    "./webapp/errors.go",
	Content: assignMsName(usualWebappErrors),
}

var usualTemplateWebappEntity = template{
	Path: "./webapp/{entity-name}.go",
	Content: assignMsName(GetUsualTemplateWebAppContent(
		Crud{true, true, true, true, true, true, true},
		Crud{true, true, true, true, true, true, true},
	)),
}

func GetUsualTemplateWebAppContent(authCrud Crud, methodsCrud Crud) string {

	var usualWebappEntity = `package webapp

import (
    "{ms-name}/core"
    "{ms-name}/logic"
    "net/http"
    "{ms-name}/mdl"
    "{ms-name}/types"
    "{ms-name}/settings"
)

    ` + getWebappFind(methodsCrud, authCrud) + `

    ` + getWebappCreate(methodsCrud, authCrud) + `

    ` + getWebappRead(methodsCrud, authCrud) + `

    ` + getWebappUpdate(methodsCrud, authCrud) + `

    ` + getWebappDelete(methodsCrud, authCrud) + `

    ` + getWebappFindOrCreate(methodsCrud, authCrud) + `

    ` + getWebappUpdateOrCreate(methodsCrud, authCrud) + `

`
	return usualWebappEntity
}

func getWebappFind(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsFind {
		c = `

func {entity-name}Find(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("Find") + `

    ` + getAuth("Find", authCrud) + `

	` + getCheckValidationError() + `
    
    // Получаем список
    data, totalRecords, err := logic.{entity-name}Find(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseFind{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
        totalRecords,
    })

    return
}`
	}

	return
}

func getErrRespIfErr() string {

	return `if err != nil {
        ErrResponse(w, err, http.StatusBadRequest, requestDto)
        return
    }`
}

func getCheckValidationError() string {

	return `if !requestDto.IsValid() {
		Bad(w, requestDto, requestDto.GetValidationError())
		return
	}
`
}

func getRequestDtoCreator(method string) string {

	return `requestDto, err := types.Get{entity-name}Filter(httpRequest, settings.FunctionType` + method + `)
	if err != nil {
		ErrResponse(w, err, http.StatusBadRequest, requestDto)
		return
	}`

}

func getWebappCreate(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsCreate {
		c = `
func {entity-name}MultiCreate(w http.ResponseWriter, httpRequest *http.Request) {


    ` + getRequestDtoCreator("MultiCreate") + `

    ` + getAuth("MultiCreate", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}MultiCreate(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseCreate{
        data,
    })

    return
}

func {entity-name}Create(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("Create") + `

    ` + getAuth("Create", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}Create(requestDto, core.Db)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseCreate{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
    })

    return
}`
	}

	return
}

func getWebappRead(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsRead {
		c = `

func {entity-name}Read(w http.ResponseWriter, httpRequest *http.Request) {


    ` + getRequestDtoCreator("Read") + `

    ` + getAuth("Read", authCrud) + `

    requestDto.PerPage = 1
    requestDto.CurrentPage = 1

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}Read(requestDto)

    // Создаём структуру ответа
    if err != nil {
        code := http.StatusBadRequest
        if err.Error() == "Not found" {
            code = http.StatusNotFound
        }
        ErrResponse(w, err, code, requestDto)
        return
    }

    ValidResponse(w, mdl.ResponseRead{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
    })

    return
}`
	}

	return
}

func getWebappUpdate(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsUpdate {
		c = `


func {entity-name}MultiUpdate(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("MultiUpdate") + `

    ` + getAuth("MultiUpdate", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}MultiUpdate(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseUpdate{
        data,
    })

    return
}

func {entity-name}Update(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("Update") + `

    ` + getAuth("Update", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}Update(requestDto, core.Db)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseUpdate{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
    })

    return
}`
	}

	return
}

func getWebappDelete(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsDelete {
		c = `

func {entity-name}MultiDelete(w http.ResponseWriter, httpRequest *http.Request) {


    ` + getRequestDtoCreator("MultiDelete") + `

    ` + getAuth("MultiDelete", authCrud) + `

	` + getCheckValidationError() + `

    isOk, err := logic.{entity-name}MultiDelete(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseDelete{
        isOk,
    })

    return
}

func {entity-name}Delete(w http.ResponseWriter, httpRequest *http.Request) {


    ` + getRequestDtoCreator("Delete") + `

    ` + getAuth("Delete", authCrud) + `

	` + getCheckValidationError() + `

    isOk, err := logic.{entity-name}Delete(requestDto, core.Db)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseDelete{
        isOk,
    })

    return
}`
	}

	return
}

func getWebappFindOrCreate(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsFindOrCreate {
		c = `

func {entity-name}FindOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("FindOrCreate") + `

    ` + getAuth("FindOrCreate", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}FindOrCreate(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseFindOrCreate{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
    })

    return
}`
	}

	return
}

func getWebappUpdateOrCreate(methodCrud Crud, authCrud Crud) (c string) {

	if methodCrud.IsUpdateOrCreate {
		c = `

func {entity-name}UpdateOrCreate(w http.ResponseWriter, httpRequest *http.Request) {

    ` + getRequestDtoCreator("UpdateOrCreate") + `

    ` + getAuth("UpdateOrCreate", authCrud) + `

	` + getCheckValidationError() + `

    data, err := logic.{entity-name}UpdateOrCreate(requestDto)

    ` + getErrRespIfErr() + `

    ValidResponse(w, mdl.ResponseUpdateOrCreate{
        ApplyFieldsFilterToData(requestDto.GetFields(), data),
    })

    return
}`
	}

	return
}

func getAuth(method string, crud Crud) (auth string) {

	switch method {

	case "Find":
		if !crud.IsFind {
			return
		}
		break

	case "Create":
		if !crud.IsCreate {
			return
		}
		break

	case "Read":
		if !crud.IsRead {
			return
		}
		break

	case "Update":
		if !crud.IsUpdate {
			return
		}
		break

	case "Delete":
		if !crud.IsDelete {
			return
		}
		break

	case "FindOrCreate":
		if !crud.IsFindOrCreate {
			return
		}
		break

	case "UpdateOrCreate":
		if !crud.IsUpdateOrCreate {
			return
		}
		break
	}

	return `if !requestDto.IsAuthorized() {
		AuthErr(w, requestDto)
		return
	}`
}
