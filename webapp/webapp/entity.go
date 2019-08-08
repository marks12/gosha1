package webapp

import (
	"net/http"
	"gosha/settings"
	"gosha/webapp/mdl"
	"gosha/webapp/types"
	"gosha/cmd"
	"strings"
)

func EntityFind(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetEntityFilter(httpRequest, settings.FunctionTypeFind)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in FindEntity", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, totalRecords, err := logicEntityFind(requestDto)

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

func EntityCreate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetEntityFilter(httpRequest, settings.FunctionTypeCreate)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in EntityCreate", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicEntityCreate(requestDto)

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

func EntityRead(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetEntityFilter(httpRequest, settings.FunctionTypeRead)

	requestDto.PerPage = 1
	requestDto.CurrentPage = 1


	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in EntityRead", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicEntityRead(requestDto)

	// Создаём структуру ответа
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ValidResponse(w, mdl.ResponseRead{
		data,
	})

	return
}


func EntityUpdate(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetEntityFilter(httpRequest, settings.FunctionTypeUpdate)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in EntityUpdate", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	data, err := logicEntityUpdate(requestDto)

	// Создаём структуру ответа
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ValidResponse(w, mdl.ResponseUpdate{
		data,
	})

	return
}

func EntityDelete(w http.ResponseWriter, httpRequest *http.Request) {

	requestDto := types.GetEntityFilter(httpRequest, settings.FunctionTypeDelete)

	if !requestDto.IsAuthorized() {
		ErrResponse(w, "Invalid authorize in EntityDelete", http.StatusForbidden)
		return
	}

	if !requestDto.IsValid() {
		ErrResponse(w, requestDto.GetValidationErrors(), http.StatusBadRequest)
		return
	}

	// Получаем список
	isOk, err := logicEntityDelete(requestDto)

	// Создаём структуру ответа
	if err != nil {
		ErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ValidResponse(w, mdl.ResponseDelete{
		isOk,
	})

	return
}

func logicEntityFind(filter types.EntityFilter) (result []types.Entity, totalRecords int, err error) {

	existsTypes := cmd.GetExistsTypes()
	typeNames := cmd.GetModelsList(existsTypes)

	existsModels := cmd.GetExistsModels()
	modelsNames := cmd.GetModelsList(existsModels)

	resModels := []types.Entity{}
	resTypes := []types.Entity{}

	id := 1
	for _, t := range typeNames {

		typeFields := []cmd.Field{}

		for _, field := range existsTypes.GetFields(t, []cmd.Field{}) {

			if field.Name == strings.Title(field.Name) {
				typeFields = append(typeFields, field)
			}
		}

		resTypes = append(resTypes, types.Entity{
			Id: id,
			Name: t,
			TypeFields: typeFields,
		})
		id++
	}

	id = 1
	for _, t := range modelsNames {

		modelFields := []cmd.Field{}

		for _, field := range existsModels.GetFields(t, []cmd.Field{}) {

			if field.Name == strings.Title(field.Name) {
				modelFields = append(modelFields, field)
			}
		}

		resModels = append(resModels, types.Entity{
			Id: id,
			Name: t,
			ModelFields: modelFields,
		})
		id++
	}

	//for _, et := range resTypes {
	//
	//	typeFields := []cmd.Field{}
	//	modelFields := []cmd.Field{}
	//
	//	em := getExistsModel(et.Name, resModels)
	//
	//	if len(em.Name) > 0 {
	//
	//		for _, etf := range et.TypeFields {
	//
	//			emf := getExistsField(etf.Name, em.ModelFields)
	//
	//
	//		}
	//
	//	} else {
	//		result = append(result, et)
	//	}
	//}

	return
}

func getExistsModel(s string, entities []types.Entity) (types.Entity) {

	for _, ent := range entities {
		if ent.Name == s {
			return ent
		}
	}

	return types.Entity{}
}

func logicEntityRead(filter types.EntityFilter) (data types.Entity, err error) {

	return
}

func logicEntityCreate(filter types.EntityFilter) (data types.Entity, err error) {

	return
}

func logicEntityUpdate(filter types.EntityFilter) (data types.Entity, err error) {
	return
}

func logicEntityDelete(filter types.EntityFilter) (isOk bool, err error) {

	return
}
