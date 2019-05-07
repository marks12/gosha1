package cmd

const usualWebappEntityDbModels = `package dbmodels

import (
    "time"
)

type {Entity} struct {

    ID        int       ` + "`" + `gorm:"primary_key"` + "`" + `
    //{Entity} ` + removeLineComment + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func ({entity} *{Entity}) Validate() {
    //Validate ` + removeLineComment + `
}

`

var usualTemplateWebappEntityDbModels = template{
    Path:    "",
    Content: assignMsName(usualWebappEntityDbModels),
}
