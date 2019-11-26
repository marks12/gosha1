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

func ComponentTemplateFind(filter types.ComponentTemplateFilter)  (result []types.ComponentTemplate, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.ComponentTemplate{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.ComponentTemplate{})

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
    //            if core.Db.NewScope(&dbmodels.ComponentTemplate{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.ComponentTemplate{}).Count(&count)

    if q.Error != nil {
       log.Println("FindComponentTemplate > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.ComponentTemplate{}).HasColumn(Field) {
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
       return []types.ComponentTemplate{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignComponentTemplateTypeFromDb(item))
    }

    return result, count, nil
}


func ComponentTemplateMultiCreate(filter types.ComponentTemplateFilter)  (data []types.ComponentTemplate, err error) {

    typeModelList, err := filter.GetComponentTemplateModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentTemplateModel(typeModel)
        item, e := ComponentTemplateCreate(filter, tx)

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

func ComponentTemplateCreate(filter types.ComponentTemplateFilter, query *gorm.DB)  (data types.ComponentTemplate, err error) {

    typeModel := filter.GetComponentTemplateModel()
    dbModel := AssignComponentTemplateDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("ComponentTemplateCreate > Create ComponentTemplate error:", dbModel)
        return types.ComponentTemplate{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("ComponentTemplateCreate > Create ComponentTemplate error:", query.Error)
        return types.ComponentTemplate{}, errors.New("cant create ComponentTemplate")
    }

    return AssignComponentTemplateTypeFromDb(dbModel), nil
}

func ComponentTemplateRead(filter types.ComponentTemplateFilter)  (data types.ComponentTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := ComponentTemplateFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.ComponentTemplate{}, errors.New("Not found")
}


func ComponentTemplateMultiUpdate(filter types.ComponentTemplateFilter)  (data []types.ComponentTemplate, err error) {

    typeModelList, err := filter.GetComponentTemplateModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentTemplateModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := ComponentTemplateUpdate(filter, tx)

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

func ComponentTemplateUpdate(filter types.ComponentTemplateFilter, query *gorm.DB)  (data types.ComponentTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ComponentTemplateRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ComponentTemplate not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetComponentTemplateModel()

    updateModel := AssignComponentTemplateDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.Code = newModel.Code
	updateModel.Path = newModel.Path
	updateModel.GroupCode = newModel.GroupCode
	updateModel.GroupId = newModel.GroupId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.ComponentTemplate{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignComponentTemplateTypeFromDb(updateModel)
    return
}


func ComponentTemplateMultiDelete(filter types.ComponentTemplateFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetComponentTemplateModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetComponentTemplateModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := ComponentTemplateDelete(filter, tx)

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

func ComponentTemplateDelete(filter types.ComponentTemplateFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ComponentTemplateRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ComponentTemplate not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.ComponentTemplate{}).Where(dbmodels.ComponentTemplate{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func ComponentTemplateFindOrCreate(filter types.ComponentTemplateFilter)  (data types.ComponentTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignComponentTemplateDbFromType(filter.GetComponentTemplateModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.ComponentTemplate{}).Where(dbmodels.ComponentTemplate{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignComponentTemplateTypeFromDb(findOrCreateModel)
    return
}

