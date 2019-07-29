package mdl

type Pagination struct {
	TotalRecords	int
	TotalPages		int
	CurrentPage		int
	PerPage			int
}

func (pagination Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}