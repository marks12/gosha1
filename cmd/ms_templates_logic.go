package cmd
//
//const msLogicAssigner = `package logic
//
//import (
//	"{ms-name}/dbmodels"
//	"golang.org/x/crypto/bcrypt"
//	"{ms-name}/types"
//	"{ms-name}/settings"
//)
//
//// add all assign functions
//`

const msLogicEntity = `
package logic

import (
    "msrpc/api"
    "log"
    "{ms-name}/dbmodels"
    "{ms-name}/core"
    "github.com/google/uuid"
    "fmt"
    "errors"
    "msrpc/mdl"
)

func Find{entity-name}(app mdl.Application, filter api.{entity-name}Filter) (products []api.{entity-name}, totalRecords int, err error) {

    appGuid := app.GetGuid()

    data 		:= []dbmodels.{entity-name}{}
    result 		:= []api.{entity-name}{}
    entityIds 	:= []uuid.UUID{}
    limit       := filter.PerPage
    offset      := filter.GetOffset()

    var count int

    criteria := core.Db.Where(dbmodels.{entity-name}{AppId: appGuid})

    q := criteria.Model(dbmodels.{entity-name}{}).Count(&count)

    if q.Error != nil {
        log.Println("Find{entity-name} > Ошибка получения данных1:", q.Error)
        return []api.{entity-name}{}, 0, nil
    }

    q = criteria.Limit(limit).Offset(offset).Find(&data)

    if q.Error != nil {
        log.Println("Find{entity-name} > Ошибка получения данных2:", q.Error)
        return []api.{entity-name}{}, 0, nil
    }

    for _, item := range data {
        entityIds = append(entityIds, item.ID)
    }

    for _, item := range data {

        result = append(result, fillApi{entity-name}FromDb(item))
    }

    return result, count, nil
}

func Create{entity-name}(app mdl.Application, {entity-name}Data api.{entity-name}) (api.{entity-name}, error) {

    appGuid := app.GetGuid()

    query := core.Db

    db{entity-name} := dbmodels.{entity-name}{}

    db{entity-name}.AppId = appGuid
    db{entity-name}.FillFromRequest({entity-name}Data)

    db{entity-name}.Validate()

    if db{entity-name}.IsValid() {

        query = core.Db.Create(&db{entity-name})

    } else {

        fmt.Println("Create{entity-name} > Create {entity-name} error:", db{entity-name})
        return api.{entity-name}{}, errors.New(db{entity-name}.GetValidationErrors())
    }

    if query.Error != nil {
        fmt.Println("Create{entity-name} > Create {entity-name} error:", query.Error)
        return api.{entity-name}{}, errors.New("cant create user")
    }

    return fillApi{entity-name}FromDb(db{entity-name}), nil
}


func Read{entity-name}(app mdl.Application, filter api.{entity-name}Filter) (product api.{entity-name}, err error) {

    appGuid := app.GetGuid()

    data        := dbmodels.{entity-name}{}
    guid        := filter.GetGuid()

    criteria := core.Db.Where(dbmodels.{entity-name}{AppId: appGuid, ID: guid})

    q := criteria.First(&data)

    if q.Error != nil {
        log.Println("Read{entity-name} > Ошибка получения данных2:", q.Error)
        return api.{entity-name}{}, errors.New("{entity-name} с ID " + data.ID.String() + " не существует")
    }

    return fillApi{entity-name}FromDb(data), nil
}

func Update{entity-name}(app mdl.Application, {entity-name}Data api.{entity-name}) (api.{entity-name}, error) {

    appGuid := app.GetGuid()

    db{entity-name} := dbmodels.{entity-name}{}

    criteria := core.Db.Where(dbmodels.{entity-name}{AppId: appGuid, ID: {entity-name}Data.GetGuid()})

    q := criteria.First(&db{entity-name})

    if q.Error != nil {
        log.Println("Read{entity-name} > Ошибка получения данных2:", q.Error)
        return api.{entity-name}{}, errors.New("{entity-name} с ID " + {entity-name}Data.ID + " не существует")
    }

    db{entity-name}.AppId         = appGuid
    db{entity-name}.FillFromRequest({entity-name}Data)

    db{entity-name}.Validate()

    if db{entity-name}.IsValid() {

        query := core.Db.Save(&db{entity-name})

        if query.Error != nil {
            log.Println("Read{entity-name} > Ошибка сохранения данных3:", q.Error)
            return api.{entity-name}{}, errors.New("Ощибка сохранения {entity-name} с ID " + {entity-name}Data.ID)
        }

    } else {

        fmt.Println("Update{entity-name} > Update product error:", db{entity-name})
        return api.{entity-name}{}, errors.New(db{entity-name}.GetValidationErrors())
    }

    return fillApi{entity-name}FromDb(db{entity-name}), nil
}

func Delete{entity-name}(app mdl.Application, filter api.{entity-name}Filter) error {

    appGuid := app.GetGuid()

    data        := dbmodels.{entity-name}{}
    guid        := filter.GetGuid()

    criteria := core.Db.Where(dbmodels.{entity-name}{AppId: appGuid, ID: guid})

    q := criteria.First(&data)

    if q.Error != nil {

        log.Println("Read{entity-name} > Ошибка получения данных2:", q.Error)
        return errors.New("{entity-name} с ID " + guid.String() + " не существует")

    } else {

        q = criteria.Delete(&data)

        if q.Error != nil {

            log.Println("Read{entity-name} > Ошибка получения данных2:", q.Error)
            return errors.New("Ошибка удаления {entity-name} с ID " + guid.String())

        }
    }

    return nil
}
`

const msLogicAssignEntity = `

func fillApi{entity-name}FromDb(db{entity-name} dbmodels.{entity-name}) api.{entity-name} {

    return api.{entity-name}{
        Model: api.Model{
            ID:               db{entity-name}.ID.String(),

            CreatedAt:        db{entity-name}.CreatedAt.String(),
            UpdatedAt:        db{entity-name}.UpdatedAt.String(),
            DeletedAt:        db{entity-name}.DeletedAt.String(),
        },
    }
}

`

//var msTemplateLogicAssigner = template{
//    Path:    "./logic/assigner.go",
//    Content: assignMsName(msLogicAssigner),
//}

var msTemplateLogicEntity = template{
    Path:    "",
    Content: msLogicEntity,
}

//var msTemplateLogicAssignEntity = template{
//   Path:    "./logic/assigner.go",
//   Content: msLogicAssignEntity,
//}
