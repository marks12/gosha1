package cmd

import "gosha/mode"

var usualTemplateLogicAssignerEntity = template{
    Path:    "",
    Content: GetUsualTemplateAssignContent(TypeConfig{true, false}),
}

func GetUsualTemplateAssignContent(config TypeConfig) string {

    if config.IsUuid == false {
        config.IsUuid = mode.GetUuidMode()
    }

    var usualWebappEntityAssigner = `
// add all assign functions

func Assign{Entity}TypeFromDb(db{Entity} dbmodels.{Entity}) types.{Entity} {

    ` + getRemoveLine("Assign{Entity}TypeFromDb predefine") + `

    return types.{Entity}{
        ` + getIdFromDB(config) + `
        ` + getRemoveLine("Assign{Entity}TypeFromDb.Field") + `
    }
}

func Assign{Entity}DbFromType(typeModel types.{Entity}) dbmodels.{Entity} {

    ` + getRemoveLine("Assign{Entity}DbFromType predefine") + `
    
    return dbmodels.{Entity}{
        ` + getIdFromType(config) + `
        ` + getRemoveLine("Assign{Entity}DbFromType.Field") + `
    }
}

`
    return usualWebappEntityAssigner
}

func getIdFromType(cfg TypeConfig) (r string) {

    if cfg.IsId {
        r = assignMsName(`ID: typeModel.Id,`)
    }

    return
}

func getIdFromDB(cfg TypeConfig) (r string) {

    if cfg.IsId {
        r = assignMsName(`Id: db{Entity}.ID,`)
    }

    return
}
