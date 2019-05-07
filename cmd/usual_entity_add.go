package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "strings"
)

func usualEntityAdd(c *ishell.Context) {

    yellow := color.New(color.FgYellow).SprintFunc()
    c.Println(yellow("Start creating new entity and api"))

    entity, err := getEntity(c)

    if err !=nil {
        return
    }

    CamelCase := strings.Title(entity)
    snakeCase := getLowerCase(entity)
    firstLowerCase := getFirstLowerCase(entity)

    sourceFile := "./router/router.go"
    destinationFile := "./router/router.go"

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
        []string{usualTemplateLogicAssignerEntity.Content, CamelCase, firstLowerCase},
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

    CreateFile(sourceFile, usualTemplateWebappEntityType.Content, c)

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

    CreateFile(sourceFile, usualTemplateEntityLogic.Content, c)

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

    c.Println("New entity " + CamelCase + " was created" )
}

func getWebAppContent() (webappContent string) {

    webappContent = usualTemplateWebappEntity.Content
    crudArgs, _ := GetOsArgument("check-auth")

    if len(crudArgs.StringResult) > 0 {

        crudParams := Crud{}
        crudParams.IsFind = strings.Contains(crudArgs.StringResult, "f")
        crudParams.IsCreate = strings.Contains(crudArgs.StringResult, "c")
        crudParams.IsRead = strings.Contains(crudArgs.StringResult, "r")
        crudParams.IsUpdate = strings.Contains(crudArgs.StringResult, "u")
        crudParams.IsDelete = strings.Contains(crudArgs.StringResult, "d")

        webappContent = assignMsName(GetUsualTemplateWebAppContent(crudParams))
    }
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

        routeContent = GetUsualTemplateRouteEntity(crudParams)
    }

    return routeContent
}
