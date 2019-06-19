package cmd

const usualMdlPagination = `package mdl

type Pagination struct {
	TotalRecords	int
	TotalPages		int
	CurrentPage		int
	PerPage			int
}

func (pagination Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}
`

const usualMdlRequest = `package mdl

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
`

const usualMdlResponse = `package mdl

type Response struct {
	Status			int
	Error			bool
	ErrorMessage	string
	Data		interface{}
}

type ResponseFind struct {
	List 	interface{}
	Total	int
}

type ResponseCreate struct {
	Model 	interface{}
}

type ResponseFindOrCreate struct {
	Model 	interface{}
}

type ResponseRead struct {
	Model 	interface{}
}

type ResponseUpdate struct {
	Model 	interface{}
}

type ResponseDelete struct {
	IsSuccess	bool
}

type ResponseFindOrCreate struct {
	Model 	interface{}
}

`

var usualTemplateMdlPagination = template{
	Path:    "./mdl/pagination.go",
	Content: assignMsName(usualMdlPagination),
}

var usualTemplateMdlRequest = template{
	Path:    "./mdl/request.go",
	Content: assignMsName(usualMdlRequest),
}

var usualTemplateMdlResponse = template{
	Path:    "./mdl/response.go",
	Content: assignMsName(usualMdlResponse),
}
