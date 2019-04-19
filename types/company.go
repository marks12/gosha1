package types

import (
	"net/http"
	"strconv"
)

type Company struct {
	Id      int
	Name    string
	OffName string
	Email   string
}

func (company *Company) Validate() {

}

type CompanyFilter struct {
	TestFilter int
	model      Company
	AbstractFilter
}

func GetCompanyFilter(request *http.Request, functionType string) CompanyFilter {

	var filter CompanyFilter

	filter.request = request
	filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *CompanyFilter) GetCompanyModel() Company {

	filter.model.Validate()

	return filter.model
}
