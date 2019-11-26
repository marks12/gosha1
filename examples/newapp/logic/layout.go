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

func LayoutFind(filter types.LayoutFilter)  (result []types.Layout, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.Layout{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.Layout{})

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
    //            if core.Db.NewScope(&dbmodels.Layout{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.Layout{}).Count(&count)

    if q.Error != nil {
       log.Println("FindLayout > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.Layout{}).HasColumn(Field) {
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
       return []types.Layout{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignLayoutTypeFromDb(item))
    }

    return result, count, nil
}


func LayoutMultiCreate(filter types.LayoutFilter)  (data []types.Layout, err error) {

    typeModelList, err := filter.GetLayoutModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutModel(typeModel)
        item, e := LayoutCreate(filter, tx)

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

func LayoutCreate(filter types.LayoutFilter, query *gorm.DB)  (data types.Layout, err error) {

    typeModel := filter.GetLayoutModel()
    dbModel := AssignLayoutDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("LayoutCreate > Create Layout error:", dbModel)
        return types.Layout{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("LayoutCreate > Create Layout error:", query.Error)
        return types.Layout{}, errors.New("cant create Layout")
    }

    return AssignLayoutTypeFromDb(dbModel), nil
}

func LayoutRead(filter types.LayoutFilter)  (data types.Layout, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := LayoutFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.Layout{}, errors.New("Not found")
}


func LayoutMultiUpdate(filter types.LayoutFilter)  (data []types.Layout, err error) {

    typeModelList, err := filter.GetLayoutModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := LayoutUpdate(filter, tx)

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

func LayoutUpdate(filter types.LayoutFilter, query *gorm.DB)  (data types.Layout, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := LayoutRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Layout not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetLayoutModel()

    updateModel := AssignLayoutDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.RegionId = newModel.RegionId
	updateModel.LanguageId = newModel.LanguageId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.Layout{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignLayoutTypeFromDb(updateModel)
    return
}


func LayoutMultiDelete(filter types.LayoutFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetLayoutModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := LayoutDelete(filter, tx)

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

func LayoutDelete(filter types.LayoutFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := LayoutRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Layout not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.Layout{}).Where(dbmodels.Layout{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func LayoutFindOrCreate(filter types.LayoutFilter)  (data types.Layout, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignLayoutDbFromType(filter.GetLayoutModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Layout{}).Where(dbmodels.Layout{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignLayoutTypeFromDb(findOrCreateModel)
    return
}

