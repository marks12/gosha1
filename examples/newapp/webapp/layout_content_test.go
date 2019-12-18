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

var validModelLayoutContent = types.LayoutContent{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelLayoutContent = types.LayoutContent{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveLayoutContent = []int{}

func validateFieldsLayoutContent(t *testing.T, testModel types.LayoutContent, validModelLayoutContent types.LayoutContent, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new LayoutContent", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelLayoutContent.Name {
	//	t.Error("Fail test creating new LayoutContent", "expect Name =", validModelLayoutContent.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelLayoutContent types.LayoutContent

var testCreateFuncLayoutContent = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	LayoutContentCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestLayoutContent = tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent)

var testCasesLayoutContent = []tests.WebTest{
    {
		Name:         "Find LayoutContents as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncLayoutContent(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.LayoutContentRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, LayoutContentFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getLayoutContentParsedList(response)

			if total < 1 {
				t.Error("Error in find LayoutContent. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsLayoutContent(t, item, validModelLayoutContent, response)
			}
		},
	},
    {
		Name:         "Create new LayoutContent as admin",
		Request:      createAdminRequestLayoutContent,
		ResponseCode: 201,
		TestFunc:     testCreateFuncLayoutContent,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelLayoutContent = getLayoutContentParsedModel(response)

			idsForRemoveLayoutContent = append(idsForRemoveLayoutContent, createdModelLayoutContent.Id)
			validateFieldsLayoutContent(t, createdModelLayoutContent, validModelLayoutContent, response)
		},
	},
	{
		Name:            "Create new LayoutContent as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.LayoutContentRoute, validModelLayoutContent),
		ResponseCode:    403,
		TestFunc:        testCreateFuncLayoutContent,
	},
    {
		Name:         "Read LayoutContent as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.LayoutContentRoute + "/" + strconv.Itoa(updateModelLayoutContent.Id))
			return tests.SendRequest(settings.LayoutContentRoute + "/{id}", request, LayoutContentRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read LayoutContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.LayoutContentRoute, updateModelLayoutContent)
			responseCreate, err := testCreateFuncLayoutContent(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getLayoutContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayoutContent = append(idsForRemoveLayoutContent, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.LayoutContentRoute + "/" + id)

			return tests.SendRequest(settings.LayoutContentRoute + "/{id}", request, LayoutContentRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelLayoutContent = getLayoutContentParsedModel(response)
			validateFieldsLayoutContent(t, createdModelLayoutContent, updateModelLayoutContent, response)
		},
	},
    {
		Name:         "Update LayoutContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent)
			responseCreate, err := testCreateFuncLayoutContent(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getLayoutContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayoutContent = append(idsForRemoveLayoutContent, responseData.Id)

			updateModelLayoutContent.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.LayoutContentRoute + "/" + id, updateModelLayoutContent)

			return tests.SendRequest(settings.LayoutContentRoute + "/{id}", request, LayoutContentUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getLayoutContentParsedModel(response)
			validateFieldsLayoutContent(t, model, updateModelLayoutContent, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.LayoutContentRoute + "/" + id)
			readResponse := tests.SendRequest(settings.LayoutContentRoute + "/{id}", request, LayoutContentRead, http.MethodGet)

			model = getLayoutContentParsedModel(readResponse)
			validateFieldsLayoutContent(t, model, updateModelLayoutContent, readResponse)
		},
	},
    {
		Name: "Delete LayoutContent as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			LayoutContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent))

			responseData, err := getLayoutContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveLayoutContent = append(idsForRemoveLayoutContent, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.LayoutContentRoute + "/" + id)

			return tests.SendRequest(settings.LayoutContentRoute+"/{id}", tt.Request, LayoutContentDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete LayoutContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create LayoutContent for next delete
			responseCreate := httptest.NewRecorder()
			LayoutContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent))

			responseData, err := getLayoutContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.LayoutContentRoute + "/" + id)

			return tests.SendRequest(settings.LayoutContentRoute+"/{id}", req, LayoutContentDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete LayoutContent as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create LayoutContent for next delete
			responseCreate := httptest.NewRecorder()
			LayoutContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.LayoutContentRoute, validModelLayoutContent))

			responseData, err := getLayoutContentResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.LayoutContentRoute + "/" + id)
			tests.SendRequest(settings.LayoutContentRoute+"/{id}", req, LayoutContentDelete, http.MethodDelete)

			return tests.SendRequest(settings.LayoutContentRoute+"/{id}", req, LayoutContentDelete, http.MethodDelete), nil
		},
	},
    

}

func getLayoutContentResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.LayoutContent, error) {

	model := getLayoutContentParsedModel(response)

	if model.Id < 1 {
		return types.LayoutContent{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getLayoutContentParsedList(response *httptest.ResponseRecorder) (list []types.LayoutContent, total int) {

	responseData := struct{
		List []types.LayoutContent
		Total int
	} {
		List: []types.LayoutContent{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getLayoutContentParsedModel(response *httptest.ResponseRecorder) types.LayoutContent {

	responseData := struct{
		Model types.LayoutContent
	} {
		Model:types.LayoutContent{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestLayoutContent(t *testing.T) {

	for _, tt := range testCasesLayoutContent {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveLayoutContent {
		tests.SendRequest(settings.LayoutContentRoute + "/{id}", tests.GetDeleteAdminRequest(settings.LayoutContentRoute + "/" + strconv.Itoa(id)), LayoutContentDelete, http.MethodDelete)
	}
}

