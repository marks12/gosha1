package tests

import (
	"newapp/core"
	"newapp/dbmodels"
	"fmt"
	"os"
)

type userType struct {
	RoleId int
	IsActive bool
}

func getAuth(uType userType) dbmodels.Auth {

	var auth []dbmodels.Auth

	core.Db.Model(dbmodels.Auth{}).
		Joins("INNER join users on auths.user_id = users.id and auths.is_active = true").
		Joins("INNER join user_roles on user_roles.user_id = users.id and user_roles.role_id = ? and users.is_active = ?", uType.RoleId, uType.IsActive).
		Where(dbmodels.Auth{IsActive: true}).
		Limit(1).
		Find(&auth)

	if len(auth) < 1 {
		fmt.Println("User with valid token in auth not found in database with roleId ", uType.RoleId, " and active ", uType.IsActive)
		os.Exit(1)
	}

	return auth[0]
}

func getRole(name string) dbmodels.Role {

	var role dbmodels.Role

	critery := core.Db.Model(dbmodels.Role{}).Where(dbmodels.Role{
		Name:    name,
	})

	critery.First(&role)

	return role
}

func GetAdminAuth(IsActive bool) dbmodels.Auth {

	role := getRole("Admin")

	t := userType{
		RoleId:   role.ID,
		IsActive: IsActive,
	}

	return getAuth(t)
}

func GetNonAdminUser(IsActive bool) dbmodels.Auth {

	role := getRole("User")

	t := userType{
		RoleId:   role.ID,
		IsActive: IsActive,
	}

	return getAuth(t)
}
