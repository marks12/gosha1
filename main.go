package main

import (
    "gosha/cmd"
    "os"
    "gosha/webapp"
)

func main() {


    if len(os.Args) > 1 && os.Args[1] == "cmd" {
        cmd.Run()
        return
    }

    webapp.Run()
}
