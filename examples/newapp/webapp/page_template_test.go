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

var validModelPageTemplate = types.PageTemplate{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelPageTemplate = types.PageTemplate{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemovePageTemplate = []int{}

func validateFieldsPageTemplate(t *testing.T, testModel types.PageTemplate, validModelPageTemplate types.PageTemplate, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new PageTemplate", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelPageTemplate.Name {
	//	t.Error("Fail test creating new PageTemplate", "expect Name =", validModelPageTemplate.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelPageTemplate types.PageTemplate

var testCreateFuncPageTemplate = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	PageTemplateCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestPageTemplate = tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate)

var testCasesPageTemplate = []tests.WebTest{
    {
		Name:         "Find PageTemplates as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncPageTemplate(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.PageTemplateRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, PageTemplateFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getPageTemplateParsedList(response)

			if total < 1 {
				t.Error("Error in find PageTemplate. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsPageTemplate(t, item, validModelPageTemplate, response)
			}
		},
	},
    {
		Name:         "Create new PageTemplate as admin",
		Request:      createAdminRequestPageTemplate,
		ResponseCode: 201,
		TestFunc:     testCreateFuncPageTemplate,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelPageTemplate = getPageTemplateParsedModel(response)

			idsForRemovePageTemplate = append(idsForRemovePageTemplate, createdModelPageTemplate.Id)
			validateFieldsPageTemplate(t, createdModelPageTemplate, validModelPageTemplate, response)
		},
	},
	{
		Name:            "Create new PageTemplate as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.PageTemplateRoute, validModelPageTemplate),
		ResponseCode:    403,
		TestFunc:        testCreateFuncPageTemplate,
	},
    {
		Name:         "Read PageTemplate as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.PageTemplateRoute + "/" + strconv.Itoa(updateModelPageTemplate.Id))
			return tests.SendRequest(settings.PageTemplateRoute + "/{id}", request, PageTemplateRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read PageTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageTemplateRoute, updateModelPageTemplate)
			responseCreate, err := testCreateFuncPageTemplate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageTemplate = append(idsForRemovePageTemplate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.PageTemplateRoute + "/" + id)

			return tests.SendRequest(settings.PageTemplateRoute + "/{id}", request, PageTemplateRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelPageTemplate = getPageTemplateParsedModel(response)
			validateFieldsPageTemplate(t, createdModelPageTemplate, updateModelPageTemplate, response)
		},
	},
    {
		Name:         "Update PageTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate)
			responseCreate, err := testCreateFuncPageTemplate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageTemplate = append(idsForRemovePageTemplate, responseData.Id)

			updateModelPageTemplate.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.PageTemplateRoute + "/" + id, updateModelPageTemplate)

			return tests.SendRequest(settings.PageTemplateRoute + "/{id}", request, PageTemplateUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getPageTemplateParsedModel(response)
			validateFieldsPageTemplate(t, model, updateModelPageTemplate, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.PageTemplateRoute + "/" + id)
			readResponse := tests.SendRequest(settings.PageTemplateRoute + "/{id}", request, PageTemplateRead, http.MethodGet)

			model = getPageTemplateParsedModel(readResponse)
			validateFieldsPageTemplate(t, model, updateModelPageTemplate, readResponse)
		},
	},
    {
		Name: "Delete PageTemplate as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			PageTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate))

			responseData, err := getPageTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageTemplate = append(idsForRemovePageTemplate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.PageTemplateRoute + "/" + id)

			return tests.SendRequest(settings.PageTemplateRoute+"/{id}", tt.Request, PageTemplateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageTemplate for next delete
			responseCreate := httptest.NewRecorder()
			PageTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate))

			responseData, err := getPageTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageTemplateRoute + "/" + id)

			return tests.SendRequest(settings.PageTemplateRoute+"/{id}", req, PageTemplateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageTemplate as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageTemplate for next delete
			responseCreate := httptest.NewRecorder()
			PageTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageTemplateRoute, validModelPageTemplate))

			responseData, err := getPageTemplateResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageTemplateRoute + "/" + id)
			tests.SendRequest(settings.PageTemplateRoute+"/{id}", req, PageTemplateDelete, http.MethodDelete)

			return tests.SendRequest(settings.PageTemplateRoute+"/{id}", req, PageTemplateDelete, http.MethodDelete), nil
		},
	},
    

}

func getPageTemplateResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.PageTemplate, error) {

	model := getPageTemplateParsedModel(response)

	if model.Id < 1 {
		return types.PageTemplate{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getPageTemplateParsedList(response *httptest.ResponseRecorder) (list []types.PageTemplate, total int) {

	responseData := struct{
		List []types.PageTemplate
		Total int
	} {
		List: []types.PageTemplate{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getPageTemplateParsedModel(response *httptest.ResponseRecorder) types.PageTemplate {

	responseData := struct{
		Model types.PageTemplate
	} {
		Model:types.PageTemplate{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestPageTemplate(t *testing.T) {

	for _, tt := range testCasesPageTemplate {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemovePageTemplate {
		tests.SendRequest(settings.PageTemplateRoute + "/{id}", tests.GetDeleteAdminRequest(settings.PageTemplateRoute + "/" + strconv.Itoa(id)), PageTemplateDelete, http.MethodDelete)
	}
}

