package logic

import (
	"newapp/dbmodels"
	"golang.org/x/crypto/bcrypt"
	"newapp/types"
	"newapp/settings"
)










// add all assign functions

func AssignEntityTypeFromDb(dbEntity dbmodels.Entity) types.Entity {

    //AssignEntityTypeFromDb predefine remove this line for disable generator functionality

    return types.Entity{
        Id: dbEntity.ID,
        Name: dbEntity.Name,
		Description: dbEntity.Description,
		//AssignEntityTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignEntityDbFromType(typeModel types.Entity) dbmodels.Entity {

    //AssignEntityDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.Entity{
        ID: typeModel.Id,
        Name: typeModel.Name,
		Description: typeModel.Description,
		//AssignEntityDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignCurrentUserTypeFromDb(dbCurrentUser dbmodels.CurrentUser) types.CurrentUser {

    //AssignCurrentUserTypeFromDb predefine remove this line for disable generator functionality

    return types.CurrentUser{
        Id: dbCurrentUser.ID,
        //AssignCurrentUserTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignCurrentUserDbFromType(typeModel types.CurrentUser) dbmodels.CurrentUser {

    //AssignCurrentUserDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.CurrentUser{
        ID: typeModel.Id,
        //AssignCurrentUserDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignAuthTypeFromDb(dbAuth dbmodels.Auth) types.Auth {

    //AssignAuthTypeFromDb predefine remove this line for disable generator functionality

    return types.Auth{
        
        //AssignAuthTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignAuthDbFromType(typeModel types.Auth) dbmodels.Auth {

    //AssignAuthDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.Auth{
        
        Email: typeModel.Email,
        Password: typeModel.Password,
		//AssignAuthDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignUserRoleTypeFromDb(dbUserRole dbmodels.UserRole) types.UserRole {

    //AssignUserRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.UserRole{
        Id: dbUserRole.ID,
        UserId: dbUserRole.UserId,
		RoleId: dbUserRole.RoleId,
		//AssignUserRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignUserRoleDbFromType(typeModel types.UserRole) dbmodels.UserRole {

    //AssignUserRoleDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.UserRole{
        ID: typeModel.Id,
        UserId: typeModel.UserId,
		RoleId: typeModel.RoleId,
		//AssignUserRoleDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignResourceTypeTypeFromDb(dbResourceType dbmodels.ResourceType) types.ResourceType {

    //AssignResourceTypeTypeFromDb predefine remove this line for disable generator functionality

    return types.ResourceType{
        Id: dbResourceType.ID,
        Name: dbResourceType.Name,
		//AssignResourceTypeTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignResourceTypeDbFromType(typeModel types.ResourceType) dbmodels.ResourceType {

    //AssignResourceTypeDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.ResourceType{
        ID: typeModel.Id,
        Name: typeModel.Name,
		//AssignResourceTypeDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignResourceTypeFromDb(dbResource dbmodels.Resource) types.Resource {

    //AssignResourceTypeFromDb predefine remove this line for disable generator functionality

    return types.Resource{
        Id: dbResource.ID,
        Name: dbResource.Name,
		Code: dbResource.Code,
		TypeId: dbResource.TypeId,
		//AssignResourceTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignResourceDbFromType(typeModel types.Resource) dbmodels.Resource {

    //AssignResourceDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.Resource{
        ID: typeModel.Id,
        Name: typeModel.Name,
		Code: typeModel.Code,
		TypeId: typeModel.TypeId,
		//AssignResourceDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignRoleResourceTypeFromDb(dbRoleResource dbmodels.RoleResource) types.RoleResource {

    //AssignRoleResourceTypeFromDb predefine remove this line for disable generator functionality

    return types.RoleResource{
        Id: dbRoleResource.ID,
        RoleId: dbRoleResource.RoleId,
		ResourceId: dbRoleResource.ResourceId,
		Find: dbRoleResource.Find,
		Read: dbRoleResource.Read,
		Create: dbRoleResource.Create,
		Update: dbRoleResource.Update,
		Delete: dbRoleResource.Delete,
		FindOrCreate: dbRoleResource.FindOrCreate,
		//AssignRoleResourceTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignRoleResourceDbFromType(typeModel types.RoleResource) dbmodels.RoleResource {

    //AssignRoleResourceDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.RoleResource{
        ID: typeModel.Id,
        RoleId: typeModel.RoleId,
		ResourceId: typeModel.ResourceId,
		Find: typeModel.Find,
		Read: typeModel.Read,
		Create: typeModel.Create,
		Update: typeModel.Update,
		Delete: typeModel.Delete,
		FindOrCreate: typeModel.FindOrCreate,
		//AssignRoleResourceDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignRoleTypeFromDb(dbRole dbmodels.Role) types.Role {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.Role{
        Id: dbRole.ID,
        Name: dbRole.Name,
		Description: dbRole.Description,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignRoleDbFromType(typeModel types.Role) dbmodels.Role {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.Role{
        ID: typeModel.Id,
        Name: typeModel.Name,
		Description: typeModel.Description,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}



func AssignUserTypeFromDb(dbUser dbmodels.User) types.User {

    //AssignUserTypeFromDb predefine remove this line for disable generator functionality

    return types.User{
        Id: dbUser.ID,
        FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		MobilePhone: dbUser.MobilePhone,
		Email:       dbUser.Email,
		Password:    "*****",
		IsActive:    dbUser.IsActive,
		Token:       dbUser.Token,
    	//AssignUserTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignUserDbFromType(typeModel types.User) dbmodels.User {

    password := []byte(typeModel.Password + settings.PASSWORD_SALT)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    //AssignUserDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.User{
        ID: typeModel.Id,
        FirstName:   typeModel.FirstName,
		LastName:    typeModel.LastName,
		MobilePhone: typeModel.MobilePhone,
		Email:       typeModel.Email,
		Password:    string(hashedPassword),
		IsActive:    typeModel.IsActive,
	    //AssignUserDbFromType.Field remove this line for disable generator functionality
    }
}


