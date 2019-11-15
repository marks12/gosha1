package dbmodels

import (
	"newapp/common"
    "time"
)

type Auth struct {

    ID        int       `gorm:"primary_key"`
    Email     string
    Password  string
    Token     string
    IsActive bool
    UserId   int
    //Auth remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (auth *Auth) Validate() {
    if len(auth.Email) < 3 || ! common.ValidateEmail(auth.Email)  {
        auth.validationErrors = append(auth.validationErrors, "User email not valid")
    }

    if len(auth.Password) < 1 {
        auth.validationErrors = append(auth.validationErrors, "User password is empty")
    }

    //Validate remove this line for disable generator functionality
}

