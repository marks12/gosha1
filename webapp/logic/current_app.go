package logic

import (
    "gosha/webapp/types"
    "errors"
    "gosha/cmd"
)

func CurrentAppFind(filter types.CurrentAppFilter)  (result []types.CurrentApp, totalRecords int, err error) {

    app := cmd.GetCurrentApp()

    result = append(result, types.CurrentApp{
        IsValidStructure: app.IsAppExists,
    })

    return result, len(result), nil
}

func CurrentAppCreate(filter types.CurrentAppFilter)  (data types.CurrentApp, err error) {

    return
}

func CurrentAppRead(filter types.CurrentAppFilter)  (data types.CurrentApp, err error) {

    findData, _, err := CurrentAppFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.CurrentApp{}, errors.New("Not found")
}
