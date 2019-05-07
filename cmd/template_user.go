package cmd

const dbmodelsAuth = `package dbmodels

type Auth struct {

    Email     string
    Password  string  `+"`"+`json:"-"`+"`"+` 
    Token     string

    validator
}

func validateEmail(email string) bool {
    Re := regexp.MustCompile(`+"`"+`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`+"`"+`)
    return Re.MatchString(email)
}

func ValidateMobile(phone string) bool {
    Re := regexp.MustCompile(`+"`"+`^[+][0-9]{11,}`+"`"+`)
    return Re.MatchString(phone)
}

func (auth *Auth) Validate()  {

    if len(auth.Email) < 3 || ! validateEmail(auth.Email)  {
        auth.validationErrors = append(auth.validationErrors, "User email not valid")
    }

    if len(auth.Password) < 1 {
        auth.validationErrors = append(auth.validationErrors, "User password is empty")
    }
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
    Password    string  `+"`"+`json:"-"`+"`"+`
    Token       string

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `+"`"+`sql:"index" json:"-"`+"`"+`

    validator
}

func (user *User) Validate() (isValid bool) {

    isValid = true

    if len(user.FirstName) < 1 {
        isValid = false
        user.validationErrors = append(user.validationErrors, "User first name is empty")
    }

    if len(user.LastName) < 1 {
        isValid = false
        user.validationErrors = append(user.validationErrors, "User last name is empty")
    }

    if len(user.Email) < 3 || ! ValidateEmail(user.Email)  {
        isValid = false
        user.validationErrors = append(user.validationErrors, "User email not valid")
    }

    if len(user.MobilePhone) > 3 &&  ! ValidateMobile(user.MobilePhone)  {
        isValid = false
        user.validationErrors = append(user.validationErrors, "User mobile phone should be valid or empty. Format +0123456789... ")
    }

    return
}
`

const dbmodelsRole = `package dbmodels

import (
	"time"
)

// Роли пользователей (менеджер компании, менеджер покупателя, админ...)
type Role struct {
	ID          int `+"`"+`gorm:"primary_key"`+"`"+`

	Description string
	Name        string

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