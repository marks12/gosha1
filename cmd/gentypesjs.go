package cmd

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/abiosoft/ishell.v2"
)

type store struct {
	name   string
	jscode string
}

type jsRoute struct {
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

	//map-namespace
	storeNamespace, _ := GetOsArgument("map-namespace")

	folder := "./jstypes"
	folderStore := folder + "/store"
	folderRoute := folder + "/route"
	folderData := folder + "/data"
	file := folder + "/apiModel.js"
	apiFile := folder + "/api.js"
	apiCSRFile := folder + "/apiCSR.js"
	apiSSRFile := folder + "/apiSSR.js"
	commonFile := folder + "/common.js"

	os.RemoveAll(folder)
	os.MkdirAll(folder, 0755)
	os.MkdirAll(folderStore, 0755)
	os.MkdirAll(folderData, 0755)
	os.MkdirAll(folderRoute, 0755)

	modelContent, stores, _, data, routes := getTypesJs(storeNamespace.StringResult)

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

	for _, d := range data {
		ioutil.WriteFile(folderData+"/"+d.name+"Data.js", []byte(d.jscode), 0644)
	}

	for _, r := range routes {
		ioutil.WriteFile(folderRoute+"/"+r.name+".js", []byte(r.jscode), 0644)
	}
}

func getTypesJs(storeNameSpace string) (string, []store, []vueComponent, []vueComponentData, []jsRoute) {

	existsTypes := GetExistsTypes()
	typeNames := GetModelsList(existsTypes)

	return getFileContent(existsTypes, typeNames, storeNameSpace)

}

func getFileContent(repository ModelRepository, typeNames []string, storeNameSpace string) (content string, stores []store, vueComponentTemplates []vueComponent, vueData []vueComponentData, routes []jsRoute) {

	content = `
	//some common actions
`

	for _, t := range typeNames {

		fields := []string{}

		validFieldsCount := 0

		isInvaliFilters, _ := InArray(t, []string{"AbstractFilter", "GoshaSearchFilter", "GoshaOrderFilter", "FilterIds"})

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
			jscode: getTemplateJsCode(t, storeNameSpace),
		})

		vueData = append(vueData, vueComponentData{
			name:   t,
			jscode: getTemplateDataJsCode(t),
		})

		routes = append(routes, jsRoute{
			name:   t,
			jscode: getTemplateRouteJsCode(t),
		})

		content += "\nexport function " + t + "() {\n\n"

		//sort.Sort(ByCase(fields))
		content += strings.Join(fields, "")

		content += "\n    return this;\n}\n"
	}

	return
}

func getStoreJsCode(entity string) string {
	return AssignVar(
		AssignVar(storeTemplate, "{Entity}", entity),
		"{entity}", GetFirstLowerCase(entity))
}

func getTemplateJsCode(entity string, storeNameSpace string) string {

	template := ""

	if len(storeNameSpace) > 0 {
		template = AssignVar(usualEntityVueComponent, "{namespace}", "'"+storeNameSpace+"', ")
	} else {
		template = AssignVar(usualEntityVueComponent, "{namespace}", "")
	}

	return AssignVar(
		AssignVar(template, "{Entity}", entity),
		"{entity}", GetFirstLowerCase(entity))
}

func getTemplateDataJsCode(entity string) string {

	return AssignVar(
		AssignVar(usualEntityVueComponentData, "{Entity}", entity),
		"{entity}", GetFirstLowerCase(entity))
}

func getTemplateRouteJsCode(entity string) string {

	return AssignVar(
		AssignVar(usualEntityJsRoute, "{Entity}", entity),
		"{entity}", GetFirstLowerCase(entity))
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

	case "uuid.uuid":
		val = "null"
		break

	case "*uuid.uuid":
		val = "null"
		break

	case "time.time":
		val = "null"
		break

	case "*time.time":
		val = "null"
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
