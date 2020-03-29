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
const GoogleTrackingId = "G-SOMETRACKINGID"

`

const DbHostpostgres = "127.0.0.1"

const dbHostMySql = "172.17.0.2"

const DbUserpostgres = "{ms-name}"

const DbUserMySql = "root"

const DbPortpostgres = "35432"

const DbPortMySql = "3306"

const msSettingsDb = `package settings

const DbHost = "dbhost"

const DbPort = "dbport"

const DbUser = "dbuser"

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

var msTemplateSettingsWssContent = assignMsName(msSettingsWss)

var msTemplateSettingsApp = template{
    Path:    "./settings/app.go",
    Content: msTemplateSettingsAppContent,
}

var msTemplateSettingsGoogle = template{
    Path:    "./settings/google.go",
    Content: msTemplateSettingsGoogleContent,
}

func getMsTemplateSettingsDb(dbtype DatabaseType) template {

    content :=  assignPass(msSettingsDb, dbPass)

    if dbtype.IsMysql {
        content = assignVar(content, "dbhost", dbHostMySql)
        content = assignVar(content, "dbport", DbPortMySql)
        content = assignVar(content, "dbuser", DbUserMySql)
    } else if dbtype.IsPostgres {
        content = assignVar(content, "dbhost", DbHostpostgres)
        content = assignVar(content, "dbport", DbPortpostgres)
        content = assignVar(content, "dbuser", DbUserpostgres)
    }

    content = assignMsName(content)

    return template{
        Path:    "./settings/db.go",
        Content: content,
    }
}

var msTemplateSettingsWss = template{
    Path:    "./settings/wss.go",
    Content: msTemplateSettingsWssContent,
}

var msTemplateSettingsWebApp = template{
    Path:    "./settings/web_app.go",
    Content: msTemplageSettingsWebAppContent,
}

