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

func PageContentFind(filter types.PageContentFilter)  (result []types.PageContent, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.PageContent{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.PageContent{})

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
    //            if core.Db.NewScope(&dbmodels.PageContent{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.PageContent{}).Count(&count)

    if q.Error != nil {
       log.Println("FindPageContent > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.PageContent{}).HasColumn(Field) {
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
       return []types.PageContent{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignPageContentTypeFromDb(item))
    }

    return result, count, nil
}


func PageContentMultiCreate(filter types.PageContentFilter)  (data []types.PageContent, err error) {

    typeModelList, err := filter.GetPageContentModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageContentModel(typeModel)
        item, e := PageContentCreate(filter, tx)

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

func PageContentCreate(filter types.PageContentFilter, query *gorm.DB)  (data types.PageContent, err error) {

    typeModel := filter.GetPageContentModel()
    dbModel := AssignPageContentDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("PageContentCreate > Create PageContent error:", dbModel)
        return types.PageContent{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("PageContentCreate > Create PageContent error:", query.Error)
        return types.PageContent{}, errors.New("cant create PageContent")
    }

    return AssignPageContentTypeFromDb(dbModel), nil
}

func PageContentRead(filter types.PageContentFilter)  (data types.PageContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := PageContentFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.PageContent{}, errors.New("Not found")
}


func PageContentMultiUpdate(filter types.PageContentFilter)  (data []types.PageContent, err error) {

    typeModelList, err := filter.GetPageContentModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageContentModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := PageContentUpdate(filter, tx)

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

func PageContentUpdate(filter types.PageContentFilter, query *gorm.DB)  (data types.PageContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageContentRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageContent not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetPageContentModel()

    updateModel := AssignPageContentDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.Name = newModel.Name
	updateModel.PageTemplateId = newModel.PageTemplateId
	updateModel.IsActive = newModel.IsActive
	updateModel.RegionId = newModel.RegionId
	updateModel.LanguageId = newModel.LanguageId
	updateModel.Position = newModel.Position
	updateModel.ComponentTemplateCode = newModel.ComponentTemplateCode
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.PageContent{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageContentTypeFromDb(updateModel)
    return
}


func PageContentMultiDelete(filter types.PageContentFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetPageContentModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageContentModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := PageContentDelete(filter, tx)

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

func PageContentDelete(filter types.PageContentFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageContentRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageContent not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.PageContent{}).Where(dbmodels.PageContent{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func PageContentFindOrCreate(filter types.PageContentFilter)  (data types.PageContent, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignPageContentDbFromType(filter.GetPageContentModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.PageContent{}).Where(dbmodels.PageContent{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageContentTypeFromDb(findOrCreateModel)
    return
}

