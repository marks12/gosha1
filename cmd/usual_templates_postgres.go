package cmd

const usualDockerCallPs = `#!/bin/bash
docker run --rm --name pg-{ms-name} -e POSTGRES_DB={ms-name} -e POSTGRES_USER={ms-name} -e POSTGRES_PASSWORD={new-pass} -d -p 35432:5432 -v "$(pwd)/postgres:/var/lib/postgresql/data" postgres
`

const usualDockerCallMy = `#!/bin/bash
docker run --name my-{ms-name} -e MYSQL_DATABASE={ms-name} -e MYSQL_ROOT_PASSWORD={new-pass} mariadb/server:10.3 --log-bin --binlog-format=MIXED
docker start my-openplatform
#mysql -h 172.17.0.2 -u root -p
#` + "const DbConnectString = settings.DbUser + `:` + settings.DbPass + `@tcp(` + settings.DbHost + `:` + settings.DbPort + `)/` + settings.DbName + `?parseTime=true`" + `
#import 	_ "github.com/jinzhu/gorm/dialects/mysql"
#var Db, DbErr = gorm.Open("mysql", DbConnectString)
`
var usualDockerPs =
        assignMsName(
            assignPass(
                usualDockerCallPs, dbPass))


var usualDockerMy =
        assignMsName(
            assignPass(
                usualDockerCallMy, dbPass))


var usualTemplateDockerPs = template{
    Path:    "./psql-docker.sh",
    Content: usualDockerPs,
}

var usualTemplateDockerMy = template{
    Path:    "./mysql-docker.sh",
    Content: usualDockerMy,
}
