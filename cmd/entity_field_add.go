package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/mode"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"gosha/settings"
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
}

type Relation struct {
	ModelName string
	FieldName string
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

	sourceFile := "./types/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + " " + dataType + "\n\t" + getRemoveLine(CamelCase)},
		nil)

	sourceFile = "./dbmodels/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		sourceFile,
		[]string{getRemoveLine(CamelCase)},
		[]string{fieldName + " " + dataType + "\n\t" + getRemoveLine(CamelCase)},
		nil)

	sourceFile = "./logic/assigner.go"

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

	return
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

					structDecl := md.(*ast.TypeSpec).Type.(*ast.StructType)

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
										}
									}
								case *ast.SelectorExpr:
									{
										if ident, ok := x.X.(*ast.Ident); ok { // allow subsequent panic to provide a more descriptive error
											typeString = ident.Name
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

	return fields
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
		os.Exit(1)
	}

	existsModels.ShowFields(entity)

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

	dataType, err := getDataType(c)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	existsModels.addField(entity, field, dataType)

	defer clearEntities(existsModels)
}

func getDataType(c *ishell.Context) (dataType string, err error) {

	var choice int
	var exists bool
	var reg RegularFind

	dataTypes := settings.SupportedFieldDataTypes

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
	f, err := parser.ParseFile(fset, "", src, 0)

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
