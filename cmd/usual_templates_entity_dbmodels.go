package cmd

var usualWebappEntityDbModels = `package dbmodels

import (
    "time"
    {IdImport}
)

type {Entity} struct {

    {ID}
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

func getDbModelContent(isUuid bool) string {

    idImport := ""
    idField := `ID        int       ` + "`" + `gorm:"primary_key"` + "`"

    if isUuid {
        idImport = "\"github.com/google/uuid\""
        idField = `ID   uuid.UUID ` + "`" + `gorm:"primary_key, default: uuid()"` + "`"
    }

    content := template{
        Path:    "",
        Content: AssignVar(AssignVar(assignMsName(usualWebappEntityDbModels), "{ID}", idField), "{IdImport}", idImport),
    }

    return content.Content

}
