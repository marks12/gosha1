package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "gosha/mode"
    "fmt"
    "os"
    "strings"
    "gosha/settings"
    "go/ast"
    "regexp"
)

func GetExistsFilters() (repo ModelRepository) {
    return getPackageModels("/types")
}

func entityFilterdAdd(c *ishell.Context) {

    var filter string
    var entity string
    var err error
    var reg RegularFind

    green := color.New(color.FgGreen).SprintFunc()
    yellow := color.New(color.FgYellow).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    InteractiveEcho([]string{
        green("Hello we start adding new filter to entity"),
    })

    existsModels := GetExistsFilters()

    filtersList := GetFiltersList(existsModels)

    if mode.IsInteractive() {

        for _, m := range filtersList {

            shell.AddCmd(&ishell.Cmd{
                Name: m,
                Help: m + " entity",
                Func: func(c *ishell.Context) {
                    fmt.Println("select predefined function", m)
                },
            })
        }

        if len(filtersList) > 0 {
            c.Println(yellow("Found some entities:"))
            c.Println(filtersList)
        }

        entity, err = getName(c, true, "Entity")

    } else {

        reg, err = GetOsArgument("entity")
        entity = reg.StringResult
    }

    if err != nil {
        return
    }

    InteractiveEcho([]string{
        green("select model: ", entity),
    })

    if !existsModels.IsStructExists(entity) {

        c.Println(red("You select entity: '", entity, "', but this struct is not exists"))

        if mode.IsInteractive() {
            os.Exit(1)
        }
    }

    if mode.IsInteractive() {
        existsModels.ShowFields(entity)
    }

    InteractiveEcho([]string{
        "Please enter name of new filter:",
    })

    if mode.IsNonInteractive() {

        reg, err = GetOsArgument("filter")
        filter = reg.StringResult

    } else {

        filter, err = getName(c, false, "filter")
    }

    if err != nil {
        return
    }

    dataType, err := getDataType(c, settings.SupportedFilterDataTypes)

    if err != nil {

        fmt.Println(err.Error())
        return
    }

    existsModels.addFilter(entity, filter, dataType)

    defer clearEntities(existsModels)
}

func (mr *ModelRepository) addFilter(modelName string, fieldName string, dataType string) (err error) {

    fmt.Println("adding new filter to model: ", modelName, "filter: ", fieldName, "type:", dataType)

    CamelCase := strings.Title(modelName)
    snakeCase := getLowerCase(modelName)

    base_snake := strings.Replace(snakeCase, "_filter", "", 1)

    sourceFile := "./types/" + base_snake + ".go"

    CopyFile(
        sourceFile,
        sourceFile,
        []string{getRemoveLine(CamelCase)},
        []string{fieldName + " " + dataType + "\n\t" + getRemoveLine(CamelCase)},
        nil)

    CopyFile(
        sourceFile,
        sourceFile,
        []string{getRemoveLine("Get" + CamelCase)},
        []string{getFilterGetter(fieldName, dataType) + "\n\t" + getRemoveLine("Get"+CamelCase)},
        nil)

    switch dataType {

        case settings.DataTypeTimeLink, settings.DataTypeTime:
            addImportIfNeed(sourceFile, "time")
            break

        case settings.DataTypeFloat64, settings.DataTypeInt:
            addImportIfNeed(sourceFile, "strconv")
            break

        case settings.DataTypeBool:
            addImportIfNeed(sourceFile, "strconv")
            break

        case settings.DataTypeUuid:
            addImportIfNeed(sourceFile, "github.com/google/uuid")
            break

        case settings.DataTypeArrayInt:
            addImportIfNeed(sourceFile, "net/url")
            addImportIfNeed(sourceFile, "strconv")
            addImportIfNeed(sourceFile, "time")
            break
        case settings.DataTypeArrayString:
            addImportIfNeed(sourceFile, "net/url")
            break
    }

    return
}

func getFilterGetter(field, dataType string) (res string) {

    switch dataType {

    case settings.DataTypeString:
        return fmt.Sprintf("filter.%s = request.FormValue(\"%s\")", field, field)
        break

    case settings.DataTypeFloat64:
        return fmt.Sprintf("filter.%s, _ = strconv.ParseFloat(request.FormValue(\"%s\"), 64)", field, field)
        break

    case settings.DataTypeInt:
        return fmt.Sprintf("filter.%s, _ = strconv.Atoi(request.FormValue(\"%s\"))", field, field)
        break

    case settings.DataTypeBool:
        return fmt.Sprintf("filter.%s, _ = strconv.ParseBool(request.FormValue(\"%s\"))", field, field)
        break

    case settings.DataTypeUuid:
        return fmt.Sprintf("filter.%s, _ = uuid.Parse(request.FormValue(\"%s\"))", field, field)
        break

    case settings.DataTypeTime:
        return fmt.Sprintf("filter.%s, _ = time.Parse(time.RFC3339, request.FormValue(\"%s\"))", field, field)
        break

    case settings.DataTypeArrayInt:

    return fmt.Sprintf(`
    arr%s, _ := url.ParseQuery(request.URL.RawQuery)
    for _, f := range arr%s["%s[]"] {
        v, _ := strconv.Atoi(f)
        filter.%s = append(filter.%s, v)
    }
    `, field, field, field, field, field)

        break

    case settings.DataTypeArrayString:

    return fmt.Sprintf(`
    arr%s, _ := url.ParseQuery(request.URL.RawQuery)
    for _, f := range arr%s["%s[]"] {
        filter.%s = append(filter.%s, f)
    }`, field, field, field, field, field)

        break

    default:

            break

    }

    return
}

func GetFiltersList(repository ModelRepository) (list []string) {

    for _, model := range repository.list {

        for _, spec := range model.Structures {

            tp, ok := spec.(*ast.GenDecl)

            if !ok {
                continue
            }

            for _, md := range tp.Specs {

                _, ok := md.(*ast.TypeSpec)
                if !ok {
                    continue
                }

                structDecl := md.(*ast.TypeSpec)

                name := structDecl.Name.String()

                isFilter, _ := regexp.Match("Filter", []byte(name))

                if isFilter {
                    list = append(list, name)
                }
            }
        }
    }

    return
}
