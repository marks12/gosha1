package cmd

const msDbmodel = `package dbmodels

import (
    "time"
    "strings"
    "github.com/google/uuid"
)

type Entity struct {
    ID          		uuid.UUID	` + "`" + `sql:"primary_key;type:uuid;default:uuid_generate_v4()"` + "`" + `
    AppId          		uuid.UUID

    validationErrors	[]string

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index"` + "`" + `
}


func (entity *Entity) Validate()  {

}

func (entity *Entity) IsValid() bool {
    return len(entity.validationErrors) < 1
}

func (entity *Entity) GetValidationErrors() string {

    return strings.Join(entity.validationErrors, ". ")
}`

var msTemplateDbmodelsEntity = template{
    Path:    "./dbmodels/entity.go",
    Content: msDbmodel,
}

