package logic

import (
	"github.com/pkg/errors"
	"gosha/cmd"
	"gosha/webapp/types"
	"os"
	"regexp"
	"strings"
)

func EntityFind(filter types.EntityFilter) (result []types.Entity, totalRecords int, err error) {
	existsTypes := cmd.GetExistsTypes()
	typeNames := cmd.GetModelsList(existsTypes)
	existsModels := cmd.GetExistsModels()
	modelsNames := cmd.GetModelsList(existsModels)
	modelsMethods := cmd.GetModelsMethods(existsTypes)

	resModels := []types.Entity{}
	res := []types.Entity{}

	id := 1
	for _, t := range typeNames {

		if t != strings.Title(t) {
			continue
		}

		fields := []types.Field{}

		for _, field := range existsTypes.GetFields(t, []cmd.Field{}) {

			if !filter.WithHiddenFields && field.Name != strings.Title(field.Name) {
				continue
			}

			fields = append(fields, types.Field{
				Name:        field.Name,
				Type:        field.Type,
				IsType:      true,
				CommentType: field.Comment,
			})
		}

		isFilter, _ := regexp.Match("Filter", []byte(t))

		res = append(res, types.Entity{
			Id:          id,
			Name:        t,
			Fields:      fields,
			IsFilter:    isFilter,
			CommentType: existsTypes.GetModelComment(t),
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
			Id:          id,
			Name:        t,
			ModelFields: modelFields,
			CommentDb:   existsModels.GetModelComment(t),
		})
		id++
	}

	for _, em := range resModels {

		eres, index := getExistsModel(em.Name, res)

		if len(eres.Name) > 0 {

			for _, emf := range em.ModelFields {

				etf, i := getExistsFieldIndex(emf.Name, eres.Fields)

				if len(etf.Name) < 1 {
					res[index].Fields = append(res[index].Fields, types.Field{
						Name:      emf.Name,
						Type:      emf.Type,
						IsDb:      true,
						CommentDb: emf.Comment,
					})

				} else {
					res[index].Fields[i].IsDb = true
					res[index].Fields[i].CommentDb = emf.Comment
				}
			}

			mres, _ := getExistsModel(em.Name, resModels)
			res[index].CommentDb = mres.CommentDb

		} else {

			for _, mf := range em.ModelFields {

				em.Fields = append(em.Fields, types.Field{
					Name:      mf.Name,
					Type:      mf.Type,
					IsDb:      true,
					CommentDb: mf.Comment,
				})
			}

			res = append(res, em)
		}
	}

	if len(filter.Search) > 0 {

		filtered := []types.Entity{}

		for _, entity := range res {
			if strings.Contains(strings.ToUpper(entity.Name), strings.ToUpper(filter.Search)) {
				filtered = append(filtered, entity)
			} else if strings.Contains(strings.ToUpper(entity.CommentType), strings.ToUpper(filter.Search)) {
				filtered = append(filtered, entity)
			} else if strings.Contains(strings.ToUpper(entity.CommentDb), strings.ToUpper(filter.Search)) {
				filtered = append(filtered, entity)
			}
		}
		result = filtered
	} else {
		result = res
	}

	if !filter.WithFilter {
		filtered := []types.Entity{}
		for _, entity := range result {
			matched, _ := regexp.Match(`Filter`, []byte(entity.Name))
			if !matched {
				filtered = append(filtered, entity)
			}
		}
		result = filtered
	}

	if filter.IsExactMatch {
		filtered := []types.Entity{}
		for _, entity := range result {
			if entity.Name == filter.Search {
				filtered = append(filtered, entity)
			}
		}
		result = filtered
	}

	totalRecords = len(result)

	for k, model := range result {
		result[k].HttpMethods.IsFind = modelsMethods[model.Name].IsFind
		result[k].HttpMethods.IsCreate = modelsMethods[model.Name].IsCreate
		result[k].HttpMethods.IsRead = modelsMethods[model.Name].IsRead
		result[k].HttpMethods.IsUpdate = modelsMethods[model.Name].IsUpdate
		result[k].HttpMethods.IsDelete = modelsMethods[model.Name].IsDelete
		result[k].HttpMethods.IsFindOrCreate = modelsMethods[model.Name].IsFindOrCreate
		result[k].HttpMethods.IsUpdateOrCreate = modelsMethods[model.Name].IsUpdateOrCreate

		result[k].HttpRoutes.Find = cmd.AssignVar(cmd.FindUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.Create = cmd.AssignVar(cmd.CreateUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.Read = cmd.AssignVar(cmd.ReadUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.Update = cmd.AssignVar(cmd.UpdateUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.Delete = cmd.AssignVar(cmd.DeleteUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.FindOrCreate = cmd.AssignVar(cmd.FindOrCreateUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
		result[k].HttpRoutes.UpdateOrCreate = cmd.AssignVar(cmd.UpdateOrCreateUrl, "{entity}", cmd.GetFirstLowerCase(model.Name))
	}
	return
}

func getExistsFieldIndex(name string, fields []types.Field) (types.Field, int) {

	for i, f := range fields {
		if strings.ToLower(f.Name) == strings.ToLower(name) {
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

func EntityRead(_ types.EntityFilter) (data types.Entity, err error) {

	return
}

func EntityCreate(filter types.EntityFilter) (data types.Entity, err error) {

	argsBak := os.Args
	defer func() {
		if filter.IsRegenerateJsTypes {
			args := []string{"", "exit", "gen:types:js"}
			os.Args = args
			cmd.RunShell()
			os.Args = argsBak
		}
	}()

	e := filter.GetEntityModel()

	fcrudax := ""

	if e.HttpMethods.IsFind {
		fcrudax += "f"
	}
	if e.HttpMethods.IsCreate {
		fcrudax += "c"
	}
	if e.HttpMethods.IsRead {
		fcrudax += "r"
	}
	if e.HttpMethods.IsUpdate {
		fcrudax += "u"
	}
	if e.HttpMethods.IsDelete {
		fcrudax += "d"
	}
	if e.HttpMethods.IsFindOrCreate {
		fcrudax += "a"
	}
	if e.HttpMethods.IsUpdateOrCreate {
		fcrudax += "x"
	}

	args := []string{"",
		cmd.Exit.CliArgument(),
		cmd.SetAppType.CliArgument(),
		cmd.Type.CliArgument("Usual"),
		cmd.USUAL_ENTITY_ADD,
		cmd.Entity.CliArgument(e.Name),
		cmd.CRUD.CliArgument(fcrudax),
		cmd.CheckAuth.CliArgument("fcrudax"),
		cmd.CheckAuth.CliArgument("fcrudax"),
	}

	if e.Structures.WithoutDbModel {
		args = append(args, cmd.WithoutDbModels.CliArgument("true"))
	}

	if filter.IsUuidMode {
		args = append(args, cmd.UuidAsPk.CliArgument("true"))
	}

	if filter.IsViewMode {
		args = append(args, cmd.ViewMode.CliArgument("true"))
	}

	if filter.IsSoftDelete {
		args = append(args, cmd.SoftDelete.CliArgument("true"))
	}

	os.Args = args
	cmd.RunShell()

	addFields(e.Name, e.Fields)

	defer func() {
		shell := cmd.GetShell()
		shell.Close()
	}()

	findFilter := types.EntityFilter{}
	findFilter.Search = e.Name
	findFilter.WithFilter = false

	res, records, err := EntityFind(findFilter)

	if records < 0 {
		err = errors.New("Creating entity error")
		return
	}

	for _, fe := range res {
		if fe.Name == e.Name {
			data = fe
			return
		}
	}

	err = errors.New("Creating entity error")

	return
}

func addFields(entityName string, fields []types.Field) {

	var args []string

	for _, f := range fields {
		args = []string{"", "exit", "setAppType", "--type=Usual", cmd.ENTITY_ADD_FIELD, "--entity=" + entityName, "--Field=" + f.Name, "--data-type=" + f.Type}
		os.Args = args
		cmd.RunShell()
	}
}

func addFilters(entityName string, fields []types.Field) {

	var args []string

	for _, f := range fields {
		args = []string{"", "exit", "setAppType", "--type=Usual", cmd.ENTITY_ADD_FILTER, "--entity=" + entityName, "--filter=" + f.Name, "--data-type=" + f.Type}
		os.Args = args
		cmd.RunShell()
	}
}

func EntityUpdate(filter types.EntityFilter) (data types.Entity, err error) {

	argsBak := os.Args
	defer func() {
		if filter.IsRegenerateJsTypes {
			args := []string{"", "exit", "gen:types:js"}
			os.Args = args
			cmd.RunShell()
			os.Args = argsBak
		}
	}()

	e := filter.GetEntityModel()

	isFilter, _ := regexp.Match("Filter", []byte(e.Name))

	if isFilter {
		addFilters(e.Name, e.Fields)
	} else {
		addFields(e.Name, e.Fields)
	}

	defer func() {
		shell := cmd.GetShell()
		shell.Close()
	}()

	findFilter := types.EntityFilter{}
	findFilter.Search = e.Name
	findFilter.WithFilter = isFilter

	res, records, err := EntityFind(findFilter)

	if records < 0 {
		err = errors.New("Updating entity error")
		return
	}

	for _, fe := range res {
		if fe.Name == e.Name {
			data = fe
			return
		}
	}

	err = errors.New("Update entity error")

	return
}

func EntityDelete(_ types.EntityFilter) (isOk bool, err error) {

	return
}
