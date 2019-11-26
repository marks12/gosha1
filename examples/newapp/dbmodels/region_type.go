package dbmodels

import (
    "time"
)

type RegionType struct {

    ID        int       `gorm:"primary_key"`
    Name string
	//RegionType remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (regionType *RegionType) Validate() {
    //Validate remove this line for disable generator functionality
}

