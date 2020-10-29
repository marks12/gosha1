package cmd

const usualSettingsRoutesConstEntity = `
const {Entity}Route = "/api/v1/{entity}"

// route-constant-generator here dont touch this line`

const usualSettingsRoutesListEntity = `	{Entity}Route,
    // router-list-generator here dont touch this line`

var usualTemplateSettingsRoutesConstEntity = template{
	Path:    "./path_error.txt",
	Content: usualSettingsRoutesConstEntity,
}

var usualTemplateSettingsRoutesListEntity = template{
	Path:    "./path_error.txt",
	Content: usualSettingsRoutesListEntity,
}
