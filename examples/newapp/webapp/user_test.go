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

var validModelUser = types.User{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelUser = types.User{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveUser = []int{}

func validateFieldsUser(t *testing.T, testModel types.User, validModelUser types.User, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new User", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelUser.Name {
	//	t.Error("Fail test creating new User", "expect Name =", validModelUser.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelUser types.User

var testCreateFuncUser = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	UserCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestUser = tests.GetCreateAdminRequest(settings.UserRoute, validModelUser)

var testCasesUser = []tests.WebTest{
    {
		Name:         "Find Users as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.UserRoute, validModelUser)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncUser(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.UserRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, UserFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getUserParsedList(response)

			if total < 1 {
				t.Error("Error in find User. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsUser(t, item, validModelUser, response)
			}
		},
	},
    {
		Name:         "Create new User as admin",
		Request:      createAdminRequestUser,
		ResponseCode: 201,
		TestFunc:     testCreateFuncUser,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelUser = getUserParsedModel(response)

			idsForRemoveUser = append(idsForRemoveUser, createdModelUser.Id)
			validateFieldsUser(t, createdModelUser, validModelUser, response)
		},
	},
	{
		Name:            "Create new User as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.UserRoute, validModelUser),
		ResponseCode:    403,
		TestFunc:        testCreateFuncUser,
	},
    {
		Name:         "Read User as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.UserRoute + "/" + strconv.Itoa(updateModelUser.Id))
			return tests.SendRequest(settings.UserRoute + "/{id}", request, UserRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read User as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.UserRoute, updateModelUser)
			responseCreate, err := testCreateFuncUser(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUser = append(idsForRemoveUser, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.UserRoute + "/" + id)

			return tests.SendRequest(settings.UserRoute + "/{id}", request, UserRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelUser = getUserParsedModel(response)
			validateFieldsUser(t, createdModelUser, updateModelUser, response)
		},
	},
    {
		Name:         "Update User as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.UserRoute, validModelUser)
			responseCreate, err := testCreateFuncUser(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUser = append(idsForRemoveUser, responseData.Id)

			updateModelUser.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.UserRoute + "/" + id, updateModelUser)

			return tests.SendRequest(settings.UserRoute + "/{id}", request, UserUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getUserParsedModel(response)
			validateFieldsUser(t, model, updateModelUser, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.UserRoute + "/" + id)
			readResponse := tests.SendRequest(settings.UserRoute + "/{id}", request, UserRead, http.MethodGet)

			model = getUserParsedModel(readResponse)
			validateFieldsUser(t, model, updateModelUser, readResponse)
		},
	},
    {
		Name: "Delete User as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			UserCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoute, validModelUser))

			responseData, err := getUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveUser = append(idsForRemoveUser, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.UserRoute + "/" + id)

			return tests.SendRequest(settings.UserRoute+"/{id}", tt.Request, UserDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete User as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create User for next delete
			responseCreate := httptest.NewRecorder()
			UserCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoute, validModelUser))

			responseData, err := getUserResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.UserRoute + "/" + id)

			return tests.SendRequest(settings.UserRoute+"/{id}", req, UserDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete User as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create User for next delete
			responseCreate := httptest.NewRecorder()
			UserCreate(responseCreate, tests.GetCreateAdminRequest(settings.UserRoute, validModelUser))

			responseData, err := getUserResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.UserRoute + "/" + id)
			tests.SendRequest(settings.UserRoute+"/{id}", req, UserDelete, http.MethodDelete)

			return tests.SendRequest(settings.UserRoute+"/{id}", req, UserDelete, http.MethodDelete), nil
		},
	},
    

}

func getUserResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.User, error) {

	model := getUserParsedModel(response)

	if model.Id < 1 {
		return types.User{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getUserParsedList(response *httptest.ResponseRecorder) (list []types.User, total int) {

	responseData := struct{
		List []types.User
		Total int
	} {
		List: []types.User{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getUserParsedModel(response *httptest.ResponseRecorder) types.User {

	responseData := struct{
		Model types.User
	} {
		Model:types.User{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestUser(t *testing.T) {

	for _, tt := range testCasesUser {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveUser {
		tests.SendRequest(settings.UserRoute + "/{id}", tests.GetDeleteAdminRequest(settings.UserRoute + "/" + strconv.Itoa(id)), UserDelete, http.MethodDelete)
	}
}

