package cmd

const usualSettingsResource = `package settings

const WsResource = "ws"

var Resources = []string{
    WsResource,
}
`

var usualTemplateSettingsResource = template{
	Path:    "./settings/resource.go",
	Content: usualSettingsResource,
}
