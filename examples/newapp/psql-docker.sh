#!/bin/bash
docker run --rm --name pg-newapp -e POSTGRES_DB=newapp -e POSTGRES_USER=newapp -e POSTGRES_PASSWORD=OxMcsAuK -d -p 35432:5432 -v "$(pwd)/postgres:/var/lib/postgresql/data" postgres
