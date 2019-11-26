package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
    "github.com/jinzhu/gorm"
    "strconv"
)

func ComponentGroupFind(filter types.ComponentGroupFilter)  (result []types.ComponentGroup, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.ComponentGroup{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.ComponentGroup{})

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
    //            if core.Db.NewScope(&dbmodels.ComponentGroup{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.ComponentGroup{}).Count(&count)

    if q.Error != nil {
       log.Println("FindComponentGroup > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.ComponentGroup{}).HasColumn(Field) {
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
       return []types.ComponentGroup{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignComponentGroupTypeFromDb(item))
    }

    return result, count, nil
}


func ComponentGroupMultiCreate(filter types.ComponentGroupFilter)  (data []types.ComponentGroup, err error) {

    typeModelList, err := filter.GetComponentGroupModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentGroupModel(typeModel)
        item, e := ComponentGroupCreate(filter, tx)

        if e != nil {
            err = e
            data = nil
            break
        }

        data = append(data, item)
    }

    if err == nil {
        tx.Commit()
    } else {
        tx.Rollback()
    }

    return
}

func ComponentGroupCreate(filter types.ComponentGroupFilter, query *gorm.DB)  (data types.ComponentGroup, err error) {

    typeModel := filter.GetComponentGroupModel()
    dbModel := AssignComponentGroupDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("ComponentGroupCreate > Create ComponentGroup error:", dbModel)
        return types.ComponentGroup{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("ComponentGroupCreate > Create ComponentGroup error:", query.Error)
        return types.ComponentGroup{}, errors.New("cant create ComponentGroup")
    }

    return AssignComponentGroupTypeFromDb(dbModel), nil
}

func ComponentGroupRead(filter types.ComponentGroupFilter)  (data types.ComponentGroup, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := ComponentGroupFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.ComponentGroup{}, errors.New("Not found")
}


func ComponentGroupMultiUpdate(filter types.ComponentGroupFilter)  (data []types.ComponentGroup, err error) {

    typeModelList, err := filter.GetComponentGroupModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentGroupModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := ComponentGroupUpdate(filter, tx)

        if e != nil {
            err = e
            data = nil
            break
        }

        data = append(data, item)
    }

    if err == nil {
        tx.Commit()
    } else {
        tx.Rollback()
    }

    return data, nil
}

func ComponentGroupUpdate(filter types.ComponentGroupFilter, query *gorm.DB)  (data types.ComponentGroup, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ComponentGroupRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ComponentGroup not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetComponentGroupModel()

    updateModel := AssignComponentGroupDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.Code = newModel.Code
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.ComponentGroup{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignComponentGroupTypeFromDb(updateModel)
    return
}


func ComponentGroupMultiDelete(filter types.ComponentGroupFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetComponentGroupModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentGroupModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := ComponentGroupDelete(filter, tx)

        if e != nil {
            err = e
            isOk = false
            break
        }
    }

    if err == nil {
        tx.Commit()
    } else {
        tx.Rollback()
    }

    return isOk, err
}

func ComponentGroupDelete(filter types.ComponentGroupFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ComponentGroupRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ComponentGroup not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.ComponentGroup{}).Where(dbmodels.ComponentGroup{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func ComponentGroupFindOrCreate(filter types.ComponentGroupFilter)  (data types.ComponentGroup, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignComponentGroupDbFromType(filter.GetComponentGroupModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.ComponentGroup{}).Where(dbmodels.ComponentGroup{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignComponentGroupTypeFromDb(findOrCreateModel)
    return
}

