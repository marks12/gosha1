package types

import (
    "net/http"
)

type FieldType struct {
    Id   int
    Name   string
    //FieldType remove this line for disable generator functionality
}

func (fieldType *FieldType) Validate()  {
    //Validate remove this line for disable generator functionality
}

type FieldTypeFilter struct {
    model FieldType
    //FieldTypeFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetFieldTypeFilter(request *http.Request, functionType string) FieldTypeFilter {

    var filter FieldTypeFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetFieldTypeFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *FieldTypeFilter) GetFieldTypeModel() FieldType {

    filter.model.Validate()

    return  filter.model
}

func (filter *FieldTypeFilter) SetFieldTypeModel(typeModel FieldType) {

    filter.model = typeModel
}
