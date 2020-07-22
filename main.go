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
	"syscall"
)

func main() {
	fmt.Println("Current version:", settings.CurrentReleaseTag)
	if settings.CurrentReleaseTag != settings.TegPlaceholderName {
		isRestart, err := updater.MakeUpdate()
		if err != nil {
			fmt.Println("Error in AutoUpdate:", err.Error())
		}

		if isRestart {
			binPath, _ := os.Executable()
			c := exec.Command(binPath, os.Args[1:]...)
			c.SysProcAttr = &syscall.SysProcAttr{
				Pdeathsig: syscall.SIGKILL,
			}
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
