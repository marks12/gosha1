package main

import (
    "gosha/cmd"
    "os"
    "gosha/webapp"
)

func main() {


    if len(os.Args) > 1 && os.Args[1] == "serve" {
        webapp.Run()
        return
    }

    cmd.Run()
}
