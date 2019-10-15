package logic

import (
    "gosha/webapp/types"
    "errors"
    "gosha/cmd"
    "os"
)

func CurrentAppFind(filter types.CurrentAppFilter)  (result []types.CurrentApp, totalRecords int, err error) {

    app := cmd.GetCurrentApp()

    result = append(result, types.CurrentApp{
        IsValidStructure: app.IsAppExists,
    })

    return result, len(result), nil
}

func CurrentAppCreate(filter types.CurrentAppFilter)  (data types.CurrentApp, err error) {

    argsBak := os.Args
    defer func(){os.Args = argsBak}()

    args := []string{"", "exit", "setAppType", "--type=Usual", "usual:create"}

    os.Args = args
    cmd.RunShell()

    defer func(){
        shell := cmd.GetShell()
        shell.Close()
    }()

    return
}

func CurrentAppRead(filter types.CurrentAppFilter)  (data types.CurrentApp, err error) {

    findData, _, err := CurrentAppFind(filter)

    if len(findData) > 0 {
        return findData[0], nil
    }

    return types.CurrentApp{}, errors.New("Not found")
}
