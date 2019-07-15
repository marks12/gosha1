#!/bin/bash

env GOOS=linux GOARCH=arm go build -o ./gosha-mac -v ./main.go

go build -v