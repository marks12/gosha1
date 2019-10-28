package cmd

const msSettingsApp = `package settings

import (
    "os"
    "fmt"
    "regexp"
    "{ms-name}/flags"
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
    return matchDev.Match([]byte(os.Args[0])) || *flags.IsDev
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

const msSettingsGoogle = `package settings

const GoogleAnalyticsUrl = "https://www.google-analytics.com/collect"
const GoogleTrackingId = "G-6EQD27PJGY"

`

const msSettingsDb = `package settings

const DbHost = "127.0.0.1"

const DbPort = "35432"

const DbUser = "{ms-name}"

const DbPass = "{new-pass}"

const DbName = "{ms-name}"
`

const msSettingsWss = `package settings

const WssPortDev = "5600"

const WssPortProd = "5050"

func GetWssPort() string {

    if IsDev() {
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

var msTemplateSettingsGoogleContent =
        assignMsName(
            assignGuid(
                assignPass(
                    msSettingsGoogle, dbPass)))

var msTemplateSettingsDbContent =
        assignMsName(
            assignPass(
                msSettingsDb, dbPass))

var msTemplateSettingsWssContent = assignMsName(msSettingsWss)

var msTemplateSettingsApp = template{
    Path:    "./settings/app.go",
    Content: msTemplateSettingsAppContent,
}

var msTemplateSettingsGoogle = template{
    Path:    "./settings/google.go",
    Content: msTemplateSettingsGoogleContent,
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

