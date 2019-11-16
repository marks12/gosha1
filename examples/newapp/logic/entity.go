package logic

import (
    "github.com/jinzhu/gorm"
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func EntityFind(filter types.EntityFilter)  (result []types.Entity, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.Entity{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.Entity{})

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
    //            if core.Db.NewScope(&dbmodels.Entity{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.Entity{}).Count(&count)

    if q.Error != nil {
       log.Println("FindEntity > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.Entity{}).HasColumn(Field) {
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
       return []types.Entity{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignEntityTypeFromDb(item))
    }

    return result, count, nil
}

func EntityCreate(filter types.EntityFilter,  query *gorm.DB)  (data types.Entity, err error) {

    typeModel := filter.GetEntityModel()
    dbModel := AssignEntityDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("EntityCreate > Create Entity error:", dbModel)
        return types.Entity{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("EntityCreate > Create Entity error:", query.Error)
        return types.Entity{}, errors.New("cant create Entity")
    }

    return AssignEntityTypeFromDb(dbModel), nil
}

func EntityMultiCreate(filter types.EntityFilter)  (data []types.Entity, err error) {

    typeModelList, err := filter.GetEntityModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetEntityModel(typeModel)
        item, e := EntityCreate(filter, tx)

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

func EntityRead(filter types.EntityFilter)  (data types.Entity, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := EntityFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.Entity{}, errors.New("Not found")
}

func EntityUpdate(filter types.EntityFilter, query *gorm.DB)  (data types.Entity, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := EntityRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Entity not found in db")
        }
        return
    }

    newModel := filter.GetEntityModel()

    updateModel := AssignEntityDbFromType(newModel)
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

    q := query.Model(dbmodels.Entity{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignEntityTypeFromDb(updateModel)
    return
}

func EntityMultiUpdate(filter types.EntityFilter)  (data []types.Entity, err error) {

    typeModelList, err := filter.GetEntityModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetEntityModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := EntityUpdate(filter, tx)

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

func EntityMultiDelete(filter types.EntityFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetEntityModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetEntityModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := EntityDelete(filter, tx)

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

    return true, nil
}

func EntityDelete(filter types.EntityFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := EntityRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Entity not found in db")
        }
        return
    }

    q := query.Model(dbmodels.Entity{}).Where(dbmodels.Entity{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func EntityFindOrCreate(filter types.EntityFilter)  (data types.Entity, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignEntityDbFromType(filter.GetEntityModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Entity{}).Where(dbmodels.Entity{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignEntityTypeFromDb(findOrCreateModel)
    return
}

