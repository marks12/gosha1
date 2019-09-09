package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"io/ioutil"
	"os"
	"strings"
)

type store struct {
	name   string
	jscode string
}

type vueComponent struct {
	name   string
	jscode string
}

type vueComponentData struct {
	name   string
	jscode string
}

func genTypesJs(c *ishell.Context) {

	folder := "./jstypes"
	folderStore := folder + "/store"
	folderComponent := folder + "/components"
	folderData := folder + "/data"
	file := folder + "/apiModel.js"
	apiFile := folder + "/api.js"
	apiCSRFile := folder + "/apiCSR.js"
	apiSSRFile := folder + "/apiSSR.js"
	commonFile := folder + "/common.js"

	os.RemoveAll(folder)
	os.MkdirAll(folder, 0755)
	os.MkdirAll(folderStore, 0755)
	os.MkdirAll(folderComponent, 0755)
	os.MkdirAll(folderData, 0755)

	modelContent, stores, tpls, data := getTypesJs()

	cb := []byte(modelContent)
	ioutil.WriteFile(file, cb, 0644)

	api := []byte(apiContent)
	ioutil.WriteFile(apiFile, api, 0644)

	apiCSR := []byte(apiCSRContent)
	ioutil.WriteFile(apiCSRFile, apiCSR, 0644)

	apiSSR := []byte(apiSSRContent)
	ioutil.WriteFile(apiSSRFile, apiSSR, 0644)

	common := []byte(commonContent)
	ioutil.WriteFile(commonFile, common, 0644)

	for _, store := range stores {
		ioutil.WriteFile(folderStore+"/"+store.name+".js", []byte(store.jscode), 0644)
	}

	for _, tpl := range tpls {
		ioutil.WriteFile(folderComponent+"/"+tpl.name+"Gen.vue", []byte(tpl.jscode), 0644)
	}

	for _, d := range data {
		ioutil.WriteFile(folderData+"/"+d.name+"Data.js", []byte(d.jscode), 0644)
	}
}

func getTypesJs() (string, []store, []vueComponent, []vueComponentData) {

	existsTypes := GetExistsTypes()
	typeNames := GetModelsList(existsTypes)

	return getFileContent(existsTypes, typeNames)

}

func getFileContent(repository ModelRepository, typeNames []string) (content string, stores []store, vueComponentTemplates []vueComponent, vueData []vueComponentData) {

	for _, t := range typeNames {

		fields := []string{}

		validFieldsCount := 0

		isInvaliFilters, _ := InArray(t, []string{"AbstractFilter", "SearchFilter"})

		// except token
		if isInvaliFilters == true {
			continue
		}

		fieldsList := repository.GetFields(t, []Field{})

		for _, field := range fieldsList {

			if field.Name[0:1] == strings.ToUpper(field.Name[0:1]) {

				isInvalidField, _ := InArray(field.Name, []string{"TotalRecords", "Token", "TotalPages"})

				// except token
				if isInvalidField == true {
					continue
				}

				validFieldsCount = validFieldsCount + 1
				fields = append(fields, "    this."+field.Name+" = "+getFiledJsVal(field.Type, typeNames)+";\n")

			}
		}

		if validFieldsCount < 1 {
			continue
		}

		stores = append(stores, store{
			name:   t,
			jscode: getStoreJsCode(t),
		})

		vueComponentTemplates = append(vueComponentTemplates, vueComponent{
			name:   t,
			jscode: getTemplateJsCode(t),
		})

		vueData = append(vueData, vueComponentData{
			name:   t,
			jscode: getTemplateDataJsCode(t),
		})

		content += "\nexport function " + t + "() {\n\n"

		//sort.Sort(ByCase(fields))
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

func getTemplateJsCode(entity string) string {

	return assignVar(
		assignVar(usualEntityVueComponent, "{Entity}", entity),
		"{entity}", getFirstLowerCase(entity))
}

func getTemplateDataJsCode(entity string) string {

	return assignVar(
		assignVar(usualEntityVueComponentData, "{Entity}", entity),
		"{entity}", getFirstLowerCase(entity))
}

func getFiledJsVal(s string, typeNames []string) (val string) {

	s = strings.ToLower(s)

	switch s {
	case "array":
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

	case "*int":
		val = "null"
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

			val = "new " + typeNames[index] + "()"

		} else {

			val = "{}"
		}

		break
	}

	return
}
