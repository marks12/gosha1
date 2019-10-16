package settings

const ServerPort = "4343"

const FunctionTypeFind		= "00000000-0000-0000-0000-000000000001"
const FunctionTypeRead		= "00000000-0000-0000-0000-000000000002"
const FunctionTypeUpdate	= "00000000-0000-0000-0000-000000000003"
const FunctionTypeDelete	= "00000000-0000-0000-0000-000000000004"
const FunctionTypeCreate	= "00000000-0000-0000-0000-000000000005"

const TimeLinkDataType = "*time.Time"
const TimeDataType = "time.Time"

var SupportedModelFieldDataTypes = []string{
	"string",
	"int",
	TimeDataType,
	TimeLinkDataType,
	"float64",
	"bool",
	"uuid.UUID",
	"*int",
}

var SupportedFilterDataTypes = []string{
	"string",
	"int",
	"float64",
	"bool",
	"uuid.UUID",
	"Object",
	"Array",
}

var UsualDefaultStructure = []string{
	"logic",
	"mdl",
	"ms",
	"router",
	"services",
	"settings",
	"static",
	"types",
	"webapp",
}

const FieldTypeModel = "model"
const FieldTypeFilter = "filter"

