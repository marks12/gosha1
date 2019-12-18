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

func RegionTypeFind(filter types.RegionTypeFilter)  (result []types.RegionType, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.RegionType{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.RegionType{})

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
    //            if core.Db.NewScope(&dbmodels.RegionType{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.RegionType{}).Count(&count)

    if q.Error != nil {
       log.Println("FindRegionType > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.RegionType{}).HasColumn(Field) {
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
       return []types.RegionType{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignRegionTypeTypeFromDb(item))
    }

    return result, count, nil
}


func RegionTypeMultiCreate(filter types.RegionTypeFilter)  (data []types.RegionType, err error) {

    typeModelList, err := filter.GetRegionTypeModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionTypeModel(typeModel)
        item, e := RegionTypeCreate(filter, tx)

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

func RegionTypeCreate(filter types.RegionTypeFilter, query *gorm.DB)  (data types.RegionType, err error) {

    typeModel := filter.GetRegionTypeModel()
    dbModel := AssignRegionTypeDbFromTp(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("RegionTypeCreate > Create RegionType error:", dbModel)
        return types.RegionType{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("RegionTypeCreate > Create RegionType error:", query.Error)
        return types.RegionType{}, errors.New("cant create RegionType")
    }

    return AssignRegionTypeTypeFromDb(dbModel), nil
}

func RegionTypeRead(filter types.RegionTypeFilter)  (data types.RegionType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := RegionTypeFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.RegionType{}, errors.New("Not found")
}


func RegionTypeMultiUpdate(filter types.RegionTypeFilter)  (data []types.RegionType, err error) {

    typeModelList, err := filter.GetRegionTypeModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionTypeModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := RegionTypeUpdate(filter, tx)

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

func RegionTypeUpdate(filter types.RegionTypeFilter, query *gorm.DB)  (data types.RegionType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RegionTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("RegionType not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetRegionTypeModel()

    updateModel := AssignRegionTypeDbFromTp(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.RegionType{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRegionTypeTypeFromDb(updateModel)
    return
}


func RegionTypeMultiDelete(filter types.RegionTypeFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetRegionTypeModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionTypeModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := RegionTypeDelete(filter, tx)

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

func RegionTypeDelete(filter types.RegionTypeFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RegionTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("RegionType not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.RegionType{}).Where(dbmodels.RegionType{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func RegionTypeFindOrCreate(filter types.RegionTypeFilter)  (data types.RegionType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignRegionTypeDbFromTp(filter.GetRegionTypeModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.RegionType{}).Where(dbmodels.RegionType{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRegionTypeTypeFromDb(findOrCreateModel)
    return
}

