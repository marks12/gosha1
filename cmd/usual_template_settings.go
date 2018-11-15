package cmd

const usualSettingsWebApp = `package settings

const ServerPort = "7005"

`

var usualTemplateSettingsWebapp = template{
    Path:    "./settings/web_app.go",
    Content: usualSettingsWebApp,
}
