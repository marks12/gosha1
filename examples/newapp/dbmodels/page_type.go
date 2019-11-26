package dbmodels

import (
    "time"
)

type PageType struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Code string
	//PageType remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (pageType *PageType) Validate() {
    //Validate remove this line for disable generator functionality
}

