package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"strings"
)

func genTestsWebapp(c *ishell.Context) {

	existsTypes := GetExistsModels()
	names := GetModelsList(existsTypes)

	for _, entity := range names {

		switch entity {
		case "Auth", "User":
			continue
		}

		CamelCase := strings.Title(entity)
		snakeCase := getLowerCase(entity)
		firstLowerCase := GetFirstLowerCase(entity)

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
