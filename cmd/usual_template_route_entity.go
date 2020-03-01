package cmd

type Crud struct {
	IsFind         bool
	IsCreate       bool
	IsRead         bool
	IsUpdate       bool
	IsDelete       bool
	IsFindOrCreate bool
	IsUpdateOrCreate bool
}

const usualRouteEntityComment = `[ {Entity} ]`

const usualRouteEntityFind = `router.HandleFunc(settings.{Entity}Route,         webapp.{Entity}Find).Methods("GET")`

const usualRouteEntityCreate = `router.HandleFunc(settings.{Entity}Route,         webapp.{Entity}Create).Methods("POST")`
const usualRouteEntityMultiCreate = `router.HandleFunc(settings.{Entity}Route+"/list", webapp.{Entity}MultiCreate).Methods("POST")`

const usualRouteEntityRead = `router.HandleFunc(settings.{Entity}Route+"/{id}", webapp.{Entity}Read).Methods("GET")`

const usualRouteEntityUpdate = `router.HandleFunc(settings.{Entity}Route+"/{id}", webapp.{Entity}Update).Methods("PUT")`
const usualRouteEntityMultiUpdate = `router.HandleFunc(settings.{Entity}Route+"/list", webapp.{Entity}MultiUpdate).Methods("PUT")`

const usualRouteEntityDelete = `router.HandleFunc(settings.{Entity}Route+"/{id}", webapp.{Entity}Delete).Methods("DELETE")`
const usualRouteEntityMultiDelete = `router.HandleFunc(settings.{Entity}Route+"/list", webapp.{Entity}MultiDelete).Methods("DELETE")`

const usualRouteEntityFindOrCreate = `router.HandleFunc(settings.{Entity}Route,         webapp.{Entity}FindOrCreate).Methods("PUT")`

const usualRouteEntityUpdateOrCreate = `router.HandleFunc(settings.{Entity}Route,         webapp.{Entity}UpdateOrCreate).Methods("PATCH")`

const usualRouteEntityGen = `

    //router-generator here dont touch this line`

var usualTemplateRouteEntity = template{
	Path:    "./path_error.txt",
	Content: GetUsualTemplateRouteEntity(Crud{true, true, true, true, true, true, true}),
}

func GetUsualTemplateRouteEntity(c Crud) (res string) {

	res = usualRouteEntityComment

	if c.IsFind {
		res += "\n    " + usualRouteEntityFind
	} else {
		res += "\n    //" + usualRouteEntityFind
	}

	if c.IsCreate {
		res += "\n    " + usualRouteEntityCreate
		res += "\n    " + usualRouteEntityMultiCreate
	} else {
		res += "\n    //" + usualRouteEntityCreate
		res += "\n    //" + usualRouteEntityMultiCreate
	}

	if c.IsRead {
		res += "\n    " + usualRouteEntityRead
	} else {
		res += "\n    //" + usualRouteEntityRead
	}

	if c.IsUpdate {
		res += "\n    " + usualRouteEntityUpdate
		res += "\n    " + usualRouteEntityMultiUpdate
	} else {
		res += "\n    //" + usualRouteEntityUpdate
		res += "\n    //" + usualRouteEntityMultiUpdate
	}

	if c.IsDelete {
		res += "\n    " + usualRouteEntityDelete
		res += "\n    " + usualRouteEntityMultiDelete
	} else {
		res += "\n    //" + usualRouteEntityDelete
		res += "\n    //" + usualRouteEntityMultiDelete
	}

	if c.IsFindOrCreate {
		res += "\n    " + usualRouteEntityFindOrCreate
	} else {
		res += "\n    //" + usualRouteEntityFindOrCreate
	}

	if c.IsUpdateOrCreate {
		res += "\n    " + usualRouteEntityUpdateOrCreate
	} else {
		res += "\n    //" + usualRouteEntityUpdateOrCreate
	}

	res += usualRouteEntityGen

	return res
}
