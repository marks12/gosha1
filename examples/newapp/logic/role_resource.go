package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func RoleResourceFind(filter types.RoleResourceFilter)  (result []types.RoleResource, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.RoleResource{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.RoleResource{})

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
    //            if core.Db.NewScope(&dbmodels.RoleResource{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.RoleResource{}).Count(&count)

    if q.Error != nil {
       log.Println("FindRoleResource > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.RoleResource{}).HasColumn(Field) {
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
       return []types.RoleResource{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignRoleResourceTypeFromDb(item))
    }

    return result, count, nil
}

func RoleResourceCreate(filter types.RoleResourceFilter)  (data types.RoleResource, err error) {

    query := core.Db

    typeModel := filter.GetRoleResourceModel()
    dbModel := AssignRoleResourceDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("RoleResourceCreate > Create RoleResource error:", dbModel)
        return types.RoleResource{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("RoleResourceCreate > Create RoleResource error:", query.Error)
        return types.RoleResource{}, errors.New("cant create RoleResource")
    }

    return AssignRoleResourceTypeFromDb(dbModel), nil
}

func RoleResourceRead(filter types.RoleResourceFilter)  (data types.RoleResource, err error) {

    findData, _, err := RoleResourceFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.RoleResource{}, errors.New("Not found")
}

func RoleResourceUpdate(filter types.RoleResourceFilter)  (data types.RoleResource, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RoleResourceRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("RoleResource not found in db")
        }
        return
    }

    newModel := filter.GetRoleResourceModel()

    updateModel := AssignRoleResourceDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.RoleId = newModel.RoleId
	updateModel.ResourceId = newModel.ResourceId
	updateModel.Find = newModel.Find
	updateModel.Read = newModel.Read
	updateModel.Create = newModel.Create
	updateModel.Update = newModel.Update
	updateModel.Delete = newModel.Delete
	updateModel.FindOrCreate = newModel.FindOrCreate
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.RoleResource{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRoleResourceTypeFromDb(updateModel)
    return
}


func RoleResourceDelete(filter types.RoleResourceFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RoleResourceRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("RoleResource not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.RoleResource{}).Where(dbmodels.RoleResource{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func RoleResourceFindOrCreate(filter types.RoleResourceFilter)  (data types.RoleResource, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignRoleResourceDbFromType(filter.GetRoleResourceModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.RoleResource{}).Where(dbmodels.RoleResource{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRoleResourceTypeFromDb(findOrCreateModel)
    return
}

