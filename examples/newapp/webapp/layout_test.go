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

var validModelLayout = types.Layout{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelLayout = types.Layout{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveLayout = []int{}

func validateFieldsLayout(t *testing.T, testModel types.Layout, validModelLayout types.Layout, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new Layout", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelLayout.Name {
	//	t.Error("Fail test creating new Layout", "expect Name =", validModelLayout.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelLayout types.Layout

var testCreateFuncLayout = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	LayoutCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestLayout = tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout)

var testCasesLayout = []tests.WebTest{
    {
		Name:         "Find Layouts as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncLayout(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.LayoutRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, LayoutFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getLayoutParsedList(response)

			if total < 1 {
				t.Error("Error in find Layout. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsLayout(t, item, validModelLayout, response)
			}
		},
	},
    {
		Name:         "Create new Layout as admin",
		Request:      createAdminRequestLayout,
		ResponseCode: 201,
		TestFunc:     testCreateFuncLayout,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelLayout = getLayoutParsedModel(response)

			idsForRemoveLayout = append(idsForRemoveLayout, createdModelLayout.Id)
			validateFieldsLayout(t, createdModelLayout, validModelLayout, response)
		},
	},
	{
		Name:            "Create new Layout as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.LayoutRoute, validModelLayout),
		ResponseCode:    403,
		TestFunc:        testCreateFuncLayout,
	},
    {
		Name:         "Read Layout as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.LayoutRoute + "/" + strconv.Itoa(updateModelLayout.Id))
			return tests.SendRequest(settings.LayoutRoute + "/{id}", request, LayoutRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read Layout as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.LayoutRoute, updateModelLayout)
			responseCreate, err := testCreateFuncLayout(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getLayoutResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayout = append(idsForRemoveLayout, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.LayoutRoute + "/" + id)

			return tests.SendRequest(settings.LayoutRoute + "/{id}", request, LayoutRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelLayout = getLayoutParsedModel(response)
			validateFieldsLayout(t, createdModelLayout, updateModelLayout, response)
		},
	},
    {
		Name:         "Update Layout as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout)
			responseCreate, err := testCreateFuncLayout(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getLayoutResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayout = append(idsForRemoveLayout, responseData.Id)

			updateModelLayout.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.LayoutRoute + "/" + id, updateModelLayout)

			return tests.SendRequest(settings.LayoutRoute + "/{id}", request, LayoutUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getLayoutParsedModel(response)
			validateFieldsLayout(t, model, updateModelLayout, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.LayoutRoute + "/" + id)
			readResponse := tests.SendRequest(settings.LayoutRoute + "/{id}", request, LayoutRead, http.MethodGet)

			model = getLayoutParsedModel(readResponse)
			validateFieldsLayout(t, model, updateModelLayout, readResponse)
		},
	},
    {
		Name: "Delete Layout as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			LayoutCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout))

			responseData, err := getLayoutResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayout = append(idsForRemoveLayout, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.LayoutRoute + "/" + id)

			return tests.SendRequest(settings.LayoutRoute+"/{id}", tt.Request, LayoutDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Layout as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Layout for next delete
			responseCreate := httptest.NewRecorder()
			LayoutCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout))

			responseData, err := getLayoutResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.LayoutRoute + "/" + id)

			return tests.SendRequest(settings.LayoutRoute+"/{id}", req, LayoutDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Layout as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Layout for next delete
			responseCreate := httptest.NewRecorder()
			LayoutCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutRoute, validModelLayout))

			responseData, err := getLayoutResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.LayoutRoute + "/" + id)
			tests.SendRequest(settings.LayoutRoute+"/{id}", req, LayoutDelete, http.MethodDelete)

			return tests.SendRequest(settings.LayoutRoute+"/{id}", req, LayoutDelete, http.MethodDelete), nil
		},
	},
    

}

func getLayoutResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Layout, error) {

	model := getLayoutParsedModel(response)

	if model.Id < 1 {
		return types.Layout{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getLayoutParsedList(response *httptest.ResponseRecorder) (list []types.Layout, total int) {

	responseData := struct{
		List []types.Layout
		Total int
	} {
		List: []types.Layout{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getLayoutParsedModel(response *httptest.ResponseRecorder) types.Layout {

	responseData := struct{
		Model types.Layout
	} {
		Model:types.Layout{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestLayout(t *testing.T) {

	for _, tt := range testCasesLayout {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveLayout {
		tests.SendRequest(settings.LayoutRoute + "/{id}", tests.GetDeleteAdminRequest(settings.LayoutRoute + "/" + strconv.Itoa(id)), LayoutDelete, http.MethodDelete)
	}
}

