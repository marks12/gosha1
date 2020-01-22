package cmd

const msInsertDataToDb = `package bootstrap

import (
	"{ms-name}/dbmodels"
	"{ms-name}/types"
	"fmt"
	"os"
	"{ms-name}/core"
	"{ms-name}/logic"
	"{ms-name}/settings"
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
		)

		fmt.Println("All tables removed")
		os.Exit(1)
	}

	core.Db.AutoMigrate(
        //generator insert entity
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
			Email:       "{email}",
			FirstName:   "Superuser",
			IsActive:    true,
			LastName:    "Admin",
			MobilePhone: "",
			Password:    "{password}",
		})
		core.Db.Model(dbmodels.User{}).Save(&user)


		userRole := dbmodels.UserRole{
			UserId:    user.ID,
			RoleId:    adminRole.ID,
		}
		core.Db.Model(dbmodels.UserRole{}).Save(&userRole)

		resourceType := logic.AssignResourceTypeDbFromType(types.ResourceType{
			Id: settings.HttpRouteResourceType,
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

    for _, route := range settings.ExtResources {
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

`

var msTemplateInsertDataToDb = template{
    Path:    "./bootstrap/insert_data_to_db.go",
    Content: assignMsName(msInsertDataToDb),
}

