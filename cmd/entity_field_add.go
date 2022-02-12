package cmd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gosha/mode"
	"gosha/settings"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/abiosoft/ishell/v2"
)

type ModelRepository struct {
	list []model
}

type model struct {
	FilePath   string
	ParsedFile *ast.File
	Structures []ast.Decl
}

type Field struct {
	Name     string
	Type     string
	Relation Relation
	Comment  string
}

type Relation struct {
	ModelName string
	FieldName string
}

type HttpMethods struct {
	IsFind           bool
	IsCreate         bool
	IsRead           bool
	IsUpdate         bool
	IsDelete         bool
	IsFindOrCreate   bool
	IsUpdateOrCreate bool
}

type HttpRoutes struct {
	Find           string
	Create         string
	Read           string
	Update         string
	Delete         string
	FindOrCreate   string
	UpdateOrCreate string
}

type Structures struct {
	WithoutDbModel bool
}

func (mr *ModelRepository) GetModels() (models []string) {

	for _, m := range mr.list {
		models = append(models, m.FilePath)
	}

	return
}

func (mr *ModelRepository) GetFieldRelation(field *ast.Field) (relation Relation) {

	//fmt.Println("structDecl", ft.Tag.Value)
	//fmt.Printf("%+v\n\n", ft.Tag.Value)
	//
	//st := reflect.StructField{}
	//st.Tag = reflect.StructTag(strings.Trim(ft.Tag.Value, "`"))
	//
	//f, ok := st.Tag.Lookup("gorm"); if ok {
	//    fmt.Printf("gorm is %+v\n\n", f)
	//}
	//
	//os.Exit(0)

	return
}

func (mr *ModelRepository) addField(modelName string, fieldName string, dataType string) (err error) {

	fmt.Println("adding new Field to model: ", modelName, "Field: ", fieldName, "type:", dataType)

	CamelCase := strings.Title(modelName)
	snakeCase := getLowerCase(modelName)
	firstLowerCase := GetFirstLowerCase(modelName)

	sourceFile := "./types/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + " " + dataType + "\n\t" + getRemoveLine(CamelCase)},
		nil)

	if isDate(dataType) {
		addImportIfNeed(sourceFile, "time")
	}
	if dataType == settings.DataTypeUuid {
		addImportIfNeed(sourceFile, "github.com/google/uuid")
	}

	sourceFile = "./dbmodels/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + " " + dataType + "\n\t" + getRemoveLine(CamelCase)},
		nil)

	if isDate(dataType) {
		addImportIfNeed(sourceFile, "time")
	}

	if dataType == settings.DataTypeUuid {
		addImportIfNeed(sourceFile, "github.com/google/uuid")
	}

	//sourceFile = "./logic/assigner.go"
	sourceFile = "./logic/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine("Assign" + CamelCase + "TypeFromDb.Field"), getRemoveLine("Assign" + CamelCase + "DbFromType.Field")},
		[]string{
			fieldName + ": db" + CamelCase + "." + fieldName + ",\n\t\t" + getRemoveLine("Assign"+CamelCase+"TypeFromDb.Field"),
			fieldName + ": typeModel." + fieldName + ",\n\t\t" + getRemoveLine("Assign"+CamelCase+"DbFromType.Field"),
		},
		nil)

	sourceFile = "./logic/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine("updateModel.Field")},
		[]string{"updateModel." + fieldName + " = newModel." + fieldName + "\n\t" + getRemoveLine("updateModel.Field")},
		nil)

	if _, err := os.Stat("./view"); os.IsNotExist(err) {
		fmt.Println("view folder not exists cant create bs4 template")
	} else {

		sourceFile = "./view/form/" + snakeCase + ".go"
		destinationFile := "./view/form/" + snakeCase + ".go"

		CreateFileIfNotExists(sourceFile, getEntityBs4vView(), nil)

		CopyFile(
			sourceFile,
			destinationFile,
			[]string{"{entity-name}", "{Entity}", "{entity}"},
			[]string{CamelCase, CamelCase, firstLowerCase},
			nil)

		sourceFile = "./view/form/" + snakeCase + ".go"
		CopyFile(
			sourceFile,
			sourceFile,
			[]string{"{entity-name}", "{Entity}", "{entity}"},
			[]string{CamelCase, CamelCase, firstLowerCase},
			nil)

		CopyFile(
			sourceFile,
			sourceFile,
			[]string{getRemoveLine(CamelCase)},
			[]string{GetFormTemplateField(modelName, fieldName, dataType) + "\n            " + getRemoveLine(CamelCase)},
			nil)

		CopyFile(
			sourceFile,
			sourceFile,
			[]string{getRemoveLine(CamelCase + "-collector")},
			[]string{GetFormFieldCollector(modelName, fieldName, dataType) + "\n                  " + getRemoveLine(CamelCase+"-collector")},
			nil)

		CopyFile(
			sourceFile,
			sourceFile,
			[]string{getRemoveLine(CamelCase + "-row-field")},
			[]string{GetRowFieldLine(modelName, fieldName, dataType) + "\n                  " + getRemoveLine(CamelCase+"-row-field")},
			nil)
	}

	//
	//CopyFile(
	//	sourceFile,
	//	sourceFile,
	//	[]string{getRemoveLine("Assign" + CamelCase + "TypeFromDb.Field"), getRemoveLine("Assign" + CamelCase + "DbFromType.Field")},
	//	[]string{
	//		fieldName + ": db" + CamelCase + "." + fieldName + ",\n\t\t" + getRemoveLine("Assign"+CamelCase+"TypeFromDb.Field"),
	//		fieldName + ": typeModel." + fieldName + ",\n\t\t" + getRemoveLine("Assign"+CamelCase+"DbFromType.Field"),
	//	},
	//	nil)

	CreateFileIfNotExists(usualTemplateGen.Path, usualTemplateGen.Content, nil)
	sourceFile = "./generator/" + snakeCase + ".go"

	if strings.ToLower(dataType) == settings.DataTypeString || strings.ToLower(dataType) == settings.DataTypeArrayString {
		addImportIfNeed(sourceFile, "strings")
	}

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + ": " + getGeneratorByDataType(dataType) + "\n\t\t" + getRemoveLine(CamelCase)},
		nil)

	// добавляем автофильтры для int полей содержащих id
	if strings.ToLower(dataType) == settings.DataTypeInt && (strings.Contains(fieldName, "Id") || strings.Contains(fieldName, "ID")) {
		err = mr.addFilter(modelName+"Filter", fieldName, dataType, "", true)
	}

	return
}

func isDate(dataType string) bool {

	switch dataType {
	case settings.DataTypeTimeLink,
		settings.DataTypeDuration,
		settings.DataTypeTime:
		return true
		break
	}

	return false

}

func GetFormTemplateField(modelName string, fieldName string, dataType string) string {

	firstLowerCase := GetFirstLowerCase(modelName)

	switch dataType {

	case settings.DataTypeString, settings.DataTypeUuid:
		return GetStringFormField(firstLowerCase, fieldName)
	case settings.DataTypeInt:
		return GetIntFormField(firstLowerCase, fieldName)
	case settings.DataTypeFloat64:
		return GetFloat64FormField(firstLowerCase, fieldName)
	case settings.DataTypeBool:
		return GetBoolFormField(firstLowerCase, fieldName)
	}

	return fmt.Sprintf("//unsupported data type %s for field %s", dataType, fieldName)
}

func GetFormFieldCollector(modelName string, fieldName string, dataType string) string {

	firstLowerCase := GetFirstLowerCase(modelName)

	switch dataType {

	case settings.DataTypeString, settings.DataTypeInt, settings.DataTypeFloat64, settings.DataTypeUuid:
		return GetStringFormFieldCollector(firstLowerCase, fieldName)
	case settings.DataTypeBool:
		return GetBoolFormFieldCollector(firstLowerCase, fieldName)
	}

	return fmt.Sprintf("//unsupported data type %s for fieldCollector %s", dataType, fieldName)
}

func GetRowFieldLine(modelName string, fieldName string, dataType string) string {

	firstLowerCase := GetFirstLowerCase(modelName)

	switch dataType {

	case settings.DataTypeInt:
		return fmt.Sprintf(
			"&bs4.Div{Text: strconv.Itoa(%s.%s), Classes: col, DataField: common.GetFieldName(&%s, &%s.%s)},", firstLowerCase, fieldName, firstLowerCase, firstLowerCase, fieldName)
	case settings.DataTypeFloat64:
		return fmt.Sprintf(
			"&bs4.Div{Text: fmt.Sprintf(\"%v\", %s.%s), Classes: col, DataField: common.GetFieldName(&%s, &%s.%s)},", firstLowerCase, fieldName, firstLowerCase, firstLowerCase, fieldName)
	case settings.DataTypeString, settings.DataTypeUuid:
		return fmt.Sprintf(
			"&bs4.Div{Text: %s.%s, Classes: col, DataField: common.GetFieldName(&%s, &%s.%s)},", firstLowerCase, fieldName, firstLowerCase, firstLowerCase, fieldName)
	case settings.DataTypeBool:
		return fmt.Sprintf(
			"&bs4.Div{Text: strconv.FormatBool(%s.%s), Classes: col, DataField: common.GetFieldName(&%s, &%s.%s)},", firstLowerCase, fieldName, firstLowerCase, firstLowerCase, fieldName)
	}

	return fmt.Sprintf("//unsupported datatype %s,", dataType)
}

func GetStringFormField(modelName string, fieldName string) string {

	return fmt.Sprintf(`view.GetStringFieldTemplate(view.FieldConfig{
				Id: common.GetFieldName(&%s, &%s.%s),
				Title: common.GetFieldName(&%s, &%s.%s),
				DefaultValueStr: defaultValue.%s,
			}),`, modelName, modelName, fieldName, modelName, modelName, fieldName, fieldName)

}

func GetIntFormField(modelName string, fieldName string) string {

	return fmt.Sprintf(`view.GetStringFieldTemplate(view.FieldConfig{
				Id: common.GetFieldName(&%s, &%s.%s),
				Title: common.GetFieldName(&%s, &%s.%s),
				DefaultValueInt: defaultValue.%s,
			}),`, modelName, modelName, fieldName, modelName, modelName, fieldName, fieldName)

}

func GetFloat64FormField(modelName string, fieldName string) string {

	return fmt.Sprintf(`view.GetStringFieldTemplate(view.FieldConfig{
				Id: common.GetFieldName(&%s, &%s.%s),
				Title: common.GetFieldName(&%s, &%s.%s),
				DefaultValueFloat64: defaultValue.%s,
			}),`, modelName, modelName, fieldName, modelName, modelName, fieldName, fieldName)

}

func GetBoolFormField(modelName string, fieldName string) string {

	return fmt.Sprintf(`view.GetBoolFieldTemplate(view.FieldConfig{
				Id: common.GetFieldName(&%s, &%s.%s),
				Title: common.GetFieldName(&%s, &%s.%s),
				IsChecked: defaultValue.%s,
			}),`, modelName, modelName, fieldName, modelName, modelName, fieldName, fieldName)

}

func GetStringFormFieldCollector(modelName string, fieldName string) string {

	return fmt.Sprintf(`{
					Id: common.GetFieldName(&%s, &%s.%s),
					CollectorKey: dataKey,
				},`, modelName, modelName, fieldName)

}

func GetBoolFormFieldCollector(modelName string, fieldName string) string {

	return fmt.Sprintf(`{
					Id: common.GetFieldName(&%s, &%s.%s),
					CollectorKey: dataKey,
				},`, modelName, modelName, fieldName)

}

func getGeneratorByDataType(dataType string) string {

	switch strings.ToLower(dataType) {
	case settings.DataTypeString:
		return "strings.Title(Babbler2.Babble()),"
	case settings.DataTypeInt:
		return "rand.Intn(100500),"
	case settings.DataTypeArrayInt:
		return "int{rand.Intn(100500), rand.Intn(100500), rand.Intn(100500)},"
	case settings.DataTypeArrayString:
		return "string{strings.Title(Babbler1.Babble()), Babbler2.Babble(), Babbler3.Babble()},"
	case settings.DataTypeFloat64:
		return "Float64n(4),"
	case settings.DataTypeBool:
		return "(rand.Intn(100500) % 2 > 0),"
	case strings.ToLower(settings.DataTypeTime), "time":
		return "randate(),"
	case strings.ToLower(settings.DataTypeDuration):
		return "randDuration(),"
	case settings.DataTypeIntLink:
		return "new(int),"
	case settings.DataTypeArrayBytes:
		return "[]byte,"
	default:
		return "nil,"
	}
}

func addImportIfNeed(file string, module string) {

	module = strings.Replace(module, "\"", "", -1)
	module = "\"" + module + "\""

	fset := token.NewFileSet()

	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	srcCode := string(input)

	f, err := parser.ParseFile(fset, "", srcCode, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := parser.ParseFile(fset, "", srcCode, parser.PackageClauseOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	imports := []string{}

	for _, s := range f.Imports {

		if s.Path.Value == module {
			return
		}

		imports = append(imports, "    "+s.Path.Value)
	}

	imports = append(imports, "    "+module)
	newImports := strings.Join(imports, "\n")

	newFileContent := "package " + p.Name.Name + "\n\nimport (\n" + string(input[1:f.Pos()]) + newImports + "\n)\n" + string(input[f.End():])

	err = ioutil.WriteFile(file, []byte(newFileContent), 0644)

	if err != nil {
		fmt.Println("Error adding import", file)
		fmt.Println(err)
	}
}

/**
 * Check structure exists in repo
 */
func (mr *ModelRepository) IsStructExists(modelName string) bool {

	for _, model := range mr.list {

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

				if structDecl.Name.String() == modelName {
					return true
				}
			}
		}
	}

	return false
}

func (mr *ModelRepository) ShowFields(modelName string) {

	for _, model := range mr.list {

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
				if structDecl.Name.String() == modelName {

					structDecl := md.(*ast.TypeSpec).Type.(*ast.StructType)

					for _, field := range structDecl.Fields.List {

						if len(field.Names) > 0 {
							fmt.Println(field.Names)
						}
					}
				}
			}
		}
	}

	return
}

func (mr *ModelRepository) GetFields(modelName string, fields []Field) []Field {

	for _, model := range mr.list {

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
				if structDecl.Name.String() == modelName {

					if structDecl, ok := md.(*ast.TypeSpec).Type.(*ast.StructType); ok {

						for _, ft := range structDecl.Fields.List {

							if len(ft.Names) > 0 {

								typeString := ""

								if ident, ok := ft.Type.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error
									typeString = ident.Name
								} else if _, ok := ft.Type.(*ast.ArrayType); ok { // allow subsequent panic to provide a more descriptive error
									typeString = "array"
								} else {

									switch x := ft.Type.(type) {
									case *ast.BasicLit:
										typeString = x.Value
									case *ast.Ident:
										typeString = x.Name
									case *ast.StarExpr:
										{
											//typeString = x.Star
											if ident, ok := x.X.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error
												typeString = "*" + ident.Name
											} else {
												// Обработка указателей на сложные типы
												x, ok := x.X.(*ast.SelectorExpr)
												if !ok {
													continue
												}
												if ident, ok := x.X.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error
													typeString = ident.Name + "." + x.Sel.Name
												}
											}
										}
									case *ast.SelectorExpr:
										{
											if ident, ok := x.X.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error
												typeString = ident.Name + "." + x.Sel.Name
											}
										}
									default:
										typeString = "unknown"
									}
								}

								fields = append(fields, Field{
									Type:     strings.Title(typeString),
									Name:     ft.Names[0].Name,
									Relation: mr.GetFieldRelation(ft),
									Comment:  strings.TrimSpace(ft.Doc.Text()),
								})
							} else {

								if ident, ok := ft.Type.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error

									fields = mr.GetFields(ident.Name, fields)
								}
							}
						}

					}
				}
			}
		}
	}

	return fields
}

func (mr *ModelRepository) IsFieldExists(modelName string, field string) bool {

	for _, model := range mr.list {

		for _, spec := range model.Structures {

			tp, ok := spec.(*ast.GenDecl)

			if !ok {
				continue
			}

			for _, md := range tp.Specs {

				structDecl, ok := md.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if structDecl.Name.String() == modelName {

					if structType, ok := md.(*ast.TypeSpec).Type.(*ast.StructType); ok {

						includedStructs := []string{}

						for _, ft := range structType.Fields.List {

							if len(ft.Names) < 1 {
								includedStructs = append(includedStructs, fmt.Sprintf("%+v", ft.Type))
								continue
							}

							if ft.Names[0].Name == field {
								return true
							}
						}

						for _, includeStruct := range includedStructs {
							hasField := mr.IsFieldExists(includeStruct, field)
							if hasField {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

func (mr *ModelRepository) IsFilter(modelName string) bool {

	for _, model := range mr.list {

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
				if structDecl.Name.String() == modelName {

					if structDecl, ok := md.(*ast.TypeSpec).Type.(*ast.StructType); ok {

						for _, ft := range structDecl.Fields.List {

							if fmt.Sprintf("%+v", ft.Type) == "AbstractFilter" {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

func (mr *ModelRepository) GetModelComment(modelName string) string {

	for _, model := range mr.list {

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
				if structDecl.Name.String() == modelName {
					return strings.TrimSpace(tp.Doc.Text())
				}
			}
		}
	}
	return ""
}

func GetFieldRelation(repository *ModelRepository, field *ast.Field) (relation Relation) {

	//fmt.Println("structDecl", ft.Tag.Value)
	//fmt.Printf("%+v\n\n", ft.Tag.Value)
	//
	//st := reflect.StructField{}
	//st.Tag = reflect.StructTag(strings.Trim(ft.Tag.Value, "`"))
	//
	//f, ok := st.Tag.Lookup("gorm"); if ok {
	//    fmt.Printf("gorm is %+v\n\n", f)
	//}
	//
	//os.Exit(0)

	return
}

func (em *ModelRepository) getModelFile(name string) (fileName string) {

	for _, m := range em.list {
		if m.FilePath == name {
			fileName = m.FilePath
			break
		}
	}

	return
}

func entityFieldAdd(c *ishell.Context) {

	var field string
	var entity string
	var err error
	var reg RegularFind

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	InteractiveEcho([]string{
		green("Hello we start adding new Field to entity"),
	})

	existsModels := GetExistsModels()

	dbmodelsList := GetModelsList(existsModels)

	if mode.IsInteractive() {

		for _, m := range dbmodelsList {

			shell.AddCmd(&ishell.Cmd{
				Name: m,
				Help: m + " entity",
				Func: func(c *ishell.Context) {
					fmt.Println("select predefined function", m)
				},
			})
		}

		if len(dbmodelsList) > 0 {
			c.Println(yellow("Found some entities:"))
			c.Println(dbmodelsList)
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
		"Please enter name of new Field:",
	})

	if mode.IsNonInteractive() {

		reg, err = GetOsArgument("Field")
		field = reg.StringResult

	} else {

		field, err = getName(c, false, "Field")
	}

	if err != nil {
		return
	}

	dataType, err := getDataType(c, settings.SupportedModelFieldDataTypes)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	existsModels.addField(entity, field, dataType)

	defer clearEntities(existsModels)
}

func getDataType(c *ishell.Context, dataTypes []string) (dataType string, err error) {

	var choice int
	var exists bool
	var reg RegularFind

	if mode.IsInteractive() {

		choice = c.MultiChoice(dataTypes, "Please select data type")

	} else {

		reg, err = GetOsArgument("data-type")

		if err != nil {
			return "", err
		}

		exists, choice = InArray(reg.StringResult, dataTypes)

		if !exists {
			return "", errors.New("unsupported type " + reg.StringResult + " need: " + strings.Join(dataTypes, ","))
		}
	}

	if choice > -1 && choice < len(dataTypes) {
		return dataTypes[choice], nil
	}

	return "", errors.New("cancel")
}

func GetModelsMethods(repo ModelRepository) map[string]HttpMethods {

	list := make(map[string]HttpMethods)

	reg := regexp.MustCompile(`(//){0,}\s{0,}router\.HandleFunc\(settings\.([A-Z]{1}[A-Za-z0-9]+)+Route.*,\s{0,}webapp\.[A-Za-z0-9]+\).Methods`)

	path := "./router/router.go"
	_, e := os.Stat(path)

	if e != nil {
		return list
	}

	src, _ := ioutil.ReadFile(path)
	methods := reg.FindAllSubmatch(src, -1)

	for _, mt := range methods {

		if len(mt) == 3 && string(mt[1]) != "//" {

			row := strings.TrimSpace(string(mt[0]))
			modelName := string(mt[2])

			if _, ok := list[modelName]; !ok {
				list[modelName] = HttpMethods{}
			}

			mp := list[modelName]

			regMethod := regexp.MustCompile("webapp\\." + modelName + "([a-zA-Z]+)")
			hm := regMethod.FindAllStringSubmatch(row, -1)

			if len(hm) < 1 {
				continue
			}

			if len(hm[0]) < 2 {
				continue
			}

			switch hm[0][1] {
			case "Find":
				mp.IsFind = true
				break
			case "Read":
				mp.IsRead = true
				break
			case "Create":
				mp.IsCreate = true
				break
			case "Update":
				mp.IsUpdate = true
				break
			case "Delete":
				mp.IsDelete = true
				break
			case "FindOrCreate":
				mp.IsFindOrCreate = true
				break
			case "UpdateOrCreate":
				mp.IsUpdateOrCreate = true
				break
			}

			list[modelName] = mp
		}
	}

	return list
}

func GetModelsList(repository ModelRepository) (list []string) {

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
				list = append(list, structDecl.Name.String())
			}
		}
	}

	return
}

func GetExistsModels() (repo ModelRepository) {
	return getPackageModels("/dbmodels")
}

func GetExistsTypes() (repo ModelRepository) {
	return getPackageModels("/types")
}

func getPackageModels(path string) (repo ModelRepository) {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	workFolder := dir + path

	files, err := ioutil.ReadDir(workFolder)
	if err != nil {
		return
	}

	validFileName := regexp.MustCompile(`^[a-zA-Z0-9_]+\.go$`)

	for _, f := range files {

		isValidFile := validFileName.MatchString(f.Name())

		if f.IsDir() || !isValidFile {
			fmt.Println("invalid file " + f.Name())
			continue
		}

		filePath := workFolder + "/" + f.Name()

		models, parsedFile := getFileModels(filePath)

		repo.list = append(repo.list, model{
			FilePath:   filePath,
			ParsedFile: parsedFile,
			Structures: models,
		})
	}

	return
}

func getFileModels(filePath string) ([]ast.Decl, *ast.File) {

	b, err := ioutil.ReadFile(filePath) // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	src := string(b) // convert content to a 'string'

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	// hard coding looking these up
	typeDecl := f.Decls

	return typeDecl, f
}

func clearEntities(repo ModelRepository) {

	for _, entity := range GetModelsList(repo) {
		shell.DeleteCmd(entity)
	}
}
