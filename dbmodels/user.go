package dbmodels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {

	ID        int       `gorm:"primary_key"`
	Name      string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`
	validator
}

func (user *User) Validate()  {

}
