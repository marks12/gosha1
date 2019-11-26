package dbmodels

import (
    "time"
)

type ComponentGroup struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Code string
	//ComponentGroup remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (componentGroup *ComponentGroup) Validate() {
    //Validate remove this line for disable generator functionality
}

