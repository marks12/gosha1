package dbmodels

import (
    "time"
)

type PageTemplate struct {

    ID        int       `gorm:"primary_key"`
    PageTypeId *int
	Name string
	IsActive bool
	RegionId *int
	LanguageId *int
	RootPageId *int
	LayoutId *int
	//PageTemplate remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (pageTemplate *PageTemplate) Validate() {
    //Validate remove this line for disable generator functionality
}

