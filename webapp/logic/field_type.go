package logic

import (
    "gosha/settings"
    "gosha/webapp/types"
)

func FieldTypeFind(filter types.FieldTypeFilter)  (result []types.FieldType, totalRecords int, err error) {

    for i, t := range settings.SupportedModelFieldDataTypes {
        result = append(result, types.FieldType{
            Id: i,
            Name: t,
            Type: settings.FieldTypeModel,
        })
    }

    for i, t := range settings.SupportedFilterDataTypes {

        result = append(result, types.FieldType{
            Id: i,
            Name: t,
            Type:  settings.FieldTypeFilter,
        })
    }

    // order global criteria

    // получение связанных сущностей

    return result, len(result), nil
}
