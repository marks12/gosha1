package types

import (
    "net/http"
)

type BuLayer struct {
    Id   int
    //BuLayer remove this line for disable generator functionality
}

func (buLayer *BuLayer) Validate()  {
    //Validate remove this line for disable generator functionality
}

type BuLayerFilter struct {
    model BuLayer
    //BuLayerFilter remove this line for disable generator functionality

    AbstractFilter
}

func GetBuLayerFilter(request *http.Request, functionType string) BuLayerFilter {

    var filter BuLayerFilter

    filter.request = request
    //filter.TestFilter, _ = strconv.Atoi(request.FormValue("TestFilter"))

    //GetBuLayerFilter remove this line for disable generator functionality

    ReadJSON(request, &filter.model)

    filter.AbstractFilter = GetAbstractFilter(request, functionType)

    return  filter
}


func (filter *BuLayerFilter) GetBuLayerModel() BuLayer {

    filter.model.Validate()

    return  filter.model
}

func (filter *BuLayerFilter) SetBuLayerModel(typeModel BuLayer) {

    filter.model = typeModel
}
