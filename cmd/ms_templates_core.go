package cmd

const postgresConnectionString = `const DbConnectString =
    "host='" +          settings.DbHost +
    "' port='" +        settings.DbPort +
    "' user='" +	    settings.DbUser +
    "' password='" +    settings.DbPass +
    "' dbname='" +    	settings.DbName +
    "' sslmode='disable'"
`

const mysqlConnectionString =  "const DbConnectString = settings.DbUser + `:` + settings.DbPass + `@tcp(` + settings.DbHost + `:` + settings.DbPort + `)/` + settings.DbName + `?parseTime=true`"

const defaultConnectionString = "const DbConnectString = ``"

const msCoreDb = `package core

import (
	"github.com/jinzhu/gorm"
	"{ms-name}/settings"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

`+  defaultConnectionString + `

var Db, DbErr = gorm.Open("postgres", DbConnectString)

func EnableSqlLog() {
	Db.LogMode(true)
}

func DisableSqlLog() {
	Db.LogMode(false)
}

func GetTableName(dbmodel interface{}) string {
	return Db.NewScope(dbmodel).TableName()
}
`

func getTemplateCoreDb(dbtype DatabaseType) template {

    conString := postgresConnectionString

    if dbtype.IsMysql {
        conString = mysqlConnectionString
    }

    return template{
        Path:    "./core/db.go",
        Content: AssignVar(AssignVar(assignMsName(msCoreDb), "postgres", dbtype.DbTypeName), defaultConnectionString, conString),
    }
}