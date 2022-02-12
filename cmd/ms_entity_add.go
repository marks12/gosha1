package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"strings"
	"github.com/fatih/color"
)

func msEntityAdd(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()

	c.Println(yellow("Hello we start creating api for new entity"))

	entity, err := getName(c, false, "Entity")

	if err != nil {
		return
	}

	CamelCase := strings.Title(entity)
	snakeCase := getLowerCase(entity)
	FirstLowerCase := GetFirstLowerCase(entity)

	sourceFile := "./dbmodels/entity.go"
	destinationFile := "./dbmodels/" + snakeCase + ".go"

	CopyFile(
		sourceFile,
		destinationFile,
		[]string{"Entity", "entity"},
		[]string{CamelCase, FirstLowerCase},
		c)

	tpl := msTemplateLogicEntity.Content

	CreateFile("./logic/"+snakeCase+".go", assignMsName(assignEntityName(tpl, CamelCase)), c)

	//AppendFile(msTemplateLogicAssignEntity.Path, assignEntityName(msTemplateLogicAssignEntity.Content, CamelCase))
	AppendFile("./logic/"+snakeCase+".go", assignEntityName(msLogicAssignEntity, CamelCase))

	CopyFile("./main.go", "./main.go", []string{"RabbitServer.Run()"}, []string{msTemplateMainEntity.Content}, c)
	CopyFile("./main.go", "./main.go", []string{"{entity-name}"}, []string{CamelCase}, c)

	AppendFile(msTemplateMsTicketEntity.Path, assignEntityName(msTemplateMsTicketEntity.Content, CamelCase))

	CreateFile("./rpcapp/"+snakeCase+".go", assignEntityName(msTemplateRpcappEntity.Content, CamelCase), c)
}
