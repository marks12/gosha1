package cmd

var usualTemplateWebappTestEntity = template{
    Path:    "./webapp/{entity-name}.go",
    Content: assignMsName(GetUsualTemplateWebAppTestContent(
        Crud{true, true, true, true, true, true},
    )),
}

func GetUsualTemplateWebAppTestContent(methodsCrud Crud) string {

    var usualWebappTestEntity = `package webapp

import (
	"{ms-name}/settings"
	"{ms-name}/tests"
	"{ms-name}/types"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var validModel{entity-name} = types.{entity-name}{
		Id:   0,
		//Name: "Some Name",
	}

var updateModel{entity-name} = types.{entity-name}{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemove{entity-name} = []int{}

func validateFields{entity-name}(t *testing.T, testModel types.{entity-name}, validModel{entity-name} types.{entity-name}, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new {entity-name}", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModel{entity-name}.Name {
	//	t.Error("Fail test creating new {entity-name}", "expect Name =", validModel{entity-name}.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModel{entity-name} types.{entity-name}

var testCreateFunc{entity-name} = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	{entity-name}Create(response, tt.Request)
	return response, nil
}

var createAdminRequest{entity-name} = tests.GetCreateAdminRequest(settings.{entity-name}Route, validModel{entity-name})

var testCases{entity-name} = []tests.WebTest{

    ` + getWebappTestFind(methodsCrud) + `

    ` + getWebappTestCreate(methodsCrud) + `

    ` + getWebappTestRead(methodsCrud) + `

    ` + getWebappTestUpdate(methodsCrud) + `

    ` + getWebappTestDelete(methodsCrud) + `

    ` + getWebappTestFindOrCreate(methodsCrud) + `
}

func get{entity-name}ResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.{entity-name}, error) {

	model := get{entity-name}ParsedModel(response)

	if model.Id < 1 {
		return types.{entity-name}{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func get{entity-name}ParsedList(response *httptest.ResponseRecorder) (list []types.{entity-name}, total int) {

	responseData := struct{
		List []types.{entity-name}
		Total int
	} {
		List: []types.{entity-name}{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func get{entity-name}ParsedModel(response *httptest.ResponseRecorder) types.{entity-name} {

	responseData := struct{
		Model types.{entity-name}
	} {
		Model:types.{entity-name}{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func Test{entity-name}(t *testing.T) {

	for _, tt := range testCases{entity-name} {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemove{entity-name} {
		tests.SendRequest(settings.{entity-name}Route + "/{id}", tests.GetDeleteAdminRequest(settings.{entity-name}Route + "/" + strconv.Itoa(id)), {entity-name}Delete, http.MethodDelete)
	}
}

`
    return usualWebappTestEntity
}

func getWebappTestFind(methodCrud Crud) (c string) {

    if methodCrud.IsFind {
        c = `
	{
		Name:         "Find {entity-name}s as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			route := settings.{entity-name}Route
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, {entity-name}Find, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := get{entity-name}ParsedList(response)

			if total < 1 {
				t.Error("Error in find {entity-name}. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFields{entity-name}(t, item, validModel{entity-name}, response)
			}
		},
	},`
    }

    return
}

func getWebappTestCreate(methodCrud Crud) (c string) {

    if methodCrud.IsCreate {
        c = `
	{
		Name:         "Create new {entity-name} as admin",
		Request:      createAdminRequest{entity-name},
		ResponseCode: 201,
		TestFunc:     testCreateFunc{entity-name},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModel{entity-name} = get{entity-name}ParsedModel(response)

			idsForRemove{entity-name} = append(idsForRemove{entity-name}, createdModel{entity-name}.Id)
			validateFields{entity-name}(t, createdModel{entity-name}, validModel{entity-name}, response)
		},
	},
	{
		Name:            "Create new {entity-name} as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.{entity-name}Route, validModel{entity-name}),
		ResponseCode:    403,
		TestFunc:        testCreateFunc{entity-name},
	},`
    }

    return
}

func getWebappTestRead(methodCrud Crud) (c string) {

    if methodCrud.IsRead {
        c = `
	{
		Name:         "Read {entity-name} as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.{entity-name}Route + "/" + strconv.Itoa(updateModel{entity-name}.Id))
			return tests.SendRequest(settings.{entity-name}Route + "/{id}", request, {entity-name}Read, http.MethodGet), nil
		},
	},
	{
		Name:         "Read {entity-name} as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.{entity-name}Route, updateModel{entity-name})
			responseCreate, err := testCreateFunc{entity-name}(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := get{entity-name}ResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemove{entity-name} = append(idsForRemove{entity-name}, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.{entity-name}Route + "/" + id)

			return tests.SendRequest(settings.{entity-name}Route + "/{id}", request, {entity-name}Read, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModel{entity-name} = get{entity-name}ParsedModel(response)
			validateFields{entity-name}(t, createdModel{entity-name}, updateModel{entity-name}, response)
		},
	},`
    }

    return
}

func getWebappTestUpdate(methodCrud Crud) (c string) {

    if methodCrud.IsUpdate {
        c = `
    {
		Name:         "Update {entity-name} as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.{entity-name}Route, validModel{entity-name})
			responseCreate, err := testCreateFunc{entity-name}(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := get{entity-name}ResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemove{entity-name} = append(idsForRemove{entity-name}, responseData.Id)

			updateModel{entity-name}.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.{entity-name}Route + "/" + id, updateModel{entity-name})

			return tests.SendRequest(settings.{entity-name}Route + "/{id}", request, {entity-name}Update, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := get{entity-name}ParsedModel(response)
			validateFields{entity-name}(t, model, updateModel{entity-name}, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.{entity-name}Route + "/" + id)
			readResponse := tests.SendRequest(settings.{entity-name}Route + "/{id}", request, {entity-name}Read, http.MethodGet)

			model = get{entity-name}ParsedModel(readResponse)
			validateFields{entity-name}(t, model, updateModel{entity-name}, readResponse)
		},
	},`
    }

    return
}

func getWebappTestDelete(methodCrud Crud) (c string) {

    if methodCrud.IsDelete {
        c = `
	{
		Name: "Delete {entity-name} as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			{entity-name}Create(responseCreate, tests.GetCreateAdminRequest(settings.{entity-name}Route, validModel{entity-name}))

			responseData, err := get{entity-name}ResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemove{entity-name} = append(idsForRemove{entity-name}, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.{entity-name}Route + "/" + id)

			return tests.SendRequest(settings.{entity-name}Route+"/{id}", tt.Request, {entity-name}Delete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete {entity-name} as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create {entity-name} for next delete
			responseCreate := httptest.NewRecorder()
			{entity-name}Create(responseCreate, createAdminRequest{entity-name})

			responseData, err := get{entity-name}ResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.{entity-name}Route + "/" + id)

			return tests.SendRequest(settings.{entity-name}Route+"/{id}", req, {entity-name}Delete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete {entity-name} as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create {entity-name} for next delete
			responseCreate := httptest.NewRecorder()
			{entity-name}Create(responseCreate, createAdminRequest{entity-name})

			responseData, err := get{entity-name}ResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.{entity-name}Route + "/" + id)
			tests.SendRequest(settings.{entity-name}Route+"/{id}", req, {entity-name}Delete, http.MethodDelete)

			return tests.SendRequest(settings.{entity-name}Route+"/{id}", req, {entity-name}Delete, http.MethodDelete), nil
		},
	},`
    }

    return
}

func getWebappTestFindOrCreate(methodCrud Crud) (c string) {

    if methodCrud.IsFindOrCreate {
        c = `
`
    }

    return
}