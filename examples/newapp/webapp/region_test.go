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

var validModelRegion = types.Region{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelRegion = types.Region{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveRegion = []int{}

func validateFieldsRegion(t *testing.T, testModel types.Region, validModelRegion types.Region, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new Region", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelRegion.Name {
	//	t.Error("Fail test creating new Region", "expect Name =", validModelRegion.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelRegion types.Region

var testCreateFuncRegion = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	RegionCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestRegion = tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion)

var testCasesRegion = []tests.WebTest{
    {
		Name:         "Find Regions as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncRegion(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.RegionRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, RegionFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getRegionParsedList(response)

			if total < 1 {
				t.Error("Error in find Region. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsRegion(t, item, validModelRegion, response)
			}
		},
	},
    {
		Name:         "Create new Region as admin",
		Request:      createAdminRequestRegion,
		ResponseCode: 201,
		TestFunc:     testCreateFuncRegion,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelRegion = getRegionParsedModel(response)

			idsForRemoveRegion = append(idsForRemoveRegion, createdModelRegion.Id)
			validateFieldsRegion(t, createdModelRegion, validModelRegion, response)
		},
	},
	{
		Name:            "Create new Region as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.RegionRoute, validModelRegion),
		ResponseCode:    403,
		TestFunc:        testCreateFuncRegion,
	},
    {
		Name:         "Read Region as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.RegionRoute + "/" + strconv.Itoa(updateModelRegion.Id))
			return tests.SendRequest(settings.RegionRoute + "/{id}", request, RegionRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read Region as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RegionRoute, updateModelRegion)
			responseCreate, err := testCreateFuncRegion(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRegionResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegion = append(idsForRemoveRegion, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.RegionRoute + "/" + id)

			return tests.SendRequest(settings.RegionRoute + "/{id}", request, RegionRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelRegion = getRegionParsedModel(response)
			validateFieldsRegion(t, createdModelRegion, updateModelRegion, response)
		},
	},
    {
		Name:         "Update Region as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion)
			responseCreate, err := testCreateFuncRegion(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getRegionResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegion = append(idsForRemoveRegion, responseData.Id)

			updateModelRegion.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.RegionRoute + "/" + id, updateModelRegion)

			return tests.SendRequest(settings.RegionRoute + "/{id}", request, RegionUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getRegionParsedModel(response)
			validateFieldsRegion(t, model, updateModelRegion, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.RegionRoute + "/" + id)
			readResponse := tests.SendRequest(settings.RegionRoute + "/{id}", request, RegionRead, http.MethodGet)

			model = getRegionParsedModel(readResponse)
			validateFieldsRegion(t, model, updateModelRegion, readResponse)
		},
	},
    {
		Name: "Delete Region as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			RegionCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion))

			responseData, err := getRegionResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveRegion = append(idsForRemoveRegion, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.RegionRoute + "/" + id)

			return tests.SendRequest(settings.RegionRoute+"/{id}", tt.Request, RegionDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Region as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Region for next delete
			responseCreate := httptest.NewRecorder()
			RegionCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion))

			responseData, err := getRegionResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RegionRoute + "/" + id)

			return tests.SendRequest(settings.RegionRoute+"/{id}", req, RegionDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Region as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Region for next delete
			responseCreate := httptest.NewRecorder()
			RegionCreate(responseCreate, tests.GetCreateAdminRequest(settings.RegionRoute, validModelRegion))

			responseData, err := getRegionResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.RegionRoute + "/" + id)
			tests.SendRequest(settings.RegionRoute+"/{id}", req, RegionDelete, http.MethodDelete)

			return tests.SendRequest(settings.RegionRoute+"/{id}", req, RegionDelete, http.MethodDelete), nil
		},
	},
    

}

func getRegionResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Region, error) {

	model := getRegionParsedModel(response)

	if model.Id < 1 {
		return types.Region{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getRegionParsedList(response *httptest.ResponseRecorder) (list []types.Region, total int) {

	responseData := struct{
		List []types.Region
		Total int
	} {
		List: []types.Region{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getRegionParsedModel(response *httptest.ResponseRecorder) types.Region {

	responseData := struct{
		Model types.Region
	} {
		Model:types.Region{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestRegion(t *testing.T) {

	for _, tt := range testCasesRegion {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveRegion {
		tests.SendRequest(settings.RegionRoute + "/{id}", tests.GetDeleteAdminRequest(settings.RegionRoute + "/" + strconv.Itoa(id)), RegionDelete, http.MethodDelete)
	}
}

