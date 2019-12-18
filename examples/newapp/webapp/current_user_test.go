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

var validModelCurrentUser = types.CurrentUser{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelCurrentUser = types.CurrentUser{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveCurrentUser = []int{}

func validateFieldsCurrentUser(t *testing.T, testModel types.CurrentUser, validModelCurrentUser types.CurrentUser, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new CurrentUser", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelCurrentUser.Name {
	//	t.Error("Fail test creating new CurrentUser", "expect Name =", validModelCurrentUser.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelCurrentUser types.CurrentUser

var testCreateFuncCurrentUser = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	CurrentUserCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestCurrentUser = tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser)

var testCasesCurrentUser = []tests.WebTest{
    {
		Name:         "Find CurrentUsers as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncCurrentUser(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.CurrentUserRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, CurrentUserFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getCurrentUserParsedList(response)

			if total < 1 {
				t.Error("Error in find CurrentUser. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsCurrentUser(t, item, validModelCurrentUser, response)
			}
		},
	},
    {
		Name:         "Create new CurrentUser as admin",
		Request:      createAdminRequestCurrentUser,
		ResponseCode: 201,
		TestFunc:     testCreateFuncCurrentUser,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelCurrentUser = getCurrentUserParsedModel(response)

			idsForRemoveCurrentUser = append(idsForRemoveCurrentUser, createdModelCurrentUser.Id)
			validateFieldsCurrentUser(t, createdModelCurrentUser, validModelCurrentUser, response)
		},
	},
	{
		Name:            "Create new CurrentUser as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.CurrentUserRoute, validModelCurrentUser),
		ResponseCode:    403,
		TestFunc:        testCreateFuncCurrentUser,
	},
    {
		Name:         "Read CurrentUser as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.CurrentUserRoute + "/" + strconv.Itoa(updateModelCurrentUser.Id))
			return tests.SendRequest(settings.CurrentUserRoute + "/{id}", request, CurrentUserRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read CurrentUser as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.CurrentUserRoute, updateModelCurrentUser)
			responseCreate, err := testCreateFuncCurrentUser(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getCurrentUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveCurrentUser = append(idsForRemoveCurrentUser, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.CurrentUserRoute + "/" + id)

			return tests.SendRequest(settings.CurrentUserRoute + "/{id}", request, CurrentUserRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelCurrentUser = getCurrentUserParsedModel(response)
			validateFieldsCurrentUser(t, createdModelCurrentUser, updateModelCurrentUser, response)
		},
	},
    {
		Name:         "Update CurrentUser as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser)
			responseCreate, err := testCreateFuncCurrentUser(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getCurrentUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveCurrentUser = append(idsForRemoveCurrentUser, responseData.Id)

			updateModelCurrentUser.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.CurrentUserRoute + "/" + id, updateModelCurrentUser)

			return tests.SendRequest(settings.CurrentUserRoute + "/{id}", request, CurrentUserUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getCurrentUserParsedModel(response)
			validateFieldsCurrentUser(t, model, updateModelCurrentUser, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.CurrentUserRoute + "/" + id)
			readResponse := tests.SendRequest(settings.CurrentUserRoute + "/{id}", request, CurrentUserRead, http.MethodGet)

			model = getCurrentUserParsedModel(readResponse)
			validateFieldsCurrentUser(t, model, updateModelCurrentUser, readResponse)
		},
	},
    {
		Name: "Delete CurrentUser as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			CurrentUserCreate(responseCreate, tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser))

			responseData, err := getCurrentUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveCurrentUser = append(idsForRemoveCurrentUser, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.CurrentUserRoute + "/" + id)

			return tests.SendRequest(settings.CurrentUserRoute+"/{id}", tt.Request, CurrentUserDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete CurrentUser as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create CurrentUser for next delete
			responseCreate := httptest.NewRecorder()
			CurrentUserCreate(responseCreate, tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser))

			responseData, err := getCurrentUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.CurrentUserRoute + "/" + id)

			return tests.SendRequest(settings.CurrentUserRoute+"/{id}", req, CurrentUserDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete CurrentUser as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create CurrentUser for next delete
			responseCreate := httptest.NewRecorder()
			CurrentUserCreate(responseCreate, tests.GetCreateAdminRequest(settings.CurrentUserRoute, validModelCurrentUser))

			responseData, err := getCurrentUserResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.CurrentUserRoute + "/" + id)
			tests.SendRequest(settings.CurrentUserRoute+"/{id}", req, CurrentUserDelete, http.MethodDelete)

			return tests.SendRequest(settings.CurrentUserRoute+"/{id}", req, CurrentUserDelete, http.MethodDelete), nil
		},
	},
    

}

func getCurrentUserResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.CurrentUser, error) {

	model := getCurrentUserParsedModel(response)

	if model.Id < 1 {
		return types.CurrentUser{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getCurrentUserParsedList(response *httptest.ResponseRecorder) (list []types.CurrentUser, total int) {

	responseData := struct{
		List []types.CurrentUser
		Total int
	} {
		List: []types.CurrentUser{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getCurrentUserParsedModel(response *httptest.ResponseRecorder) types.CurrentUser {

	responseData := struct{
		Model types.CurrentUser
	} {
		Model:types.CurrentUser{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestCurrentUser(t *testing.T) {

	for _, tt := range testCasesCurrentUser {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveCurrentUser {
		tests.SendRequest(settings.CurrentUserRoute + "/{id}", tests.GetDeleteAdminRequest(settings.CurrentUserRoute + "/" + strconv.Itoa(id)), CurrentUserDelete, http.MethodDelete)
	}
}

