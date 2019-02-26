package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "fmt"
)

func entityFieldAdd(c *ishell.Context) {

    var entities []string

    entities = findAllExistsEntities()

    green := color.New(color.FgGreen).SprintFunc()

    c.Println(green("Hello we start adding new field to entity"))

    shell.AddCmd(&ishell.Cmd{
        Name: "SomeEntity",
        Help: "Entity helper",
        Func: func(c *ishell.Context) {
            fmt.Println("select predefined function", "some function")
        },
    })

    entity, err := getName(c, true)

    if err != nil {
        return
    }

    fmt.Println("select model", entity)

    defer clearEntiies(entities)
}

func findAllExistsEntities() []string {

    var entities []string



    return entities
}

func clearEntiies(entities []string) {

    for _, entity := range entities {
        shell.DeleteCmd(entity)
    }
}