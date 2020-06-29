package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "gosha/mode"
    "os"
)

func msInit(c *ishell.Context) {

    green := color.New(color.FgCyan).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    c.Println("Hello new microservice", green(getCurrentDirName()))

    msCreateMain(c)

    dbtype := DatabaseType{DbTypeName:dbTypeMysql, IsMysql:true, IsPostgres:false}

    msTemplateInsertDataToDb := GetMsTemplateInsertDataToDb()
    //bootstrap
    CreateFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Content, c)
    if mode.GetUuidMode() {
        addImportIfNeed(msTemplateInsertDataToDb.Path, "github.com/google/uuid")
    }

    //core
    msTemplateCoreDb := getTemplateCoreDb(dbtype)
    CreateFile(msTemplateCoreDb.Path, msTemplateCoreDb.Content, c)

    //settings
    CreateFile(msTemplateSettingsApp.Path, msTemplateSettingsApp.Content, c)

    msTemplateSettingsDb := getMsTemplateSettingsDb(dbtype)
    CreateFile(msTemplateSettingsDb.Path, msTemplateSettingsDb.Content, c)

    //dbmodels
    CreateFile(msTemplateDbmodelsEntity.Path, msTemplateDbmodelsEntity.Content, c)
    CreateFile(msTemplateDbmodelsValidator.Path, msTemplateDbmodelsValidator.Content, c)

    //rpcapp
    CreateFile(msTemplateRpcappErrors.Path, msTemplateRpcappErrors.Content, c)

    //logic
    //CreateFile(msTemplateLogicAssigner.Path, msTemplateLogicAssigner.Content, c)

    //ms forlder
    CreateFile(msTemplateMsTicket.Path, msTemplateMsTicket.Content, c)

    CreateFile("./.gitignore", "./\\.idea\n", c)

    c.Println(red("New app with microservice structure created"))
}

func msCreateMain(c *ishell.Context) {

    choice := c.MultiChoice([]string{"Yes", "No"}, "Do you wont rewrite mian.go ?")

    if choice == 0 {

        CreateFile(msTemplateMain.Path, msTemplateMain.Content, c)

        for _, folder := range []string{
            "core",
            "settings",
            "dbmodels",
            "bootstrap",
            "logic",
            "ms",
            "rpcapp",
        } {
            if _, err := os.Stat(folder); os.IsNotExist(err) {
                os.Mkdir(folder, 0755)
            }
        }
    }
}