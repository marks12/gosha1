package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"fmt"
	"strings"
	"regexp"
	"github.com/jinzhu/gorm"
	"os"
)

type modelEntity struct {
	Id int
	Name     string
	TypeFields	 []Field
	ModelFields	 []Field
	Fields	 []Field
	IsFilter bool

	Structures Structures
	HttpMethods HttpMethods
	HttpRoutes HttpRoutes
}

func createViewForms(c *ishell.Context) {

	entities, _, _ := GetEntities()

	for _, ent := range entities {

		CamelCase := strings.Title(ent.Name)
		snakeCase := getLowerCase(ent.Name)
		firstLowerCase := GetFirstLowerCase(ent.Name)

		file := fmt.Sprintf("./view/form/%s.go",snakeCase)

		if _, err := os.Stat(file); ! os.IsNotExist(err) {
			fmt.Println("file allready exists", file)
			continue
		}
		fmt.Println("creating file", file)

		CreateFileIfNotExists(file, getEntityBs4vView(), c)

		CopyFile(
			file,
			file,
			[]string{"{entity-name}", "{Entity}", "{entity}"},
			[]string{CamelCase, CamelCase, firstLowerCase},
			nil)

		for _, field := range ent.Fields {

			//exclude private fields
			if field.Name == GetFirstLowerCase(field.Name) || field.Name == "Id" || field.Name == "CreatedAt" || field.Name == "UpdatedAt" || field.Name == "DeletedAt" {
				continue
			}

			CopyFile(
				file,
				file,
				[]string{getRemoveLine(CamelCase)},
				[]string{GetFormTemplateField(ent.Name, field.Name, strings.ToLower(field.Type)) + "\n            " + getRemoveLine(CamelCase)},
				nil)
		}
	}
}

func GetEntities() (result []modelEntity, totalRecords int, err error) {

	existsTypes := GetExistsTypes()
	typeNames := GetModelsList(existsTypes)
	existsModels := GetExistsModels()
	modelsNames := GetModelsList(existsModels)

	resModels := []modelEntity{}
	res := []modelEntity{}

	id := 1
	for _, t := range typeNames {

		if t != strings.Title(t) {
			continue
		}

		fields := []Field{}

		for _, field := range existsTypes.GetFields(t, []Field{}) {

			fields = append(fields, Field{
				Name:   field.Name,
				Type:   field.Type,
			})
		}

		isFilter, _ := regexp.Match("Filter", []byte(t))

		res = append(res, modelEntity{
			Id:       id,
			Name:     t,
			Fields:   fields,
			IsFilter: isFilter,
		})
		id++
	}

	id = 1
	for _, t := range modelsNames {

		if t != strings.Title(t) {
			continue
		}

		modelFields := []Field{}

		for _, field := range existsModels.GetFields(t, []Field{}) {

			if field.Name == strings.Title(field.Name) {
				modelFields = append(modelFields, field)
			}
		}

		resModels = append(resModels, modelEntity{
			Id:          id,
			Name:        t,
			ModelFields: modelFields,
		})
		id++
	}

	for _, em := range resModels {

		eres, index := getExistsModel(em.Name, res)

		if len(eres.Name) > 0 {

			for _, emf := range em.ModelFields {

				etf, _ := getExistsFieldIndex(emf.Name, eres.Fields)

				if len(etf.Name) < 1 {
					res[index].Fields = append(res[index].Fields, Field{
						Name: emf.Name,
						Type: emf.Type,
					})
				} else {
				}
			}

		} else {

			for _, mf := range em.ModelFields {

				em.Fields = append(em.Fields, Field{
					Name: mf.Name,
					Type: mf.Type,
				})
			}

			res = append(res, em)
		}
	}



	filtered := []modelEntity{}

	for _, entity := range res {

		matched, _ := regexp.Match(`Filter`, []byte(entity.Name))

		if !matched {
			filtered = append(filtered, entity)
		}
	}

	result = filtered

	totalRecords = len(result)

	return 
}


func getExistsFieldIndex(name string, fields []Field) (Field, int) {

	for i, f := range fields {
		if gorm.ToColumnName(f.Name) == gorm.ToColumnName(name) {
			return f, i
		}
	}

	return Field{}, -1
}

func getExistsModel(s string, entities []modelEntity) (modelEntity, int) {

	for i, ent := range entities {
		if ent.Name == s {
			return ent, i
		}
	}

	return modelEntity{}, 0
}