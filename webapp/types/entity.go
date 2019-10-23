package types

import (
	"net/http"
	"strconv"
	"gosha/cmd"
	"strings"
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

	cmd.HttpMethods
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
	IsFind bool
	IsCreate bool
	IsRead bool
	IsUpdate bool
	IsDelete bool
	IsFindOrCreate bool
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

	wf := request.FormValue("WithFilter")
	whf := request.FormValue("WithHiddenFields")

	filter.WithFilter = strings.ToLower(wf) != "false" && len(wf) > 0
	filter.WithHiddenFields = strings.ToLower(whf) != "false" && len(whf) > 0

	filter.IsFind = strings.ToLower(request.FormValue("IsFind")) != "false" && len(request.FormValue("IsFind")) > 0
	filter.IsCreate = strings.ToLower(request.FormValue("IsCreate")) != "false" && len(request.FormValue("IsCreate")) > 0
	filter.IsRead = strings.ToLower(request.FormValue("IsRead")) != "false" && len(request.FormValue("IsRead")) > 0
	filter.IsUpdate = strings.ToLower(request.FormValue("IsUpdate")) != "false" && len(request.FormValue("IsUpdate")) > 0
	filter.IsDelete = strings.ToLower(request.FormValue("IsDelete")) != "false" && len(request.FormValue("IsDelete")) > 0
	filter.IsFindOrCreate = strings.ToLower(request.FormValue("IsFindOrCreate")) != "false" && len(request.FormValue("IsFindOrCreate")) > 0
	//GetEntityFilter remove this line for disable generator functionality

	ReadJSON(request, &filter.model)

	filter.AbstractFilter = GetAbstractFilter(request, functionType)

	return filter
}

func (filter *EntityFilter) GetEntityModel() Entity {

	filter.model.Validate()

	return filter.model
}
