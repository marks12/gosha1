package settings

const ServerPort = "4343"

const FunctionTypeFind		= "00000000-0000-0000-0000-000000000001"
const FunctionTypeRead		= "00000000-0000-0000-0000-000000000002"
const FunctionTypeUpdate	= "00000000-0000-0000-0000-000000000003"
const FunctionTypeDelete	= "00000000-0000-0000-0000-000000000004"
const FunctionTypeCreate	= "00000000-0000-0000-0000-000000000005"

const DataTypeTimeLink = "*time.Time"
const DataTypeTime = "time.Time"
const DataTypeDuration = "time.Duration"
const DataTypeInt = "int"
const DataTypeIntLink = "*int"
const DataTypeArrayInt = "[]int"
const DataTypeFloat64 = "float64"
const DataTypeBool = "bool"
const DataTypeUuid = "uuid.UUID"
const DataTypeString = "string"
const DataTypeArrayString = "[]string"
const DataTypeArrayBytes = "[]byte"

var SupportedModelFieldDataTypes = []string{
    DataTypeString,
    DataTypeInt,
    DataTypeTime,
	DataTypeDuration,
    DataTypeTimeLink,
    DataTypeFloat64,
    DataTypeBool,
    DataTypeUuid,
    DataTypeIntLink,
	DataTypeArrayBytes,
}

var SupportedFilterDataTypes = []string{
	DataTypeString,
    DataTypeInt,
    DataTypeFloat64,
	DataTypeBool,
	DataTypeUuid,
    DataTypeArrayInt,
	DataTypeArrayString,
	DataTypeTime,
}

var UsualDefaultStructure = []string{
	"logic",
	"mdl",
	"router",
	"services",
	"settings",
	"static",
	"types",
	"webapp",
}

const FieldTypeModel = "model"
const FieldTypeFilter = "filter"

