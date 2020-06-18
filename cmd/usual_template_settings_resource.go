package cmd

const usualSettingsResource = `package settings

var ExtResources = []string{
    "ws",
}

const HttpResource ConfigId = 1
`

var usualTemplateSettingsResource = template{
	Path:    "./settings/resource.go",
	Content: usualSettingsResource,
}
