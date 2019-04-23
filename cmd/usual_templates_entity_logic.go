package cmd

const usualEntityLogic = `package logic

import (
    "{ms-name}/types"
    "{ms-name}/dbmodels"
    "{ms-name}/core"
    "log"
    "errors"
    "fmt"
)

func {Entity}Find(filter types.{Entity}Filter)  (result []types.{Entity}, totalRecords int, err error) {

    foundIds 	:= []int{}
    dbmodelData	:= []dbmodels.{Entity}{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()

    var count int

    critery := core.Db.Where(dbmodels.{Entity}{})

    if len(filterIds) > 0 {
        critery = critery.Where("id in (?)", filterIds)
    }

    q := critery.Model(dbmodels.{Entity}{}).Count(&count)

    if q.Error != nil {
       log.Println("Find{Entity} > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    q = critery.Limit(limit).Offset(offset).Find(&dbmodelData)

    if q.Error != nil {
       log.Println("FindProduct > Ошибка получения данных2:", q.Error)
       return []types.{Entity}{}, 0, nil
    }

    // подготовка id для получения связанных сущностей
    for _, item := range dbmodelData {
        foundIds = append(foundIds, item.ID)
    }

    // получение связнаных сущностей

    //формирование результатов
    for _, item := range dbmodelData {
       result = append(result, assign{Entity}TypeFromDb(item))
    }

    return result, count, nil
}


func {Entity}Read(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    findData, _, err := {Entity}Find(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.{Entity}{}, errors.New("Not found")
}

func {Entity}Create(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    query := core.Db

    typeModel := filter.Get{Entity}Model()
    dbModel := assign{Entity}DbFromType(typeModel)
    dbModel.ID = 0

    dbModel.Validate()

    if dbModel.IsValid() {

        query = core.Db.Create(&dbModel)

    } else {

        fmt.Println("{Entity}Create > Create {Entity} error:", dbModel)
        return types.{Entity}{}, errors.New(dbModel.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("{Entity}Create > Create {Entity} error:", query.Error)
        return types.{Entity}{}, errors.New("cant create {Entity}")
    }

    return assign{Entity}TypeFromDb(dbModel), nil
}


func {Entity}Update(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := {Entity}Read(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("{Entity} not found in db")
        }
        return
    }

    updateModel := assign{Entity}DbFromType(filter.Get{Entity}Model())
    updateModel.ID = existsModel.Id

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Where(dbmodels.{Entity}{ID: updateModel.ID}).Update(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = assign{Entity}TypeFromDb(updateModel)
    return
}

func {Entity}Delete(filter types.{Entity}Filter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := {Entity}Read(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("{Entity} not found in db")
        }
        return
    }

    q := core.Db.Where(dbmodels.{Entity}{ID: existsModel.Id}).Delete(&existsModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

`

var usualTemplateEntityLogic = template{
    Path:    "",
    Content: assignMsName(usualEntityLogic),
}
