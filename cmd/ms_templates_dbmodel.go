package cmd

const msDbmodelEntity = `package dbmodels

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Entity struct {
    ID          		uuid.UUID	` + "`" + `sql:"primary_key;type:uuid;default:uuid_generate_v4()"` + "`" + `
    AppId          		uuid.UUID

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt ` + "`" + `sql:"index"` + "`" + `

    validator
}


func (entity *Entity) FillFromRequest(request api.Entity) {
    
}

func (entity *Entity) Validate()  {

}
`

const msDbmodelValidator = `package dbmodels

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

var msTemplateDbmodelsEntity = template{
    Path:    "./dbmodels/entity.go",
    Content: msDbmodelEntity,
}

var msTemplateDbmodelsValidator = template{
    Path:    "./dbmodels/validator.go",
    Content: msDbmodelValidator,
}

