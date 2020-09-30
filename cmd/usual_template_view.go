package cmd


const usualEntityBs4FormTemplate = `package generator

import (
	"{ms-name}/types"
	"{ms-name}/view"
    "{ms-name}/view/bs4"
    "{ms-name}/common"
)

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

import "{ms-name}/view/bs4"

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

`

var usualTemplateBs4EntityForms = template{
	Path:    "./view/form/{entity}.go",
	Content: assignMsName(usualEntityBs4FormTemplate),
}

var usualTemplateBs4ViewFields = template{
	Path:    "./view/fields.go",
	Content: assignMsName(usualEntityBs4FieldsTemplate),
}
