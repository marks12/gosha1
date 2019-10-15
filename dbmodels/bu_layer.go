package dbmodels

import (
    "time"
)

type BuLayer struct {

    ID        int       `gorm:"primary_key"`
    //BuLayer remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (buLayer *BuLayer) Validate() {
    //Validate remove this line for disable generator functionality
}

