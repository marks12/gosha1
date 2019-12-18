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

var validModelPageContent = types.PageContent{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelPageContent = types.PageContent{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemovePageContent = []int{}

func validateFieldsPageContent(t *testing.T, testModel types.PageContent, validModelPageContent types.PageContent, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new PageContent", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelPageContent.Name {
	//	t.Error("Fail test creating new PageContent", "expect Name =", validModelPageContent.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelPageContent types.PageContent

var testCreateFuncPageContent = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	PageContentCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestPageContent = tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent)

var testCasesPageContent = []tests.WebTest{
    {
		Name:         "Find PageContents as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncPageContent(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.PageContentRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, PageContentFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getPageContentParsedList(response)

			if total < 1 {
				t.Error("Error in find PageContent. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsPageContent(t, item, validModelPageContent, response)
			}
		},
	},
    {
		Name:         "Create new PageContent as admin",
		Request:      createAdminRequestPageContent,
		ResponseCode: 201,
		TestFunc:     testCreateFuncPageContent,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelPageContent = getPageContentParsedModel(response)

			idsForRemovePageContent = append(idsForRemovePageContent, createdModelPageContent.Id)
			validateFieldsPageContent(t, createdModelPageContent, validModelPageContent, response)
		},
	},
	{
		Name:            "Create new PageContent as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.PageContentRoute, validModelPageContent),
		ResponseCode:    403,
		TestFunc:        testCreateFuncPageContent,
	},
    {
		Name:         "Read PageContent as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.PageContentRoute + "/" + strconv.Itoa(updateModelPageContent.Id))
			return tests.SendRequest(settings.PageContentRoute + "/{id}", request, PageContentRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read PageContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageContentRoute, updateModelPageContent)
			responseCreate, err := testCreateFuncPageContent(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageContent = append(idsForRemovePageContent, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.PageContentRoute + "/" + id)

			return tests.SendRequest(settings.PageContentRoute + "/{id}", request, PageContentRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelPageContent = getPageContentParsedModel(response)
			validateFieldsPageContent(t, createdModelPageContent, updateModelPageContent, response)
		},
	},
    {
		Name:         "Update PageContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent)
			responseCreate, err := testCreateFuncPageContent(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageContent = append(idsForRemovePageContent, responseData.Id)

			updateModelPageContent.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.PageContentRoute + "/" + id, updateModelPageContent)

			return tests.SendRequest(settings.PageContentRoute + "/{id}", request, PageContentUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getPageContentParsedModel(response)
			validateFieldsPageContent(t, model, updateModelPageContent, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.PageContentRoute + "/" + id)
			readResponse := tests.SendRequest(settings.PageContentRoute + "/{id}", request, PageContentRead, http.MethodGet)

			model = getPageContentParsedModel(readResponse)
			validateFieldsPageContent(t, model, updateModelPageContent, readResponse)
		},
	},
    {
		Name: "Delete PageContent as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			PageContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent))

			responseData, err := getPageContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageContent = append(idsForRemovePageContent, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.PageContentRoute + "/" + id)

			return tests.SendRequest(settings.PageContentRoute+"/{id}", tt.Request, PageContentDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageContent as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageContent for next delete
			responseCreate := httptest.NewRecorder()
			PageContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent))

			responseData, err := getPageContentResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageContentRoute + "/" + id)

			return tests.SendRequest(settings.PageContentRoute+"/{id}", req, PageContentDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageContent as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageContent for next delete
			responseCreate := httptest.NewRecorder()
			PageContentCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageContentRoute, validModelPageContent))

			responseData, err := getPageContentResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageContentRoute + "/" + id)
			tests.SendRequest(settings.PageContentRoute+"/{id}", req, PageContentDelete, http.MethodDelete)

			return tests.SendRequest(settings.PageContentRoute+"/{id}", req, PageContentDelete, http.MethodDelete), nil
		},
	},
    

}

func getPageContentResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.PageContent, error) {

	model := getPageContentParsedModel(response)

	if model.Id < 1 {
		return types.PageContent{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getPageContentParsedList(response *httptest.ResponseRecorder) (list []types.PageContent, total int) {

	responseData := struct{
		List []types.PageContent
		Total int
	} {
		List: []types.PageContent{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getPageContentParsedModel(response *httptest.ResponseRecorder) types.PageContent {

	responseData := struct{
		Model types.PageContent
	} {
		Model:types.PageContent{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestPageContent(t *testing.T) {

	for _, tt := range testCasesPageContent {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemovePageContent {
		tests.SendRequest(settings.PageContentRoute + "/{id}", tests.GetDeleteAdminRequest(settings.PageContentRoute + "/" + strconv.Itoa(id)), PageContentDelete, http.MethodDelete)
	}
}

