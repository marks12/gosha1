package dbmodels

import (
    "time"
)

type ResourceType struct {

    ID        int       `gorm:"primary_key"`
    Name string
	//ResourceType remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (resourceType *ResourceType) Validate() {
    //Validate remove this line for disable generator functionality
}

