package webapp

import (
	"newapp/settings"
	"newapp/tests"
	"newapp/types"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var validModelRole = types.Role{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelRole = types.Role{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveRole = []int{}

func validateFieldsRole(t *testing.T, testModel types.Role, validModelRole types.Role, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new Role", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelRole.Name {
	//	t.Error("Fail test creating new Role", "expect Name =", validModelRole.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelRole types.Role

var testCreateFuncRole = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	RoleCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestRole = tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole)

var testCasesRole = []tests.WebTest{
    {
		Name:         "Find Roles as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncRole(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.RoleRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, RoleFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getRoleParsedList(response)

			if total < 1 {
				t.Error("Error in find Role. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsRole(t, item, validModelRole, response)
			}
		},
	},
    {
		Name:         "Create new Role as admin",
		Request:      createAdminRequestRole,
		ResponseCode: 201,
		TestFunc:     testCreateFuncRole,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelRole = getRoleParsedModel(response)

			idsForRemoveRole = append(idsForRemoveRole, createdModelRole.Id)
			validateFieldsRole(t, createdModelRole, validModelRole, response)
		},
	},
	{
		Name:            "Create new Role as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.RoleRoute, validModelRole),
		ResponseCode:    403,
		TestFunc:        testCreateFuncRole,
	},
    {
		Name:         "Read Role as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.RoleRoute + "/" + strconv.Itoa(updateModelRole.Id))
			return tests.SendRequest(settings.RoleRoute + "/{id}", request, RoleRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read Role as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RoleRoute, updateModelRole)
			responseCreate, err := testCreateFuncRole(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRole = append(idsForRemoveRole, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.RoleRoute + "/" + id)

			return tests.SendRequest(settings.RoleRoute + "/{id}", request, RoleRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelRole = getRoleParsedModel(response)
			validateFieldsRole(t, createdModelRole, updateModelRole, response)
		},
	},
    {
		Name:         "Update Role as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole)
			responseCreate, err := testCreateFuncRole(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRole = append(idsForRemoveRole, responseData.Id)

			updateModelRole.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.RoleRoute + "/" + id, updateModelRole)

			return tests.SendRequest(settings.RoleRoute + "/{id}", request, RoleUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getRoleParsedModel(response)
			validateFieldsRole(t, model, updateModelRole, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.RoleRoute + "/" + id)
			readResponse := tests.SendRequest(settings.RoleRoute + "/{id}", request, RoleRead, http.MethodGet)

			model = getRoleParsedModel(readResponse)
			validateFieldsRole(t, model, updateModelRole, readResponse)
		},
	},
    {
		Name: "Delete Role as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			RoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole))

			responseData, err := getRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRole = append(idsForRemoveRole, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.RoleRoute + "/" + id)

			return tests.SendRequest(settings.RoleRoute+"/{id}", tt.Request, RoleDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Role as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Role for next delete
			responseCreate := httptest.NewRecorder()
			RoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole))

			responseData, err := getRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RoleRoute + "/" + id)

			return tests.SendRequest(settings.RoleRoute+"/{id}", req, RoleDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Role as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Role for next delete
			responseCreate := httptest.NewRecorder()
			RoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleRoute, validModelRole))

			responseData, err := getRoleResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RoleRoute + "/" + id)
			tests.SendRequest(settings.RoleRoute+"/{id}", req, RoleDelete, http.MethodDelete)

			return tests.SendRequest(settings.RoleRoute+"/{id}", req, RoleDelete, http.MethodDelete), nil
		},
	},
    

}

func getRoleResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Role, error) {

	model := getRoleParsedModel(response)

	if model.Id < 1 {
		return types.Role{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getRoleParsedList(response *httptest.ResponseRecorder) (list []types.Role, total int) {

	responseData := struct{
		List []types.Role
		Total int
	} {
		List: []types.Role{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getRoleParsedModel(response *httptest.ResponseRecorder) types.Role {

	responseData := struct{
		Model types.Role
	} {
		Model:types.Role{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestRole(t *testing.T) {

	for _, tt := range testCasesRole {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveRole {
		tests.SendRequest(settings.RoleRoute + "/{id}", tests.GetDeleteAdminRequest(settings.RoleRoute + "/" + strconv.Itoa(id)), RoleDelete, http.MethodDelete)
	}
}

