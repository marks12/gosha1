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
	"gorm.io/gorm"
	"{ms-name}/settings"
	"gorm.io/driver/postgres"
    "gorm.io/gorm/logger"
)

`+  defaultConnectionString + `

var	config = gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

var Db, DbErr = gorm.Open(postgres.Open(DbConnectString), &config)

func EnableSqlLog() {
	Db.Logger.LogMode(logger.Info)
}

func DisableSqlLog() {
	Db.Logger.LogMode(logger.Error)
}

func GetTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: Db}
	_ = stmt.Parse(model)
	return stmt.Schema.Table
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
