package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func RoleFind(filter types.RoleFilter)  (result []types.Role, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.Role{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.Role{})

    if len(filterIds) > 0 {
        criteria = criteria.Where("id in (?)", filterIds)
    }

    //if len(filter.Search) > 0 {
    //
    //    s := ("%" + filter.Search + "%")
    //
    //    if len(filter.SearchBy) > 0 {
    //
    //        for _, field := range filter.SearchBy {
    //
    //            if core.Db.NewScope(&dbmodels.Role{}).HasColumn(field) {
    //                criteria = criteria.Where("`"+field+"`"+" like ?", s)
    //            } else {
    //                err = errors.New("Search by unknown field " + field)
    //                return
    //            }
    //        }
    //    }
    //} else {
    //    criteria = criteria.Where("name like ? or code like ?", ("%" + filter.Search + "%"), ("%" + filter.Search + "%"))
    //}

    q := criteria.Model(dbmodels.Role{}).Count(&count)

    if q.Error != nil {
       log.Println("FindRole > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.Role{}).HasColumn(Field) {
                criteria = criteria.Order("`" + Field + "`" + " " + filter.OrderDirection[index])
            } else {
                err = errors.New("Ordering by unknown Field " + Field)
                return
            }
        }
    }

    q = criteria.Limit(limit).Offset(offset).Find(&dbmodelData)

    if q.Error != nil {
       log.Println("FindProduct > Ошибка получения данных2:", q.Error)
       return []types.Role{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignRoleTypeFromDb(item))
    }

    return result, count, nil
}

func RoleCreate(filter types.RoleFilter)  (data types.Role, err error) {

    query := core.Db

    typeModel := filter.GetRoleModel()
    dbModel := AssignRoleDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("RoleCreate > Create Role error:", dbModel)
        return types.Role{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("RoleCreate > Create Role error:", query.Error)
        return types.Role{}, errors.New("cant create Role")
    }

    return AssignRoleTypeFromDb(dbModel), nil
}

func RoleRead(filter types.RoleFilter)  (data types.Role, err error) {

    findData, _, err := RoleFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.Role{}, errors.New("Not found")
}

func RoleUpdate(filter types.RoleFilter)  (data types.Role, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RoleRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Role not found in db")
        }
        return
    }

    newModel := filter.GetRoleModel()

    updateModel := AssignRoleDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.Description = newModel.Description
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Role{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRoleTypeFromDb(updateModel)
    return
}


func RoleDelete(filter types.RoleFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RoleRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Role not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.Role{}).Where(dbmodels.Role{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func RoleFindOrCreate(filter types.RoleFilter)  (data types.Role, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignRoleDbFromType(filter.GetRoleModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Role{}).Where(dbmodels.Role{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRoleTypeFromDb(findOrCreateModel)
    return
}

