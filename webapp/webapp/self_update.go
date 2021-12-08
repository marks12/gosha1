package webapp

import (
	"gosha/settings"
	"gosha/webapp/logic"
	"gosha/webapp/mdl"
	"gosha/webapp/types"
	"net/http"
)

    

func SelfUpdateFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSelfUpdateFilter(httpRequest, settings.FunctionTypeFind)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in ProjectInfoFind", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, totalRecords, err := logic.SelfUpdateFind(requestDto)

	// Создаём структуру ответа
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ValidResponse(w, mdl.ResponseFind{
		data,
		totalRecords,
	})

	return
}

func SelfUpdateCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetSelfUpdateFilter(httpRequest, settings.FunctionTypeCreate)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in EntityCreate", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logic.SelfUpdateCreate(requestDto)

	// Создаём структуру ответа
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ValidResponse(w, mdl.ResponseCreate{
		data,
	})

	return
}
