package logic

import (
    "newapp/types"
    "newapp/dbmodels"
    "newapp/core"
    "log"
    "errors"
    "fmt"
)

func ResourceFind(filter types.ResourceFilter)  (result []types.Resource, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.Resource{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.Resource{})

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
    //            if core.Db.NewScope(&dbmodels.Resource{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.Resource{}).Count(&count)

    if q.Error != nil {
       log.Println("FindResource > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.Resource{}).HasColumn(Field) {
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
       return []types.Resource{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignResourceTypeFromDb(item))
    }

    return result, count, nil
}

func ResourceCreate(filter types.ResourceFilter)  (data types.Resource, err error) {

    query := core.Db

    typeModel := filter.GetResourceModel()
    dbModel := AssignResourceDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("ResourceCreate > Create Resource error:", dbModel)
        return types.Resource{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("ResourceCreate > Create Resource error:", query.Error)
        return types.Resource{}, errors.New("cant create Resource")
    }

    return AssignResourceTypeFromDb(dbModel), nil
}

func ResourceRead(filter types.ResourceFilter)  (data types.Resource, err error) {

    findData, _, err := ResourceFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.Resource{}, errors.New("Not found")
}

func ResourceUpdate(filter types.ResourceFilter)  (data types.Resource, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ResourceRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Resource not found in db")
        }
        return
    }

    newModel := filter.GetResourceModel()

    updateModel := AssignResourceDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.Code = newModel.Code
	updateModel.TypeId = newModel.TypeId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Resource{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignResourceTypeFromDb(updateModel)
    return
}


func ResourceDelete(filter types.ResourceFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := ResourceRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("Resource not found in db")
        }
        return
    }

    q := core.Db.Model(dbmodels.Resource{}).Where(dbmodels.Resource{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func ResourceFindOrCreate(filter types.ResourceFilter)  (data types.Resource, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignResourceDbFromType(filter.GetResourceModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.Resource{}).Where(dbmodels.Resource{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignResourceTypeFromDb(findOrCreateModel)
    return
}

