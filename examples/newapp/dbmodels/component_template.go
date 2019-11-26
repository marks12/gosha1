package dbmodels

import (
    "time"
)

type ComponentTemplate struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Code string
	Path string
	GroupCode string
	GroupId *int
	//ComponentTemplate remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (componentTemplate *ComponentTemplate) Validate() {
    //Validate remove this line for disable generator functionality
}

