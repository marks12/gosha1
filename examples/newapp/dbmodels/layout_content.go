package dbmodels

import (
    "time"
)

type LayoutContent struct {

    ID        int       `gorm:"primary_key"`
    LayoutId *int
	ComponentTemplateCode string
	Position int
	IsActive bool
	RegionId *int
	LanguageId *int
	Name string
	//LayoutContent remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (layoutContent *LayoutContent) Validate() {
    //Validate remove this line for disable generator functionality
}

