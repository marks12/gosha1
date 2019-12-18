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

var validModelComponentGroup = types.ComponentGroup{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelComponentGroup = types.ComponentGroup{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveComponentGroup = []int{}

func validateFieldsComponentGroup(t *testing.T, testModel types.ComponentGroup, validModelComponentGroup types.ComponentGroup, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new ComponentGroup", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelComponentGroup.Name {
	//	t.Error("Fail test creating new ComponentGroup", "expect Name =", validModelComponentGroup.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelComponentGroup types.ComponentGroup

var testCreateFuncComponentGroup = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	ComponentGroupCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestComponentGroup = tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup)

var testCasesComponentGroup = []tests.WebTest{
    {
		Name:         "Find ComponentGroups as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncComponentGroup(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.ComponentGroupRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, ComponentGroupFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getComponentGroupParsedList(response)

			if total < 1 {
				t.Error("Error in find ComponentGroup. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsComponentGroup(t, item, validModelComponentGroup, response)
			}
		},
	},
    {
		Name:         "Create new ComponentGroup as admin",
		Request:      createAdminRequestComponentGroup,
		ResponseCode: 201,
		TestFunc:     testCreateFuncComponentGroup,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelComponentGroup = getComponentGroupParsedModel(response)

			idsForRemoveComponentGroup = append(idsForRemoveComponentGroup, createdModelComponentGroup.Id)
			validateFieldsComponentGroup(t, createdModelComponentGroup, validModelComponentGroup, response)
		},
	},
	{
		Name:            "Create new ComponentGroup as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.ComponentGroupRoute, validModelComponentGroup),
		ResponseCode:    403,
		TestFunc:        testCreateFuncComponentGroup,
	},
    {
		Name:         "Read ComponentGroup as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.ComponentGroupRoute + "/" + strconv.Itoa(updateModelComponentGroup.Id))
			return tests.SendRequest(settings.ComponentGroupRoute + "/{id}", request, ComponentGroupRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read ComponentGroup as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ComponentGroupRoute, updateModelComponentGroup)
			responseCreate, err := testCreateFuncComponentGroup(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getComponentGroupResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentGroup = append(idsForRemoveComponentGroup, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.ComponentGroupRoute + "/" + id)

			return tests.SendRequest(settings.ComponentGroupRoute + "/{id}", request, ComponentGroupRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelComponentGroup = getComponentGroupParsedModel(response)
			validateFieldsComponentGroup(t, createdModelComponentGroup, updateModelComponentGroup, response)
		},
	},
    {
		Name:         "Update ComponentGroup as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup)
			responseCreate, err := testCreateFuncComponentGroup(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getComponentGroupResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentGroup = append(idsForRemoveComponentGroup, responseData.Id)

			updateModelComponentGroup.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.ComponentGroupRoute + "/" + id, updateModelComponentGroup)

			return tests.SendRequest(settings.ComponentGroupRoute + "/{id}", request, ComponentGroupUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getComponentGroupParsedModel(response)
			validateFieldsComponentGroup(t, model, updateModelComponentGroup, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.ComponentGroupRoute + "/" + id)
			readResponse := tests.SendRequest(settings.ComponentGroupRoute + "/{id}", request, ComponentGroupRead, http.MethodGet)

			model = getComponentGroupParsedModel(readResponse)
			validateFieldsComponentGroup(t, model, updateModelComponentGroup, readResponse)
		},
	},
    {
		Name: "Delete ComponentGroup as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			ComponentGroupCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup))

			responseData, err := getComponentGroupResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveComponentGroup = append(idsForRemoveComponentGroup, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.ComponentGroupRoute + "/" + id)

			return tests.SendRequest(settings.ComponentGroupRoute+"/{id}", tt.Request, ComponentGroupDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ComponentGroup as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ComponentGroup for next delete
			responseCreate := httptest.NewRecorder()
			ComponentGroupCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup))

			responseData, err := getComponentGroupResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ComponentGroupRoute + "/" + id)

			return tests.SendRequest(settings.ComponentGroupRoute+"/{id}", req, ComponentGroupDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete ComponentGroup as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create ComponentGroup for next delete
			responseCreate := httptest.NewRecorder()
			ComponentGroupCreate(responseCreate, tests.GetCreateAdminRequest(settings.ComponentGroupRoute, validModelComponentGroup))

			responseData, err := getComponentGroupResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.ComponentGroupRoute + "/" + id)
			tests.SendRequest(settings.ComponentGroupRoute+"/{id}", req, ComponentGroupDelete, http.MethodDelete)

			return tests.SendRequest(settings.ComponentGroupRoute+"/{id}", req, ComponentGroupDelete, http.MethodDelete), nil
		},
	},
    

}

func getComponentGroupResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.ComponentGroup, error) {

	model := getComponentGroupParsedModel(response)

	if model.Id < 1 {
		return types.ComponentGroup{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getComponentGroupParsedList(response *httptest.ResponseRecorder) (list []types.ComponentGroup, total int) {

	responseData := struct{
		List []types.ComponentGroup
		Total int
	} {
		List: []types.ComponentGroup{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getComponentGroupParsedModel(response *httptest.ResponseRecorder) types.ComponentGroup {

	responseData := struct{
		Model types.ComponentGroup
	} {
		Model:types.ComponentGroup{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestComponentGroup(t *testing.T) {

	for _, tt := range testCasesComponentGroup {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveComponentGroup {
		tests.SendRequest(settings.ComponentGroupRoute + "/{id}", tests.GetDeleteAdminRequest(settings.ComponentGroupRoute + "/" + strconv.Itoa(id)), ComponentGroupDelete, http.MethodDelete)
	}
}

