package cmd

var usualDbmodelEntity = `package dbmodels

import (
    "time"
    "../types"
    "gorm.io/gorm"
)

type Entity struct {
    ID    int
    ` + getRemoveLine("Entity") + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func (entity *Entity) Validate()  {
    ` + getRemoveLine("Validate") + `
}

func (entity *Entity) FillFromRequest(request types.Entity) {
    ` + getRemoveLine("FillFromRequest") + `
}
`

const usualDbmodelValidator = `package dbmodels

import (
	"{ms-name}/errors"
)

type validator struct {
	validationError errors.ValidatorError
}

func (val *validator) IsValid() bool {

	return val.validationError.IsEmpty()
}

func (val *validator) GetValidationError() errors.ValidatorErrorInterface {
	return &val.validationError
}

func (val *validator) AddValidationError(err string, code errors.ErrorCode, field string) {
	val.validationError.AddError(errors.NewErrorWithCode(err, code, field))
}
`

/*var usualTemplateDbmodelsEntity = template{
    Path:    "./dbmodels/entity.go",
    Content: usualDbmodelEntity,
}*/

var usualTemplateDbmodelsValidator = template{
	Path:    "./dbmodels/validator.go",
	Content: usualDbmodelValidator,
}
