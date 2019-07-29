package mdl

type Response struct {
	Status       int
	Error        bool
	ErrorMessage string
	Data         interface{}
}

type ResponseFind struct {
	List  interface{}
	Total int
}

type ResponseCreate struct {
	Model interface{}
}

type ResponseRead struct {
	Model interface{}
}

type ResponseUpdate struct {
	Model interface{}
}

type ResponseDelete struct {
	IsSuccess bool
}

type ResponseFindOrCreate struct {
	Model interface{}
}
