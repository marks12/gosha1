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

	core.Db.Model(dbmodels.User{}).Count(&count)

	if count < 1 {

		user := logic.AssignUserDbFromType(types.User{
			Id:          0,
			Email:       "tsv@serveon.ru",
			FirstName:   "Superuser",
			IsActive:    true,
			LastName:    "Admin",
			MobilePhone: "",
			Password:    "virq3t4",
			Token:       "",
		})
		core.Db.Model(dbmodels.User{}).Save(&user)

		role := dbmodels.Role{
			Name:        "Admin",
			Description: "Administrator",
		}
		core.Db.Model(dbmodels.Role{}).Save(&role)

		userRole := dbmodels.UserRole{
			UserId:    user.ID,
			RoleId:    role.ID,
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
			RoleId:       role.ID,
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

    AddResource()
}

func AddResource() {
    for _, route := range settings.RoutesArray {
        dbModel := dbmodels.Resource{
            Name:   strings.Split(route, "/")[3],
            Code:   route,
            TypeId: 1,
        }
        core.Db.Where(dbModel).FirstOrCreate(&dbModel)
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

`

var msTemplateInsertDataToDb = template{
    Path:    "./bootstrap/insert_data_to_db.go",
    Content: assignMsName(msInsertDataToDb),
}

