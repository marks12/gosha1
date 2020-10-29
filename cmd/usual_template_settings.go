package cmd

const usualSettingsWebApp = `package settings

const ServerPort = "7005"

const FunctionTypeFind		= "00000000-0000-0000-0000-000000000001"
const FunctionTypeRead		= "00000000-0000-0000-0000-000000000002"
const FunctionTypeUpdate	= "00000000-0000-0000-0000-000000000003"
const FunctionTypeMultiUpdate	= "00000000-0000-0000-0000-000000000023"
const FunctionTypeDelete	= "00000000-0000-0000-0000-000000000004"
const FunctionTypeMultiDelete	= "00000000-0000-0000-0000-000000000024"
const FunctionTypeCreate	= "00000000-0000-0000-0000-000000000005"
const FunctionTypeMultiCreate	= "00000000-0000-0000-0000-000000000025"
const FunctionTypeFindOrCreate	= "00000000-0000-0000-0000-000000000006"
const FunctionTypeUpdateOrCreate	= "00000000-0000-0000-0000-000000000007"
const FunctionTypeMultiFindOrCreate	= "00000000-0000-0000-0000-000000000026"

`

const usualSettingsRoutes = `package settings

const HomePageRoute = "/api"

// routes as app resource
const HttpRouteResourceType ConfigId = 1
// web socket resource type
const WsResourceType ConfigId = 2
// html template resource type
const HtmlResourceType ConfigId = 3

// route-constant-generator here dont touch this line

var RoutesArray = []string{

    // router-list-generator here dont touch this line
}
`

var usualTemplateSettingsWebapp = template{
	Path:    "./settings/web_app.go",
	Content: usualSettingsWebApp,
}

var usualTemplateSettingsRoutes = template{
	Path:    "./settings/routes.go",
	Content: usualSettingsRoutes,
}
