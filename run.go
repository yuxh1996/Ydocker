package main

import (
	"github.com/yuxh1996/Ydocker/cgroups"
	"github.com/yuxh1996/Ydocker/cgroups/subsystems"
	"os"
	"strings"

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

	cgroupManager := cgroups.NewCgroupManager("Ydocker-cgroup")
	defer cgroupManager.Destory()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	sendInitCommand(cmdArr, writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
