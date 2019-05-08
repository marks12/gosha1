package cmd

var usualTemplateLogicAssignerEntity = template{
    Path:    "",
    Content: GetUsualTemplateAssignContent(TypeConfig{true}),
}

func GetUsualTemplateAssignContent(config TypeConfig) string {

    var usualWebappEntityAssigner = `
// add all assign functions

func Assign{Entity}TypeFromDb(db{Entity} dbmodels.{Entity}) types.{Entity} {

    //Assign{Entity}TypeFromDb predefine ` + removeLineComment + `

    return types.{Entity}{
        ` + getIdFromDB(config) + `
        //Assign{Entity}TypeFromDb.Field ` + removeLineComment + `
    }
}

func Assign{Entity}DbFromType(typeModel types.{Entity}) dbmodels.{Entity} {

    //Assign{Entity}DbFromType predefine ` + removeLineComment + `
    
    return dbmodels.{Entity}{
        ` + getIdFromType(config) + `
        //Assign{Entity}DbFromType.Field ` + removeLineComment + `
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
