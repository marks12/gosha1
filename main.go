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
	isRestart, err := updater.MakeUpdate()
	if err != nil {
		fmt.Println("Error in AutoUpdate:", err.Error())
		return
	}
	fmt.Println("Current version:", settings.CurrentReleaseTag, isRestart)

	if isRestart {
		binPath, _ := os.Executable()
		os.Args = append(os.Args, "childProcess")
		cmd := exec.Command(binPath, os.Args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGKILL,
		}
		cmd.Start()

		select {}

	} else {
		if len(os.Args) > 1 {
			cmd.Run()
			return
		}
		mode.SetNonInteractiveMode()
		fmt.Println("Run server")
		webapp.Run()
	}
}
