package dbmodels

import (
    "time"
)

type Layout struct {

    ID        int       `gorm:"primary_key"`
    Name string
	RegionId int
	LanguageId int
	//Layout remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (layout *Layout) Validate() {
    //Validate remove this line for disable generator functionality
}

