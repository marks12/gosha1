package webapp

import (
	"webapp/settings"
	"webapp/tests"
	"webapp/types"
	"webapp/core"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var validModelSelfUpdate = types.SelfUpdate{
		Id:   1,
		//Name: "Some Name",
	}

var updateModelSelfUpdate = types.SelfUpdate{
		Id:   1,
		//Name: "Some Another Name",
	}

var idsForRemoveSelfUpdate = []int{}

func validateFieldsSelfUpdate(t *testing.T, testModel types.SelfUpdate, validModelSelfUpdate types.SelfUpdate, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new SelfUpdate", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelSelfUpdate.Name {
	//	t.Error("Fail test creating new SelfUpdate", "expect Name =", validModelSelfUpdate.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelSelfUpdate types.SelfUpdate

var testCreateFuncSelfUpdate = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	SelfUpdateCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestSelfUpdate = tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate)

var testCasesSelfUpdate = []tests.WebTest{
    {
		Name:         "Find SelfUpdates as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncSelfUpdate(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.SelfUpdateRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, SelfUpdateFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getSelfUpdateParsedList(response)

			if total < 1 {
				t.Error("Error in find SelfUpdate. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsSelfUpdate(t, item, validModelSelfUpdate, response)
			}
		},
	},
    {
		Name:         "Create new SelfUpdate as admin",
		Request:      createAdminRequestSelfUpdate,
		ResponseCode: 201,
		TestFunc:     testCreateFuncSelfUpdate,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelSelfUpdate = getSelfUpdateParsedModel(response)

			idsForRemoveSelfUpdate = append(idsForRemoveSelfUpdate, createdModelSelfUpdate.Id)
			validateFieldsSelfUpdate(t, createdModelSelfUpdate, validModelSelfUpdate, response)
		},
	},
	{
		Name:            "Create new SelfUpdate as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.SelfUpdateRoute, validModelSelfUpdate),
		ResponseCode:    403,
		TestFunc:        testCreateFuncSelfUpdate,
	},
    {
		Name:         "Read SelfUpdate as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.SelfUpdateRoute + "/" + strconv.Itoa(updateModelSelfUpdate.Id))
			return tests.SendRequest(settings.SelfUpdateRoute + "/{id}", request, SelfUpdateRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read SelfUpdate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.SelfUpdateRoute, updateModelSelfUpdate)
			responseCreate, err := testCreateFuncSelfUpdate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getSelfUpdateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveSelfUpdate = append(idsForRemoveSelfUpdate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.SelfUpdateRoute + "/" + id)

			return tests.SendRequest(settings.SelfUpdateRoute + "/{id}", request, SelfUpdateRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelSelfUpdate = getSelfUpdateParsedModel(response)
			validateFieldsSelfUpdate(t, createdModelSelfUpdate, updateModelSelfUpdate, response)
		},
	},
    {
		Name:         "Update SelfUpdate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate)
			responseCreate, err := testCreateFuncSelfUpdate(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getSelfUpdateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveSelfUpdate = append(idsForRemoveSelfUpdate, responseData.Id)

			updateModelSelfUpdate.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.SelfUpdateRoute + "/" + id, updateModelSelfUpdate)

			return tests.SendRequest(settings.SelfUpdateRoute + "/{id}", request, SelfUpdateUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getSelfUpdateParsedModel(response)
			validateFieldsSelfUpdate(t, model, updateModelSelfUpdate, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.SelfUpdateRoute + "/" + id)
			readResponse := tests.SendRequest(settings.SelfUpdateRoute + "/{id}", request, SelfUpdateRead, http.MethodGet)

			model = getSelfUpdateParsedModel(readResponse)
			validateFieldsSelfUpdate(t, model, updateModelSelfUpdate, readResponse)
		},
	},
    {
		Name: "Delete SelfUpdate as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			SelfUpdateCreate(responseCreate, tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate))

			responseData, err := getSelfUpdateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveSelfUpdate = append(idsForRemoveSelfUpdate, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.SelfUpdateRoute + "/" + id)

			return tests.SendRequest(settings.SelfUpdateRoute+"/{id}", tt.Request, SelfUpdateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete SelfUpdate as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create SelfUpdate for next delete
			responseCreate := httptest.NewRecorder()
			SelfUpdateCreate(responseCreate, tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate))

			responseData, err := getSelfUpdateResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.SelfUpdateRoute + "/" + id)

			return tests.SendRequest(settings.SelfUpdateRoute+"/{id}", req, SelfUpdateDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete SelfUpdate as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create SelfUpdate for next delete
			responseCreate := httptest.NewRecorder()
			SelfUpdateCreate(responseCreate, tests.GetCreateAdminRequest(settings.SelfUpdateRoute, validModelSelfUpdate))

			responseData, err := getSelfUpdateResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.SelfUpdateRoute + "/" + id)
			tests.SendRequest(settings.SelfUpdateRoute+"/{id}", req, SelfUpdateDelete, http.MethodDelete)

			return tests.SendRequest(settings.SelfUpdateRoute+"/{id}", req, SelfUpdateDelete, http.MethodDelete), nil
		},
	},
    

    

}

func getSelfUpdateResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.SelfUpdate, error) {

	model := getSelfUpdateParsedModel(response)

	if model.Id < 1 {
		return types.SelfUpdate{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getSelfUpdateParsedList(response *httptest.ResponseRecorder) (list []types.SelfUpdate, total int) {

	responseData := struct{
		List []types.SelfUpdate
		Total int
	} {
		List: []types.SelfUpdate{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getSelfUpdateParsedModel(response *httptest.ResponseRecorder) types.SelfUpdate {

	responseData := struct{
		Model types.SelfUpdate
	} {
		Model:types.SelfUpdate{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestSelfUpdate(t *testing.T) {

	tmpDb := core.Db
	core.Db = tmpDb.Begin()

	for _, tt := range testCasesSelfUpdate {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	defer func() {
		core.Db.Rollback()
		core.Db = tmpDb
	}()
}

