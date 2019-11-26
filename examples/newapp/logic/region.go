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

func RegionFind(filter types.RegionFilter)  (result []types.Region, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.Region{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.Region{})

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
    //            if core.Db.NewScope(&dbmodels.Region{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.Region{}).Count(&count)

    if q.Error != nil {
       log.Println("FindRegion > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.Region{}).HasColumn(Field) {
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
       return []types.Region{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignRegionTypeFromDb(item))
    }

    return result, count, nil
}


func RegionMultiCreate(filter types.RegionFilter)  (data []types.Region, err error) {

    typeModelList, err := filter.GetRegionModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionModel(typeModel)
        item, e := RegionCreate(filter, tx)

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

func RegionCreate(filter types.RegionFilter, query *gorm.DB)  (data types.Region, err error) {

    typeModel := filter.GetRegionModel()
    dbModel := AssignRegionDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("RegionCreate > Create Region error:", dbModel)
        return types.Region{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("RegionCreate > Create Region error:", query.Error)
        return types.Region{}, errors.New("cant create Region")
    }

    return AssignRegionTypeFromDb(dbModel), nil
}

func RegionRead(filter types.RegionFilter)  (data types.Region, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := RegionFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.Region{}, errors.New("Not found")
}


func RegionMultiUpdate(filter types.RegionFilter)  (data []types.Region, err error) {

    typeModelList, err := filter.GetRegionModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := RegionUpdate(filter, tx)

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

func RegionUpdate(filter types.RegionFilter, query *gorm.DB)  (data types.Region, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RegionRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Region not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetRegionModel()

    updateModel := AssignRegionDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.Code = newModel.Code
	updateModel.TypeId = newModel.TypeId
	updateModel.EntityId = newModel.EntityId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.Region{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRegionTypeFromDb(updateModel)
    return
}


func RegionMultiDelete(filter types.RegionFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetRegionModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetRegionModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := RegionDelete(filter, tx)

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

func RegionDelete(filter types.RegionFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := RegionRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Region not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.Region{}).Where(dbmodels.Region{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func RegionFindOrCreate(filter types.RegionFilter)  (data types.Region, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignRegionDbFromType(filter.GetRegionModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Region{}).Where(dbmodels.Region{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignRegionTypeFromDb(findOrCreateModel)
    return
}

