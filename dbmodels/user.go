package dbmodels

import (
"time"
)

type User struct {

	ID        int       `gorm:"primary_key"`
	Name      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
	validator
}

func (user *User) Validate()  {

}
