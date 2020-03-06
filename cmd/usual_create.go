package cmd

import (
	"github.com/fatih/color"
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/mode"
	"os"
)

func UsualAppInit(c *ishell.Context) {

	var choice int
	var email string
	var password string

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	c.Println(yellow("Hello we start creating new usual app in current directory"))

	if mode.IsNonInteractive() {
		choice = 1
	} else {
		choice = c.MultiChoice([]string{
			"No",
			"Yes",
		}, "Continue ?")
	}

	email, _ = getEmail(c)
	password, _ = getPassword(c)

	switch choice {

	case 1:

		if mode.IsInteractive() {
			c.Println(green("Creating app structure"))
		}

		usualCreate(c, email, password)
		usualAuthAdd(c)

		break

	}

	return
}

func getEmail(c *ishell.Context) (email string, err error) {

	var arguments RegularFind
	arguments, err = GetOsArgument("adminMail")
	if len(arguments.StringResult) < 1 || err != nil {
		return getName(c, false, "Email")
	}
	email = arguments.StringResult
	return
}

func getPassword(c *ishell.Context) (password string, err error) {

	var arguments RegularFind
	arguments, err = GetOsArgument("adminPassword")
	if len(arguments.StringResult) < 1 || err != nil {
		return getName(c, false, "Password")
	}
	password = arguments.StringResult
	return
}

func usualCreate(c *ishell.Context, email, password string) {

	green := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	c.Println("Hello new app", green(getCurrentDirName()))

	usualCreateMain(c)

	//bootstrap
	CreateFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Content, c)
	CopyFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Path,
		[]string{"{email}", "{password}"},
		[]string{email, password},
		c)

	//core
	CreateFile(msTemplateCoreDb.Path, msTemplateCoreDb.Content, c)

	//dbmodels
	//CreateFile(usualTemplateDbmodelsEntity.Path, usualTemplateDbmodelsEntity.Content, c)
	CreateFile(usualTemplateDbmodelsValidator.Path, usualTemplateDbmodelsValidator.Content, c)

	//logic
	CreateFile(msTemplateLogicAssigner.Path, msTemplateLogicAssigner.Content, c)

	//router
	CreateFile(usualTemplateRouter.Path, usualTemplateRouter.Content, c)
	CreateFile(usualTemplateWssRouter.Path, usualTemplateWssRouter.Content, c)

	//wss
	CreateFile(usualTemplateWsserver.Path, usualTemplateWsserver.Content, c)
	CreateFile(usualTemplateWssHandlers.Path, usualTemplateWssHandlers.Content, c)

	//google
	CreateFile(usualTemplateGoogleAnalytics.Path, usualTemplateGoogleAnalytics.Content, c)

	//service
	CreateFile(usualTemplateServicesCaller.Path, usualTemplateServicesCaller.Content, c)
	CreateFile(usualTemplateServicesTicket.Path, usualTemplateServicesTicket.Content, c)


	//settings
	CreateFile(msTemplateSettingsApp.Path, msTemplateSettingsApp.Content, c)
	CreateFile(msTemplateSettingsGoogle.Path, msTemplateSettingsGoogle.Content, c)
	CreateFile(msTemplateSettingsDb.Path, msTemplateSettingsDb.Content, c)
	CreateFile(usualTemplateSettingsWebapp.Path, usualTemplateSettingsWebapp.Content, c)
	CreateFile(usualTemplateSettingsRoutes.Path, usualTemplateSettingsRoutes.Content, c)
	CreateFile(usualTemplateTestsRequest.Path, usualTemplateTestsRequest.Content, c)
	CreateFile(usualTemplateTestsUsers.Path, usualTemplateTestsUsers.Content, c)
	CreateFile(msTemplateSettingsWss.Path, msTemplateSettingsWss.Content, c)
	CreateFile(usualTemplateSettingsResource.Path, usualTemplateSettingsResource.Content, c)

	//generator
	CreateFile(usualTemplateGen.Path, usualTemplateGen.Content, c)

	//types
	CreateFile(usualTemplateTypesAuthenticator.Path, usualTemplateTypesAuthenticator.Content, c)
	//CreateFile(usualTemplateTypesEntity.Path, usualTemplateTypesEntity.Content, c)
	CreateFile(usualTemplateTypesFilter.Path, usualTemplateTypesFilter.Content, c)
	CreateFile(usualTemplateTypesRequest.Path, usualTemplateTypesRequest.Content, c)
	CreateFile(usualTemplateTypesValidator.Path, usualTemplateTypesValidator.Content, c)
	CreateFile(usualTemplateTypesResponse.Path, usualTemplateTypesResponse.Content, c)

	//mdl
	CreateFile(usualTemplateMdlPagination.Path, usualTemplateMdlPagination.Content, c)
	CreateFile(usualTemplateMdlRequest.Path, usualTemplateMdlRequest.Content, c)
	CreateFile(usualTemplateMdlResponse.Path, usualTemplateMdlResponse.Content, c)

	//webapp
	CreateFile(usualTemplateWebappErrors.Path, usualTemplateWebappErrors.Content, c)

	//common
	CreateFile(usualTemplateCommonValidator.Path, usualTemplateCommonValidator.Content, c)
	CreateFile(usualTemplateCommonGenerator.Path, usualTemplateCommonGenerator.Content, c)

	//flags
	CreateFile(usualTemplateFlags.Path, usualTemplateFlags.Content, c)

	////ms folder
	//CreateFile(msTemplateMsTicket.Path, msTemplateMsTicket.Content, c)

	CreateFile("./.gitignore", "./\\.idea\n", c)

	//docker
	CreateFile(usualTemplateDockerPs.Path, usualTemplateDockerPs.Content, c)

	//docker
	CreateFile(usualTemplateDockerMy.Path, usualTemplateDockerMy.Content, c)

	c.Println(red("New app with usual structure created"))
}

func usualCreateMain(c *ishell.Context) {

	CreateFile(usualTemplateMain.Path, usualTemplateMain.Content, c)

	for _, folder := range []string{
		"postgres",
		"bootstrap",
		"core",
		"dbmodels",
		"logic",
		"ms",
		"router",
		"services",
		"settings",
		"tests",
		"generator",
		"google",
		"static",
		"types",
		"webapp",
		"flags",
		"mdl",
		"common",
		"wsserver",
	} {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
	}
}
