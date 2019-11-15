package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func CurrentUserFind(filter types.CurrentUserFilter)  (result []types.CurrentUser, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.CurrentUser{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.CurrentUser{})

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
    //            if core.Db.NewScope(&dbmodels.CurrentUser{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.CurrentUser{}).Count(&count)

    if q.Error != nil {
       log.Println("FindCurrentUser > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.CurrentUser{}).HasColumn(Field) {
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
       return []types.CurrentUser{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignCurrentUserTypeFromDb(item))
    }

    return result, count, nil
}

func CurrentUserCreate(filter types.CurrentUserFilter)  (data types.CurrentUser, err error) {

    query := core.Db

    typeModel := filter.GetCurrentUserModel()
    dbModel := AssignCurrentUserDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("CurrentUserCreate > Create CurrentUser error:", dbModel)
        return types.CurrentUser{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("CurrentUserCreate > Create CurrentUser error:", query.Error)
        return types.CurrentUser{}, errors.New("cant create CurrentUser")
    }

    return AssignCurrentUserTypeFromDb(dbModel), nil
}

func CurrentUserRead(filter types.CurrentUserFilter)  (data types.CurrentUser, err error) {

    findData, _, err := CurrentUserFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.CurrentUser{}, errors.New("Not found")
}

func CurrentUserUpdate(filter types.CurrentUserFilter)  (data types.CurrentUser, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := CurrentUserRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("CurrentUser not found in db")
        }
        return
    }

    newModel := filter.GetCurrentUserModel()

    updateModel := AssignCurrentUserDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    //updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.CurrentUser{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignCurrentUserTypeFromDb(updateModel)
    return
}


func CurrentUserDelete(filter types.CurrentUserFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := CurrentUserRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("CurrentUser not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.CurrentUser{}).Where(dbmodels.CurrentUser{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func CurrentUserFindOrCreate(filter types.CurrentUserFilter)  (data types.CurrentUser, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignCurrentUserDbFromType(filter.GetCurrentUserModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.CurrentUser{}).Where(dbmodels.CurrentUser{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignCurrentUserTypeFromDb(findOrCreateModel)
    return
}

