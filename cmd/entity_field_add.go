package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "fmt"
    "io/ioutil"
    "os"
    "go/parser"
    "go/token"
    "go/ast"
    "regexp"
    "go/printer"
    "bufio"
)

type modelRepository struct {
    list []model
}

type model struct {
    FilePath string
    ParsedFile *ast.File
    Structures []ast.Decl
}

func (mr *modelRepository) getModels() (models []string) {

    for _, m := range mr.list {
        models = append(models, m.FilePath)
    }

    return
}

func (mr *modelRepository) addField(modelName string, fieldName string, dataType string) (err error) {

    for _, model := range mr.list {

        for _, spec := range model.Structures {

            tp, ok := spec.(*ast.GenDecl)

            if !ok {
                continue
            }

            for _,md := range tp.Specs {

                _, ok := md.(*ast.TypeSpec)
                if !ok {
                    continue
                }

                structDecl := md.(*ast.TypeSpec)
                if structDecl.Name.String() == modelName {

                    fset := token.NewFileSet()

                    structDecl := md.(*ast.TypeSpec).Type.(*ast.StructType)

                    field := &ast.Field{
                        Doc:   nil,
                        Names: []*ast.Ident{ast.NewIdent(fieldName)},
                        Type:  ast.NewIdent(dataType),
                    }

                    structDecl.Fields.List = append(structDecl.Fields.List, field)

                    f, _ := os.OpenFile(model.FilePath, os.O_WRONLY, 0666)
                    w := bufio.NewWriter(f)

                    printer.Fprint(w, fset, model.ParsedFile)
                    w.Flush()

                    fmt.Println("Field " + fieldName + " was added to model " + modelName)
                }
            }
        }
    }

    return
}

func (mr *modelRepository) ShowFields(modelName string) {

    for _, model := range mr.list {

        for _, spec := range model.Structures {

            tp, ok := spec.(*ast.GenDecl)

            if !ok {
                continue
            }

            for _,md := range tp.Specs {

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

func (em *modelRepository) getModelFile(name string) (fileName string) {

    for _, m := range em.list {
        if m.FilePath == name {
            fileName = m.FilePath
            break
        }
    }

    return
}

func entityFieldAdd(c *ishell.Context) {

    green := color.New(color.FgGreen).SprintFunc()
    yellow := color.New(color.FgYellow).SprintFunc()

    c.Println(green("Hello we start adding new field to entity"))

    existsModels := getExistsModels()

    dbmodelsList := getDbmodelsList(existsModels)

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

    entity, err := getName(c, true, "Entity")

    if err != nil {
        return
    }

    fmt.Println(green("select model: ", entity))

    existsModels.ShowFields(entity)

    fmt.Println("Please enter name of new field:")

    field, err := getName(c, false, "Field")

    if err != nil {
        return
    }

    dataType, err := getDataType(c)

    if err != nil {
        return
    }

    existsModels.addField(entity, field, dataType)

    defer clearEntities(existsModels)
}
func getDataType(c *ishell.Context) (dataType string, err error) {

    dataTypes := []string{
        "string",
        "int",
        "time.Time",
        "float",
        "uuid.UUID",
    }

    choice := c.MultiChoice(dataTypes, "Please select data type")

    return dataTypes[choice], nil
}

func getDbmodelsList(repository modelRepository) (list []string) {

    for _, model := range repository.list {

        for _, spec := range model.Structures {

            tp, ok := spec.(*ast.GenDecl)

            if !ok {
                continue
            }

            for _,md := range tp.Specs {

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

func getExistsModels() (repo modelRepository) {

    dir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    workFolder := dir + "/dbmodels"

    files, err := ioutil.ReadDir(workFolder)
    if err != nil {
        return
    }

    validFileName := regexp.MustCompile(`^[a-zA-Z0-9_]+\.go$`)

    for _, f := range files {

        isValidFile := validFileName.MatchString(f.Name())

        if f.IsDir() || ! isValidFile {
            fmt.Println("invalid file " + f.Name())
            continue
        }

        filePath := workFolder + "/" + f.Name()

        models, parsedFile := getFileModels(filePath)

        repo.list = append(repo.list, model{
            FilePath: filePath,
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

func clearEntities(repo modelRepository) {

    for _, entity := range getDbmodelsList(repo) {
       shell.DeleteCmd(entity)
    }
}
