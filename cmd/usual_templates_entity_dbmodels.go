package cmd

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

var usualWebappEntityDbModels = `package dbmodels

import (
    "gorm.io/gorm"
    "time"
    {IdImport}
)

type {Entity} struct {

    {ID}
    ` + getRemoveLine("{Entity}") + `

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt ` + "`" + `sql:"index" json:"-"` + "`" + `

    validator
}
{beforeCreate}
func ({entity} *{Entity}) Validate() {
    ` + getRemoveLine("Validate") + `
}

`

func getDbModelContent(isUuid bool) string {

    idImport := ""
    idField := `ID        int       ` + "`" + `gorm:"primary_key"` + "`"

    if isUuid {

        idImport = "\"github.com/google/uuid\""

        if IsPostgres() {
            idField = "ID   uuid.UUID `sql:\"primary_key;type:uuid;default:uuid_generate_v4()\"`"
        } else {
            idField = "ID   uuid.UUID `sql:\"primary_key;default:uuid()\"`"
            //idField = `ID   uuid.UUID ` + "`" + `gorm:"primary_key, default: uuid()"` + "`"
        }
    }

    beforeCreate := ""
    if ! IsPostgres() && isUuid {
        beforeCreate = `
func (model *{Entity}) BeforeCreate(scope *gorm.Scope) error {
    if model.ID == uuid.Nil {
        model.ID = uuid.New()
    }
    return scope.SetColumn("ID", model.ID)
}
`
    }

    usualWebappEntityDbModels = AssignVar(usualWebappEntityDbModels, "{beforeCreate}", beforeCreate)
    cont := AssignVar(AssignVar(assignMsName(usualWebappEntityDbModels), "{ID}", idField), "{IdImport}", idImport)
    content := template{
        Path:    "",
        Content: cont,
    }

    return content.Content

}

func IsPostgres() bool {

    dat, err := ioutil.ReadFile("core/db.go")
    if err != nil {
        fmt.Println("Cant read file \"core/db.go\" for detect database type")
        os.Exit(1)
    }

    if strings.Contains(string(dat), "postgres") && ! strings.Contains(string(dat), "mysql") {
        return true
    }

    if ! strings.Contains(string(dat), "postgres") && strings.Contains(string(dat), "mysql") {
        return false
    }

    fmt.Println("Cant detect db type in file \"core/db.go\"")
    os.Exit(1)

    return false
}
