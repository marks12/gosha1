#!/bin/bash

env GOOS=darwin GOARCH=amd64 go build -o ./gosha-mac -v ./main.go

go build -v