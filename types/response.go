package types

import (
	"lucy/settings"
)

type APIStatus struct {
	Status string
}

type APIError struct {
	Error bool
	ErrorMessage string
}

type Pagination struct {

	TotalRecords	int
	TotalPages		int
	CurrentPage		int
	PerPage			int

	validator
}

func (pagination *Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}

func (pagination *Pagination) Validate(functionType string) {

	switch functionType {

	case settings.FunctionTypeFind:

		if pagination.CurrentPage < 1 {
			pagination.validator.validationErrors = append(pagination.validator.validationErrors, "CurrentPage parameter is not set")
		}

		if pagination.PerPage < 1 {
			pagination.validator.validationErrors = append(pagination.validator.validationErrors, "PerPage parameter is not set")
		}

		break
	case settings.FunctionTypeCreate:



		break
	case settings.FunctionTypeRead:



		break
	case settings.FunctionTypeUpdate:



		break
	case settings.FunctionTypeDelete:

		break

	default:
		pagination.validator.validationErrors = append(pagination.validator.validationErrors, "Usupported function type: " + functionType)
		break
	}
}
