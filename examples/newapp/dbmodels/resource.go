package dbmodels

import (
    "time"
)

type Resource struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Code string
	TypeId int
	//Resource remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (resource *Resource) Validate() {
    //Validate remove this line for disable generator functionality
}

