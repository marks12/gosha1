package cmd

const usualDockerCallPs = `#!/bin/bash
docker run --rm --name pg-{ms-name} -e POSTGRES_DB={ms-name} -e POSTGRES_USER={ms-name} -e POSTGRES_PASSWORD={new-pass} -d -p 35432:5432 -v "$(pwd)/.postgres:/var/lib/postgresql/data" postgres
`

const usualDockerCallMy = `#!/bin/bash
docker run --name my-{ms-name} -e MYSQL_DATABASE={ms-name} -e MYSQL_ROOT_PASSWORD={new-pass} mariadb/server:10.4 --log-bin --binlog-format=MIXED
docker start my-{ms-name} &
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
