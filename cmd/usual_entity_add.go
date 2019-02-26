package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "strings"
)

func usualEntityAdd(c *ishell.Context) {

    yellow := color.New(color.FgYellow).SprintFunc()

    c.Println(yellow("Hello we start creating new entity and api"))

    entity, err := getName(c, false)

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
        []string{usualTemplateRouteEntity.Content, CamelCase, firstLowerCase},
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

    CreateFile(sourceFile, usualTemplateWebappEntity.Content, c)

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

    //
    //CopyFile(
    //    sourceFile,
    //    destinationFile2,
    //    []string{"NewEntity", "newEntity", "QueueNameCore"},
    //    []string{camelCase, FirstLowerCase, QueueConstant},
    //    c)
    //
    //if isNew {
    //    AppendFile("./cnf/queue.go", "\n// generated code\nconst " + QueueConstant + " = \"" + strings.ToLower(camelCase) + "\"")
    //}
    //
    //c.Println("Created files:" )
    //
    //c.Println(destinationFile1)
    //c.Println(destinationFile2)

    c.Println("Конгретулэйшенс плеер уан, Ю уин. New entity " + CamelCase + " was created. Bye" )

}
