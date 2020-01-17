package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/settings"
	"strings"
)

func genGenerators(c *ishell.Context) {

	existsTypes := GetExistsTypes()
	typeNames := GetModelsList(existsTypes)

	CreateFileIfNotExists(usualTemplateGen.Path, usualTemplateGen.Content, nil)

	for _, modelName := range typeNames {

		CamelCase := strings.Title(modelName)
		snakeCase := getLowerCase(modelName)
		firstLowerCase := getFirstLowerCase(modelName)

		if strings.Replace(snakeCase + ".go", "_filter.go", "", 1) != snakeCase + ".go" {
			continue
		}

		if isInvalid, _ := InArray(CamelCase, []string{"APIError", "APIStatus", "Authenticator", "FilterIds", "GoshaOrderFilter", "GoshaSearchFilter", "Pagination", "Validator"}); isInvalid {
			continue
		}

		sourceFile := "./generator/" + snakeCase + ".go"
		destinationFile := "./generator/" + snakeCase + ".go"
		CreateFile(sourceFile, getEntityGenContent(), c)
		CopyFile(
			sourceFile,
			destinationFile,
			[]string{"{entity-name}", "{Entity}", "{entity}"},
			[]string{CamelCase, CamelCase, firstLowerCase},
			c)

		fields := existsTypes.GetFields(modelName, []Field{})

		for _, field := range fields {

			if isInvalid, _ := InArray(field.Name, []string{"Id", "validationErrors"}); ! isInvalid {
				addGenerator(field.Type, modelName, field.Name)
			}
		}
	}
}

func addGenerator(dataType, modelName, fieldName string) {

	CamelCase := strings.Title(modelName)
	snakeCase := getLowerCase(modelName)

	sourceFile := "./generator/" + snakeCase + ".go"
	if strings.ToLower(dataType) == settings.DataTypeString {
		addImportIfNeed(sourceFile, "strings")
	}

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + ": " + getGeneratorByDataType(dataType) + "\n\t\t" + getRemoveLine(CamelCase)},
		nil)
}
