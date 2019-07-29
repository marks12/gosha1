package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
	"github.com/fatih/color"
	"os"
	"gosha/common"
)

func usualAuthAdd(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()
	c.Println(yellow("Hello we start adding auth to app"))

	os.Args = append(os.Args, "--entity=User")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-1]

	os.Args = append(os.Args, "--entity=Role")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-1]

	os.Args = append(os.Args, "--entity=UserRole")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-1]

	os.Args = append(os.Args,"--entity=Auth")
	os.Args = append(os.Args,"--crud=cd")
	os.Args = append(os.Args,"--check-auth=d")
	os.Args = append(os.Args,"--no-id")
	os.Args = append(os.Args,"--no-assign")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-5]

	c.Println("Models created success")


	// fill fields
	fillUser(c)
	fillRole(c)
	fillUserRole(c)
	fillAuth(c)
}

func fillUserRole(c *ishell.Context) {
	CopyFile(
		"types/user_role.go",
		"types/user_role.go",
		[]string{getRemoveLine("UserRole")},
		[]string{
			`RoleID int
	UserID int
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
    `+ getRemoveLine("Role")},
		c)
}

func fillUser(c *ishell.Context) {

	CreateFile("settings/user.go", `package settings
	const PASSWORD_SALT = "` + common.RandomString(10) + `"`, c)

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
    Password    string  ` + "`" + `json:"-"` + "`" + `
    Token       string
    `+ getRemoveLine("User")},
		c)

	CopyFile(
		"logic/assigner.go",
		"logic/assigner.go",
		[]string{getRemoveLine("AssignUserDbFromType predefine")},
		[]string{
			`password := []byte(typeModel.Password + settings.PASSWORD_SALT)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    ` + getRemoveLine("AssignUserDbFromType predefine")},
		c)

	CopyFile(
		"logic/assigner.go",
		"logic/assigner.go",
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
		"logic/assigner.go",
		"logic/assigner.go",
		[]string{getRemoveLine("AssignUserTypeFromDb.Field")},
		[]string{
			`FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		MobilePhone: dbUser.MobilePhone,
		Email:       dbUser.Email,
		Password:    "*****",
		IsActive:    dbUser.IsActive,
		Token:       dbUser.Token,
    	` + getRemoveLine("AssignUserTypeFromDb.Field")},
		c)

	CopyFile(
		"dbmodels/user.go",
		"dbmodels/user.go",
		[]string{getRemoveLine("User")},
		[]string{
			`Email       string
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string
    Token       string
    ` + getRemoveLine("User")},
		c)

	CopyFile(
		"dbmodels/user.go",
		"dbmodels/user.go",
		[]string{getRemoveLine("Validate")},
		[]string{`

    if len(user.FirstName) < 1 {
        user.validationErrors = append(user.validationErrors, "User first name is empty")
    }

    if len(user.LastName) < 1 {
        user.validationErrors = append(user.validationErrors, "User last name is empty")
    }

    if len(user.Email) < 3 || ! common.ValidateEmail(user.Email)  {
        user.validationErrors = append(user.validationErrors, "User email not valid")
    }

    if len(user.MobilePhone) > 3 &&  ! common.ValidateMobile(user.MobilePhone)  {
        user.validationErrors = append(user.validationErrors, "User mobile phone should be valid or empty. Format +0123456789... ")
    }

    `+ getRemoveLine("Validate") + `
`},
		c)
}

func fillAuth(c *ishell.Context) {

	CreateFile(
		"logic/auth.go",
		usualTemplateAuthLogic.Content,
		c)

	CopyFile(
		"types/auth.go",
		"types/auth.go",
		[]string{getRemoveLine("Auth")},
		[]string{
			`Email     string
    Password  string	` + "`" + `json:"-"` + "`" + `
    Token     string
    `+ getRemoveLine("Auth")},
		c)

	CopyFile(
		"dbmodels/auth.go",
		"dbmodels/auth.go",
		[]string{getRemoveLine("Auth")},
		[]string{
			`Email     string
    Password  string
    Token     string
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

	CopyFile(
		"dbmodels/auth.go",
		"dbmodels/auth.go",
		[]string{getRemoveLine("Validate")},
		[]string{
			`if len(auth.Email) < 3 || ! common.ValidateEmail(auth.Email)  {
        auth.validationErrors = append(auth.validationErrors, "User email not valid")
    }

    if len(auth.Password) < 1 {
        auth.validationErrors = append(auth.validationErrors, "User password is empty")
    }

    `+ getRemoveLine("Validate")},
		c)


}
