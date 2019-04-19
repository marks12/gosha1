package types

import (
	"net/http"
	"strconv"
)

type Comment struct {
	Id          int
	AuthorID    int
	RecepientID int
	OrderID     int
	OrderListID int
}

func (comment *Comment) Validate() {

}

type CommentFilter struct {
	TestFilter int
	model      Comment
	AbstractFilter
}

func GetCommentFilter(request *http.Request, functionType string) CommentFilter {

	var filter CommentFilter

	filter.request = request
	filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *CommentFilter) GetCommentModel() Comment {

	filter.model.Validate()

	return filter.model
}
