package cmd

var usualWebappEntityDbModels = `package dbmodels

import (
    "time"
)

type {Entity} struct {

    ID        int       ` + "`" + `gorm:"primary_key"` + "`" + `
    ` + getRemoveLine("{Entity}") + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}

func ({entity} *{Entity}) Validate() {
    ` + getRemoveLine("Validate") + `
}

`

var usualTemplateWebappEntityDbModels = template{
    Path:    "",
    Content: assignMsName(usualWebappEntityDbModels),
}
