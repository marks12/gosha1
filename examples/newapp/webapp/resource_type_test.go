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

var validModelResourceType = types.ResourceType{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelResourceType = types.ResourceType{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveResourceType = []int{}

func validateFieldsResourceType(t *testing.T, testModel types.ResourceType, validModelResourceType types.ResourceType, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new ResourceType", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelResourceType.Name {
	//	t.Error("Fail test creating new ResourceType", "expect Name =", validModelResourceType.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelResourceType types.ResourceType

var testCreateFuncResourceType = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	ResourceTypeCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestResourceType = tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType)

var testCasesResourceType = []tests.WebTest{
    {
		Name:         "Find ResourceTypes as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncResourceType(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.ResourceTypeRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, ResourceTypeFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getResourceTypeParsedList(response)

			if total < 1 {
				t.Error("Error in find ResourceType. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsResourceType(t, item, validModelResourceType, response)
			}
		},
	},
    {
		Name:         "Create new ResourceType as admin",
		Request:      createAdminRequestResourceType,
		ResponseCode: 201,
		TestFunc:     testCreateFuncResourceType,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelResourceType = getResourceTypeParsedModel(response)

			idsForRemoveResourceType = append(idsForRemoveResourceType, createdModelResourceType.Id)
			validateFieldsResourceType(t, createdModelResourceType, validModelResourceType, response)
		},
	},
	{
		Name:            "Create new ResourceType as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.ResourceTypeRoute, validModelResourceType),
		ResponseCode:    403,
		TestFunc:        testCreateFuncResourceType,
	},
    {
		Name:         "Read ResourceType as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.ResourceTypeRoute + "/" + strconv.Itoa(updateModelResourceType.Id))
			return tests.SendRequest(settings.ResourceTypeRoute + "/{id}", request, ResourceTypeRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read ResourceType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ResourceTypeRoute, updateModelResourceType)
			responseCreate, err := testCreateFuncResourceType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getResourceTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResourceType = append(idsForRemoveResourceType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.ResourceTypeRoute + "/" + id)

			return tests.SendRequest(settings.ResourceTypeRoute + "/{id}", request, ResourceTypeRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelResourceType = getResourceTypeParsedModel(response)
			validateFieldsResourceType(t, createdModelResourceType, updateModelResourceType, response)
		},
	},
    {
		Name:         "Update ResourceType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType)
			responseCreate, err := testCreateFuncResourceType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getResourceTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResourceType = append(idsForRemoveResourceType, responseData.Id)

			updateModelResourceType.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.ResourceTypeRoute + "/" + id, updateModelResourceType)

			return tests.SendRequest(settings.ResourceTypeRoute + "/{id}", request, ResourceTypeUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getResourceTypeParsedModel(response)
			validateFieldsResourceType(t, model, updateModelResourceType, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.ResourceTypeRoute + "/" + id)
			readResponse := tests.SendRequest(settings.ResourceTypeRoute + "/{id}", request, ResourceTypeRead, http.MethodGet)

			model = getResourceTypeParsedModel(readResponse)
			validateFieldsResourceType(t, model, updateModelResourceType, readResponse)
		},
	},
    {
		Name: "Delete ResourceType as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			ResourceTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType))

			responseData, err := getResourceTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveResourceType = append(idsForRemoveResourceType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.ResourceTypeRoute + "/" + id)

			return tests.SendRequest(settings.ResourceTypeRoute+"/{id}", tt.Request, ResourceTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ResourceType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ResourceType for next delete
			responseCreate := httptest.NewRecorder()
			ResourceTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType))

			responseData, err := getResourceTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ResourceTypeRoute + "/" + id)

			return tests.SendRequest(settings.ResourceTypeRoute+"/{id}", req, ResourceTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ResourceType as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ResourceType for next delete
			responseCreate := httptest.NewRecorder()
			ResourceTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.ResourceTypeRoute, validModelResourceType))

			responseData, err := getResourceTypeResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ResourceTypeRoute + "/" + id)
			tests.SendRequest(settings.ResourceTypeRoute+"/{id}", req, ResourceTypeDelete, http.MethodDelete)

			return tests.SendRequest(settings.ResourceTypeRoute+"/{id}", req, ResourceTypeDelete, http.MethodDelete), nil
		},
	},
    

}

func getResourceTypeResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.ResourceType, error) {

	model := getResourceTypeParsedModel(response)

	if model.Id < 1 {
		return types.ResourceType{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getResourceTypeParsedList(response *httptest.ResponseRecorder) (list []types.ResourceType, total int) {

	responseData := struct{
		List []types.ResourceType
		Total int
	} {
		List: []types.ResourceType{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getResourceTypeParsedModel(response *httptest.ResponseRecorder) types.ResourceType {

	responseData := struct{
		Model types.ResourceType
	} {
		Model:types.ResourceType{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestResourceType(t *testing.T) {

	for _, tt := range testCasesResourceType {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveResourceType {
		tests.SendRequest(settings.ResourceTypeRoute + "/{id}", tests.GetDeleteAdminRequest(settings.ResourceTypeRoute + "/" + strconv.Itoa(id)), ResourceTypeDelete, http.MethodDelete)
	}
}

