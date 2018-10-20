package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "os"
)

func msInit(c *ishell.Context) {

    green := color.New(color.FgCyan).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    c.Println("Hello new microservice", green(getCurrentDirName()))

    createMain(c)
    CreateFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Content)
    CreateFile(msTemplateCoreDb.Path, msTemplateCoreDb.Content)
    CreateFile(msTemplateSettingsApp.Path, msTemplateSettingsApp.Content)
    CreateFile(msTemplateSettingsDb.Path, msTemplateSettingsDb.Content)
    CreateFile(msTemplateDbmodelsEntity.Path, msTemplateDbmodelsEntity.Content)
    CreateFile(msTemplateRpcappErrors.Path, msTemplateRpcappErrors.Content)
    CreateFile("./.gitignore", ".idea\n")

    c.Println(red("New app with microservice structure created"))
}

func createMain(c *ishell.Context) {

    choice := c.MultiChoice([]string{"Yes", "No"}, "Do you wont rewrite mian.go ?")

    if choice == 0 {

        CreateFile(msTemplateMain.Path, msTemplateMain.Content)

        for _, folder := range []string{
            "core",
            "settings",
            "dbmodels",
            "bootstrap",
            "ms",
            "rpcapp",
        } {
            if _, err := os.Stat(folder); os.IsNotExist(err) {
                os.Mkdir(folder, 0755)
            }
        }
    }
}