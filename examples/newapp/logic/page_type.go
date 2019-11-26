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

func PageTypeFind(filter types.PageTypeFilter)  (result []types.PageType, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.PageType{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.PageType{})

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
    //            if core.Db.NewScope(&dbmodels.PageType{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.PageType{}).Count(&count)

    if q.Error != nil {
       log.Println("FindPageType > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.PageType{}).HasColumn(Field) {
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
       return []types.PageType{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignPageTypeTypeFromDb(item))
    }

    return result, count, nil
}


func PageTypeMultiCreate(filter types.PageTypeFilter)  (data []types.PageType, err error) {

    typeModelList, err := filter.GetPageTypeModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTypeModel(typeModel)
        item, e := PageTypeCreate(filter, tx)

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

func PageTypeCreate(filter types.PageTypeFilter, query *gorm.DB)  (data types.PageType, err error) {

    typeModel := filter.GetPageTypeModel()
    dbModel := AssignPageTypeDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("PageTypeCreate > Create PageType error:", dbModel)
        return types.PageType{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("PageTypeCreate > Create PageType error:", query.Error)
        return types.PageType{}, errors.New("cant create PageType")
    }

    return AssignPageTypeTypeFromDb(dbModel), nil
}

func PageTypeRead(filter types.PageTypeFilter)  (data types.PageType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := PageTypeFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.PageType{}, errors.New("Not found")
}


func PageTypeMultiUpdate(filter types.PageTypeFilter)  (data []types.PageType, err error) {

    typeModelList, err := filter.GetPageTypeModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTypeModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := PageTypeUpdate(filter, tx)

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

func PageTypeUpdate(filter types.PageTypeFilter, query *gorm.DB)  (data types.PageType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageType not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetPageTypeModel()

    updateModel := AssignPageTypeDbFromType(newModel)
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

    q := query.Model(dbmodels.PageType{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageTypeTypeFromDb(updateModel)
    return
}


func PageTypeMultiDelete(filter types.PageTypeFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetPageTypeModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTypeModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := PageTypeDelete(filter, tx)

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

func PageTypeDelete(filter types.PageTypeFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageType not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.PageType{}).Where(dbmodels.PageType{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func PageTypeFindOrCreate(filter types.PageTypeFilter)  (data types.PageType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignPageTypeDbFromType(filter.GetPageTypeModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.PageType{}).Where(dbmodels.PageType{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageTypeTypeFromDb(findOrCreateModel)
    return
}

