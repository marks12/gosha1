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

var validModelPageType = types.PageType{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelPageType = types.PageType{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemovePageType = []int{}

func validateFieldsPageType(t *testing.T, testModel types.PageType, validModelPageType types.PageType, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new PageType", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelPageType.Name {
	//	t.Error("Fail test creating new PageType", "expect Name =", validModelPageType.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelPageType types.PageType

var testCreateFuncPageType = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	PageTypeCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestPageType = tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType)

var testCasesPageType = []tests.WebTest{
    {
		Name:         "Find PageTypes as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncPageType(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.PageTypeRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, PageTypeFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getPageTypeParsedList(response)

			if total < 1 {
				t.Error("Error in find PageType. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsPageType(t, item, validModelPageType, response)
			}
		},
	},
    {
		Name:         "Create new PageType as admin",
		Request:      createAdminRequestPageType,
		ResponseCode: 201,
		TestFunc:     testCreateFuncPageType,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelPageType = getPageTypeParsedModel(response)

			idsForRemovePageType = append(idsForRemovePageType, createdModelPageType.Id)
			validateFieldsPageType(t, createdModelPageType, validModelPageType, response)
		},
	},
	{
		Name:            "Create new PageType as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.PageTypeRoute, validModelPageType),
		ResponseCode:    403,
		TestFunc:        testCreateFuncPageType,
	},
    {
		Name:         "Read PageType as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.PageTypeRoute + "/" + strconv.Itoa(updateModelPageType.Id))
			return tests.SendRequest(settings.PageTypeRoute + "/{id}", request, PageTypeRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read PageType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageTypeRoute, updateModelPageType)
			responseCreate, err := testCreateFuncPageType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageType = append(idsForRemovePageType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.PageTypeRoute + "/" + id)

			return tests.SendRequest(settings.PageTypeRoute + "/{id}", request, PageTypeRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelPageType = getPageTypeParsedModel(response)
			validateFieldsPageType(t, createdModelPageType, updateModelPageType, response)
		},
	},
    {
		Name:         "Update PageType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType)
			responseCreate, err := testCreateFuncPageType(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageType = append(idsForRemovePageType, responseData.Id)

			updateModelPageType.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.PageTypeRoute + "/" + id, updateModelPageType)

			return tests.SendRequest(settings.PageTypeRoute + "/{id}", request, PageTypeUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getPageTypeParsedModel(response)
			validateFieldsPageType(t, model, updateModelPageType, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.PageTypeRoute + "/" + id)
			readResponse := tests.SendRequest(settings.PageTypeRoute + "/{id}", request, PageTypeRead, http.MethodGet)

			model = getPageTypeParsedModel(readResponse)
			validateFieldsPageType(t, model, updateModelPageType, readResponse)
		},
	},
    {
		Name: "Delete PageType as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			PageTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType))

			responseData, err := getPageTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageType = append(idsForRemovePageType, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.PageTypeRoute + "/" + id)

			return tests.SendRequest(settings.PageTypeRoute+"/{id}", tt.Request, PageTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageType as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageType for next delete
			responseCreate := httptest.NewRecorder()
			PageTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType))

			responseData, err := getPageTypeResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageTypeRoute + "/" + id)

			return tests.SendRequest(settings.PageTypeRoute+"/{id}", req, PageTypeDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageType as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageType for next delete
			responseCreate := httptest.NewRecorder()
			PageTypeCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTypeRoute, validModelPageType))

			responseData, err := getPageTypeResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageTypeRoute + "/" + id)
			tests.SendRequest(settings.PageTypeRoute+"/{id}", req, PageTypeDelete, http.MethodDelete)

			return tests.SendRequest(settings.PageTypeRoute+"/{id}", req, PageTypeDelete, http.MethodDelete), nil
		},
	},
    

}

func getPageTypeResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.PageType, error) {

	model := getPageTypeParsedModel(response)

	if model.Id < 1 {
		return types.PageType{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getPageTypeParsedList(response *httptest.ResponseRecorder) (list []types.PageType, total int) {

	responseData := struct{
		List []types.PageType
		Total int
	} {
		List: []types.PageType{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getPageTypeParsedModel(response *httptest.ResponseRecorder) types.PageType {

	responseData := struct{
		Model types.PageType
	} {
		Model:types.PageType{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestPageType(t *testing.T) {

	for _, tt := range testCasesPageType {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemovePageType {
		tests.SendRequest(settings.PageTypeRoute + "/{id}", tests.GetDeleteAdminRequest(settings.PageTypeRoute + "/" + strconv.Itoa(id)), PageTypeDelete, http.MethodDelete)
	}
}

