package cmd

const usualHtmlTemplateReadLogic = `
	fmt.Println("route")

	if err != nil {
		return
	}

	key := view.TemplateKey{}
	e := key.Parse(filter.Key)

	if e != nil {
		return data, e
	}

	tpl, e := 	view.Get(key)
	if e != nil {
		return data, e
	}

	data = types.HtmlTemplate{
		Id:   0,
		Key:  tpl.Key,
		Html: string(tpl.HtmlText),
	}
`

const usualHtmlTemplateFindLogic = `
	all, err := view.GetAll()

	if err != nil {
		fmt.Println("err", err)
		return
	}

	for i, tpl := range all {

		isValidNamespace := len(filter.Namespace) < 1 || tpl.TemplateKey.Namespace == filter.Namespace
		isValidKey := len(filter.Key) < 1 || tpl.TemplateKey.Value == filter.Key

		if i <= filter.PerPage && isValidNamespace && isValidKey {
			result = append(result, types.HtmlTemplate{
				Id:   0,
				Key:  tpl.Key,
				Html: string(tpl.HtmlText),
			})
		}
	}`

const usualViewStore = `package view

import (
	"{ms-name}/view/bs4"
	"{ms-name}/settings"
	"{ms-name}/types"
	"{ms-name}/core"
	"{ms-name}/common"
	"bytes"
	"fmt"
	"errors"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

type Namespace string

type NamespaceInterface interface {
	GetNamespace() Namespace
}

func (namespace Namespace) ToString() string {
	return string(namespace)
}

func (namespace Namespace) GetNamespace() Namespace {
	return namespace
}

const Empty Namespace = ""
const StringField Namespace = "StringField"
const BoolField Namespace = "BoolField"
const Panel Namespace = "Panel"
const RightPanelContainer Namespace = "RightPanelContainer"

// need to check on search template. Always add new Namespace to arr
var Namespaces = []string{
	Empty.ToString(),
	Panel.ToString(),
	RightPanelContainer.ToString(),
	StringField.ToString(),
	BoolField.ToString(),
	//Namespace remove this line for disable generator functionality
}

const KeySeparator = "/"

type TemplateKey struct {
	Namespace string
	Value     string
}

type Template struct {
	TemplateKey TemplateKey
	Key         string
	HtmlObject  bs4.HtmlInterface
	HtmlText    []byte
}

func GetHtmlTemplateUrl(ns NamespaceInterface, dbmodel interface{}) string {

	val := ""

	if v, ok := dbmodel.(NamespaceInterface); ok == true {
		val = v.GetNamespace().ToString()
	} else {
		val = core.GetTableName(dbmodel)
	}

	template := TemplateKey{
		Namespace: ns.GetNamespace().ToString(),
		Value: val,
	}

	filter := types.HtmlTemplateFilter{}

	return settings.HtmlTemplateRoute + "/0?" + common.GetFieldName(&filter, &filter.Key) + "=" + template.GetKey()
}

func (key *TemplateKey) Parse(k string) error {

	keys := strings.Split(k, KeySeparator)

	if len(keys) != 2 {
		return errors.New("Wrong key " + k + ". It should have format Namespace/Key")
	}

	for _, ns := range Namespaces {
		if ns == keys[0] {
			key.Namespace = ns
		}
	}

	key.Value = keys[1]

	return nil

}

func (key TemplateKey) GetKey() string {
	return key.Namespace + KeySeparator + key.Value
}

var instance *sync.Map
var once sync.Once

func getStore() *sync.Map {

	once.Do(func() {
		instance = &sync.Map{}
	})

	return instance
}

func Get(dataKey TemplateKey) (data Template, err error) {

	if err != nil {
		fmt.Println("Cant create template key for: ", dataKey.GetKey())
	}
	res, ok := getStore().Load(dataKey.GetKey())
	if ok != true {
		err = errors.New("Template not found by key: " + dataKey.GetKey())
		return
	}
	data = res.(Template)
	return
}

func GetAll() (data []Template, err error) {

	getStore().Range(func(key, value interface{}) bool {
		data = append(data, value.(Template))
		return true
	})

	return
}

func renderTemplate(htmlInterface bs4.HtmlInterface) []byte {

	node := htmlInterface.GetNode()

	b := new(bytes.Buffer)
	if err := html.Render(b, node); err != nil {
		fmt.Println("Error render document", err)
	}

	return b.Bytes()

}

func Save(dataKey TemplateKey, bs4Doc bs4.HtmlInterface) (err error) {

	value := Template{
		TemplateKey: dataKey,
		Key:         dataKey.GetKey(),
		HtmlObject:  bs4Doc,
		HtmlText:    renderTemplate(bs4Doc),
	}

	getStore().Store(dataKey.GetKey(), value)
	return
}

func Remove(dataKey TemplateKey) {
	getStore().Delete(dataKey.GetKey())
}

`
const usualEntityBs4FormTemplate = `package form

import (
	"{ms-name}/types"
	"{ms-name}/view"
    "{ms-name}/view/bs4"
    "{ms-name}/common"
    "{ms-name}/wsserver"
    "{ms-name}/settings"
    "net/http"
)

var Template{Entity}CreateForm = view.TemplateKey{
	Namespace: common.GetTypeName(types.{Entity}{}),
	Value: "form",
}

func Store{Entity}CreateForm() {
	view.Save(Template{Entity}CreateForm, Get{Entity}CreateFormTemplate())
}

func Get{Entity}CreateInstructions() (inst []wsserver.Instruction) {

	{entity} := types.{Entity}{}
	dataKey := "types.{Entity}-fields"

	return []wsserver.Instruction{
		{
			EventName: dataKey,
			CollectFields: []wsserver.FormFields{
				//{Entity}-collector remove this line for disable generator functionality
			},
			RestApiRequests: []wsserver.ApiRequest{
				{
					Url:               settings.{Entity}Route,
					Type:              http.MethodPost,
					DataVar:           dataKey,
					SuccessRespStatus: http.StatusCreated,
					SuccessInstructions: []wsserver.Instruction{
						{
							ConsoleLog: []wsserver.KeyVal{
								{Key: "success", Val: "Save"},
							},
						},
					},
				},
			},
		},
	}
}

func Get{Entity}CreateFormTemplate() (createForm bs4.HtmlInterface) {

	{entity} := types.{Entity}{}

	createForm = &bs4.Form{
		Children: []bs4.HtmlInterface{
			view.GetStringFieldTemplate(view.FieldConfig{
				Id: common.GetFieldName(&{entity}, &{entity}.Id),
			}),
			//{Entity} remove this line for disable generator functionality
		},
	}

	return
}
`

const usualEntityBs4FieldsTemplate = `package view

import (
	"{ms-name}/view/bs4"
	"{ms-name}/view/bs4/css"
)

type FieldConfig struct {
	Id string
	Placeholder string
	DefaultValue string
	IsDisabled bool
}

func GetStringFieldTemplate(config FieldConfig) (field bs4.HtmlInterface) {

	field = &bs4.Block{
		Children: []bs4.HtmlInterface{
			&bs4.InputText{
				Id: config.Id,
				Value: config.DefaultValue,
				Placeholder: config.Placeholder,
				IsDisabled: true,
			},
		},
	}

	return
}

func GetBoolFieldTemplate(config FieldConfig) (field bs4.HtmlInterface) {

	field = &bs4.Block{
		Children: []bs4.HtmlInterface{
			&bs4.Div{
				Classes: []css.Class{css.FormGroup, css.FormCheck, css.Mt3},
				Children: []bs4.HtmlInterface{
					&bs4.Input{
						Id: config.Id,
						Type:"checkbox",
						IsDisabled: config.IsDisabled,
						Classes: []css.Class{css.FormCheckInput},
					},
					&bs4.Label{Text: config.Placeholder, Classes: []css.Class{css.FormCheckLabel}},
				},
			},
		},
	}

	return
}

`

var usualTemplateBs4EntityForms = template{
	Path:    "./view/form/{entity}.go",
	Content: assignMsName(usualEntityBs4FormTemplate),
}

var usualTemplateBs4ViewFields = template{
	Path:    "./view/fields.go",
	Content: assignMsName(usualEntityBs4FieldsTemplate),
}

var usualTemplateViewStore = template{
	Path:    "./view/store.go",
	Content: assignMsName(usualViewStore),
}

var usualTemplateHtmlTemplateFind = template{
	Path:    "./logic/html_template.go",
	Content: usualHtmlTemplateFindLogic,
}

var usualTemplateHtmlTemplateRead = template{
	Path:    "./logic/html_template.go",
	Content: usualHtmlTemplateReadLogic,
}
