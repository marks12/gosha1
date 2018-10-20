package cmd

const msCoreDb = `package core

import (
	"github.com/jinzhu/gorm"
	"msproduct/settings"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const DbConnectString =
    "host='" +          settings.DbHost +
    "' port='" +        settings.DbPort +
    "' user='" +	    settings.DbUser +
    "' password='" +    settings.DbPass +
    "' dbname='" +    	settings.DbName +
    "' sslmode='disable'"

var Db, DbErr = gorm.Open("postgres", DbConnectString)`

var msTemplateCoreDb = template{
    Path:    "./core/db.go",
    Content: microserviceNameRegexp.ReplaceAllString(msCoreDb, getCurrentDirName()),
}
