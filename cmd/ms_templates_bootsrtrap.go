package cmd

const msInsertDataToDb = `package bootstrap

import (
	"{ms-name}/dbmodels"
	"{ms-name}/types"
	"{ms-name}/core"
	"{ms-name}/logic"
	"{ms-name}/settings"
	"strings"
	"fmt"
	"os"
	"errors"
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
	addRouteType()
	addUser()

}

func addRouteType() {

	resourceType := logic.AssignResourceTypeDbFromType(types.ResourceType{
		Id: settings.HttpRouteResourceType,
		Name: "Route",
	})
	core.Db.Model(dbmodels.ResourceType{}).FirstOrCreate(&resourceType)
}

func addUser() {

	var count int

	adminRole := dbmodels.Role{
		ID:			 settings.AdminRoleId,
		Name:        "Admin",
		Description: "Administrator",
	}
	core.Db.Where(adminRole).FirstOrCreate(&adminRole)

	userRole := dbmodels.Role{
		ID:			 settings.UserRoleId,
		Name:        "User",
		Description: "Application user",
	}
	core.Db.Where(userRole).FirstOrCreate(&userRole)
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
		core.Db.Model(dbmodels.User{}).FirstOrCreate(&user)

		setRole(user.ID, settings.AdminRoleId)
		setRole(user.ID, settings.UserRoleId)
	}

    AddAdminResources(adminRole.ID)
	AddUserResources(userRole.ID)
}


func setRole(userId int, roleId int) {

	userRole := dbmodels.UserRole{
		UserId:    userId,
		RoleId:    roleId,
	}
	core.Db.Model(dbmodels.UserRole{}).FirstOrCreate(&userRole, "user_id = ? AND role_id = ?", userId, roleId)
}

func AddAdminResources(adminRoleId int) {

    // ADD SPECIAL ADMIN ROLES HERE. Otherwise admin roles will take full access to routes  

    for _, route := range settings.RoutesArray {
		err := setRoleAccess(adminRoleId, route, types.Access{
			Find:         true,
			Read:         true,
			Create:       true,
			Update:       true,
			Delete:       true,
			FindOrCreate: true,
		})

		if err != nil {
			fmt.Println(err)
		}
    }

    for _, route := range settings.ExtResources {
        name := route
        if strings.Count(name, "/") > 2 {
            name = strings.ToLower(strings.Split(route, "/")[3])
        }
        dbModel := dbmodels.Resource{
            Name:   name,
            Code:   route,
            TypeId: 1,
        }
        core.Db.Where(dbModel).FirstOrCreate(&dbModel)
    }
}

func AddUserResources(userRoleId int) {

	// access user to me route
	err := setRoleAccess(userRoleId, settings.CurrentUserRoute, types.Access{
		Read:			true,
	})

	if err != nil {
		fmt.Println("Cannot create user resources access CurrentUserRoute")
	}
}

func setRoleAccess(roleId int, route string, access types.Access) error {

	strArr := strings.Split(route, "/")

	if len(strArr) > 2 {
		dbModel := dbmodels.Resource{
			Name:   strArr[3],
			Code:   route,
			TypeId: 1,
		}

		core.Db.Where(dbModel).FirstOrCreate(&dbModel)

		roleResource := logic.AssignRoleResourceDbFromType(types.RoleResource{
			RoleId:       roleId,
			ResourceId:   dbModel.ID,
			Find:         access.Find,
			Read:         access.Read,
			Create:       access.Create,
			Update:       access.Update,
			Delete:       access.Delete,
			FindOrCreate: access.FindOrCreate,
		})
		core.Db.Model(dbmodels.RoleResource{}).FirstOrCreate(&roleResource, "role_id = ? AND resource_id = ?", roleId, dbModel.ID)

		return nil
	}

	return errors.New("Wrong route length. Cant set access for route: " + route)
}
`

var msTemplateInsertDataToDb = template{
    Path:    "./bootstrap/insert_data_to_db.go",
    Content: assignMsName(msInsertDataToDb),
}

