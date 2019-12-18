package bootstrap

import (
	"newapp/dbmodels"
	"newapp/types"
	"fmt"
	"os"
	"newapp/core"
	"newapp/logic"
	"newapp/settings"
	"strings"
)

func FillDBTestData()  {

	if core.DbErr != nil {
		fmt.Println("Error dabatabse connect", core.DbErr.Error())
		os.Exit(0)
	}

	isDropTables := false

	if (len(os.Args) > 1 && os.Args[1] == "drop") ||
		(len(os.Args) > 2 && os.Args[2] == "drop") {
		isDropTables = true
	}

	if isDropTables == true {

		core.Db.DropTableIfExists(
          //generator insert entity
          &dbmodels.Entity{},
          &dbmodels.CurrentUser{},
          &dbmodels.Auth{},
          &dbmodels.UserRole{},
          &dbmodels.ResourceType{},
          &dbmodels.Resource{},
          &dbmodels.RoleResource{},
          &dbmodels.Role{},
          &dbmodels.User{},
          &dbmodels.PageContent{},
          &dbmodels.PageInfo{},
          &dbmodels.PageTemplate{},
          &dbmodels.PageType{},
          &dbmodels.Layout{},
          &dbmodels.LayoutContent{},
          &dbmodels.RegionType{},
			&dbmodels.Region{},
			&dbmodels.ComponentTemplate{},
			&dbmodels.ComponentGroup{},
		)

		fmt.Println("All tables removed")
		os.Exit(1)
	}

	core.Db.AutoMigrate(
        //generator insert entity
          &dbmodels.Entity{},
          &dbmodels.CurrentUser{},
          &dbmodels.Auth{},
          &dbmodels.UserRole{},
          &dbmodels.ResourceType{},
          &dbmodels.Resource{},
          &dbmodels.RoleResource{},
          &dbmodels.Role{},
			&dbmodels.User{},
			&dbmodels.PageContent{},
			&dbmodels.PageInfo{},
			&dbmodels.PageTemplate{},
			&dbmodels.PageType{},
			&dbmodels.Layout{},
			&dbmodels.LayoutContent{},
			&dbmodels.RegionType{},
			&dbmodels.Region{},
		&dbmodels.ComponentTemplate{},
		&dbmodels.ComponentGroup{},

	)

	// add fixtures
	addUser()

}

func addUser() {

	var count int

	adminRole := dbmodels.Role{
		Name:        "Admin",
		Description: "Administrator",
	}
	core.Db.Where(adminRole).FirstOrCreate(&adminRole)

	core.Db.Model(dbmodels.User{}).Count(&count)
	if count < 1 {

		user := logic.AssignUserDbFromType(types.User{
			Id:          0,
			Email:       "123@140140.ru",
			FirstName:   "Superuser",
			IsActive:    true,
			LastName:    "Admin",
			MobilePhone: "",
			Password:    "123123",
			Token:       "",
		})
		core.Db.Model(dbmodels.User{}).Save(&user)

		userRole := dbmodels.UserRole{
			UserId:    user.ID,
			RoleId:    adminRole.ID,
		}
		core.Db.Model(dbmodels.UserRole{}).Save(&userRole)

		resourceType := logic.AssignResourceTypeDbFromType(types.ResourceType{
			Name: "Route",
		})
		core.Db.Model(dbmodels.ResourceType{}).Save(&resourceType)

		resource := logic.AssignResourceDbFromType(types.Resource{
			Name:   "User",
			Code:   "/api/v1/user",
			TypeId: resourceType.ID,
		})
		core.Db.Model(dbmodels.Resource{}).Save(&resource)

		roleResource := logic.AssignRoleResourceDbFromType(types.RoleResource{
			RoleId:       adminRole.ID,
			ResourceId:   resource.ID,
			Find:         true,
			Read:         true,
			Create:       true,
			Update:       true,
			Delete:       true,
			FindOrCreate: false,
		})
		core.Db.Model(dbmodels.RoleResource{}).Save(&roleResource)
	}

    AddResource(adminRole.ID)
}

func AddResource(adminRoleId int) {

    for _, route := range settings.RoutesArray {

    	strArr := strings.Split(route, "/")

    	if len(strArr) > 3 {

			dbModel := dbmodels.Resource{
				Name:   strArr[3],
				Code:   route,
				TypeId: 1,
			}
			core.Db.Where(dbModel).FirstOrCreate(&dbModel)

			roleResource := logic.AssignRoleResourceDbFromType(types.RoleResource{
				RoleId:       adminRoleId,
				ResourceId:   dbModel.ID,
				Find:         true,
				Read:         true,
				Create:       true,
				Update:       true,
				Delete:       true,
				FindOrCreate: true,
			})
			core.Db.Model(dbmodels.RoleResource{}).Save(&roleResource)
		}
    }

    for _, route := range settings.Resources {
        name := route
        if strings.Count(name, "/") > 2 {
            name = strings.Split(route, "/")[3]
        }
        dbModel := dbmodels.Resource{
            Name:   name,
            Code:   route,
            TypeId: 1,
        }
        core.Db.Where(dbModel).FirstOrCreate(&dbModel)
    }

}

