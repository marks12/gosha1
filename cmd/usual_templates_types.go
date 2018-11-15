package cmd

const usualTypesAuthenticator = `package types

type Authenticator struct {
    Token string
}

func (auth *Authenticator) IsAuthorized() bool {
    return  true
}
`

const usualTypesEntity = `package types

import (
    "time"
)
// default entity will used when create new entity
type Entity struct {
    ID        uint       ` + "`" + `gorm:"primary_key"` + "`" + `
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func (entity *Entity) Validate()  {

}
`

const usualTypesValidator = `package types

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

const usualTypesResponse = `package types

type APIStatus struct {
	Status string
}

type APIError struct {
	Error bool
	ErrorMessage string
}

type Pagination struct {
	TotalRecords int
	TotalPages int
	CurrentPage int
	PerPage int
}
`

var usualTemplateTypesAuthenticator = template{
    Path:    "./types/authenticator.go",
    Content: usualTypesAuthenticator,
}

var usualTemplateTypesEntity = template{
    Path:    "./types/entity.go",
    Content: usualTypesEntity,
}

var usualTemplateTypesValidator = template{
    Path:    "./types/validator.go",
    Content: usualTypesValidator,
}

var usualTemplateTypesResponse = template{
    Path:    "./types/response.go",
    Content: usualTypesResponse,
}

