package main

import (
	"fmt"
	"gosha/cmd"
	"gosha/mode"
	"gosha/settings"
	"gosha/updater"
	"gosha/webapp"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Current version:", settings.CurrentReleaseTag, "OS:", runtime.GOOS)
	if settings.CurrentReleaseTag != settings.TegPlaceholderName {
		isRestart, err := updater.MakeUpdate()
		if err != nil {
			fmt.Println("Error in AutoUpdate:", err.Error())
		}

		if isRestart {
			binPath, _ := os.Executable()
			c := exec.Command(binPath, os.Args[1:]...)

			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Start()

			select {}

		}
	}
	if len(os.Args) > 1 {
		cmd.Run()
		return
	}
	mode.SetNonInteractiveMode()
	webapp.Run()

}
