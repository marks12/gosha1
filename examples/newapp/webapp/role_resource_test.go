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

var validModelRoleResource = types.RoleResource{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelRoleResource = types.RoleResource{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveRoleResource = []int{}

func validateFieldsRoleResource(t *testing.T, testModel types.RoleResource, validModelRoleResource types.RoleResource, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new RoleResource", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelRoleResource.Name {
	//	t.Error("Fail test creating new RoleResource", "expect Name =", validModelRoleResource.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelRoleResource types.RoleResource

var testCreateFuncRoleResource = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	RoleResourceCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestRoleResource = tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource)

var testCasesRoleResource = []tests.WebTest{
    {
		Name:         "Find RoleResources as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncRoleResource(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.RoleResourceRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, RoleResourceFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getRoleResourceParsedList(response)

			if total < 1 {
				t.Error("Error in find RoleResource. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsRoleResource(t, item, validModelRoleResource, response)
			}
		},
	},
    {
		Name:         "Create new RoleResource as admin",
		Request:      createAdminRequestRoleResource,
		ResponseCode: 201,
		TestFunc:     testCreateFuncRoleResource,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelRoleResource = getRoleResourceParsedModel(response)

			idsForRemoveRoleResource = append(idsForRemoveRoleResource, createdModelRoleResource.Id)
			validateFieldsRoleResource(t, createdModelRoleResource, validModelRoleResource, response)
		},
	},
	{
		Name:            "Create new RoleResource as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.RoleResourceRoute, validModelRoleResource),
		ResponseCode:    403,
		TestFunc:        testCreateFuncRoleResource,
	},
    {
		Name:         "Read RoleResource as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.RoleResourceRoute + "/" + strconv.Itoa(updateModelRoleResource.Id))
			return tests.SendRequest(settings.RoleResourceRoute + "/{id}", request, RoleResourceRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read RoleResource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RoleResourceRoute, updateModelRoleResource)
			responseCreate, err := testCreateFuncRoleResource(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRoleResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRoleResource = append(idsForRemoveRoleResource, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.RoleResourceRoute + "/" + id)

			return tests.SendRequest(settings.RoleResourceRoute + "/{id}", request, RoleResourceRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelRoleResource = getRoleResourceParsedModel(response)
			validateFieldsRoleResource(t, createdModelRoleResource, updateModelRoleResource, response)
		},
	},
    {
		Name:         "Update RoleResource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource)
			responseCreate, err := testCreateFuncRoleResource(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRoleResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRoleResource = append(idsForRemoveRoleResource, responseData.Id)

			updateModelRoleResource.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.RoleResourceRoute + "/" + id, updateModelRoleResource)

			return tests.SendRequest(settings.RoleResourceRoute + "/{id}", request, RoleResourceUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getRoleResourceParsedModel(response)
			validateFieldsRoleResource(t, model, updateModelRoleResource, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.RoleResourceRoute + "/" + id)
			readResponse := tests.SendRequest(settings.RoleResourceRoute + "/{id}", request, RoleResourceRead, http.MethodGet)

			model = getRoleResourceParsedModel(readResponse)
			validateFieldsRoleResource(t, model, updateModelRoleResource, readResponse)
		},
	},
    {
		Name: "Delete RoleResource as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			RoleResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource))

			responseData, err := getRoleResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRoleResource = append(idsForRemoveRoleResource, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.RoleResourceRoute + "/" + id)

			return tests.SendRequest(settings.RoleResourceRoute+"/{id}", tt.Request, RoleResourceDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete RoleResource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create RoleResource for next delete
			responseCreate := httptest.NewRecorder()
			RoleResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource))

			responseData, err := getRoleResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RoleResourceRoute + "/" + id)

			return tests.SendRequest(settings.RoleResourceRoute+"/{id}", req, RoleResourceDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete RoleResource as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create RoleResource for next delete
			responseCreate := httptest.NewRecorder()
			RoleResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.RoleResourceRoute, validModelRoleResource))

			responseData, err := getRoleResourceResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RoleResourceRoute + "/" + id)
			tests.SendRequest(settings.RoleResourceRoute+"/{id}", req, RoleResourceDelete, http.MethodDelete)

			return tests.SendRequest(settings.RoleResourceRoute+"/{id}", req, RoleResourceDelete, http.MethodDelete), nil
		},
	},
    

}

func getRoleResourceResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.RoleResource, error) {

	model := getRoleResourceParsedModel(response)

	if model.Id < 1 {
		return types.RoleResource{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getRoleResourceParsedList(response *httptest.ResponseRecorder) (list []types.RoleResource, total int) {

	responseData := struct{
		List []types.RoleResource
		Total int
	} {
		List: []types.RoleResource{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getRoleResourceParsedModel(response *httptest.ResponseRecorder) types.RoleResource {

	responseData := struct{
		Model types.RoleResource
	} {
		Model:types.RoleResource{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestRoleResource(t *testing.T) {

	for _, tt := range testCasesRoleResource {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveRoleResource {
		tests.SendRequest(settings.RoleResourceRoute + "/{id}", tests.GetDeleteAdminRequest(settings.RoleResourceRoute + "/" + strconv.Itoa(id)), RoleResourceDelete, http.MethodDelete)
	}
}

