package settings

const ServerPort = "4343"

const FunctionTypeFind		= "00000000-0000-0000-0000-000000000001"
const FunctionTypeRead		= "00000000-0000-0000-0000-000000000002"
const FunctionTypeUpdate	= "00000000-0000-0000-0000-000000000003"
const FunctionTypeDelete	= "00000000-0000-0000-0000-000000000004"
const FunctionTypeCreate	= "00000000-0000-0000-0000-000000000005"

const DataTypeTimeLink = "*time.Time"
const DataTypeTime = "time.Time"
const DataTypeInt = "int"
const DataTypeIntLink = "*int"
const DataTypeArrayInt = "[]int"
const DataTypeFloat64 = "float64"
const DataTypeBool = "bool"
const DataTypeUuid = "uuid.UUID"
const DataTypeString = "string"
const DataTypeArrayString = "[]string"

var SupportedModelFieldDataTypes = []string{
    DataTypeString,
    DataTypeInt,
    DataTypeTime,
    DataTypeTimeLink,
    DataTypeFloat64,
    DataTypeBool,
    DataTypeUuid,
    DataTypeIntLink,
}

var SupportedFilterDataTypes = []string{
	DataTypeString,
    DataTypeInt,
    DataTypeFloat64,
	DataTypeBool,
	DataTypeUuid,
    DataTypeArrayInt,
	DataTypeArrayString,
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

