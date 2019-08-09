package webapp

import (
	"net/http"
	"gosha/settings"
	"gosha/webapp/mdl"
	"gosha/webapp/types"
	"gosha/cmd"
	"strings"
	"github.com/jinzhu/gorm"
	"regexp"
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
	res := []types.Entity{}

	id := 1
	for _, t := range typeNames {

		if t != strings.Title(t) {
			continue
		}

		fields := []types.Field{}

		for _, field := range existsTypes.GetFields(t, []cmd.Field{}) {

			fields = append(fields, types.Field{
				Name: field.Name,
				Type: field.Type,
				IsType: true,
			})
		}

		res = append(res, types.Entity{
			Id: id,
			Name: t,
			Fields: fields,
		})
		id++
	}

	id = 1
	for _, t := range modelsNames {

		if t != strings.Title(t) {
			continue
		}

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

	for _, em := range resModels {

		eres, index := getExistsModel(em.Name, res)

		if len(eres.Name) > 0 {

			for _, emf := range  em.ModelFields {

				etf, i := getExistsFieldIndex(emf.Name, eres.Fields)

				if len(etf.Name) < 1 {
					res[index].Fields = append(res[index].Fields, types.Field{
						Name: emf.Name,
						Type: emf.Type,
						IsDb: true,
					})
				} else {
					res[index].Fields[i].IsDb = true
				}
			}

		} else {

			for _, mf := range em.ModelFields {
				em.Fields = append(em.Fields, types.Field{
					Name: mf.Name,
					Type: mf.Type,
					IsDb: true,
				})
			}

			result = append(result, em)
		}
	}

	if len(filter.Search) > 0 {

		filtered := []types.Entity{}

		for _, entity := range res {

			matched, _ := regexp.Match(`[a-zA-Z0-9]*` + filter.Search + `[a-zA-Z0-9]*`, []byte(entity.Name))

			if matched {
				filtered = append(filtered, entity)
			}
		}

		result = filtered
	} else {
		result = res
	}

	return
}

func logicEntityFind2(filter types.EntityFilter) (result []types.Entity, totalRecords int, err error) {

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

	for _, et := range resTypes {

		em, _ := getExistsModel(et.Name, resModels)

		if len(em.Name) > 0 {

			for _, etf := range et.TypeFields {
				emf := getExistsField(etf.Name, em.ModelFields)
				et.ModelFields = append(et.ModelFields, emf)
			}

		} else {

			for i := 0; i < len(et.TypeFields); i++ {
				et.ModelFields = append(et.ModelFields, cmd.Field{})
			}
		}

		result = append(result, et)
	}

	for _, em := range resModels {

		eres, index := getExistsModel(em.Name, result)

		if len(eres.Name) > 0 {

			for _, emf := range  em.ModelFields {

				etf := getExistsField(emf.Name, eres.TypeFields)

				if len(etf.Name) < 1 {
					result[index].TypeFields = append(result[index].TypeFields, cmd.Field{})
					result[index].ModelFields = append(result[index].ModelFields, emf)
				}
			}

		} else {

			for i := 0; i < len(em.ModelFields); i++ {
				em.TypeFields = append(em.TypeFields, cmd.Field{})
			}

			result = append(result, em)
		}
	}

	if len(filter.Search) > 0 {

		filtered := []types.Entity{}

		for _, entity := range result {

			matched, _ := regexp.Match(`[a-zA-Z0-9]*` + filter.Search + `[a-zA-Z0-9]*`, []byte(entity.Name))

			if matched {
				filtered = append(filtered, entity)
			}
		}

		result = filtered
	}

	return
}

func getExistsField(name string, fields []cmd.Field) cmd.Field {

	for _, f := range fields {
		if gorm.ToColumnName(f.Name) == gorm.ToColumnName(name) {
			return f
		}
	}

	return cmd.Field{}
}

func getExistsFieldIndex(name string, fields []types.Field) (types.Field, int) {

	for i, f := range fields {
		if gorm.ToColumnName(f.Name) == gorm.ToColumnName(name) {
			return f, i
		}
	}

	return types.Field{}, -1
}

func getExistsModel(s string, entities []types.Entity) (types.Entity, int) {

	for i, ent := range entities {
		if ent.Name == s {
			return ent, i
		}
	}

	return types.Entity{}, 0
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
