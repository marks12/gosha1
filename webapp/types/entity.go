package types

import (
	"net/http"
	"strconv"
	"gosha/cmd"
)


type Field struct {
	Name string
	Type string
	IsDb bool
	IsType bool
}

type Entity struct {
	Id int
	Name     string
	TypeFields	 []cmd.Field
	ModelFields	 []cmd.Field
	Fields	 []Field
	IsFilter bool
	//Entity remove this line for disable generator functionality
}

func (unit *Entity) Validate() {

	//Validate remove this line for disable generator functionality
}

type EntityFilter struct {
	Cats  []int
	model Entity
	WithFilter bool
	WithHiddenFields bool
	//EntityFilter remove this line for disable generator functionality

	AbstractFilter
}

func GetEntityFilter(request *http.Request, functionType string) EntityFilter {

	var filter EntityFilter

	filter.request = request

	arr, _ := request.URL.Query()["Cats[]"]

	for _, id := range arr {
		newID, _ := strconv.Atoi(id)
		filter.Cats = append(filter.Cats, newID)
	}

	v := request.FormValue("WithFilter")
	filter.WithFilter = v != "false" && len(v) > 0

	//GetEntityFilter remove this line for disable generator functionality

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *EntityFilter) GetEntityModel() Entity {

	filter.model.Validate()

	return filter.model
}
