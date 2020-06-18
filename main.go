package main

import (
    "gosha/cmd"
    "gosha/mode"
    "gosha/webapp"
    "os"
)

func main() {

    if len(os.Args) > 1 {
        cmd.Run()
        return
    }

    mode.SetNonInteractiveMode()
    webapp.Run()
}
