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

const msSettingsWss = `package settings

import (
    "{ms-name}/flags"
)

const WssPortDev = "5500"

const WssPortProd = "5050"

func GetWssPort() string {

    if flags.IsDev {
        return WssPortDev
    }

    return WssPortProd
}
`

const msTemplageSettingsWebAppContent = `package settings

const ServerPort = "7005"

`
var dbPass = generatePassword(8)

var msTemplateSettingsAppContent =
        assignMsName(
            assignGuid(
                assignPass(
                    msSettingsApp, dbPass)))

var msTemplateSettingsDbContent =
        assignMsName(
            assignPass(
                msSettingsDb, dbPass))

var msTemplateSettingsWssContent = assignMsName(msSettingsWss)

var msTemplateSettingsApp = template{
    Path:    "./settings/app.go",
    Content: msTemplateSettingsAppContent,
}

var msTemplateSettingsDb = template{
    Path:    "./settings/db.go",
    Content: msTemplateSettingsDbContent,
}

var msTemplateSettingsWss = template{
    Path:    "./settings/wss.go",
    Content: msTemplateSettingsWssContent,
}

var msTemplateSettingsWebApp = template{
    Path:    "./settings/web_app.go",
    Content: msTemplageSettingsWebAppContent,
}

