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

func PageTemplateFind(filter types.PageTemplateFilter)  (result []types.PageTemplate, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.PageTemplate{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    criteria := core.Db.Where(dbmodels.PageTemplate{})

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
    //            if core.Db.NewScope(&dbmodels.PageTemplate{}).HasColumn(field) {
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

    q := criteria.Model(dbmodels.PageTemplate{}).Count(&count)

    if q.Error != nil {
       log.Println("FindPageTemplate > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.NewScope(&dbmodels.PageTemplate{}).HasColumn(Field) {
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
       return []types.PageTemplate{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, AssignPageTemplateTypeFromDb(item))
    }

    return result, count, nil
}


func PageTemplateMultiCreate(filter types.PageTemplateFilter)  (data []types.PageTemplate, err error) {

    typeModelList, err := filter.GetPageTemplateModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTemplateModel(typeModel)
        item, e := PageTemplateCreate(filter, tx)

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

func PageTemplateCreate(filter types.PageTemplateFilter, query *gorm.DB)  (data types.PageTemplate, err error) {

    typeModel := filter.GetPageTemplateModel()
    dbModel := AssignPageTemplateDbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = query.Create(&dbModel)

    } else {

        fmt.Println("PageTemplateCreate > Create PageTemplate error:", dbModel)
        return types.PageTemplate{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("PageTemplateCreate > Create PageTemplate error:", query.Error)
        return types.PageTemplate{}, errors.New("cant create PageTemplate")
    }

    return AssignPageTemplateTypeFromDb(dbModel), nil
}

func PageTemplateRead(filter types.PageTemplateFilter)  (data types.PageTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := PageTemplateFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.PageTemplate{}, errors.New("Not found")
}


func PageTemplateMultiUpdate(filter types.PageTemplateFilter)  (data []types.PageTemplate, err error) {

    typeModelList, err := filter.GetPageTemplateModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTemplateModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := PageTemplateUpdate(filter, tx)

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

func PageTemplateUpdate(filter types.PageTemplateFilter, query *gorm.DB)  (data types.PageTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageTemplateRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageTemplate not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    newModel := filter.GetPageTemplateModel()

    updateModel := AssignPageTemplateDbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    updateModel.PageTypeId = newModel.PageTypeId
	updateModel.Name = newModel.Name
	updateModel.IsActive = newModel.IsActive
	updateModel.RegionId = newModel.RegionId
	updateModel.LanguageId = newModel.LanguageId
	updateModel.RootPageId = newModel.RootPageId
	updateModel.LayoutId = newModel.LayoutId
	//updateModel.Field remove this line for disable generator functionality

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := query.Model(dbmodels.PageTemplate{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageTemplateTypeFromDb(updateModel)
    return
}


func PageTemplateMultiDelete(filter types.PageTemplateFilter)  (isOk bool, err error) {

    typeModelList, err := filter.GetPageTemplateModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.SetPageTemplateModel(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := PageTemplateDelete(filter, tx)

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

func PageTemplateDelete(filter types.PageTemplateFilter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := PageTemplateRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("PageTemplate not found in db with id: " + strconv.Itoa(filter.GetCurrentId()))
        }
        return
    }

    q := query.Model(dbmodels.PageTemplate{}).Where(dbmodels.PageTemplate{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func PageTemplateFindOrCreate(filter types.PageTemplateFilter)  (data types.PageTemplate, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := AssignPageTemplateDbFromType(filter.GetPageTemplateModel())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = errors.New(findOrCreateModel.GetValidationErrors())
        return
    }

    q := core.Db.Model(dbmodels.PageTemplate{}).Where(dbmodels.PageTemplate{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = AssignPageTemplateTypeFromDb(findOrCreateModel)
    return
}

