package cmd

const usualWebappEntityAssigner = `
// add all assign functions

func assign{Entity}TypeFromDb(db{Entity} dbmodels.{Entity}) types.{Entity} {

    return types.{Entity}{
        Id: db{Entity}.ID,
        Name: db{Entity}.Name,
    }
}

func assign{Entity}DbFromType(typeModel types.{Entity}) dbmodels.{Entity} {

    return dbmodels.{Entity}{
        ID: typeModel.Id,
        Name: typeModel.Name,
    }
}

`

var usualTemplateLogicAssignerEntity = template{
    Path:    "",
    Content: usualWebappEntityAssigner,
}
