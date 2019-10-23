package cmd

import (
	"github.com/fatih/color"
	"gopkg.in/abiosoft/ishell.v2"
	"strings"
)

func usualEntityAdd(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()
	c.Println(yellow("Start creating new entity and api"))

	entity, err := getEntity(c)

	if err != nil {
		return
	}

	CamelCase := strings.Title(entity)
	snakeCase := getLowerCase(entity)
	firstLowerCase := getFirstLowerCase(entity)

	sourceFile := "./settings/routes.go"
	destinationFile := "./settings/routes.go"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"\n// route-constant-generator here dont touch this line", "{Entity}", "{entity}"},
		[]string{usualTemplateSettingsRoutesConstEntity.Content, CamelCase, firstLowerCase},
		c)

	sourceFile = "./settings/routes.go"
	destinationFile = "./settings/routes.go"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"\n    // router-list-generator here dont touch this line", "{Entity}", "{entity}"},
		[]string{usualTemplateSettingsRoutesListEntity.Content, CamelCase, firstLowerCase},
		c)

	sourceFile = "./router/router.go"
	destinationFile = "./router/router.go"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"router-generator here dont touch this line", "{Entity}", "{entity}"},
		[]string{getRouteContent(), CamelCase, firstLowerCase},
		c)

	sourceFile = "./logic/assigner.go"
	destinationFile = "./logic/assigner.go"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"// add all assign functions", "{Entity}", "{entity}"},
		[]string{getAssignContent(), CamelCase, firstLowerCase},
		c)

	sourceFile = "./webapp/" + snakeCase + ".go"
	destinationFile = "./webapp/" + snakeCase + ".go"

	CreateFile(sourceFile, getWebAppContent(), c)

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"{entity-name}", "{Entity}", "{entity}"},
		[]string{CamelCase, CamelCase, firstLowerCase},
		c)

	sourceFile = "./types/" + snakeCase + ".go"
	destinationFile = "./types/" + snakeCase + ".go"

	CreateFile(sourceFile, getTypeContent(), c)

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"{entity-name}", "{Entity}", "{entity}"},
		[]string{CamelCase, CamelCase, firstLowerCase},
		c)

	sourceFile = "./dbmodels/" + snakeCase + ".go"
	destinationFile = "./dbmodels/" + snakeCase + ".go"

	CreateFile(sourceFile, usualTemplateWebappEntityDbModels.Content, c)

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"{entity-name}", "{Entity}", "{entity}"},
		[]string{CamelCase, CamelCase, firstLowerCase},
		c)

	sourceFile = "./logic/" + snakeCase + ".go"
	destinationFile = "./logic/" + snakeCase + ".go"

	CreateFile(sourceFile, getLogicContent(), c)

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"{entity-name}", "{Entity}", "{entity}"},
		[]string{CamelCase, CamelCase, firstLowerCase},
		c)

	sourceFile = "./bootstrap/insert_data_to_db.go"
	destinationFile = "./bootstrap/insert_data_to_db.go"

	replaceFrom := `//generator insert entity`
	replaceTo := `//generator insert entity
          ` + "&dbmodels." + CamelCase + "{},"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{replaceFrom},
		[]string{replaceTo},
		c)

	c.Println("New entity " + CamelCase + " was created")
}

func getLogicContent() (c string) {

	// DISABLED crudParams. Just comment in ROUTES for routes not in logic

	//crudArgs, _ := GetOsArgument("crud")
	crudParams := Crud{}

	//if len(crudArgs.StringResult) > 0 {
	//
	//	crudParams.IsFind = strings.Contains(crudArgs.StringResult, "f")
	//	crudParams.IsCreate = strings.Contains(crudArgs.StringResult, "c")
	//	crudParams.IsRead = strings.Contains(crudArgs.StringResult, "r")
	//	crudParams.IsUpdate = strings.Contains(crudArgs.StringResult, "u")
	//	crudParams.IsDelete = strings.Contains(crudArgs.StringResult, "d")
	//	crudParams.IsFindOrCreate = strings.Contains(crudArgs.StringResult, "a")
	//} else {
		crudParams = Crud{true, true, true, true, true, true}
	//}

	c = GetUsualTemplateLogicContent(crudParams)
	return
}

func getAssignContent() (c string) {

	cfg := TypeConfig{}

	hasId, _ := GetOsArgument("no-id")
	cfg.IsId = !hasId.BoolResult

	c = GetUsualTemplateAssignContent(cfg)
	return
}

func getTypeContent() (c string) {

	cfg := TypeConfig{}

	hasId, _ := GetOsArgument("no-id")
	cfg.IsId = !hasId.BoolResult

	c = GetUsualTemplateTypeContent(cfg)
	return
}

func getWebAppContent() (webappContent string) {

	webappContent = usualTemplateWebappEntity.Content

	AuthcrudArgs, _ := GetOsArgument("check-auth")
	authParams := Crud{true, true, true, true, true, true}

	if len(AuthcrudArgs.StringResult) > 0 {

		authParams.IsFind = strings.Contains(AuthcrudArgs.StringResult, "f")
		authParams.IsCreate = strings.Contains(AuthcrudArgs.StringResult, "c")
		authParams.IsRead = strings.Contains(AuthcrudArgs.StringResult, "r")
		authParams.IsUpdate = strings.Contains(AuthcrudArgs.StringResult, "u")
		authParams.IsDelete = strings.Contains(AuthcrudArgs.StringResult, "d")
		authParams.IsFindOrCreate = strings.Contains(AuthcrudArgs.StringResult, "a")

	}

	// DISABLED crudParams. Just comment in ROUTES for routes not in logic

	//methodCrudArgs, _ := GetOsArgument("crud")
	methodCrudParams := Crud{true, true, true, true, true, true}
	//
	//if len(methodCrudArgs.StringResult) > 0 {
	//
	//	methodCrudParams.IsFind = strings.Contains(methodCrudArgs.StringResult, "f")
	//	methodCrudParams.IsCreate = strings.Contains(methodCrudArgs.StringResult, "c")
	//	methodCrudParams.IsRead = strings.Contains(methodCrudArgs.StringResult, "r")
	//	methodCrudParams.IsUpdate = strings.Contains(methodCrudArgs.StringResult, "u")
	//	methodCrudParams.IsDelete = strings.Contains(methodCrudArgs.StringResult, "d")
	//	methodCrudParams.IsFindOrCreate = strings.Contains(methodCrudArgs.StringResult, "a")
	//
	//}

	webappContent = assignMsName(GetUsualTemplateWebAppContent(authParams, methodCrudParams))

	return
}

func getEntity(c *ishell.Context) (entity string, err error) {

	var arguments RegularFind

	arguments, err = GetOsArgument("entity")

	if len(arguments.StringResult) < 1 || err != nil {
		return getName(c, false, "Entity")
	}

	entity = arguments.StringResult

	return
}

func getRouteContent() string {

	routeContent := usualTemplateRouteEntity.Content

	crudArgs, _ := GetOsArgument("crud")

	if len(crudArgs.StringResult) > 0 {

		crudParams := Crud{}
		crudParams.IsFind = strings.Contains(crudArgs.StringResult, "f")
		crudParams.IsCreate = strings.Contains(crudArgs.StringResult, "c")
		crudParams.IsRead = strings.Contains(crudArgs.StringResult, "r")
		crudParams.IsUpdate = strings.Contains(crudArgs.StringResult, "u")
		crudParams.IsDelete = strings.Contains(crudArgs.StringResult, "d")
		crudParams.IsFindOrCreate = strings.Contains(crudArgs.StringResult, "a")

		routeContent = GetUsualTemplateRouteEntity(crudParams)
	}

	return routeContent
}
