package main

import (
    "gosha/cmd"
    "os"
    "gosha/webapp"
    "gosha/mode"
)

func main() {


    if len(os.Args) > 1 {
        cmd.Run()
        return
    }

    mode.SetNonInteractiveMode()
    webapp.Run()
}
