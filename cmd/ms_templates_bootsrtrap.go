package cmd

const msInsertDataToDb = `package bootstrap

import (
	"msproduct/dbmodels"
	"fmt"
	"os"
	"msproduct/core"
)

func FillDBTestData()  {

	isDropTables := false

	if (len(os.Args) > 1 && os.Args[1] == "drop") ||
		(len(os.Args) > 2 && os.Args[2] == "drop") {
		isDropTables = true
	}

	if isDropTables == true {

		core.Db.DropTableIfExists(
//          dbmodels.Product{},
		)

		fmt.Println("All tables removed")
		os.Exit(1)
	}

	core.Db.AutoMigrate(
//		&dbmodels.Product{},
	)

	// add fixtures

}`

var msTemplateInsertDataToDb = template{
    Path:    "./bootstrap/insert_data_to_db.go",
    Content: microserviceNameRegexp.ReplaceAllString(msInsertDataToDb, getCurrentDirName()),
}

