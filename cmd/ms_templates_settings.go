package cmd

const msSettingsApp = `package settings

import (
    "os"
    "fmt"
    "regexp"
)

const RabbitServerPassword = "{new-pass}"
const RabbitServerLogin = "{ms-name}"
const RabbitServer = "core.140140.ru"
const RabbitServerPort = "5672"
const rabbitServerVirtualhost = "/microservices"
const rabbitServerVirtualhostTest = "/microservices-test"

const MicroserviceAuthKey = "{new-guid}"

func IsDev() bool {
    var matchDev = regexp.MustCompile("^/tmp/go-build")
    return matchDev.Match([]byte(os.Args[0]))
}

func GetVirtualHost() string {

    for _, param := range os.Args {

        if param == "dev" {
            fmt.Println("RPC use DEV environment")
            return rabbitServerVirtualhostTest
        }
    }

    if IsDev() {
        fmt.Println("RPC use DEV environment")
        return rabbitServerVirtualhostTest
    }

    fmt.Println("RPC use PRODUCTION environment")
    return rabbitServerVirtualhost
}`

const msSettingsDb = `package settings

const DbHost = "127.0.0.1"

const DbPort = "5432"

const DbUser = "{ms-name}"

const DbPass = "{new-pass}"

const DbName = "{ms-name}"
`

var msTemplageSettingsAppContent =
        assignMsName(
            assignGuid(
                assignPass(
                    msSettingsApp)))

var msTemplageSettingsDbContent =
        assignMsName(
            assignPass(
                msSettingsDb))

var msTemplateSettingsApp = template{
    Path:    "./settings/app.go",
    Content: msTemplageSettingsAppContent,
}

var msTemplateSettingsDb = template{
    Path:    "./settings/db.go",
    Content: msTemplageSettingsDbContent,
}

