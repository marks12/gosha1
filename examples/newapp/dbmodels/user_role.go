package dbmodels

import (
    "time"
)

type UserRole struct {

    ID        int       `gorm:"primary_key"`
    UserId int
	RoleId int
	//UserRole remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (userRole *UserRole) Validate() {
    //Validate remove this line for disable generator functionality
}

