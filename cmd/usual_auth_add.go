package cmd

import (
	"github.com/fatih/color"
	"gopkg.in/abiosoft/ishell.v2"
	"gosha/common"
	"os"
)

func usualAuthAdd(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()
	c.Println(yellow("Hello we start adding auth to app"))

	os.Args = append(os.Args, "--entity=User")
	os.Args = append(os.Args,"--check-auth=fruda")
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

	os.Args = append(os.Args,"--entity=Auth")
	os.Args = append(os.Args,"--crud=cd")
	os.Args = append(os.Args,"--check-auth=d")
	os.Args = append(os.Args,"--no-id")
	os.Args = append(os.Args,"--no-assign")

	usualEntityAdd(c)

	os.Args = os.Args[:len(os.Args)-5]

	os.Args = append(os.Args,"--entity=CurrentUser")
	os.Args = append(os.Args,"--crud=f")
	os.Args = append(os.Args,"--check-auth=fcrudax")

	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-3]

	c.Println("Models created success")

    argsBak := os.Args

    os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=ResourceType", "--Field=Name", "--data-type=string"}
    entityFieldAdd(c)

    os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=UserRole", "--Field=UserId", "--data-type=int"}
    entityFieldAdd(c)
    os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=UserRole", "--Field=RoleId", "--data-type=int"}
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

    os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=RoleId", "--data-type=int"}
    entityFieldAdd(c)
    os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=RoleResource", "--Field=ResourceId", "--data-type=int"}
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
	fillUser(c)
	//fillRole(c)
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
	const PASSWORD_SALT = "` + common.RandomString(10) + `"

	const AdminRoleId = 1
	const UserRoleId = 2
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
    Token       string
    `+ getRemoveLine("User")},
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
    	` + getRemoveLine("AssignUserTypeFromDb.Field")},
		c)

	CopyFile(
		"dbmodels/user.go",
		"dbmodels/user.go",
		[]string{getRemoveLine("User")},
		[]string{
			`Email       string  `+"`"+`gorm:"type:varchar(100);unique_index"`+"`"+`
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string
    ` + getRemoveLine("User")},
		c)

	//CopyFile(
		//"dbmodels/role.go",
		//"dbmodels/role.go",
		//[]string{getRemoveLine("Role")},
		//[]string{
		//	`Name       string
    //Description   string	`+"`"+`gorm:"type:text"`+"`"+`
    //` + getRemoveLine("Role")},
		//c)

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
    Password  string
    Token     string
    UserId   int
    Id   int
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
    IsActive bool
    UserId   int
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


	CopyFile(
		"logic/assigner.go",
		"logic/assigner.go",
		[]string{getRemoveLine("AssignAuthDbFromType.Field")},
		[]string{`Email: typeModel.Email,
        Password: typeModel.Password,
		` + getRemoveLine("AssignAuthDbFromType.Field")},
		c)

}
