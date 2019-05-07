package cmd

const usualDbmodelEntity = `package dbmodels

import (
    "time"
    "../types"
)

type Entity struct {
    ID        int       ` + "`" + `gorm:"primary_key"` + "`" + `
    //Entity ` + removeLineComment + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func (entity *Entity) Validate()  {
    //Validate ` + removeLineComment + `
}

func (entity *Entity) FillFromRequest(request types.Entity) {
    //FillFromRequest ` + removeLineComment + `
}
`

const usualDbmodelValidator = `package dbmodels

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

var usualTemplateDbmodelsEntity = template{
    Path:    "./dbmodels/entity.go",
    Content: usualDbmodelEntity,
}

var usualTemplateDbmodelsValidator = template{
    Path:    "./dbmodels/validator.go",
    Content: usualDbmodelValidator,
}

