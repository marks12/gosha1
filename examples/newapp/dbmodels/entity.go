package dbmodels

import (
    "time"
)

type Entity struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Description string
	//Entity remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (entity *Entity) Validate() {
    //Validate remove this line for disable generator functionality
}

