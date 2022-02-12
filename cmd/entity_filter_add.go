package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"go/ast"
	"github.com/abiosoft/ishell/v2"
	"gosha/mode"
	"gosha/settings"
	"os"
	"regexp"
	"strings"
)

func GetExistsFilters() (repo ModelRepository) {
	return getPackageModels("/types")
}

func entityFilterdAdd(c *ishell.Context) {

	fmt.Println("add filter")

	var filter string
	var sqlField string
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

	if mode.IsNonInteractive() {

		regSql, _ := GetOsArgument("sql-field")
		sqlField = regSql.StringResult

	} else {
		sqlField, _ = getName(c, false, "sql-field")
	}

	if err != nil {
		return
	}

	dataType, err := getDataType(c, settings.SupportedFilterDataTypes)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	existsModels.addFilter(entity, filter, dataType, sqlField, false)

	defer clearEntities(existsModels)
}

func (mr *ModelRepository) addFilter(modelName string, fieldName string, dataType string, specialSqlField string, useAbstractIfExists bool) (err error) {

	fmt.Println("adding new filter to model: ", modelName, "filter: ", fieldName, "type:", dataType, "specialSqlField", specialSqlField, "useAbstractIfExists", useAbstractIfExists)

	CamelCase := strings.Title(modelName)
	snakeCase := getLowerCase(modelName)

	base_snake_go := strings.Replace(snakeCase+".go", "_filter.go", ".go", 1)

	BaseCamel := ""

	re := regexp.MustCompile("FilterFilter$")
	if re.Match([]byte(CamelCase)) {
		BaseCamel = strings.Replace(CamelCase, "FilterFilter", "Filter", 1)
	} else {
		ref := regexp.MustCompile("(Filter)$")
		BaseCamel = string(ref.ReplaceAll([]byte(CamelCase), []byte("")))
	}

	sourceFile := "./types/" + base_snake_go

	// в случае если useAbstractIfExists, проверяем наличие поля в абстрактном фильтре,
	// и если оно там есть не добавляем поле

	isAddToFilterStruct := true
	if useAbstractIfExists {

		types := GetExistsTypes()
		existsInAbstract := types.IsFieldExists("AbstractFilter", fieldName)

		isAddToFilterStruct = !existsInAbstract
	}

	if isAddToFilterStruct {

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
	}

	filterCode := ""

	switch dataType {
	case settings.DataTypeArrayInt, settings.DataTypeArrayBytes, settings.DataTypeArrayString, settings.DataTypeString, settings.DataTypeUuid:

		fieldSnakeCase := getLowerCase(fieldName)
		preComment := "//"

		if len(specialSqlField) > 0 {
			preComment = ""
			fieldSnakeCase = specialSqlField
		}

		filterCode = `if len(filter.` + fieldName + `) > 0 {
        ` + preComment + `criteria = criteria.Where("` + fieldSnakeCase + ` in (?)", filter.` + fieldName + `)
    }`

		break
	case settings.DataTypeInt, settings.DataTypeFloat64:

		filterCode = `if filter.` + fieldName + ` > 0 {
        criteria = criteria.Where(dbmodels.` + BaseCamel + `{` + fieldName + `: filter.` + fieldName + `})
    }`

		break

	}

	if len(filterCode) > 0 {

		path := "./logic/" + base_snake_go
		line := getRemoveLine(BaseCamel + ".FindFilterCode")

		CopyFile(
			path,
			path,
			[]string{line},
			[]string{filterCode + "\n\t" + line},
			nil)
	}

	if isAddToFilterStruct {

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
			break
		case settings.DataTypeArrayString:
			addImportIfNeed(sourceFile, "net/url")
			break
		}
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

				//isFilter, _ := regexp.Match("Filter", []byte(name))
				isFilter := repository.IsFilter(name)

				if isFilter {
					list = append(list, name)
				}
			}
		}
	}

	return
}
