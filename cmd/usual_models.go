package cmd

const usualModelsInit = `package entity

func InitEntities() {

	//initEntity remove this line for disable generator functionality

}
`

const usualModelsStore = `package models

import (
	"sync"
	"fmt"
	"errors"
	"{ms-name}/wsserver"
	"{ms-name}/view/bs4"
)

type Entity struct {
	EntityName            string
	DbModel               interface{}
	TypeModel             interface{}
	GetCreateInstructions func() []wsserver.Instruction
	FormTemplate          func(interface{}) bs4.HtmlInterface
	Route                 string
}

var instance *sync.Map
var once sync.Once

func getStore() *sync.Map {

	once.Do(func() {
		instance = &sync.Map{}
	})

	return instance
}

func Get(entityName string) (data Entity, err error) {

	if err != nil {
		fmt.Println("Cant get Entity key for: ", entityName)
	}
	res, ok := getStore().Load(entityName)
	if ok != true {
		err = errors.New("Entity not found by key: " + entityName)
		return
	}
	data = res.(Entity)
	return
}

func GetAll() (data []Entity, err error) {

	getStore().Range(func(key, value interface{}) bool {
		data = append(data, value.(Entity))
		return true
	})

	return
}

// Saving Entity to memory store
func Save(entities ...Entity) (err error) {

	for _, ent := range entities {
		getStore().Store(ent.EntityName, ent)
	}

	return
}

func Remove(entityName string) {
	getStore().Delete(entityName)
}
`

const usualEntityModel = `package entity

import (
	"{ms-name}/models"
	"{ms-name}/dbmodels"
	"{ms-name}/types"
	"{ms-name}/common"
	"{ms-name}/view/form"
	"{ms-name}/settings"
)

func init{Entity}() {

    type Entity struct {
	    EntityName            string
    	DbModel               interface{}
	    TypeModel             interface{}
	    GetCreateInstructions func() []wsserver.Instruction
	    GetFormTemplate       func(interface{}) bs4.HtmlInterface
	    Route                 string
    }

	form.Store{Entity}CreateForm()
	form.Store{Entity}RowTemplate()
}
`

var usualTemplateModelsInit = template{
	Path:    "./models/entity/init.go",
	Content: assignMsName(usualModelsInit),
}

var usualTemplateModelsStore = template{
	Path:    "./models/store.go",
	Content: assignMsName(usualModelsStore),
}

var usualTemplateModelsEntity = template{
	Path:    "./models/entity/{entity}.go",
	Content: assignMsName(usualEntityModel),
}