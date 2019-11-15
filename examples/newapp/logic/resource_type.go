package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func ResourceTypeFind(filter types.ResourceTypeFilter)  (result []types.ResourceType, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.ResourceType{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.ResourceType{})

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
    //            if core.Db.NewScope(&dbmodels.ResourceType{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.ResourceType{}).Count(&count)

    if q.Error != nil {
       log.Println("FindResourceType > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.ResourceType{}).HasColumn(Field) {
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
       return []types.ResourceType{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignResourceTypeTypeFromDb(item))
    }

    return result, count, nil
}

func ResourceTypeCreate(filter types.ResourceTypeFilter)  (data types.ResourceType, err error) {

    query := core.Db

    typeModel := filter.GetResourceTypeModel()
    dbModel := AssignResourceTypeDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("ResourceTypeCreate > Create ResourceType error:", dbModel)
        return types.ResourceType{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("ResourceTypeCreate > Create ResourceType error:", query.Error)
        return types.ResourceType{}, errors.New("cant create ResourceType")
    }

    return AssignResourceTypeTypeFromDb(dbModel), nil
}

func ResourceTypeRead(filter types.ResourceTypeFilter)  (data types.ResourceType, err error) {

    findData, _, err := ResourceTypeFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.ResourceType{}, errors.New("Not found")
}

func ResourceTypeUpdate(filter types.ResourceTypeFilter)  (data types.ResourceType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ResourceTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ResourceType not found in db")
        }
        return
    }

    newModel := filter.GetResourceTypeModel()

    updateModel := AssignResourceTypeDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.ResourceType{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignResourceTypeTypeFromDb(updateModel)
    return
}


func ResourceTypeDelete(filter types.ResourceTypeFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ResourceTypeRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("ResourceType not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.ResourceType{}).Where(dbmodels.ResourceType{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func ResourceTypeFindOrCreate(filter types.ResourceTypeFilter)  (data types.ResourceType, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignResourceTypeDbFromType(filter.GetResourceTypeModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.ResourceType{}).Where(dbmodels.ResourceType{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignResourceTypeTypeFromDb(findOrCreateModel)
    return
}

