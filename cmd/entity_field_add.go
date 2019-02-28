package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "fmt"
    "go/token"
    "go/ast"
    "go/parser"
)

type existsModels struct {
    list []model
}

type model struct {
    File string
    Name string
}

func (em *existsModels) getModels() (models []string) {

    for _, m := range em.list {
        models = append(models, m.Name)
    }
    return
}

func (em *existsModels) getModelFile(name string) (fileName string) {

    for _, m := range em.list {
        if m.Name == name {
            fileName = m.File
            break
        }
    }

    return
}

func entityFieldAdd(c *ishell.Context) {

    green := color.New(color.FgGreen).SprintFunc()

    c.Println(green("Hello we start adding new field to entity"))

    existsModels := getExistsModels()

    entity, err := getName(c, true)

    if err != nil {
        return
    }

    fmt.Println("select model", entity)

    defer clearEntities(existsModels)
}

func getExistsModels() {

    fs := token.NewFileSet()
    pkgs, err := parser.ParseDir(fs, "./dbmodels", nil, 0)

    if err != nil {
        fmt.Println("error reading dbmodels", err)
        return
    }

    pkg, ok := pkgs["dbmodels"]

    if ! ok {
        fmt.Println("error reading package dbmodels", err)
        return
    }

    ast.Walk(VisitorFunc(FindTypes), pkg)

    for _, model := range existsModels {

        shell.AddCmd(&ishell.Cmd{
            Name: model,
            Help: model + " entity",
            Func: func(c *ishell.Context) {
                fmt.Println("select predefined function", model)
            },
        })
    }

    return
}


func clearEntities(entities []string) {

    for _, entity := range entities {
        shell.DeleteCmd(entity)
    }
}

func FindTypes(s ast.Node) ast.Visitor {

    switch n := s.(type) {

        case *ast.Package:
            return VisitorFunc(FindTypes)
        case *ast.File:

            fmt.Println("found file in pkg")

            return VisitorFunc(FindTypes)
        case *ast.GenDecl:
            if n.Tok == token.TYPE {
                return VisitorFunc(FindTypes)
            }
        case *ast.TypeSpec:
            existsModels = append(existsModels, n.Name.Name)
    }

    return nil
}