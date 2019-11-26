package dbmodels

import (
    "time"
)

type PageInfo struct {

    ID        int       `gorm:"primary_key"`
    SeoMeta *int
	//PageInfo remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (pageInfo *PageInfo) Validate() {
    //Validate remove this line for disable generator functionality
}

