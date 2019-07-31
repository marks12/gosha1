package types

import (
	"net/http"
	"strconv"
)

type EntityField struct {
	Id int
	Name     string
	DataType string
	//EntityField remove this line for disable generator functionality
}

func (unit *EntityField) Validate() {

	//Validate remove this line for disable generator functionality
}

type EntityFieldFilter struct {
	Cats  []int
	model EntityField
	//EntityFieldFilter remove this line for disable generator functionality

	AbstractFilter
}

func GetEntityFieldFilter(request *http.Request, functionType string) EntityFieldFilter {

	var filter EntityFieldFilter

	filter.request = request

	arr, _ := request.URL.Query()["Cats[]"]

	for _, id := range arr {
		newID, _ := strconv.Atoi(id)
		filter.Cats = append(filter.Cats, newID)
	}

	//GetEntityFieldFilter remove this line for disable generator functionality

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *EntityFieldFilter) GetEntityFieldModel() EntityField {

	filter.model.Validate()

	return filter.model
}
