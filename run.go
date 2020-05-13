package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yuxh1996/Ydocker/container"
)

func Run(tty bool, cmdArr []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
