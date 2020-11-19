package dbmodels

import (
	"time"
)

// User entity example
type AclUser struct {
	ID		int	`gorm:"AUTO_INCREMENT"`
	FirstName	string
	LastName	string
	Email		string
	Active		bool
	ConfirmedAt	time.Time	`sql:"type:DATETIME"`
	RealtorId	int
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:"index" json:"-"`
	validator
	SecName		string
	SecNameLast	string
	Birthday	time.Time
	Somedate	time.Time
	SomeElseDate	time.Time
	DateForMe	time.Time
}

func (aclUser *AclUser) Validate() {
}
