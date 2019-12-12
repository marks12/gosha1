package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"strings"
)

func genTestsWebapp(c *ishell.Context) {

	existsTypes := GetExistsModels()
	names := GetModelsList(existsTypes)

	for _, entity := range names {

		CamelCase := strings.Title(entity)
		snakeCase := getLowerCase(entity)
		firstLowerCase := getFirstLowerCase(entity)

		sourceFile := "./webapp/" + snakeCase + "_test.go"
		destinationFile := "./webapp/" + snakeCase + "_test.go"

		CreateFile(sourceFile, getWebAppTestContent(), c)

		CopyFile(
			sourceFile,
			destinationFile,
			[]string{"{entity-name}", "{Entity}", "{entity}"},
			[]string{CamelCase, CamelCase, firstLowerCase},
			c)
	}

}
