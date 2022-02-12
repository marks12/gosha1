package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/fatih/color"
	"gosha/mode"
	"os"
)

func usualAuthAdd(c *ishell.Context, salt string) {

	yellow := color.New(color.FgYellow).SprintFunc()
	c.Println(yellow("Hello we start adding auth to app"))

	os.Args = append(os.Args, "--entity=User")
	os.Args = append(os.Args, "--check-auth=fcruda")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-2]
	os.Args = append(os.Args, "--entity=Role")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-1]
	os.Args = append(os.Args, "--entity=RoleResource")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-1]
	os.Args = append(os.Args, "--entity=Resource")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-1]
	os.Args = append(os.Args, "--entity=ResourceType")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-1]
	os.Args = append(os.Args, "--entity=UserRole")
	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-1]
	os.Args = append(os.Args, "--entity=Auth")
	os.Args = append(os.Args, "--crud=cd")
	os.Args = append(os.Args, "--check-auth=d")
	os.Args = append(os.Args, "--no-id")
	os.Args = append(os.Args, "--no-assign")

	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-5]

	os.Args = append(os.Args, "--entity=CurrentUser")
	os.Args = append(os.Args, "--crud=f")
	os.Args = append(os.Args, "--check-auth=fcrudax")

	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-3]

	c.Println("Models created success")

	argsBak := os.Args

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=ResourceType", "--Field=Name", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=UserRole", "--Field=UserId", "--data-type=" + GetPKType(mode.GetUuidMode())}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=UserRole", "--Field=RoleId", "--data-type=" + GetPKType(mode.GetUuidMode())}
	entityFieldAdd(c)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Role", "--Field=Name", "--data-type=string"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Role", "--Field=Description", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Resource", "--Field=Name", "--data-type=string"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Resource", "--Field=Code", "--data-type=string"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=Resource", "--Field=TypeId", "--data-type=int"}
	entityFieldAdd(c)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=RoleId", "--data-type=" + GetPKType(mode.GetUuidMode())}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=ResourceId", "--data-type=" + GetPKType(mode.GetUuidMode())}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=Find", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=Read", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=Create", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=Update", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=Delete", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=FindOrCreate", "--data-type=bool"}
	entityFieldAdd(c)
	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=UpdateOrCreate", "--data-type=bool"}
	entityFieldAdd(c)

	os.Args = argsBak

	// fill fields
	fillUser(c, salt)
	//fillRole(c)
	//fillUserRole(c)
	fillAuth(c)
}

func fillUserRole(c *ishell.Context) {
	CopyFile(
		"types/user_role.go",
		"types/user_role.go",
		[]string{getRemoveLine("UserRole")},
		[]string{
			`RoleID ` + GetPKType(mode.GetUuidMode()) + `
	UserID ` + GetPKType(mode.GetUuidMode()) + `
    ` + getRemoveLine("UserRole")},
		c)
}

func fillRole(c *ishell.Context) {
	CopyFile(
		"types/role.go",
		"types/role.go",
		[]string{getRemoveLine("Role")},
		[]string{
			`Name        string
	Description string
    ` + getRemoveLine("Role")},
		c)
}

func fillUser(c *ishell.Context, salt string) {

	CreateFile("settings/user.go", `package settings
	const PASSWORD_SALT = "`+salt+`"

	const AdminRoleId ConfigId = 1
	const UserRoleId ConfigId = 2
`, c)

	CopyFile(
		"types/user.go",
		"types/user.go",
		[]string{getRemoveLine("User")},
		[]string{
			`Email       string
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string
    ` + getRemoveLine("User")},
		c)

	CopyFile(
		"logic/user.go",
		"logic/user.go",
		[]string{getRemoveLine("AssignUserDbFromType.Field")},
		[]string{
			`Email:       typeModel.Email,
        FirstName:   typeModel.FirstName,
        IsActive:    typeModel.IsActive,
        LastName:    typeModel.LastName,
        MobilePhone: typeModel.MobilePhone,
		Password:    string(hashedPassword),
    ` + getRemoveLine("AssignUserDbFromType.Field")},
		c)

	CopyFile(
		"logic/user.go",
		"logic/user.go",
		[]string{getRemoveLine("AssignUserTypeFromDb.Field")},
		[]string{
			`Email:       dbUser.Email,
        FirstName:   dbUser.FirstName,
        IsActive:    dbUser.IsActive,
        LastName:    dbUser.LastName,
        MobilePhone: dbUser.MobilePhone,
        Password:    "******",
    ` + getRemoveLine("AssignUserTypeFromDb.Field")},
		c)

	addImportIfNeed("logic/user.go", GetCurrentAppName()+"/settings")
	addImportIfNeed("logic/user.go", "golang.org/x/crypto/bcrypt")

	CopyFile(
		"logic/user.go",
		"logic/user.go",
		[]string{getRemoveLine("AssignUserDbFromType predefine")},
		[]string{
			`password := []byte(typeModel.Password + settings.PASSWORD_SALT)
    hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    ` + getRemoveLine("AssignUserTypeFromDb predefine")},
		c)

	CopyFile(
		"logic/current_user.go",
		"logic/current_user.go",
		[]string{"criteria := core.Db.Where(dbmodels.CurrentUser{})"},
		[]string{
			`criteria := core.Db.Where(dbmodels.User{})
	    criteria = criteria.Where(dbmodels.User{Token: filter.Authenticator.Token})
    	q := criteria.Model(dbmodels.User{}).Count(&count)`}, c)

	CopyFile(
		"logic/auth.go",
		"logic/auth.go",
		[]string{getRemoveLine("AssignUserDbFromType predefine")},
		[]string{
			`password := []byte(typeModel.Password + settings.PASSWORD_SALT)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    ` + getRemoveLine("AssignUserDbFromType predefine")},
		c)

	CopyFile(
		"logic/auth.go",
		"logic/auth.go",
		[]string{getRemoveLine("AssignUserDbFromType.Field")},
		[]string{
			`FirstName:   typeModel.FirstName,
		LastName:    typeModel.LastName,
		MobilePhone: typeModel.MobilePhone,
		Email:       typeModel.Email,
		Password:    string(hashedPassword),
		IsActive:    typeModel.IsActive,
	    ` + getRemoveLine("AssignUserDbFromType.Field")},
		c)

	CopyFile(
		"logic/auth.go",
		"logic/auth.go",
		[]string{getRemoveLine("AssignUserTypeFromDb.Field")},
		[]string{
			`FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		MobilePhone: dbUser.MobilePhone,
		Email:       dbUser.Email,
		Password:    "*****",
		IsActive:    dbUser.IsActive,
    	` + getRemoveLine("AssignUserTypeFromDb.Field")},
		c)

	CopyFile(
		"dbmodels/user.go",
		"dbmodels/user.go",
		[]string{getRemoveLine("User")},
		[]string{
			`Email       string  ` + "`" + `gorm:"type:varchar(100);unique_index"` + "`" + `
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string
    ` + getRemoveLine("User")},
		c)

	path := "dbmodels/user.go"
	CopyFile(
		path,
		path,
		[]string{getRemoveLine("Validate")},
		[]string{`

    if len(user.FirstName) < 1 {
        user.AddValidationError("User first name is empty", errors.ErrorCodeFieldLengthTooShort, "FirstName")
    }

    if len(user.LastName) < 1 {
        user.AddValidationError("User last name is empty", errors.ErrorCodeFieldLengthTooShort, "LastName")
    }

    if len(user.Email) < 3 || ! common.ValidateEmail(user.Email)  {
        user.AddValidationError("User email not valid", errors.ErrorCodeNotValid, "Email")
    }

    if len(user.MobilePhone) > 3 &&  ! common.ValidateMobile(user.MobilePhone)  {
        user.AddValidationError("User mobile phone should be valid or empty. Format +0123456789... ", errors.ErrorCodeNotValid, "MobilePhone")
    }

    ` + getRemoveLine("Validate") + `
`},
		c)
	addImportIfNeed(path, GetCurrentAppName()+"/errors")

}

func fillAuth(c *ishell.Context) {

	isUuidAsPk := mode.GetUuidMode()

	CreateFile(
		"logic/auth.go",
		GetUsualTemplateAuthLogic().Content,
		c)

	CopyFile(
		"types/auth.go",
		"types/auth.go",
		[]string{getRemoveLine("Auth")},
		[]string{
			`Email     string
    Password  string
    Token     string
    UserId   ` + GetPKType(isUuidAsPk) + `
    Id   ` + GetPKType(isUuidAsPk) + `
    ` + getRemoveLine("Auth")},
		c)

	CopyFile(
		"dbmodels/auth.go",
		"dbmodels/auth.go",
		[]string{getRemoveLine("Auth")},
		[]string{
			`Email     string
    Password  string
    Token     string
    IsActive bool
    UserId   ` + GetPKType(isUuidAsPk) + `
    ` + getRemoveLine("Auth")},
		c)

	CopyFile(
		"dbmodels/auth.go",
		"dbmodels/auth.go",
		[]string{`import \(`},
		[]string{`import (
	` + assignMsName(`"{ms-name}/common"`)},
		c)

	CopyFile(
		"dbmodels/user.go",
		"dbmodels/user.go",
		[]string{`import \(`},
		[]string{`import (
	` + assignMsName(`"{ms-name}/common"`)},
		c)

	path := "dbmodels/auth.go"
	CopyFile(
		path,
		path,
		[]string{getRemoveLine("Validate")},
		[]string{
			`    if len(auth.Email) < 3 || ! common.ValidateEmail(auth.Email)  {
        auth.AddValidationError("User email not valid", errors.ErrorCodeFieldLengthTooShort, "Email")
    }

    if len(auth.Password) < 1 {
        auth.AddValidationError("User password is empty", errors.ErrorCodeNotEmpty, "Password")
    }

    ` + getRemoveLine("Validate")},
		c)
	addImportIfNeed(path, GetCurrentAppName()+"/errors")

	CopyFile(
		"logic/auth.go",
		"logic/auth.go",
		[]string{getRemoveLine("AssignAuthDbFromType.Field")},
		[]string{`Email: typeModel.Email,
        Password: typeModel.Password,
		` + getRemoveLine("AssignAuthDbFromType.Field")},
		c)

}
