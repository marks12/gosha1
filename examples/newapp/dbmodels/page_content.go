package dbmodels

import (
    "time"
)

type PageContent struct {

    ID        int       `gorm:"primary_key"`
    Name string
	PageTemplateId *int
	IsActive bool
	RegionId *int
	LanguageId *int
	Position int
	ComponentTemplateCode string
	//PageContent remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (pageContent *PageContent) Validate() {
    //Validate remove this line for disable generator functionality
}

