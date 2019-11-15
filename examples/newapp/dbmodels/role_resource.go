package dbmodels

import (
    "time"
)

type RoleResource struct {

    ID        int       `gorm:"primary_key"`
    RoleId int
	ResourceId int
	Find bool
	Read bool
	Create bool
	Update bool
	Delete bool
	FindOrCreate bool
	//RoleResource remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (roleResource *RoleResource) Validate() {
    //Validate remove this line for disable generator functionality
}

