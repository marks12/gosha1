package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/common"
	"gosha/mode"
	"os"
)

const dbTypeMysql = "mysql"
const dbTypePostgres = "postgres"

type DatabaseType struct {
	IsMysql    bool
	IsPostgres bool
	DbTypeName string
}

func (db *DatabaseType) Validate() {
	db.IsPostgres = db.DbTypeName == dbTypePostgres
	db.IsMysql = db.DbTypeName == dbTypeMysql
}

func (db *DatabaseType) IsValid() bool {
	return db.IsPostgres != db.IsMysql
}

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
	database, err := getDatabaseType(c)
	isUuidMode, err := getUuidMode(c)
	isViewMode, err := getViewMode(c)

	if isUuidMode {
		mode.SetUuidMode()
	} else {
		mode.SetNonUuidMode()
	}

	if isViewMode {
		mode.SetViewMode()
	} else {
		mode.SetNonViewMode()
	}

	if err != nil {
		fmt.Println("Creating app error: " + err.Error())
		os.Exit(1)
	}

	switch choice {

	case 1:

		if mode.IsInteractive() {
			c.Println(green("Creating app structure"))
		}

		salt := common.RandomString(10)

		usualCreate(c, email, password, salt, database, isUuidMode, isViewMode)
		usualAuthAdd(c)
		initGoModules(c)
		addTranslate(c)
		addRegionLang(c)
		if mode.GetViewMode() {
			usualCreateHtmlTemplate(c)
		}

		break

	}

	return
}

func initGoModules(c *ishell.Context) {
	content := GetGoModContent()
	CreateFile("./go.mod", content, c)
}

func addTranslate(c *ishell.Context) {

	argsBak := os.Args

	os.Args = append(os.Args, "--entity=TranslateError")
	os.Args = append(os.Args, "--check-auth=fcruda")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-2]

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=TranslateError", "--Field=Code", "--data-type=int"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=TranslateError", "--Field=LanguageCode", "--data-type=string"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=TranslateError", "--Field=Translate", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = argsBak

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FILTER, "--entity=TranslateErrorFilter", "--filter=ErrorCodes", "--data-type=[]int", "--sql-field=code"}
	entityFilterdAdd(c)

	os.Args = argsBak

	return
}

func addRegionLang(c *ishell.Context) {

	argsBak := os.Args

	os.Args = append(os.Args, "--entity=Region")
	os.Args = append(os.Args, "--check-auth=fcruda")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-2]

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Region", "--Field=Name", "--data-type=int"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Region", "--Field=Code", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = argsBak

	os.Args = append(os.Args, "--entity=Language")
	os.Args = append(os.Args, "--check-auth=fcruda")
	usualEntityAdd(c)

	os.Args = argsBak

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Language", "--Field=Name", "--data-type=int"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Language", "--Field=Code", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = argsBak

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

func getUuidMode(c *ishell.Context) (isUuid bool, err error) {

	arguments, err := GetOsArgument(UuidAsPk.ToString())
	if err != nil {

		choice := c.MultiChoice([]string{
			"Yes",
			"No",
		}, "Would you like to use UUID as primary key? Choose `No` for use integer")

		if choice == 0 {
			isUuid = true
		}

	} else {
		isUuid = arguments.BoolResult
	}

	return
}

func getViewMode(c *ishell.Context) (isViewMode bool, err error) {

	arguments, err := GetOsArgument(ViewMode.ToString())
	if err != nil {

		choice := c.MultiChoice([]string{
			"Yes",
			"No",
		}, "Would you like to use ViewMode for build html on backend? Choose `No` for disable views")

		if choice == 0 {
			isViewMode = true
		}

	} else {
		isViewMode = arguments.BoolResult
	}

	return
}

func getDatabaseType(c *ishell.Context) (dbtype DatabaseType, err error) {

	var arguments RegularFind
	arguments, err = GetOsArgument("dbType")

	if err != nil || len(arguments.StringResult) < 1 {
		return DatabaseType{}, errors.New("Invalid database. Only postgres, mysql supports " + dbTypeMysql + " " + dbTypePostgres + " ")
	}

	if isEx, _ := InArray(arguments.StringResult, []string{dbTypeMysql, dbTypePostgres}); isEx != true {
		return DatabaseType{}, errors.New("Invalid database. Only postgres, mysql supports " + dbTypeMysql + " " + dbTypePostgres + " ")
	}

	dbtype = DatabaseType{
		DbTypeName: arguments.StringResult,
	}

	dbtype.Validate()

	if !dbtype.IsValid() {
		return DatabaseType{}, errors.New("Invalid database. Only postgres, mysql supports " + dbTypeMysql + " " + dbTypePostgres + " ")
	}

	return dbtype, nil
}

func usualCreate(c *ishell.Context, email, password, salt string, databaseType DatabaseType, isUuidAsPk bool, isViewMode bool) {

	green := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	c.Println("Hello new app", green(GetCurrentAppName()))

	usualCreateMain(c)

	//core
	msTemplateCoreDb := getTemplateCoreDb(databaseType)
	CreateFile(msTemplateCoreDb.Path, msTemplateCoreDb.Content, c)

	//bootstrap
	msTemplateInsertDataToDb := GetMsTemplateInsertDataToDb()
	CreateFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Content, c)

	pass := []byte(password + salt)
	hashedPassword, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	CopyFile(msTemplateInsertDataToDb.Path, msTemplateInsertDataToDb.Path,
		[]string{"{email}", "{password}"},
		[]string{email, string(hashedPassword)},
		c)

	if mode.GetUuidMode() {
		addImportIfNeed(msTemplateInsertDataToDb.Path, "github.com/google/uuid")
	}

	//dbmodels
	//CreateFile(usualTemplateDbmodelsEntity.Path, usualTemplateDbmodelsEntity.Content, c)
	CreateFile(usualTemplateDbmodelsValidator.Path, usualTemplateDbmodelsValidator.Content, c)

	//logic
	//CreateFile(msTemplateLogicAssigner.Path, msTemplateLogicAssigner.Content, c)

	//router
	CreateFile(usualTemplateRouter.Path, usualTemplateRouter.Content, c)
	CreateFile(usualTemplateWssRouter.Path, usualTemplateWssRouter.Content, c)

	//wss
	CreateFile(usualTemplateWsserver.Path, usualTemplateWsserver.Content, c)
	if mode.GetViewMode() {
		CreateFile(usualTemplateWssHandlersView.Path, usualTemplateWssHandlersView.Content, c)
	} else {
		CreateFile(usualTemplateWssHandlersNonView.Path, usualTemplateWssHandlersNonView.Content, c)
	}

	//google
	CreateFile(usualTemplateGoogleAnalytics.Path, usualTemplateGoogleAnalytics.Content, c)

	//service
	CreateFile(usualTemplateServicesCaller.Path, usualTemplateServicesCaller.Content, c)
	CreateFile(usualTemplateServicesTicket.Path, usualTemplateServicesTicket.Content, c)

	msTemplateSettingsDb := getMsTemplateSettingsDb(databaseType)
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

	CreateFile(usualTemplateSettingsType.Path, usualTemplateSettingsType.Content, c)

	//generator
	CreateFile(usualTemplateGen.Path, usualTemplateGen.Content, c)

	//types
	usualTemplateTypesAuthenticator := getUsualTemplateTypesAuthenticator(isUuidAsPk)
	CreateFile(usualTemplateTypesAuthenticator.Path, usualTemplateTypesAuthenticator.Content, c)
	if isUuidAsPk {
		addImportIfNeed(usualTemplateTypesAuthenticator.Path, "github.com/google/uuid")
	}

	//CreateFile(usualTemplateTypesEntity.Path, usualTemplateTypesEntity.Content, c)
	usualTemplateTypesFilter := getUsualTemplateTypesFilter(isUuidAsPk)
	CreateFile(usualTemplateTypesFilter.Path, usualTemplateTypesFilter.Content, c)
	if isUuidAsPk {
		addImportIfNeed(usualTemplateTypesFilter.Path, "github.com/google/uuid")
	}

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

	//Regionality
	CreateFile(usualTemplateRegionality.Path, usualTemplateRegionality.Content, c)

	//errors
	CreateFile(usualTemplateErrorCodes.Path, usualTemplateErrorCodes.Content, c)
	CreateFile(usualTemplateError.Path, usualTemplateError.Content, c)

	////ms folder
	//CreateFile(msTemplateMsTicket.Path, msTemplateMsTicket.Content, c)

	CreateFile("./.gitignore", "./\\.idea\n.postgres\n", c)

	if IsPostgres() {
		//docker
		CreateFile(usualTemplateDockerPs.Path, usualTemplateDockerPs.Content, c)
	} else {
		//docker
		CreateFile(usualTemplateDockerMy.Path, usualTemplateDockerMy.Content, c)
	}

	c.Println(red("New app with usual structure created"))
}

func usualCreateMain(c *ishell.Context) {

	CreateFile(usualTemplateMain.Path, usualTemplateMain.Content, c)

	folders := []string{
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
	}

	if mode.GetViewMode() {
		folders = append(folders, "view", "view/form")
	}

	for _, folder := range folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
	}
}
