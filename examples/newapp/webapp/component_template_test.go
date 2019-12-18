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

var validModelComponentTemplate = types.ComponentTemplate{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelComponentTemplate = types.ComponentTemplate{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveComponentTemplate = []int{}

func validateFieldsComponentTemplate(t *testing.T, testModel types.ComponentTemplate, validModelComponentTemplate types.ComponentTemplate, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new ComponentTemplate", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelComponentTemplate.Name {
	//	t.Error("Fail test creating new ComponentTemplate", "expect Name =", validModelComponentTemplate.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelComponentTemplate types.ComponentTemplate

var testCreateFuncComponentTemplate = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	ComponentTemplateCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestComponentTemplate = tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate)

var testCasesComponentTemplate = []tests.WebTest{
    {
		Name:         "Find ComponentTemplates as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncComponentTemplate(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.ComponentTemplateRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, ComponentTemplateFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getComponentTemplateParsedList(response)

			if total < 1 {
				t.Error("Error in find ComponentTemplate. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsComponentTemplate(t, item, validModelComponentTemplate, response)
			}
		},
	},
    {
		Name:         "Create new ComponentTemplate as admin",
		Request:      createAdminRequestComponentTemplate,
		ResponseCode: 201,
		TestFunc:     testCreateFuncComponentTemplate,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelComponentTemplate = getComponentTemplateParsedModel(response)

			idsForRemoveComponentTemplate = append(idsForRemoveComponentTemplate, createdModelComponentTemplate.Id)
			validateFieldsComponentTemplate(t, createdModelComponentTemplate, validModelComponentTemplate, response)
		},
	},
	{
		Name:            "Create new ComponentTemplate as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.ComponentTemplateRoute, validModelComponentTemplate),
		ResponseCode:    403,
		TestFunc:        testCreateFuncComponentTemplate,
	},
    {
		Name:         "Read ComponentTemplate as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.ComponentTemplateRoute + "/" + strconv.Itoa(updateModelComponentTemplate.Id))
			return tests.SendRequest(settings.ComponentTemplateRoute + "/{id}", request, ComponentTemplateRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read ComponentTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, updateModelComponentTemplate)
			responseCreate, err := testCreateFuncComponentTemplate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getComponentTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentTemplate = append(idsForRemoveComponentTemplate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.ComponentTemplateRoute + "/" + id)

			return tests.SendRequest(settings.ComponentTemplateRoute + "/{id}", request, ComponentTemplateRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelComponentTemplate = getComponentTemplateParsedModel(response)
			validateFieldsComponentTemplate(t, createdModelComponentTemplate, updateModelComponentTemplate, response)
		},
	},
    {
		Name:         "Update ComponentTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate)
			responseCreate, err := testCreateFuncComponentTemplate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getComponentTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentTemplate = append(idsForRemoveComponentTemplate, responseData.Id)

			updateModelComponentTemplate.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.ComponentTemplateRoute + "/" + id, updateModelComponentTemplate)

			return tests.SendRequest(settings.ComponentTemplateRoute + "/{id}", request, ComponentTemplateUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getComponentTemplateParsedModel(response)
			validateFieldsComponentTemplate(t, model, updateModelComponentTemplate, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.ComponentTemplateRoute + "/" + id)
			readResponse := tests.SendRequest(settings.ComponentTemplateRoute + "/{id}", request, ComponentTemplateRead, http.MethodGet)

			model = getComponentTemplateParsedModel(readResponse)
			validateFieldsComponentTemplate(t, model, updateModelComponentTemplate, readResponse)
		},
	},
    {
		Name: "Delete ComponentTemplate as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			ComponentTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate))

			responseData, err := getComponentTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentTemplate = append(idsForRemoveComponentTemplate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.ComponentTemplateRoute + "/" + id)

			return tests.SendRequest(settings.ComponentTemplateRoute+"/{id}", tt.Request, ComponentTemplateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ComponentTemplate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ComponentTemplate for next delete
			responseCreate := httptest.NewRecorder()
			ComponentTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate))

			responseData, err := getComponentTemplateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ComponentTemplateRoute + "/" + id)

			return tests.SendRequest(settings.ComponentTemplateRoute+"/{id}", req, ComponentTemplateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ComponentTemplate as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ComponentTemplate for next delete
			responseCreate := httptest.NewRecorder()
			ComponentTemplateCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentTemplateRoute, validModelComponentTemplate))

			responseData, err := getComponentTemplateResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ComponentTemplateRoute + "/" + id)
			tests.SendRequest(settings.ComponentTemplateRoute+"/{id}", req, ComponentTemplateDelete, http.MethodDelete)

			return tests.SendRequest(settings.ComponentTemplateRoute+"/{id}", req, ComponentTemplateDelete, http.MethodDelete), nil
		},
	},
    

}

func getComponentTemplateResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.ComponentTemplate, error) {

	model := getComponentTemplateParsedModel(response)

	if model.Id < 1 {
		return types.ComponentTemplate{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getComponentTemplateParsedList(response *httptest.ResponseRecorder) (list []types.ComponentTemplate, total int) {

	responseData := struct{
		List []types.ComponentTemplate
		Total int
	} {
		List: []types.ComponentTemplate{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getComponentTemplateParsedModel(response *httptest.ResponseRecorder) types.ComponentTemplate {

	responseData := struct{
		Model types.ComponentTemplate
	} {
		Model:types.ComponentTemplate{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestComponentTemplate(t *testing.T) {

	for _, tt := range testCasesComponentTemplate {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveComponentTemplate {
		tests.SendRequest(settings.ComponentTemplateRoute + "/{id}", tests.GetDeleteAdminRequest(settings.ComponentTemplateRoute + "/" + strconv.Itoa(id)), ComponentTemplateDelete, http.MethodDelete)
	}
}

