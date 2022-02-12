package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/fatih/color"
	"os"
)

func usualCreateHtmlTemplate(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()
	c.Println(yellow("Hello we start adding html template to app"))

	os.Args = append(os.Args, "--entity=HtmlTemplate")
	os.Args = append(os.Args, "--check-auth=fr")
	os.Args = append(os.Args, "--crud=fr")
	os.Args = append(os.Args, "--without-db-models=true")
	usualEntityAdd(c)
	os.Args = os.Args[:len(os.Args)-4]

	argsBak := os.Args

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=HtmlTemplate", "--Field=Key", "--data-type=string"}
	entityFieldAdd(c)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FIELD, "--entity=HtmlTemplate", "--Field=Html", "--data-type=string"}
	entityFieldAdd(c)

	c.Println("Models created success")

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FILTER, "--entity=HtmlTemplateFilter", "--filter=Namespace", "--data-type=string"}
	entityFilterdAdd(nil)

	os.Args = []string{"", "exit", "setAppType", "--type=Usual", ENTITY_ADD_FILTER, "--entity=HtmlTemplateFilter", "--filter=Key", "--data-type=string"}
	entityFilterdAdd(nil)

	os.Args = argsBak

	CopyFile(
		usualTemplateHtmlTemplateFind.Path,
		usualTemplateHtmlTemplateFind.Path,
		[]string{"//HtmlTemplate Find logic code"},
		[]string{usualTemplateHtmlTemplateFind.Content},
		c)

	CopyFile(
		usualTemplateHtmlTemplateRead.Path,
		usualTemplateHtmlTemplateRead.Path,
		[]string{"//HtmlTemplate Read logic code"},
		[]string{usualTemplateHtmlTemplateRead.Content},
		c)

	addImportIfNeed(usualTemplateHtmlTemplateRead.Path, "fmt")
	addImportIfNeed(usualTemplateHtmlTemplateRead.Path, GetCurrentAppName()+"/view")

}
