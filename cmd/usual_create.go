package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "github.com/fatih/color"
    "os"
)

func usualAppInit(c *ishell.Context) {

    yellow := color.New(color.FgYellow).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()

    c.Println(yellow("Hello we start creating new usual app in current directory"))

    choice := c.MultiChoice([]string{
        "No",
        "Yes",
    }, "Continue ?")

    switch choice {

        // mscore
        case 1:

            c.Println(green("Creating app structure"))

            usualCreate(c)

            break

    }

    return
}

func usualCreate(c *ishell.Context) {

    green := color.New(color.FgCyan).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    c.Println("Hello new app", green(getCurrentDirName()))

    usualCreateMain(c)

    //bootstrap
    CreateFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Content, c)

    //core
    CreateFile(msTemplateCoreDb.Path, msTemplateCoreDb.Content, c)

    //settings
    CreateFile(msTemplateSettingsApp.Path, msTemplateSettingsApp.Content, c)
    CreateFile(msTemplateSettingsDb.Path, msTemplateSettingsDb.Content, c)
    CreateFile(usualTemplateSettingsWebapp.Path, usualTemplateSettingsWebapp.Content, c)

    //service
    CreateFile(usualTemplateServicesCaller.Path, usualTemplateServicesCaller.Content, c)
    CreateFile(usualTemplateServicesTicket.Path, usualTemplateServicesTicket.Content, c)

    //dbmodels
    CreateFile(usualTemplateDbmodelsEntity.Path, usualTemplateDbmodelsEntity.Content, c)
    CreateFile(usualTemplateDbmodelsValidator.Path, usualTemplateDbmodelsValidator.Content, c)

    //types
    CreateFile(usualTemplateTypesAuthenticator.Path, usualTemplateTypesAuthenticator.Content, c)
    CreateFile(usualTemplateTypesEntity.Path, usualTemplateTypesEntity.Content, c)
    CreateFile(usualTemplateTypesValidator.Path, usualTemplateTypesValidator.Content, c)
    CreateFile(usualTemplateTypesResponse.Path, usualTemplateTypesResponse.Content, c)

    //webapp
    CreateFile(usualTemplateWebappErrors.Path, usualTemplateWebappErrors.Content, c)

    //logic
    CreateFile(msTemplateLogicAssigner.Path, msTemplateLogicAssigner.Content, c)

    //router
    CreateFile(usualTemplateRouter.Path, usualTemplateRouter.Content, c)

    ////ms folder
    //CreateFile(msTemplateMsTicket.Path, msTemplateMsTicket.Content, c)

    CreateFile("./.gitignore", "./\\.idea\n", c)

    c.Println(red("New app with usual structure created"))
}

func usualCreateMain(c *ishell.Context) {

    CreateFile(usualTemplateMain.Path, usualTemplateMain.Content, c)

    for _, folder := range []string{
        "bootstrap",
        "core",
        "dbmodels",
        "logic",
        "ms",
        "router",
        "services",
        "settings",
        "static",
        "filters",
        "types",
        "webapp",
    } {
        if _, err := os.Stat(folder); os.IsNotExist(err) {
            os.Mkdir(folder, 0755)
        }
    }
}