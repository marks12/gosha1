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

var validModelUserRole = types.UserRole{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelUserRole = types.UserRole{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveUserRole = []int{}

func validateFieldsUserRole(t *testing.T, testModel types.UserRole, validModelUserRole types.UserRole, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new UserRole", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelUserRole.Name {
	//	t.Error("Fail test creating new UserRole", "expect Name =", validModelUserRole.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelUserRole types.UserRole

var testCreateFuncUserRole = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	UserRoleCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestUserRole = tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole)

var testCasesUserRole = []tests.WebTest{
    {
		Name:         "Find UserRoles as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncUserRole(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.UserRoleRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, UserRoleFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getUserRoleParsedList(response)

			if total < 1 {
				t.Error("Error in find UserRole. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsUserRole(t, item, validModelUserRole, response)
			}
		},
	},
    {
		Name:         "Create new UserRole as admin",
		Request:      createAdminRequestUserRole,
		ResponseCode: 201,
		TestFunc:     testCreateFuncUserRole,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelUserRole = getUserRoleParsedModel(response)

			idsForRemoveUserRole = append(idsForRemoveUserRole, createdModelUserRole.Id)
			validateFieldsUserRole(t, createdModelUserRole, validModelUserRole, response)
		},
	},
	{
		Name:            "Create new UserRole as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.UserRoleRoute, validModelUserRole),
		ResponseCode:    403,
		TestFunc:        testCreateFuncUserRole,
	},
    {
		Name:         "Read UserRole as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.UserRoleRoute + "/" + strconv.Itoa(updateModelUserRole.Id))
			return tests.SendRequest(settings.UserRoleRoute + "/{id}", request, UserRoleRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read UserRole as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.UserRoleRoute, updateModelUserRole)
			responseCreate, err := testCreateFuncUserRole(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getUserRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUserRole = append(idsForRemoveUserRole, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.UserRoleRoute + "/" + id)

			return tests.SendRequest(settings.UserRoleRoute + "/{id}", request, UserRoleRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelUserRole = getUserRoleParsedModel(response)
			validateFieldsUserRole(t, createdModelUserRole, updateModelUserRole, response)
		},
	},
    {
		Name:         "Update UserRole as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole)
			responseCreate, err := testCreateFuncUserRole(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getUserRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUserRole = append(idsForRemoveUserRole, responseData.Id)

			updateModelUserRole.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.UserRoleRoute + "/" + id, updateModelUserRole)

			return tests.SendRequest(settings.UserRoleRoute + "/{id}", request, UserRoleUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getUserRoleParsedModel(response)
			validateFieldsUserRole(t, model, updateModelUserRole, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.UserRoleRoute + "/" + id)
			readResponse := tests.SendRequest(settings.UserRoleRoute + "/{id}", request, UserRoleRead, http.MethodGet)

			model = getUserRoleParsedModel(readResponse)
			validateFieldsUserRole(t, model, updateModelUserRole, readResponse)
		},
	},
    {
		Name: "Delete UserRole as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			UserRoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole))

			responseData, err := getUserRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUserRole = append(idsForRemoveUserRole, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.UserRoleRoute + "/" + id)

			return tests.SendRequest(settings.UserRoleRoute+"/{id}", tt.Request, UserRoleDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete UserRole as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create UserRole for next delete
			responseCreate := httptest.NewRecorder()
			UserRoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole))

			responseData, err := getUserRoleResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.UserRoleRoute + "/" + id)

			return tests.SendRequest(settings.UserRoleRoute+"/{id}", req, UserRoleDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete UserRole as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create UserRole for next delete
			responseCreate := httptest.NewRecorder()
			UserRoleCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoleRoute, validModelUserRole))

			responseData, err := getUserRoleResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.UserRoleRoute + "/" + id)
			tests.SendRequest(settings.UserRoleRoute+"/{id}", req, UserRoleDelete, http.MethodDelete)

			return tests.SendRequest(settings.UserRoleRoute+"/{id}", req, UserRoleDelete, http.MethodDelete), nil
		},
	},
    

}

func getUserRoleResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.UserRole, error) {

	model := getUserRoleParsedModel(response)

	if model.Id < 1 {
		return types.UserRole{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getUserRoleParsedList(response *httptest.ResponseRecorder) (list []types.UserRole, total int) {

	responseData := struct{
		List []types.UserRole
		Total int
	} {
		List: []types.UserRole{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getUserRoleParsedModel(response *httptest.ResponseRecorder) types.UserRole {

	responseData := struct{
		Model types.UserRole
	} {
		Model:types.UserRole{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestUserRole(t *testing.T) {

	for _, tt := range testCasesUserRole {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveUserRole {
		tests.SendRequest(settings.UserRoleRoute + "/{id}", tests.GetDeleteAdminRequest(settings.UserRoleRoute + "/" + strconv.Itoa(id)), UserRoleDelete, http.MethodDelete)
	}
}

