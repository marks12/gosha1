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

var validModelRegionType = types.RegionType{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelRegionType = types.RegionType{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveRegionType = []int{}

func validateFieldsRegionType(t *testing.T, testModel types.RegionType, validModelRegionType types.RegionType, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new RegionType", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelRegionType.Name {
	//	t.Error("Fail test creating new RegionType", "expect Name =", validModelRegionType.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelRegionType types.RegionType

var testCreateFuncRegionType = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	RegionTypeCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestRegionType = tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType)

var testCasesRegionType = []tests.WebTest{
    {
		Name:         "Find RegionTypes as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncRegionType(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.RegionTypeRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, RegionTypeFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getRegionTypeParsedList(response)

			if total < 1 {
				t.Error("Error in find RegionType. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsRegionType(t, item, validModelRegionType, response)
			}
		},
	},
    {
		Name:         "Create new RegionType as admin",
		Request:      createAdminRequestRegionType,
		ResponseCode: 201,
		TestFunc:     testCreateFuncRegionType,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelRegionType = getRegionTypeParsedModel(response)

			idsForRemoveRegionType = append(idsForRemoveRegionType, createdModelRegionType.Id)
			validateFieldsRegionType(t, createdModelRegionType, validModelRegionType, response)
		},
	},
	{
		Name:            "Create new RegionType as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.RegionTypeRoute, validModelRegionType),
		ResponseCode:    403,
		TestFunc:        testCreateFuncRegionType,
	},
    {
		Name:         "Read RegionType as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.RegionTypeRoute + "/" + strconv.Itoa(updateModelRegionType.Id))
			return tests.SendRequest(settings.RegionTypeRoute + "/{id}", request, RegionTypeRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read RegionType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RegionTypeRoute, updateModelRegionType)
			responseCreate, err := testCreateFuncRegionType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRegionTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegionType = append(idsForRemoveRegionType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.RegionTypeRoute + "/" + id)

			return tests.SendRequest(settings.RegionTypeRoute + "/{id}", request, RegionTypeRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelRegionType = getRegionTypeParsedModel(response)
			validateFieldsRegionType(t, createdModelRegionType, updateModelRegionType, response)
		},
	},
    {
		Name:         "Update RegionType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType)
			responseCreate, err := testCreateFuncRegionType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRegionTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegionType = append(idsForRemoveRegionType, responseData.Id)

			updateModelRegionType.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.RegionTypeRoute + "/" + id, updateModelRegionType)

			return tests.SendRequest(settings.RegionTypeRoute + "/{id}", request, RegionTypeUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getRegionTypeParsedModel(response)
			validateFieldsRegionType(t, model, updateModelRegionType, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.RegionTypeRoute + "/" + id)
			readResponse := tests.SendRequest(settings.RegionTypeRoute + "/{id}", request, RegionTypeRead, http.MethodGet)

			model = getRegionTypeParsedModel(readResponse)
			validateFieldsRegionType(t, model, updateModelRegionType, readResponse)
		},
	},
    {
		Name: "Delete RegionType as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			RegionTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType))

			responseData, err := getRegionTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegionType = append(idsForRemoveRegionType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.RegionTypeRoute + "/" + id)

			return tests.SendRequest(settings.RegionTypeRoute+"/{id}", tt.Request, RegionTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete RegionType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create RegionType for next delete
			responseCreate := httptest.NewRecorder()
			RegionTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType))

			responseData, err := getRegionTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RegionTypeRoute + "/" + id)

			return tests.SendRequest(settings.RegionTypeRoute+"/{id}", req, RegionTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete RegionType as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create RegionType for next delete
			responseCreate := httptest.NewRecorder()
			RegionTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionTypeRoute, validModelRegionType))

			responseData, err := getRegionTypeResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RegionTypeRoute + "/" + id)
			tests.SendRequest(settings.RegionTypeRoute+"/{id}", req, RegionTypeDelete, http.MethodDelete)

			return tests.SendRequest(settings.RegionTypeRoute+"/{id}", req, RegionTypeDelete, http.MethodDelete), nil
		},
	},
    

}

func getRegionTypeResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.RegionType, error) {

	model := getRegionTypeParsedModel(response)

	if model.Id < 1 {
		return types.RegionType{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getRegionTypeParsedList(response *httptest.ResponseRecorder) (list []types.RegionType, total int) {

	responseData := struct{
		List []types.RegionType
		Total int
	} {
		List: []types.RegionType{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getRegionTypeParsedModel(response *httptest.ResponseRecorder) types.RegionType {

	responseData := struct{
		Model types.RegionType
	} {
		Model:types.RegionType{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestRegionType(t *testing.T) {

	for _, tt := range testCasesRegionType {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveRegionType {
		tests.SendRequest(settings.RegionTypeRoute + "/{id}", tests.GetDeleteAdminRequest(settings.RegionTypeRoute + "/" + strconv.Itoa(id)), RegionTypeDelete, http.MethodDelete)
	}
}

