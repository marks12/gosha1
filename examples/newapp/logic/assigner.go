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



func AssignComponentGroupTypeFromDb(dbRole dbmodels.ComponentGroup) types.ComponentGroup {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.ComponentGroup{
        Id: dbRole.ID,
        Name: dbRole.Name,
		Code: dbRole.Code,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignComponentGroupDbFromType(typeModel types.ComponentGroup) dbmodels.ComponentGroup {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.ComponentGroup{
        ID: typeModel.Id,
        Name: typeModel.Name,
		Code: typeModel.Code,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignComponentTemplateTypeFromDb(dbRole dbmodels.ComponentTemplate) types.ComponentTemplate {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.ComponentTemplate{
        Id: dbRole.ID,
        Name: dbRole.Name,
		Code: dbRole.Code,
		Path: dbRole.Path,
		GroupCode: dbRole.GroupCode,
		GroupId: dbRole.GroupId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignComponentTemplateDbFromType(typeModel types.ComponentTemplate) dbmodels.ComponentTemplate {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.ComponentTemplate{
        ID: typeModel.Id,
        Name: typeModel.Name,
		Code: typeModel.Code,
		Path: typeModel.Path,
		GroupCode: typeModel.GroupCode,
		GroupId: typeModel.GroupId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}

func AssignLayoutTypeFromDb(dbRole dbmodels.Layout) types.Layout {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.Layout{
        Id: dbRole.ID,
        Name: dbRole.Name,
        RegionId: dbRole.RegionId,
        LanguageId: dbRole.LanguageId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignLayoutDbFromType(typeModel types.Layout) dbmodels.Layout {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.Layout{
        ID: typeModel.Id,
        Name: typeModel.Name,
        RegionId: typeModel.RegionId,
        LanguageId: typeModel.LanguageId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}

func AssignLayoutContentTypeFromDb(dbRole dbmodels.LayoutContent) types.LayoutContent {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.LayoutContent{
        Id: dbRole.ID,
        LayoutId: dbRole.LayoutId,
        ComponentTemplateCode: dbRole.ComponentTemplateCode,
        Position: dbRole.Position,
        IsActive: dbRole.IsActive,
        Name: dbRole.Name,
        RegionId: dbRole.RegionId,
        LanguageId: dbRole.LanguageId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignLayoutContentDbFromType(typeModel types.LayoutContent) dbmodels.LayoutContent {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.LayoutContent{
        ID: typeModel.Id,
		LayoutId: typeModel.LayoutId,
		ComponentTemplateCode: typeModel.ComponentTemplateCode,
		Position: typeModel.Position,
		IsActive: typeModel.IsActive,
		Name: typeModel.Name,
		RegionId: typeModel.RegionId,
		LanguageId: typeModel.LanguageId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignPageContentTypeFromDb(dbRole dbmodels.PageContent) types.PageContent {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.PageContent{
        Id: dbRole.ID,
		Name: dbRole.Name,
		PageTemplateId: dbRole.PageTemplateId,
        ComponentTemplateCode: dbRole.ComponentTemplateCode,
        Position: dbRole.Position,
        IsActive: dbRole.IsActive,
        RegionId: dbRole.RegionId,
        LanguageId: dbRole.LanguageId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignPageContentDbFromType(typeModel types.PageContent) dbmodels.PageContent {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.PageContent{
        ID: typeModel.Id,
		Name: typeModel.Name,
		PageTemplateId: typeModel.PageTemplateId,
		ComponentTemplateCode: typeModel.ComponentTemplateCode,
		Position: typeModel.Position,
		IsActive: typeModel.IsActive,
		RegionId: typeModel.RegionId,
		LanguageId: typeModel.LanguageId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignPageInfoTypeFromDb(dbRole dbmodels.PageInfo) types.PageInfo {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.PageInfo{
        Id: dbRole.ID,
        SeoMeta: dbRole.SeoMeta,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignPageInfoDbFromType(typeModel types.PageInfo) dbmodels.PageInfo {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.PageInfo{
        ID: typeModel.Id,
		SeoMeta: typeModel.SeoMeta,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignPageTemplateTypeFromDb(dbRole dbmodels.PageTemplate) types.PageTemplate {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.PageTemplate{
        Id: dbRole.ID,
        PageTypeId: dbRole.PageTypeId,
        Name: dbRole.Name,
        IsActive: dbRole.IsActive,
        RegionId: dbRole.RegionId,
        LanguageId: dbRole.LanguageId,
        RootPageId: dbRole.RootPageId,
        LayoutId: dbRole.LayoutId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignPageTemplateDbFromType(typeModel types.PageTemplate) dbmodels.PageTemplate {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.PageTemplate{
        ID: typeModel.Id,
		PageTypeId: typeModel.PageTypeId,
		Name: typeModel.Name,
		IsActive: typeModel.IsActive,
		RegionId: typeModel.RegionId,
		LanguageId: typeModel.LanguageId,
		RootPageId: typeModel.RootPageId,
		LayoutId: typeModel.LayoutId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}

func AssignPageTypeTypeFromDb(dbRole dbmodels.PageType) types.PageType {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.PageType{
        Id: dbRole.ID,
        Name: dbRole.Name,
        Code: dbRole.Code,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignPageTypeDbFromType(typeModel types.PageType) dbmodels.PageType {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.PageType{
        ID: typeModel.Id,
		Name: typeModel.Name,
		Code: typeModel.Code,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignRegionTypeFromDb(dbRole dbmodels.Region) types.Region {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.Region{
        Id: dbRole.ID,
        Name: dbRole.Name,
        Code: dbRole.Code,
        TypeId: dbRole.TypeId,
        EntityId: dbRole.EntityId,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignRegionDbFromType(typeModel types.Region) dbmodels.Region {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.Region{
        ID: typeModel.Id,
		Name: typeModel.Name,
		Code: typeModel.Code,
		TypeId: typeModel.TypeId,
		EntityId: typeModel.EntityId,
		//AssignRoleDbFromType.Field remove this line for disable generator functionality
    }
}


func AssignRegionTypeTypeFromDb(dbRole dbmodels.RegionType) types.RegionType {

    //AssignRoleTypeFromDb predefine remove this line for disable generator functionality

    return types.RegionType{
        Id: dbRole.ID,
        Name: dbRole.Name,
		//AssignRoleTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignRegionTypeDbFromTp(typeModel types.RegionType) dbmodels.RegionType {

    //AssignRoleDbFromType predefine remove this line for disable generator functionality

    return dbmodels.RegionType{
        ID: typeModel.Id,
		Name: typeModel.Name,
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


