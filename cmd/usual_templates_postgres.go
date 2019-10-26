package cmd

const usualDockerCall = `#!/bin/bash
docker run --rm --name pg-{ms-name} -e POSTGRES_DB={ms-name} -e POSTGRES_USER={ms-name} -e POSTGRES_PASSWORD={new-pass} -d -p 35432:5432 -v "$(pwd)/postgres:/var/lib/postgresql/data" postgres
`

var usualDockerCallontent =
        assignMsName(
            assignPass(
                usualDockerCall, dbPass))


var usualTemplateDockerCallontent = template{
    Path:    "./psql-docker.sh",
    Content: usualDockerCallontent,
}
