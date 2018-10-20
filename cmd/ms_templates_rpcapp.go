package cmd

const msRpcappErrors = `package rpcapp

import (
    "msrpc/mdl"
)

func errMakeResponse (err string) (response mdl.Response) {

    response.Error = true
    response.ErrorMessage = err

    return
}`

var msTemplateRpcappErrors = template{
    Path:    "./rpcapp/errors.go",
    Content: msRpcappErrors,
}
