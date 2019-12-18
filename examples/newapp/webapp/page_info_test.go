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

var validModelPageInfo = types.PageInfo{
		Id:   0,
		//Name: "Some Name",
	}

var updateModelPageInfo = types.PageInfo{
		Id:   0,
		//Name: "Some Another Name",
	}

var idsForRemovePageInfo = []int{}

func validateFieldsPageInfo(t *testing.T, testModel types.PageInfo, validModelPageInfo types.PageInfo, response *httptest.ResponseRecorder) {

	if testModel.Id < 1 {
		t.Error("Fail test creating new PageInfo", "expect id > 0", "got", testModel.Id, "response:", response.Body)
	}

	//if testModel.Name != validModelPageInfo.Name {
	//	t.Error("Fail test creating new PageInfo", "expect Name =", validModelPageInfo.Name, "got", testModel.Name, "response:", response.Body)
	//}
}

var createdModelPageInfo types.PageInfo

var testCreateFuncPageInfo = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
	response := httptest.NewRecorder()
	PageInfoCreate(response, tt.Request)
	return response, nil
}

var createAdminRequestPageInfo = tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo)

var testCasesPageInfo = []tests.WebTest{
    {
		Name:         "Find PageInfos as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			req := tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo)
			webtest := tests.WebTest{Request: req}
			webtest.Name = "Creating before find"

			_, err := testCreateFuncPageInfo(webtest)

			if err != nil {
				return nil, err
			}

			route := settings.PageInfoRoute
			request := tests.GetFindAdminRequest(route, 1, 1)

			return tests.SendRequest(route, request, PageInfoFind, http.MethodGet), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			list, total := getPageInfoParsedList(response)

			if total < 1 {
				t.Error("Error in find PageInfo. Total rows must be > 0, got", total)
				return
			}

			for _, item := range list {
				validateFieldsPageInfo(t, item, validModelPageInfo, response)
			}
		},
	},
    {
		Name:         "Create new PageInfo as admin",
		Request:      createAdminRequestPageInfo,
		ResponseCode: 201,
		TestFunc:     testCreateFuncPageInfo,
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {

			createdModelPageInfo = getPageInfoParsedModel(response)

			idsForRemovePageInfo = append(idsForRemovePageInfo, createdModelPageInfo.Id)
			validateFieldsPageInfo(t, createdModelPageInfo, validModelPageInfo, response)
		},
	},
	{
		Name:            "Create new PageInfo as non authorized user",
		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.PageInfoRoute, validModelPageInfo),
		ResponseCode:    403,
		TestFunc:        testCreateFuncPageInfo,
	},
    {
		Name:         "Read PageInfo as non authorized user",
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
			request := tests.GetReadNonAuthorizedUserRequest(settings.PageInfoRoute + "/" + strconv.Itoa(updateModelPageInfo.Id))
			return tests.SendRequest(settings.PageInfoRoute + "/{id}", request, PageInfoRead, http.MethodGet), nil
		},
	},
	{
		Name:         "Read PageInfo as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageInfoRoute, updateModelPageInfo)
			responseCreate, err := testCreateFuncPageInfo(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageInfoResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageInfo = append(idsForRemovePageInfo, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			request := tests.GetReadAdminRequest(settings.PageInfoRoute + "/" + id)

			return tests.SendRequest(settings.PageInfoRoute + "/{id}", request, PageInfoRead, http.MethodGet), nil

		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			createdModelPageInfo = getPageInfoParsedModel(response)
			validateFieldsPageInfo(t, createdModelPageInfo, updateModelPageInfo, response)
		},
	},
    {
		Name:         "Update PageInfo as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			tt.Request = tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo)
			responseCreate, err := testCreateFuncPageInfo(tt)

			if err != nil {
				return nil, err
			}

			responseData, err := getPageInfoResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageInfo = append(idsForRemovePageInfo, responseData.Id)

			updateModelPageInfo.Id = responseData.Id

			id := strconv.Itoa(responseData.Id)
			request := tests.GetUpdateAdminRequest(settings.PageInfoRoute + "/" + id, updateModelPageInfo)

			return tests.SendRequest(settings.PageInfoRoute + "/{id}", request, PageInfoUpdate, http.MethodPut), nil
		},
		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
			model := getPageInfoParsedModel(response)
			validateFieldsPageInfo(t, model, updateModelPageInfo, response)

			id := strconv.Itoa(model.Id)
			request := tests.GetReadAdminRequest(settings.PageInfoRoute + "/" + id)
			readResponse := tests.SendRequest(settings.PageInfoRoute + "/{id}", request, PageInfoRead, http.MethodGet)

			model = getPageInfoParsedModel(readResponse)
			validateFieldsPageInfo(t, model, updateModelPageInfo, readResponse)
		},
	},
    {
		Name: "Delete PageInfo as unauthorized user",
		//Request: inside delete func,
		ResponseCode: 403,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			responseCreate := httptest.NewRecorder()
			PageInfoCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo))

			responseData, err := getPageInfoResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			idsForRemovePageInfo = append(idsForRemovePageInfo, responseData.Id)

			id := strconv.Itoa(responseData.Id)
			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.PageInfoRoute + "/" + id)

			return tests.SendRequest(settings.PageInfoRoute+"/{id}", tt.Request, PageInfoDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageInfo as admin",
		ResponseCode: 200,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageInfo for next delete
			responseCreate := httptest.NewRecorder()
			PageInfoCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo))

			responseData, err := getPageInfoResponseModel(tt, responseCreate)

			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageInfoRoute + "/" + id)

			return tests.SendRequest(settings.PageInfoRoute+"/{id}", req, PageInfoDelete, http.MethodDelete), nil
		},
	},
	{
		Name: "Delete PageInfo as admin two times",
		ResponseCode: 400,
		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {

			// create PageInfo for next delete
			responseCreate := httptest.NewRecorder()
			PageInfoCreate(responseCreate, tests.GetCreateAdminRequest(settings.PageInfoRoute, validModelPageInfo))

			responseData, err := getPageInfoResponseModel(tt, responseCreate)
			if err != nil {
				return nil, err
			}

			id := strconv.Itoa(responseData.Id)
			req := tests.GetDeleteAdminRequest(settings.PageInfoRoute + "/" + id)
			tests.SendRequest(settings.PageInfoRoute+"/{id}", req, PageInfoDelete, http.MethodDelete)

			return tests.SendRequest(settings.PageInfoRoute+"/{id}", req, PageInfoDelete, http.MethodDelete), nil
		},
	},
    

}

func getPageInfoResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.PageInfo, error) {

	model := getPageInfoParsedModel(response)

	if model.Id < 1 {
		return types.PageInfo{}, errors.New("Test " + tt.Name + " fail")
	}

	return model, nil
}


func getPageInfoParsedList(response *httptest.ResponseRecorder) (list []types.PageInfo, total int) {

	responseData := struct{
		List []types.PageInfo
		Total int
	} {
		List: []types.PageInfo{},
		Total: 1,
	}

	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.List, responseData.Total
}

func getPageInfoParsedModel(response *httptest.ResponseRecorder) types.PageInfo {

	responseData := struct{
		Model types.PageInfo
	} {
		Model:types.PageInfo{},
	}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	return responseData.Model
}


func TestPageInfo(t *testing.T) {

	for _, tt := range testCasesPageInfo {
		t.Run(tt.Name, func(t *testing.T) {
			tests.TestFunction(t, tt)
		})
	}

	// clear created data from database
	for _, id := range idsForRemovePageInfo {
		tests.SendRequest(settings.PageInfoRoute + "/{id}", tests.GetDeleteAdminRequest(settings.PageInfoRoute + "/" + strconv.Itoa(id)), PageInfoDelete, http.MethodDelete)
	}
}

