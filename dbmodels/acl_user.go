package dbmodels

import (
    "time"
)

type AclUser struct {

    ID          int       `gorm:"AUTO_INCREMENT"`
    FirstName   string
    LastName    string
    Email       string
    Active      bool
    ConfirmedAt time.Time `sql:"type:DATETIME"`
    RealtorId   int

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index" json:"-"`

    validator
}

func (aclUser *AclUser) Validate()  {

}

