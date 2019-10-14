package logic

import (
    "gosha/webapp/types"
    "gosha/settings"
)

func FieldTypeFind(filter types.FieldTypeFilter)  (result []types.FieldType, totalRecords int, err error) {

    for i, t := range settings.SupportedFieldDataTypes {
        result = append(result, types.FieldType{
            Id: i,
            Name: t,
        })
    }

    // order global criteria

    // получение связнаных сущностей

    return result, len(result), nil
}
