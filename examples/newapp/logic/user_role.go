package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func UserRoleFind(filter types.UserRoleFilter)  (result []types.UserRole, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.UserRole{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.UserRole{})

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
    //            if core.Db.NewScope(&dbmodels.UserRole{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.UserRole{}).Count(&count)

    if q.Error != nil {
       log.Println("FindUserRole > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.UserRole{}).HasColumn(Field) {
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
       return []types.UserRole{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignUserRoleTypeFromDb(item))
    }

    return result, count, nil
}

func UserRoleCreate(filter types.UserRoleFilter)  (data types.UserRole, err error) {

    query := core.Db

    typeModel := filter.GetUserRoleModel()
    dbModel := AssignUserRoleDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("UserRoleCreate > Create UserRole error:", dbModel)
        return types.UserRole{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("UserRoleCreate > Create UserRole error:", query.Error)
        return types.UserRole{}, errors.New("cant create UserRole")
    }

    return AssignUserRoleTypeFromDb(dbModel), nil
}

func UserRoleRead(filter types.UserRoleFilter)  (data types.UserRole, err error) {

    findData, _, err := UserRoleFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.UserRole{}, errors.New("Not found")
}

func UserRoleUpdate(filter types.UserRoleFilter)  (data types.UserRole, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := UserRoleRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("UserRole not found in db")
        }
        return
    }

    newModel := filter.GetUserRoleModel()

    updateModel := AssignUserRoleDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.UserId = newModel.UserId
	updateModel.RoleId = newModel.RoleId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.UserRole{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignUserRoleTypeFromDb(updateModel)
    return
}


func UserRoleDelete(filter types.UserRoleFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := UserRoleRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("UserRole not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.UserRole{}).Where(dbmodels.UserRole{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func UserRoleFindOrCreate(filter types.UserRoleFilter)  (data types.UserRole, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignUserRoleDbFromType(filter.GetUserRoleModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.UserRole{}).Where(dbmodels.UserRole{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignUserRoleTypeFromDb(findOrCreateModel)
    return
}

