package cmd


type osArgumentVal string
type command string

const (
	SetAppType command = "setAppType"
	Exit command = "exit"

	Type osArgumentVal = "type"
	Entity osArgumentVal = "entity"
	CRUD osArgumentVal = "crud"
	CheckAuth osArgumentVal = "check-auth"
	UuidAsPk osArgumentVal = "uuidaspk"

	WithoutDbModels osArgumentVal = "without-db-models"
)

func (arg osArgumentVal) CliArgument(value string) string {
	return " --" + string(arg) + "=" + value
}

func (arg osArgumentVal) ToString() string {
	return string(arg)
}

func (arg command) CliArgument() string {
	return string(arg)
}
