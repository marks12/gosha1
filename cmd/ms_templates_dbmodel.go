package cmd

const msDbmodelEntity = `package dbmodels

import (
    "time"
    "github.com/google/uuid"
)

type Entity struct {
    ID          		uuid.UUID	` + "`" + `sql:"primary_key;type:uuid;default:uuid_generate_v4()"` + "`" + `
    AppId          		uuid.UUID

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index"` + "`" + `

    validator
}

func (entity *Entity) Validate()  {

}
`

const msDbmodelValidator = `package dbmodels

import "strings"

type validator struct {
    validationErrors	[]string
}

func (val *validator) IsValid() bool {
    return len(val.validationErrors) < 1
}

func (val *validator) GetValidationErrors() string {

    return strings.Join(val.validationErrors, ". ")
}
`

var msTemplateDbmodelsEntity = template{
    Path:    "./dbmodels/entity.go",
    Content: msDbmodelEntity,
}

var msTemplateDbmodelsValidator = template{
    Path:    "./dbmodels/validator.go",
    Content: msDbmodelValidator,
}

