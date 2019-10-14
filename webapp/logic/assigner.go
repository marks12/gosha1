package logic

import (
	"gosha/dbmodels"
	"gosha/webapp/types"
)


// add all assign functions


func AssignBuLayerTypeFromDb(dbBuLayer dbmodels.BuLayer) types.BuLayer {

    //AssignBuLayerTypeFromDb predefine remove this line for disable generator functionality

    return types.BuLayer{
        Id: dbBuLayer.ID,
        //AssignBuLayerTypeFromDb.Field remove this line for disable generator functionality
    }
}

func AssignBuLayerDbFromType(typeModel types.BuLayer) dbmodels.BuLayer {

    //AssignBuLayerDbFromType predefine remove this line for disable generator functionality
    
    return dbmodels.BuLayer{
        ID: typeModel.Id,
        //AssignBuLayerDbFromType.Field remove this line for disable generator functionality
    }
}




