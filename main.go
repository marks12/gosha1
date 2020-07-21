package main

import (
	"fmt"
	"gosha/cmd"
	"gosha/mode"
	"gosha/settings"
	"gosha/updater"
	"gosha/webapp"
	"log"
	"os"
	"syscall"
)

func main() {
	args := []string{}
	for _, arg := range os.Args {
		if arg == "childProcess" {
			parent := syscall.Getppid()
			log.Printf("main: Killing parent pid: %v", parent)
			syscall.Kill(parent, syscall.SIGTERM)
		} else {
			args = append(args, arg)
		}
	}
	os.Args = args

	fmt.Println("Current version:", settings.CurrentReleaseTag)

	isRestart, err := updater.MakeUpdate()
	if err != nil {
		fmt.Println("Error in AutoUpdate:", err.Error())
	}

	if isRestart {
		binPath, _ := os.Executable()
		os.Args = append(os.Args, "childProcess")
		if err = syscall.Exec(binPath, os.Args, os.Environ()); err != nil {
			panic(err)
		}
	}

	if len(os.Args) > 1 {
		cmd.Run()
		return
	}

	mode.SetNonInteractiveMode()
	webapp.Run()
}
