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

func LayoutContentFind(filter types.LayoutContentFilter)  (result []types.LayoutContent, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.LayoutContent{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.LayoutContent{})

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
    //            if core.Db.NewScope(&dbmodels.LayoutContent{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.LayoutContent{}).Count(&count)

    if q.Error != nil {
       log.Println("FindLayoutContent > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.LayoutContent{}).HasColumn(Field) {
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
       return []types.LayoutContent{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignLayoutContentTypeFromDb(item))
    }

    return result, count, nil
}


func LayoutContentMultiCreate(filter types.LayoutContentFilter)  (data []types.LayoutContent, err error) {

    typeModelList, err := filter.GetLayoutContentModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutContentModel(typeModel)
        item, e := LayoutContentCreate(filter, tx)

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

func LayoutContentCreate(filter types.LayoutContentFilter, query *gorm.DB)  (data types.LayoutContent, err error) {

    typeModel := filter.GetLayoutContentModel()
    dbModel := AssignLayoutContentDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("LayoutContentCreate > Create LayoutContent error:", dbModel)
        return types.LayoutContent{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("LayoutContentCreate > Create LayoutContent error:", query.Error)
        return types.LayoutContent{}, errors.New("cant create LayoutContent")
    }

    return AssignLayoutContentTypeFromDb(dbModel), nil
}

func LayoutContentRead(filter types.LayoutContentFilter)  (data types.LayoutContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := LayoutContentFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.LayoutContent{}, errors.New("Not found")
}


func LayoutContentMultiUpdate(filter types.LayoutContentFilter)  (data []types.LayoutContent, err error) {

    typeModelList, err := filter.GetLayoutContentModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutContentModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := LayoutContentUpdate(filter, tx)

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

func LayoutContentUpdate(filter types.LayoutContentFilter, query *gorm.DB)  (data types.LayoutContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := LayoutContentRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("LayoutContent not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetLayoutContentModel()

    updateModel := AssignLayoutContentDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.LayoutId = newModel.LayoutId
	updateModel.ComponentTemplateCode = newModel.ComponentTemplateCode
	updateModel.Position = newModel.Position
	updateModel.IsActive = newModel.IsActive
	updateModel.RegionId = newModel.RegionId
	updateModel.LanguageId = newModel.LanguageId
	updateModel.Name = newModel.Name
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.LayoutContent{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignLayoutContentTypeFromDb(updateModel)
    return
}


func LayoutContentMultiDelete(filter types.LayoutContentFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetLayoutContentModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetLayoutContentModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := LayoutContentDelete(filter, tx)

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

func LayoutContentDelete(filter types.LayoutContentFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := LayoutContentRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("LayoutContent not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.LayoutContent{}).Where(dbmodels.LayoutContent{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func LayoutContentFindOrCreate(filter types.LayoutContentFilter)  (data types.LayoutContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignLayoutContentDbFromType(filter.GetLayoutContentModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.LayoutContent{}).Where(dbmodels.LayoutContent{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignLayoutContentTypeFromDb(findOrCreateModel)
    return
}

