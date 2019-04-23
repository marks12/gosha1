package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"os"
	"io/ioutil"
	"sort"
	"strings"
)

func genTypesJs (c *ishell.Context) {

	folder := "./jstypes"
	file := folder + "/apiModel.js"

	os.RemoveAll(folder)
	os.MkdirAll(folder,0755)

	content := getTypesJs()

	cb := []byte(content)
	ioutil.WriteFile(file, cb, 0644)
}

func getTypesJs () string {

	existsTypes := getExistsTypes()
	typeNames := getModelsList(existsTypes)

	return getFileContent(existsTypes, typeNames)

}

func getFileContent(repository modelRepository, typeNames []string) (content string) {

	for _, t := range typeNames {


		fields := []string{}

		validFieldsCount := 0

		for _, field := range repository.GetFields(t) {

			if field.Name[0:1] == strings.ToUpper(field.Name[0:1]) {

				validFieldsCount = validFieldsCount + 1
				fields = append(fields, "    this." + field.Name + " = " + getFiledJsVal(field.Type, typeNames) + ";\n")
			}
		}

		if validFieldsCount < 1 {
			continue
		}

		content += "\nexport function " + t + "() {\n\n"

		sort.Sort(ByCase(fields))
		content += strings.Join(fields, "")

		content += "\n    return this;\n}\n"
	}

	return
}

func getFiledJsVal(s string, typeNames []string) (val string) {

	switch s {
		case "Array":
			val = "[]"
			break

		case "int":
			val = "0"
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
