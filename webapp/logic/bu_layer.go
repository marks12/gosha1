package logic

import (
    "gosha/webapp/types"
    "errors"
)

func BuLayerFind(filter types.BuLayerFilter)  (result []types.BuLayer, totalRecords int, err error) {

//    foundIds 	:= []int{}
//    dbmodelData	:= []dbmodels.BuLayer{}
//    limit       := filter.PerPage
//    offset      := filter.GetOffset()
//
//    filterIds 	:= filter.GetIds()
//
   var count int
//
//    criteria := core.Db.Where(dbmodels.BuLayer{})
//
//    if len(filterIds) > 0 {
//        criteria = criteria.Where("id in (?)", filterIds)
//    }
//
////    if len(filter.Search) > 0 {
////        criteria = criteria.Where("name like ?", ("%" + filter.Search + "%"), filter.Search)
////    }
//
//    q := criteria.Model(dbmodels.BuLayer{}).Count(&count)
//
//    if q.Error != nil {
//       log.Println("FindBuLayer > Ошибка получения данных:", q.Error)
//       return result, 0, nil
//    }
//
//    // order global criteria
//    if len(filter.Order) > 0  {
//        for index, Field := range filter.Order {
//             if core.Db.NewScope(&dbmodels.BuLayer{}).HasColumn(Field) {
//                criteria = criteria.Order("`" + Field + "`" + " " + filter.OrderDirection[index])
//            } else {
//                err = errors.New("Ordering by unknown Field " + Field)
//                return
//            }
//        }
//    }
//
//    q = criteria.Limit(limit).Offset(offset).Find(&dbmodelData)
//
//    if q.Error != nil {
//       log.Println("FindProduct > Ошибка получения данных2:", q.Error)
//       return []types.BuLayer{}, 0, nil
//    }
//
//    // подготовка id для получения связанных сущностей
//    for _, item := range dbmodelData {
//        foundIds = append(foundIds, item.ID)
//    }
//
//    // получение связнаных сущностей
//
//    //формирование результатов
//    for _, item := range dbmodelData {
//       result = append(result, AssignBuLayerTypeFromDb(item))
//    }

    return result, count, nil
}

func BuLayerCreate(filter types.BuLayerFilter)  (data types.BuLayer, err error) {
	//
    //query := core.Db
	//
	//typeModel := filter.GetBuLayerModel()
    //dbModel := AssignBuLayerDbFromType(typeModel)
    //dbModel.ID = 0
	//
    //dbModel.Validate()
	//
    //if dbModel.IsValid() {
	//
    //    query = core.Db.Create(&dbModel)
	//
    //} else {
	//
    //    fmt.Println("BuLayerCreate > Create BuLayer error:", dbModel)
    //    return types.BuLayer{}, errors.New(dbModel.GetValidationErrors())
    //}
	//
    //if query.Error != nil {
    //    fmt.Println("BuLayerCreate > Create BuLayer error:", query.Error)
    //    return types.BuLayer{}, errors.New("cant create BuLayer")
    //}

    //return AssignBuLayerTypeFromDb(dbModel), nil

    return
}

func BuLayerRead(filter types.BuLayerFilter)  (data types.BuLayer, err error) {

    findData, _, err := BuLayerFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.BuLayer{}, errors.New("Not found")
}

func BuLayerUpdate(filter types.BuLayerFilter)  (data types.BuLayer, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := BuLayerRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("BuLayer not found in db")
        }
        return
    }

    //newModel := filter.GetBuLayerModel()

    //updateModel := AssignBuLayerDbFromType(newModel)
    //updateModel.ID = existsModel.Id

    //updateModel.Some = newModel.Some

    //updateModel.Field remove this line for disable generator functionality

    //updateModel.Validate()
	//
    //if !updateModel.IsValid() {
    //    err = errors.New(updateModel.GetValidationErrors())
    //    return
    //}

    //q := core.Db.Model(dbmodels.BuLayer{}).Save(&updateModel)
	//
    //if q.Error != nil {
    //    err = q.Error
    //    return
    //}

    //data = AssignBuLayerTypeFromDb(updateModel)
    return
}


func BuLayerDelete(filter types.BuLayerFilter)  (isOk bool, err error) {

    filter.Pagination.CurrentPage = 1
    filter.Pagination.PerPage = 1

    existsModel, err := BuLayerRead(filter)

    if existsModel.Id < 1 || err != nil {

        if err != nil {
            err = errors.New("BuLayer not found in db")
        }
        return
    }
	//
    //q := core.Db.Model(dbmodels.BuLayer{}).Where(dbmodels.BuLayer{ID: existsModel.Id}).Delete(&existsModel)
	//
    //if q.Error != nil {
    //    err = q.Error
    //    return
    //}

    isOk = true
    return
}
