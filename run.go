package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yuxh1996/Ydocker/container"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
