package cmd

const msSettingsApp = `package settings

import (
    "os"
    "fmt"
)

const RabbitServerPassword = "rabbit-pass"
const RabbitServerLogin = "msproduct"
const RabbitServer = "core.140140.ru"
const RabbitServerPort = "5672"
const rabbitServerVirtualhost = "/microservices"
const rabbitServerVirtualhostTest = "/microservices-test"

const MicroserviceAuthKey = "new-guid";


func GetVirtualHost() string {

    if len(os.Args) < 1 {
        fmt.Println("RPC use PRODUCTION environment")
        return rabbitServerVirtualhost
    }

    for _, param := range os.Args {

        if param == "dev" {
            fmt.Println("RPC use DEV environment")
            return rabbitServerVirtualhostTest
        }
    }

    fmt.Println("RPC use PRODUCTION environment")
    return rabbitServerVirtualhost
}`

const msSettingsDb = `package settings

const DbHost = "127.0.0.1"

const DbPort = "5432"

const DbUser = "msproduct"

const DbPass = "prd00dfghLLvQuw"

const DbName = "msproduct"
`

var msTemplageSettingsAppContent =
        microserviceNameRegexp.ReplaceAllString(
            newGuidRegexp.ReplaceAllString(
                newPassRegexp.ReplaceAllString(msSettingsApp, generatePassword(8)),
                getNewGuid()),
        getCurrentDirName())

var msTemplageSettingsDbContent = newPassRegexp.ReplaceAllString(msSettingsDb, generatePassword(8))

var msTemplateSettingsApp = template{
    Path:    "./settings/app.go",
    Content: msTemplageSettingsAppContent,
}

var msTemplateSettingsDb = template{
    Path:    "./settings/db.go",
    Content: msTemplageSettingsDbContent,
}

