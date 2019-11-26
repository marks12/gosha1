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

func PageInfoFind(filter types.PageInfoFilter)  (result []types.PageInfo, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.PageInfo{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.PageInfo{})

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
    //            if core.Db.NewScope(&dbmodels.PageInfo{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.PageInfo{}).Count(&count)

    if q.Error != nil {
       log.Println("FindPageInfo > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.PageInfo{}).HasColumn(Field) {
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
       return []types.PageInfo{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignPageInfoTypeFromDb(item))
    }

    return result, count, nil
}


func PageInfoMultiCreate(filter types.PageInfoFilter)  (data []types.PageInfo, err error) {

    typeModelList, err := filter.GetPageInfoModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageInfoModel(typeModel)
        item, e := PageInfoCreate(filter, tx)

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

func PageInfoCreate(filter types.PageInfoFilter, query *gorm.DB)  (data types.PageInfo, err error) {

    typeModel := filter.GetPageInfoModel()
    dbModel := AssignPageInfoDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("PageInfoCreate > Create PageInfo error:", dbModel)
        return types.PageInfo{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("PageInfoCreate > Create PageInfo error:", query.Error)
        return types.PageInfo{}, errors.New("cant create PageInfo")
    }

    return AssignPageInfoTypeFromDb(dbModel), nil
}

func PageInfoRead(filter types.PageInfoFilter)  (data types.PageInfo, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := PageInfoFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.PageInfo{}, errors.New("Not found")
}


func PageInfoMultiUpdate(filter types.PageInfoFilter)  (data []types.PageInfo, err error) {

    typeModelList, err := filter.GetPageInfoModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageInfoModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := PageInfoUpdate(filter, tx)

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

func PageInfoUpdate(filter types.PageInfoFilter, query *gorm.DB)  (data types.PageInfo, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageInfoRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageInfo not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetPageInfoModel()

    updateModel := AssignPageInfoDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.SeoMeta = newModel.SeoMeta
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.PageInfo{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageInfoTypeFromDb(updateModel)
    return
}


func PageInfoMultiDelete(filter types.PageInfoFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetPageInfoModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageInfoModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := PageInfoDelete(filter, tx)

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

func PageInfoDelete(filter types.PageInfoFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageInfoRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageInfo not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.PageInfo{}).Where(dbmodels.PageInfo{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func PageInfoFindOrCreate(filter types.PageInfoFilter)  (data types.PageInfo, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignPageInfoDbFromType(filter.GetPageInfoModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.PageInfo{}).Where(dbmodels.PageInfo{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageInfoTypeFromDb(findOrCreateModel)
    return
}

