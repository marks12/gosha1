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
)

type modelRepository struct {
    list []model
}

type model struct {
    FilePath string
    Structures []ast.Decl
}

func (mr *modelRepository) getModels() (models []string) {

    for _, m := range mr.list {
        models = append(models, m.FilePath)
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

    entity, err := getName(c, true)

    if err != nil {
        return
    }

    fmt.Println("select model", entity)

    defer clearEntities(existsModels)
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

                //structDecl := md.(*ast.TypeSpec).Type.(*ast.StructType)

                structDecl := md.(*ast.TypeSpec)
                list = append(list, structDecl.Name.String())
            }
        }
    }

    return

    //structDecl := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
    //
    //c := &ast.Comment{Text: fmt.Sprint("// ", "Some comment")}
    //cg := &ast.CommentGroup{List: []*ast.Comment{c}}
    //field := &ast.Field{
    //   Doc:   cg,
    //   Names: []*ast.Ident{ast.NewIdent("Field4")},
    //   Type:  ast.NewIdent("float"),
    //}
    //
    //structDecl.Fields.List = append(structDecl.Fields.List, field)
    //
    //return
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

        fmt.Println("Valid file " + f.Name())

        filePath := workFolder + "/" + f.Name()

        models := getFileModels(filePath)

        repo.list = append(repo.list, model{
            FilePath: filePath,
            Structures: models,
        })
    }

    return
}

func getFileModels(filePath string) []ast.Decl {

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

    return typeDecl

}

func clearEntities(repo modelRepository) {

    //for _, entity := range repo.list {
    //    shell.DeleteCmd(entity.FilePath)
    //}
}
