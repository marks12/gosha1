package cmd

const usualEntityLogicFind = `
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
       result = append(result, Assign{Entity}TypeFromDb(item))
    }

    return result, count, nil
}
`

const usualEntityLogicCreate = `
func {Entity}Create(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    query := core.Db

    typeModel := filter.Get{Entity}Model()
    dbModel := Assign{Entity}DbFromType(typeModel)
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

    return Assign{Entity}TypeFromDb(dbModel), nil
}
`

const usualEntityLogicRead = `
func {Entity}Read(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    findData, _, err := {Entity}Find(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.{Entity}{}, errors.New("Not found")
}
`

var usualEntityLogicUpdate = `
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

    updateModel := Assign{Entity}DbFromType(filter.Get{Entity}Model())
    updateModel.ID = existsModel.Id
    ` + getRemoveLine("updateModel.field") + `

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = errors.New(updateModel.GetValidationErrors())
        return
    }

    q := core.Db.Where(dbmodels.{Entity}{ID: updateModel.ID}).Update(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = Assign{Entity}TypeFromDb(updateModel)
    return
}

`
const usualEntityLogicDelete = `
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

const usualEntityLogic = `package logic

import (
    "{ms-name}/types"
    "{ms-name}/dbmodels"
    "{ms-name}/core"
    "log"
    "errors"
    "fmt"
)
`

var usualTemplateEntityLogic = template{
    Path:    "",
    Content: assignMsName(GetUsualTemplateLogicContent(Crud{true, true, true, true, true,})),
}

func GetUsualTemplateLogicContent(crud Crud) (content string) {

    content = usualEntityLogic

    if crud.IsFind {
        content += usualEntityLogicFind
    }

    if crud.IsCreate {
        content += usualEntityLogicCreate
    }

    if crud.IsRead {
        content += usualEntityLogicRead
    }

    if crud.IsUpdate {
        content += usualEntityLogicUpdate
    }

    if crud.IsDelete {
        content += usualEntityLogicDelete
    }

    content = assignMsName(content)

    return
}