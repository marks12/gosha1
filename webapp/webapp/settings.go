package webapp

import (
	"net/http"
	"gosha/settings"
	"gosha/webapp/mdl"
	"gosha/webapp/types"
)

func SettingFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSettingFilter(httpRequest, settings.FunctionTypeFind)

	if !requestDto.IsAuthorized() {
		errResponse(w, "Invalid authorize in FindSetting", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, totalRecords, err := logicSettingFind(requestDto)

	// Создаём структуру ответа
	if err != nil {
		errResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validResponse(w, mdl.ResponseFind{
		data,
		totalRecords,
	})

	return
}

func SettingCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSettingFilter(httpRequest, settings.FunctionTypeCreate)

	if !requestDto.IsAuthorized() {
		errResponse(w, "Invalid authorize in SettingCreate", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicSettingCreate(requestDto)

	// Создаём структуру ответа
	if err != nil {
		errResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validResponse(w, mdl.ResponseCreate{
		data,
	})

	return
}

func SettingRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSettingFilter(httpRequest, settings.FunctionTypeRead)

	requestDto.PerPage = 1
	requestDto.CurrentPage = 1


	if !requestDto.IsAuthorized() {
		errResponse(w, "Invalid authorize in SettingRead", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicSettingRead(requestDto)

	// Создаём структуру ответа
	if err != nil {
		errResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validResponse(w, mdl.ResponseRead{
		data,
	})

	return
}


func SettingUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSettingFilter(httpRequest, settings.FunctionTypeUpdate)

	if !requestDto.IsAuthorized() {
		errResponse(w, "Invalid authorize in SettingUpdate", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicSettingUpdate(requestDto)

	// Создаём структуру ответа
	if err != nil {
		errResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validResponse(w, mdl.ResponseUpdate{
		data,
	})

	return
}

func SettingDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSettingFilter(httpRequest, settings.FunctionTypeDelete)

	if !requestDto.IsAuthorized() {
		errResponse(w, "Invalid authorize in SettingDelete", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		errResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	isOk, err := logicSettingDelete(requestDto)

	// Создаём структуру ответа
	if err != nil {
		errResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validResponse(w, mdl.ResponseDelete{
		isOk,
	})

	return
}

func logicSettingFind(filter types.SettingFilter) (result []types.Setting, totalRecords int, err error) {

	return
}

func logicSettingRead(filter types.SettingFilter) (data types.Setting, err error) {

	return
}

func logicSettingCreate(filter types.SettingFilter) (data types.Setting, err error) {

	return
}

func logicSettingUpdate(filter types.SettingFilter) (data types.Setting, err error) {
	return
}

func logicSettingDelete(filter types.SettingFilter) (isOk bool, err error) {

	return
}
