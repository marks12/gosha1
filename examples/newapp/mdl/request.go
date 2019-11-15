package mdl

import "github.com/google/uuid"


type RequestFind struct {
	Filter		interface{}
	Pagination	Pagination
}

type RequestCreate struct {
	Model	interface{}
}

type RequestRead struct {
	GUID uuid.UUID
}

type RequestUpdate struct {
	GUID uuid.UUID
	Model	interface{}
}

type RequestDelete struct {
	GUID uuid.UUID
}

func (request *RequestFind) GetPagination() Pagination {
	return request.Pagination
}

func (request *RequestFind) GetFilter() interface{} {
	return request.Filter
}

func (request *RequestFind) CheckPagination() (offset int, page int, perPage int, isValid bool) {

	pagination := request.GetPagination()

	if pagination.CurrentPage < 1 {
		return 0,0,0, false
	}

	if pagination.PerPage < 1 {
		return 0,0,0, false
	}

	page = pagination.CurrentPage
	perPage = pagination.PerPage

	// Считаем смещение
	offset = page - 1
	offset = offset * perPage

	isValid = true

	return
}
