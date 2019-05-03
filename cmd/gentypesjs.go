package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"os"
	"io/ioutil"
	"sort"
	"strings"
)

type store struct {
	name string
	jscode string
}

func genTypesJs (c *ishell.Context) {

	folder := "./jstypes"
	folderStore := folder + "/store"
	file := folder + "/apiModel.js"
	apiFile := folder + "/api.js"

	os.RemoveAll(folder)
	os.MkdirAll(folder,0755)
	os.MkdirAll(folderStore,0755)

	modelContent, stores := getTypesJs()

	cb := []byte(modelContent)
	ioutil.WriteFile(file, cb, 0644)

	api := []byte(apiContent)
	ioutil.WriteFile(apiFile, api, 0644)

	for _, store := range stores {
		ioutil.WriteFile(folderStore + "/" + store.name + ".js" , []byte(store.jscode), 0644)
	}
}


func getTypesJs () (string, []store) {

	existsTypes := getExistsTypes()
	typeNames := getModelsList(existsTypes)

	return getFileContent(existsTypes, typeNames)

}

func getFileContent(repository modelRepository, typeNames []string) (content string, stores []store) {

	for _, t := range typeNames {

		fields := []string{}

		validFieldsCount := 0

		for _, field := range repository.GetFields(t, []field{}) {

			if field.Name[0:1] == strings.ToUpper(field.Name[0:1]) {

				validFieldsCount = validFieldsCount + 1
				fields = append(fields, "    this." + field.Name + " = " + getFiledJsVal(field.Type, typeNames) + ";\n")

			}
		}

		if validFieldsCount < 1 {
			continue
		}

		stores = append(stores, store{
			name: t,
			jscode: getStoreJsCode(t),
		})

		content += "\nexport function " + t + "() {\n\n"

		sort.Sort(ByCase(fields))
		content += strings.Join(fields, "")

		content += "\n    return this;\n}\n"
	}

	return
}

func getStoreJsCode(entity string) string {

	return assignVar(
		assignVar(storeTemplate, "{Entity}", entity),
		"{entity}", getFirstLowerCase(entity))
}

func getFiledJsVal(s string, typeNames []string) (val string) {

	switch s {
		case "Array":
			val = "[]"
			break

		case "int64":
            val = "0"
            break

		case "int32":
            val = "0"
            break

		case "int":
			val = "0"
			break

		case "float":
            val = "0"
            break

		case "float32":
            val = "0"
            break

		case "float64":
			val = "0.0"
			break

		case "string":
			val = "\"\""
			break

		case "bool":
			val = "false"
			break

		default:

			isExists, index := InArray(s, typeNames)

			if isExists {

				val = typeNames[index]

			} else {

				val = "{}"
			}

			break
	}

	return
}
