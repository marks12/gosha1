package types

import (
	"net/http"
	"strconv"
)

type Category struct {
	Id               int
	Name             string
	SourceID         int
	SourceCatID      int
	ParentID         int
	SourceParentID   int
	SourceParentName string
	SourceName       string
	IsOwn            bool
}

func (category *Category) Validate() {

}

type CategoryFilter struct {
	TestFilter int
	model      Category
	AbstractFilter
}

func GetCategoryFilter(request *http.Request, functionType string) CategoryFilter {

	var filter CategoryFilter

	filter.request = request
	filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *CategoryFilter) GetCategoryModel() Category {

	filter.model.Validate()

	return filter.model
}
