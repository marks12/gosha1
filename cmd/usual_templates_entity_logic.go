package cmd

import (
	"gosha/mode"
)

const usualEntityLogicFindWoDb = `
func {Entity}Find(filter types.{Entity}Filter)  (result []types.{Entity}, totalRecords int, err error) {

	//{Entity} Find logic code

    return
}
`

const usualEntityLogicFind = `
func {Entity}Find(filter types.{Entity}Filter)  (result []types.{Entity}, totalRecords int, err error) {

    foundIds 	:= []{PkType}{}
    dbmodelData	:= []dbmodels.{Entity}{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    filterIds 	:= filter.GetIds()
    filterExceptIds 	:= filter.GetExceptIds()

    var count int64

    criteria := core.Db.Where(dbmodels.{Entity}{})

	//{Entity}.FindFilterCode remove this line for disable generator functionality

    if len(filterIds) > 0 {
        criteria = criteria.Where("id in (?)", filterIds)
    }

    if len(filterExceptIds) > 0 {
        criteria = criteria.Where("id not in (?)", filterExceptIds)
    }

    //if len(filter.Search) > 0 {
    //
    //    s := ("%" + filter.Search + "%")
    //
    //    if len(filter.SearchBy) > 0 {
    //
    //        for _, field := range filter.SearchBy {
    //
    //            if core.Db.Migrator().HasColumn(&dbmodels.{Entity}{}, field) {
    //                criteria = criteria.Or("` + "`" + `"+field+"` + "`" + `"+" ilike ?", s)
    //            } else {
    //                err = errors.NewErrorWithCode("Search by unknown field", errors.ErrorCodeNotValid ,field)
    //                return
    //            }
    //        }
    //    } else {
    //      criteria = criteria.Where("name ilike ? or code ilike ?", ("%" + filter.Search + "%"), ("%" + filter.Search + "%"))
    //    }
    //}

    q := criteria.Model(dbmodels.{Entity}{}).Count(&count)

    if q.Error != nil {
       log.Println("Find{Entity} > Ошибка получения данных:", q.Error)
       return result, 0, nil
    }

    // order global criteria
    if len(filter.Order) > 0  {
        for index, Field := range filter.Order {
             if core.Db.Migrator().HasColumn(&dbmodels.{Entity}{}, Field) {
                criteria = criteria.Order("\"" + strings.ToLower(Field) + "\"" + " " + filter.OrderDirection[index])
            } else {
				err = errors.NewErrorWithCode("Ordering by unknown Field", errors.ErrorCodeNotValid ,Field)
                return
            }
        }
    }

    q = criteria.Limit(limit).Offset(offset).Find(&dbmodelData)

    if q.Error != nil {
       log.Println("Find{Entity} > Ошибка получения данных2:", q.Error)
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

    return result, int(count), nil
}
`

const usualEntityLogicCreateWoDb = `
func {Entity}MultiCreate(filter types.{Entity}Filter)  (data []types.{Entity}, err error) {

	//{Entity} MultiCreate logic code

    return
}

func {Entity}Create(filter types.{Entity}Filter, query *gorm.DB)  (data types.{Entity}, err error) {
    
	//{Entity} Create logic code
    return
}
`

const usualEntityLogicCreate = `

func {Entity}MultiCreate(filter types.{Entity}Filter)  (data []types.{Entity}, err error) {

    typeModelList, err := filter.Get{Entity}ModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.Set{Entity}Model(typeModel)
        item, e := {Entity}Create(filter, tx)

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

func {Entity}Create(filter types.{Entity}Filter, query *gorm.DB)  (data types.{Entity}, err error) {

    typeModel := filter.Get{Entity}Model()
    dbModel := Assign{Entity}DbFromType(typeModel)
    dbModel.ID = {PkNil}

    dbModel.Validate()

    if ! dbModel.IsValid() {
        fmt.Println("{Entity}Create > Create {Entity} error:", dbModel)
        return types.{Entity}{}, dbModel.GetValidationError()
    }

    query = query.Create(&dbModel)

    if query.Error != nil {
        fmt.Println("{Entity}Create > Create {Entity} error:", query.Error)
        return types.{Entity}{}, errors.NewErrorWithCode("cant create {Entity}", errors.ErrorCodeSqlError, "")
    }

    return Assign{Entity}TypeFromDb(dbModel), nil
}
`

const usualEntityLogicReadWoDb = `
func {Entity}Read(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

	//{Entity} Read logic code
    return
}
`

const usualEntityLogicRead = `
func {Entity}Read(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1
    filter.ClearIds()
    filter.AddId(filter.GetCurrentId())

    findData, _, err := {Entity}Find(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.{Entity}{}, errors.NewErrorWithCode("Not found", errors.ErrorCodeNotFound, "")
}
`

const usualEntityLogicUpdateWoDb = `

func {Entity}MultiUpdate(filter types.{Entity}Filter)  (data []types.{Entity}, err error) {

	//{Entity} MultiUpdate logic code
    return
}

func {Entity}Update(filter types.{Entity}Filter, query *gorm.DB)  (data types.{Entity}, err error) {

	//{Entity} Update logic code
    return
}

`

func GetUsualEntityLogicUpdate() string {
	return `

func {Entity}MultiUpdate(filter types.{Entity}Filter)  (data []types.{Entity}, err error) {

    typeModelList, err := filter.Get{Entity}ModelList()

    if err != nil {
        return
    }

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.Set{Entity}Model(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        item, e := {Entity}Update(filter, tx)

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

func {Entity}Update(filter types.{Entity}Filter, query *gorm.DB)  (data types.{Entity}, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := {Entity}Read(filter)

    if existsModel.Id ` + GetIdIsNotValidExp(mode.GetUuidMode()) + ` || err != nil {
        err = errors.NewErrorWithCode("{Entity} not found in db with id: " + ` + GetPkAsString("filter.GetCurrentId()", mode.GetUuidMode()) + `, errors.ErrorCodeNotFound, "Id")
        return
    }

    newModel := filter.Get{Entity}Model()

    updateModel := Assign{Entity}DbFromType(newModel)
    updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    ` + getRemoveLine("updateModel.Field") + `

    updateModel.Validate()

    if !updateModel.IsValid() {
        err = updateModel.GetValidationError()
        return
    }

    q := query.Model(dbmodels.{Entity}{}).Save(&updateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = Assign{Entity}TypeFromDb(updateModel)
    return
}

`
}

const usualEntityLogicDeleteWoDb = `

func {Entity}MultiDelete(filter types.{Entity}Filter)  (isOk bool, err error) {

	//{Entity} MultiDelete logic code
    return
}

func {Entity}Delete(filter types.{Entity}Filter, query *gorm.DB)  (isOk bool, err error) {

	//{Entity} Delete logic code
    return
}

`

func GetUsualEntityLogicDelete() string {

	return `
func {Entity}MultiDelete(filter types.{Entity}Filter)  (isOk bool, err error) {

    typeModelList, err := filter.Get{Entity}ModelList()

    if err != nil {
        return
    }

    isOk = true

    tx := core.Db.Begin()

    for _, typeModel := range typeModelList {

        filter.Set{Entity}Model(typeModel)
        filter.ClearIds()
        filter.SetCurrentId(typeModel.Id)

        _, e := {Entity}Delete(filter, tx)

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

func {Entity}Delete(filter types.{Entity}Filter, query *gorm.DB)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := {Entity}Read(filter)

    if existsModel.Id {GetIdIsNotValidExp} || err != nil {

        if err != nil {
            err = errors.NewErrorWithCode("{Entity} not found in db with id: " + ` + GetPkAsString("filter.GetCurrentId()", mode.GetUuidMode()) + `, errors.ErrorCodeNotFound, "")
        }
        return
    }

    dbModel := Assign{Entity}DbFromType(existsModel)
    q := query.Model(dbmodels.{Entity}{}).Where(dbmodels.{Entity}{ID: dbModel.ID}).Delete(&dbModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}
`
}

const usualEntityLogicFindOrCreateWoDb = `

func {Entity}FindOrCreate(filter types.{Entity}Filter)  (data types.{Entity}, err error) {
    
	//{Entity} FindOrCreate logic code
    return
}
`

const usualEntityLogicUpdateOrCreateWoDb = `

func {Entity}UpdateOrCreate(filter types.{Entity}Filter)  (data types.{Entity}, err error) {
    
	//{Entity} UpdateOrCreate logic code
    return
}

// add all assign functions
`

const usualEntityLogicFindOrCreate = `
func {Entity}FindOrCreate(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    findOrCreateModel := Assign{Entity}DbFromType(filter.Get{Entity}Model())
	//findOrCreateModel.Field remove this line for disable generator functionality

    findOrCreateModel.Validate()

    if !findOrCreateModel.IsValid() {
        err = findOrCreateModel.GetValidationError()
        return
    }

    q := core.Db.Model(dbmodels.{Entity}{}).Where(dbmodels.{Entity}{ID: findOrCreateModel.ID}).FirstOrCreate(&findOrCreateModel)

    if q.Error != nil {
        err = q.Error
        return
    }

    data = Assign{Entity}TypeFromDb(findOrCreateModel)
    return
}

`

const usualEntityLogicUpdateOrCreate = `
func {Entity}UpdateOrCreate(filter types.{Entity}Filter)  (data types.{Entity}, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    updateOrCreateModel := Assign{Entity}DbFromType(filter.Get{Entity}Model())
	//updateOrCreateModel.Field remove this line for disable generator functionality

    updateOrCreateModel.Validate()

    if !updateOrCreateModel.IsValid() {
        err = updateOrCreateModel.GetValidationError()
        return
    }

    //please uncomment and set criteria
    //q := core.Db.Model(dbmodels.{Entity}{}).Where(dbmodels.{Entity}{ID: updateOrCreateModel.ID}).Assign(dbmodels.{Entity}{/*PLEASE SET CRITERIA*/}).FirstOrCreate(&updateOrCreateModel)

    //if q.Error != nil {
    //    err = q.Error
    //    return
    //}

    data = Assign{Entity}TypeFromDb(updateOrCreateModel)
    return
}

// add all assign functions
`

func getUsualEntityLogicHeader(isWoModels bool) (header string) {

	top := `package logic

    import (
        "{ms-name}/types"
`

	if mode.GetUuidMode() {
		if !isWoModels {
			top += `        "github.com/google/uuid"
`
		}
	} else {
		if !isWoModels {
			top += `        "strconv"
`
		}
	}

	footer := `
        "gorm.io/gorm"
    )
`
	middle := ""

	if !isWoModels {
		middle = `
        "log"
        "fmt"
        "{ms-name}/core"
        "{ms-name}/errors"
        "{ms-name}/dbmodels"
		"strings"` + "\n"
	}

	return top + middle + footer

}

//var usualTemplateEntityLogic = template{
//    Path:    "",
//    Content: assignMsName(GetUsualTemplateLogicContent(Crud{true, true, true, true, true, true}, false)),
//}

func GetUsualTemplateLogicContent(crud Crud, isWoDbModel bool) (content string) {

	isUuidAsPk := mode.GetUuidMode()

	content = getUsualEntityLogicHeader(isWoDbModel)

	if crud.IsFind {
		if isWoDbModel {
			content += usualEntityLogicFindWoDb
		} else {
			content += usualEntityLogicFind
		}

	}

	if crud.IsCreate {
		if isWoDbModel {
			content += usualEntityLogicCreateWoDb
		} else {
			content += usualEntityLogicCreate
		}
	}

	if crud.IsRead {
		if isWoDbModel {
			content += usualEntityLogicReadWoDb
		} else {
			content += usualEntityLogicRead
		}
	}

	if crud.IsUpdate {
		if isWoDbModel {
			content += usualEntityLogicUpdateWoDb
		} else {
			content += GetUsualEntityLogicUpdate()
		}
	}

	if crud.IsDelete {
		if isWoDbModel {
			content += usualEntityLogicDeleteWoDb
		} else {
			content += GetUsualEntityLogicDelete()
		}
	}

	if crud.IsFindOrCreate {
		if isWoDbModel {
			content += usualEntityLogicFindOrCreateWoDb
		} else {
			content += usualEntityLogicFindOrCreate
		}
	}

	if crud.IsUpdateOrCreate {
		if isWoDbModel {
			content += usualEntityLogicUpdateOrCreateWoDb
		} else {
			content += usualEntityLogicUpdateOrCreate
		}
	}

	content = assignMsName(content)
	content = AssignVar(content, "{GetIdIsNotValidExp}", GetIdIsNotValidExp(isUuidAsPk))
	content = AssignVar(content, "{PkNil}", GetIdNil(isUuidAsPk))
	content = AssignVar(content, "{PkType}", GetPKType(isUuidAsPk))

	return
}
