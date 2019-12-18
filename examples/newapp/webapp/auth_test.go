package webapp

//
//var validModelAuth = types.Auth{
//		Id:   0,
//		//Name: "Some Name",
//	}
//
//var updateModelAuth = types.Auth{
//		Id:   0,
//		//Name: "Some Another Name",
//	}
//
//var idsForRemoveAuth = []int{}
//
//func validateFieldsAuth(t *testing.T, testModel types.Auth, validModelAuth types.Auth, response *httptest.ResponseRecorder) {
//
//	if testModel.Id < 1 {
//		t.Error("Fail test creating new Auth", "expect id > 0", "got", testModel.Id, "response:", response.Body)
//	}
//
//	//if testModel.Name != validModelAuth.Name {
//	//	t.Error("Fail test creating new Auth", "expect Name =", validModelAuth.Name, "got", testModel.Name, "response:", response.Body)
//	//}
//}
//
//var createdModelAuth types.Auth
//
//var testCreateFuncAuth = func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//	response := httptest.NewRecorder()
//	AuthCreate(response, tt.Request)
//	return response, nil
//}
//
//var createAdminRequestAuth = tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth)
//
//var testCasesAuth = []tests.WebTest{
//    {
//		Name:         "Find Auths as admin",
//		ResponseCode: 200,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			req := tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth)
//			webtest := tests.WebTest{Request: req}
//			webtest.Name = "Creating before find"
//
//			_, err := testCreateFuncAuth(webtest)
//
//			if err != nil {
//				return nil, err
//			}
//
//			route := settings.AuthRoute
//			request := tests.GetFindAdminRequest(route, 1, 1)
//
//			return tests.SendRequest(route, request, AuthFind, http.MethodGet), nil
//		},
//		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
//
//			list, total := getAuthParsedList(response)
//
//			if total < 1 {
//				t.Error("Error in find Auth. Total rows must be > 0, got", total)
//				return
//			}
//
//			for _, item := range list {
//				validateFieldsAuth(t, item, validModelAuth, response)
//			}
//		},
//	},
//    {
//		Name:         "Create new Auth as admin",
//		Request:      createAdminRequestAuth,
//		ResponseCode: 201,
//		TestFunc:     testCreateFuncAuth,
//		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
//
//			createdModelAuth = getAuthParsedModel(response)
//
//			idsForRemoveAuth = append(idsForRemoveAuth, createdModelAuth.Id)
//			validateFieldsAuth(t, createdModelAuth, validModelAuth, response)
//		},
//	},
//	{
//		Name:            "Create new Auth as non authorized user",
//		Request:         tests.GetCreateNonAuthorizedUserRequest(settings.AuthRoute, validModelAuth),
//		ResponseCode:    403,
//		TestFunc:        testCreateFuncAuth,
//	},
//    {
//		Name:         "Read Auth as non authorized user",
//		ResponseCode: 403,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//			request := tests.GetReadNonAuthorizedUserRequest(settings.AuthRoute + "/" + strconv.Itoa(updateModelAuth.Id))
//			return tests.SendRequest(settings.AuthRoute + "/{id}", request, AuthRead, http.MethodGet), nil
//		},
//	},
//	{
//		Name:         "Read Auth as admin",
//		ResponseCode: 200,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			tt.Request = tests.GetCreateAdminRequest(settings.AuthRoute, updateModelAuth)
//			responseCreate, err := testCreateFuncAuth(tt)
//
//			if err != nil {
//				return nil, err
//			}
//
//			responseData, err := getAuthResponseModel(tt, responseCreate)
//
//			if err != nil {
//				return nil, err
//			}
//
//			idsForRemoveAuth = append(idsForRemoveAuth, responseData.Id)
//
//			id := strconv.Itoa(responseData.Id)
//			request := tests.GetReadAdminRequest(settings.AuthRoute + "/" + id)
//
//			return tests.SendRequest(settings.AuthRoute + "/{id}", request, AuthRead, http.MethodGet), nil
//
//		},
//		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
//			createdModelAuth = getAuthParsedModel(response)
//			validateFieldsAuth(t, createdModelAuth, updateModelAuth, response)
//		},
//	},
//    {
//		Name:         "Update Auth as admin",
//		ResponseCode: 200,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			tt.Request = tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth)
//			responseCreate, err := testCreateFuncAuth(tt)
//
//			if err != nil {
//				return nil, err
//			}
//
//			responseData, err := getAuthResponseModel(tt, responseCreate)
//
//			if err != nil {
//				return nil, err
//			}
//
//			idsForRemoveAuth = append(idsForRemoveAuth, responseData.Id)
//
//			updateModelAuth.Id = responseData.Id
//
//			id := strconv.Itoa(responseData.Id)
//			request := tests.GetUpdateAdminRequest(settings.AuthRoute + "/" + id, updateModelAuth)
//
//			return tests.SendRequest(settings.AuthRoute + "/{id}", request, AuthUpdate, http.MethodPut), nil
//		},
//		ResultValidator: func(t *testing.T, response *httptest.ResponseRecorder) {
//			model := getAuthParsedModel(response)
//			validateFieldsAuth(t, model, updateModelAuth, response)
//
//			id := strconv.Itoa(model.Id)
//			request := tests.GetReadAdminRequest(settings.AuthRoute + "/" + id)
//			readResponse := tests.SendRequest(settings.AuthRoute + "/{id}", request, AuthRead, http.MethodGet)
//
//			model = getAuthParsedModel(readResponse)
//			validateFieldsAuth(t, model, updateModelAuth, readResponse)
//		},
//	},
//    {
//		Name: "Delete Auth as unauthorized user",
//		//Request: inside delete func,
//		ResponseCode: 403,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			responseCreate := httptest.NewRecorder()
//			AuthCreate(responseCreate, tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth))
//
//			responseData, err := getAuthResponseModel(tt, responseCreate)
//
//			if err != nil {
//				return nil, err
//			}
//
//			idsForRemoveAuth = append(idsForRemoveAuth, responseData.Id)
//
//			id := strconv.Itoa(responseData.Id)
//			tt.Request = tests.GetDeleteNonAuthorizedUserRequest(settings.AuthRoute + "/" + id)
//
//			return tests.SendRequest(settings.AuthRoute+"/{id}", tt.Request, AuthDelete, http.MethodDelete), nil
//		},
//	},
//	{
//		Name: "Delete Auth as admin",
//		ResponseCode: 200,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			// create Auth for next delete
//			responseCreate := httptest.NewRecorder()
//			AuthCreate(responseCreate, tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth))
//
//			responseData, err := getAuthResponseModel(tt, responseCreate)
//
//			if err != nil {
//				return nil, err
//			}
//
//			id := strconv.Itoa(responseData.Id)
//			req := tests.GetDeleteAdminRequest(settings.AuthRoute + "/" + id)
//
//			return tests.SendRequest(settings.AuthRoute+"/{id}", req, AuthDelete, http.MethodDelete), nil
//		},
//	},
//	{
//		Name: "Delete Auth as admin two times",
//		ResponseCode: 400,
//		TestFunc: func(tt tests.WebTest) (*httptest.ResponseRecorder, error) {
//
//			// create Auth for next delete
//			responseCreate := httptest.NewRecorder()
//			AuthCreate(responseCreate, tests.GetCreateAdminRequest(settings.AuthRoute, validModelAuth))
//
//			responseData, err := getAuthResponseModel(tt, responseCreate)
//			if err != nil {
//				return nil, err
//			}
//
//			id := strconv.Itoa(responseData.Id)
//			req := tests.GetDeleteAdminRequest(settings.AuthRoute + "/" + id)
//			tests.SendRequest(settings.AuthRoute+"/{id}", req, AuthDelete, http.MethodDelete)
//
//			return tests.SendRequest(settings.AuthRoute+"/{id}", req, AuthDelete, http.MethodDelete), nil
//		},
//	},
//
//
//}
//
//func getAuthResponseModel(tt tests.WebTest, response *httptest.ResponseRecorder) (types.Auth, error) {
//
//	model := getAuthParsedModel(response)
//
//	if model.Id < 1 {
//		return types.Auth{}, errors.New("Test " + tt.Name + " fail")
//	}
//
//	return model, nil
//}
//
//
//func getAuthParsedList(response *httptest.ResponseRecorder) (list []types.Auth, total int) {
//
//	responseData := struct{
//		List []types.Auth
//		Total int
//	} {
//		List: []types.Auth{},
//		Total: 1,
//	}
//
//	json.Unmarshal(response.Body.Bytes(), &responseData)
//
//	return responseData.List, responseData.Total
//}
//
//func getAuthParsedModel(response *httptest.ResponseRecorder) types.Auth {
//
//	responseData := struct{
//		Model types.Auth
//	} {
//		Model:types.Auth{},
//	}
//	json.Unmarshal(response.Body.Bytes(), &responseData)
//
//	return responseData.Model
//}
//
//
//func TestAuth(t *testing.T) {
//
//	for _, tt := range testCasesAuth {
//		t.Run(tt.Name, func(t *testing.T) {
//			tests.TestFunction(t, tt)
//		})
//	}
//
//	// clear created data from database
//	for _, id := range idsForRemoveAuth {
//		tests.SendRequest(settings.AuthRoute + "/{id}", tests.GetDeleteAdminRequest(settings.AuthRoute + "/" + strconv.Itoa(id)), AuthDelete, http.MethodDelete)
//	}
//}
//
