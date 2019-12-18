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

var validModelResource = types.Resource{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelResource = types.Resource{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveResource = []int{}

func validateFieldsResource(t *testing.T, testModel types.Resource, validModelResource types.Resource, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new Resource", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelResource.Name {
	//	t.Error("Fail test creating new Resource", "expect Name =", validModelResource.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelResource types.Resource

var testCreateFuncResource = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	ResourceCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestResource = tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource)

var testCasesResource = []tests.WebTest{
    {
		Name:         "Find Resources as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncResource(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.ResourceRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, ResourceFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getResourceParsedList(response)

			if total < 1 {
				t.Error("Error in find Resource. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsResource(t, item, validModelResource, response)
			}
		},
	},
    {
		Name:         "Create new Resource as admin",
		Request:      createAdminRequestResource,
		ResponseCode: 201,
		TestFunc:     testCreateFuncResource,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelResource = getResourceParsedModel(response)

			idsForRemoveResource = append(idsForRemoveResource, createdModelResource.Id)
			validateFieldsResource(t, createdModelResource, validModelResource, response)
		},
	},
	{
		Name:            "Create new Resource as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.ResourceRoute, validModelResource),
		ResponseCode:    403,
		TestFunc:        testCreateFuncResource,
	},
    {
		Name:         "Read Resource as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.ResourceRoute + "/" + strconv.Itoa(updateModelResource.Id))
			return tests.SendRequest(settings.ResourceRoute + "/{id}", request, ResourceRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read Resource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ResourceRoute, updateModelResource)
			responseCreate, err := testCreateFuncResource(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResource = append(idsForRemoveResource, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.ResourceRoute + "/" + id)

			return tests.SendRequest(settings.ResourceRoute + "/{id}", request, ResourceRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelResource = getResourceParsedModel(response)
			validateFieldsResource(t, createdModelResource, updateModelResource, response)
		},
	},
    {
		Name:         "Update Resource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource)
			responseCreate, err := testCreateFuncResource(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResource = append(idsForRemoveResource, responseData.Id)

			updateModelResource.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.ResourceRoute + "/" + id, updateModelResource)

			return tests.SendRequest(settings.ResourceRoute + "/{id}", request, ResourceUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getResourceParsedModel(response)
			validateFieldsResource(t, model, updateModelResource, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.ResourceRoute + "/" + id)
			readResponse := tests.SendRequest(settings.ResourceRoute + "/{id}", request, ResourceRead, http.MethodGet)

			model = getResourceParsedModel(readResponse)
			validateFieldsResource(t, model, updateModelResource, readResponse)
		},
	},
    {
		Name: "Delete Resource as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			ResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource))

			responseData, err := getResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResource = append(idsForRemoveResource, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.ResourceRoute + "/" + id)

			return tests.SendRequest(settings.ResourceRoute+"/{id}", tt.Request, ResourceDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Resource as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Resource for next delete
			responseCreate := httptest.NewRecorder()
			ResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource))

			responseData, err := getResourceResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ResourceRoute + "/" + id)

			return tests.SendRequest(settings.ResourceRoute+"/{id}", req, ResourceDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Resource as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Resource for next delete
			responseCreate := httptest.NewRecorder()
			ResourceCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceRoute, validModelResource))

			responseData, err := getResourceResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ResourceRoute + "/" + id)
			tests.SendRequest(settings.ResourceRoute+"/{id}", req, ResourceDelete, http.MethodDelete)

			return tests.SendRequest(settings.ResourceRoute+"/{id}", req, ResourceDelete, http.MethodDelete), nil
		},
	},
    

}

func getResourceResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Resource, error) {

	model := getResourceParsedModel(response)

	if model.Id < 1 {
		return types.Resource{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getResourceParsedList(response *httptest.ResponseRecorder) (list []types.Resource, total int) {

	responseData := struct{
		List []types.Resource
		Total int
	} {
		List: []types.Resource{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getResourceParsedModel(response *httptest.ResponseRecorder) types.Resource {

	responseData := struct{
		Model types.Resource
	} {
		Model:types.Resource{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestResource(t *testing.T) {

	for _, tt := range testCasesResource {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveResource {
		tests.SendRequest(settings.ResourceRoute + "/{id}", tests.GetDeleteAdminRequest(settings.ResourceRoute + "/" + strconv.Itoa(id)), ResourceDelete, http.MethodDelete)
	}
}

