package cmd

const usualSettingsResource = `package settings

var ExtResources = []string{
    "ws",
}
`

var usualTemplateSettingsResource = template{
	Path:    "./settings/resource.go",
	Content: usualSettingsResource,
}
