package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func UserFind(filter types.UserFilter)  (result []types.User, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.User{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.User{})

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
    //            if core.Db.NewScope(&dbmodels.User{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.User{}).Count(&count)

    if q.Error != nil {
       log.Println("FindUser > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.User{}).HasColumn(Field) {
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
       return []types.User{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignUserTypeFromDb(item))
    }

    return result, count, nil
}

func UserCreate(filter types.UserFilter)  (data types.User, err error) {

    query := core.Db

    typeModel := filter.GetUserModel()
    dbModel := AssignUserDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("UserCreate > Create User error:", dbModel)
        return types.User{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("UserCreate > Create User error:", query.Error)
        return types.User{}, errors.New("cant create User")
    }

    return AssignUserTypeFromDb(dbModel), nil
}

func UserRead(filter types.UserFilter)  (data types.User, err error) {

    findData, _, err := UserFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.User{}, errors.New("Not found")
}

func UserUpdate(filter types.UserFilter)  (data types.User, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := UserRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("User not found in db")
        }
        return
    }

    newModel := filter.GetUserModel()

    updateModel := AssignUserDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    //updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.User{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignUserTypeFromDb(updateModel)
    return
}


func UserDelete(filter types.UserFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := UserRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("User not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.User{}).Where(dbmodels.User{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func UserFindOrCreate(filter types.UserFilter)  (data types.User, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignUserDbFromType(filter.GetUserModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.User{}).Where(dbmodels.User{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignUserTypeFromDb(findOrCreateModel)
    return
}

