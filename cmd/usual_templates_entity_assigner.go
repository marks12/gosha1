package cmd

const usualWebappEntityAssigner = `
// add all assign functions

func Assign{Entity}TypeFromDb(db{Entity} dbmodels.{Entity}) types.{Entity} {

    //Assign{Entity}TypeFromDb predefine ` + removeLineComment + `

    return types.{Entity}{
        Id: db{Entity}.ID,
        //Assign{Entity}TypeFromDb.Field ` + removeLineComment + `
    }
}

func Assign{Entity}DbFromType(typeModel types.{Entity}) dbmodels.{Entity} {

    //Assign{Entity}DbFromType predefine ` + removeLineComment + `
    
    return dbmodels.{Entity}{
        ID: typeModel.Id,
        //Assign{Entity}DbFromType.Field ` + removeLineComment + `
    }
}

`

var usualTemplateLogicAssignerEntity = template{
    Path:    "",
    Content: usualWebappEntityAssigner,
}
