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

var validModelEntity = types.Entity{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelEntity = types.Entity{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemoveEntity = []int{}

func validateFieldsEntity(t *testing.T, testModel types.Entity, validModelEntity types.Entity, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new Entity", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelEntity.Name {
	//	t.Error("Fail test creating new Entity", "expect Name =", validModelEntity.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelEntity types.Entity

var testCreateFuncEntity = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	EntityCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestEntity = tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity)

var testCasesEntity = []tests.WebTest{
    {
		Name:         "Find Entitys as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncEntity(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.EntityRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, EntityFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getEntityParsedList(response)

			if total < 1 {
				t.Error("Error in find Entity. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsEntity(t, item, validModelEntity, response)
			}
		},
	},
    {
		Name:         "Create new Entity as admin",
		Request:      createAdminRequestEntity,
		ResponseCode: 201,
		TestFunc:     testCreateFuncEntity,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelEntity = getEntityParsedModel(response)

			idsForRemoveEntity = append(idsForRemoveEntity, createdModelEntity.Id)
			validateFieldsEntity(t, createdModelEntity, validModelEntity, response)
		},
	},
	{
		Name:            "Create new Entity as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.EntityRoute, validModelEntity),
		ResponseCode:    403,
		TestFunc:        testCreateFuncEntity,
	},
    {
		Name:         "Read Entity as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.EntityRoute + "/" + strconv.Itoa(updateModelEntity.Id))
			return tests.SendRequest(settings.EntityRoute + "/{id}", request, EntityRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read Entity as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.EntityRoute, updateModelEntity)
			responseCreate, err := testCreateFuncEntity(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getEntityResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveEntity = append(idsForRemoveEntity, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.EntityRoute + "/" + id)

			return tests.SendRequest(settings.EntityRoute + "/{id}", request, EntityRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelEntity = getEntityParsedModel(response)
			validateFieldsEntity(t, createdModelEntity, updateModelEntity, response)
		},
	},
    {
		Name:         "Update Entity as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity)
			responseCreate, err := testCreateFuncEntity(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getEntityResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveEntity = append(idsForRemoveEntity, responseData.Id)

			updateModelEntity.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.EntityRoute + "/" + id, updateModelEntity)

			return tests.SendRequest(settings.EntityRoute + "/{id}", request, EntityUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getEntityParsedModel(response)
			validateFieldsEntity(t, model, updateModelEntity, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.EntityRoute + "/" + id)
			readResponse := tests.SendRequest(settings.EntityRoute + "/{id}", request, EntityRead, http.MethodGet)

			model = getEntityParsedModel(readResponse)
			validateFieldsEntity(t, model, updateModelEntity, readResponse)
		},
	},
    {
		Name: "Delete Entity as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			EntityCreate(responseCreate, tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity))

			responseData, err := getEntityResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemoveEntity = append(idsForRemoveEntity, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.EntityRoute + "/" + id)

			return tests.SendRequest(settings.EntityRoute+"/{id}", tt.Request, EntityDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Entity as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Entity for next delete
			responseCreate := httptest.NewRecorder()
			EntityCreate(responseCreate, tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity))

			responseData, err := getEntityResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.EntityRoute + "/" + id)

			return tests.SendRequest(settings.EntityRoute+"/{id}", req, EntityDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete Entity as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create Entity for next delete
			responseCreate := httptest.NewRecorder()
			EntityCreate(responseCreate, tests.GetCreateAdminRequest(settings.EntityRoute, validModelEntity))

			responseData, err := getEntityResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.EntityRoute + "/" + id)
			tests.SendRequest(settings.EntityRoute+"/{id}", req, EntityDelete, http.MethodDelete)

			return tests.SendRequest(settings.EntityRoute+"/{id}", req, EntityDelete, http.MethodDelete), nil
		},
	},
    

}

func getEntityResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Entity, error) {

	model := getEntityParsedModel(response)

	if model.Id < 1 {
		return types.Entity{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getEntityParsedList(response *httptest.ResponseRecorder) (list []types.Entity, total int) {

	responseData := struct{
		List []types.Entity
		Total int
	} {
		List: []types.Entity{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getEntityParsedModel(response *httptest.ResponseRecorder) types.Entity {

	responseData := struct{
		Model types.Entity
	} {
		Model:types.Entity{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestEntity(t *testing.T) {

	for _, tt := range testCasesEntity {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemoveEntity {
		tests.SendRequest(settings.EntityRoute + "/{id}", tests.GetDeleteAdminRequest(settings.EntityRoute + "/" + strconv.Itoa(id)), EntityDelete, http.MethodDelete)
	}
}

