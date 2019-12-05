package cmd

const dbmodelsAuth = `package dbmodels

type Auth struct {
    ID       int ` + "`" + `gorm:"primary_key"` + "`" + `
    Email    string
    Password string ` + "`" + `gorm:"-"` + "`" + `
    Token    string
    IsActive bool
    UserId   int
    //Auth remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func (auth *Auth) Validate() {
    if len(auth.Email) < 3 || ! common.ValidateEmail(auth.Email) {
        auth.validationErrors = append(auth.validationErrors, "User email not valid")
    }

    if len(auth.Password) < 1 {
        auth.validationErrors = append(auth.validationErrors, "User password is empty")
    }

    //Validate remove this line for disable generator functionality
}


`

const dbmodelsUser = `package dbmodels

import (
    "time"
)

type User struct {

    ID          int     `+"`"+`sql:"primary_key"`+"`"+`

    Email       string
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `+"`"+`sql:"index" json:"-"`+"`"+`

    validator
}

func (user *User) Validate() {

    if len(user.FirstName) < 1 {
        user.validationErrors = append(user.validationErrors, "User first name is empty")
    }

    if len(user.LastName) < 1 {
        user.validationErrors = append(user.validationErrors, "User last name is empty")
    }

    if len(user.Email) < 3 || ! common.ValidateEmail(user.Email) {
        user.validationErrors = append(user.validationErrors, "User email not valid")
    }

    if len(user.MobilePhone) > 3 && ! common.ValidateMobile(user.MobilePhone) {
        user.validationErrors = append(user.validationErrors, "User mobile phone should be valid or empty. Format +0123456789... ")
    }

    //Validate remove this line for disable generator functionality

}
`

const dbmodelsRole = `package dbmodels

import (
	"time"
)

// Роли пользователей (менеджер компании, менеджер покупателя, админ...)
type Role struct {
    ID        int       `+"`"+`gorm:"primary_key"`+"`"+`
    Name        string
    Description string `+"`"+`gorm:"type:text"`+"`"+`
    //Role remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `+"`"+`sql:"index" json:"-"`+"`"+`

    validator
}

func (role *Role) Validate() (isValid bool) {

	if len(role.Name) < 1 {
		isValid = false
		role.validationErrors = append(role.validationErrors, "Role name is too short")
	}

	if len(role.Description) < 5 {
		isValid = false
		role.validationErrors = append(role.validationErrors, "Role description is too short")
	}

	return
}
`

const dbmodelsUserRole = `package dbmodels

import (
	"time"
)

type UserRole struct {
	ID     int `+"`"+`gorm:"primary_key"`+"`"+`

	RoleID int
	UserID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `+"`"+`sql:"index" json:"-"`+"`"+`

	validator
}

func (userRole *UserRole) Validate() (isValid bool) {

	if userRole.RoleID < 1 {
		isValid = false
		userRole.validationErrors = append(userRole.validationErrors, "RoleId is wrong")
	}

	if userRole.UserID < 1 {
		isValid = false
		userRole.validationErrors = append(userRole.validationErrors, "UserId is wrong")
	}
	
	return
}

`