package cgroups

import (
	"github.com/sirupsen/logrus"
	"github.com/yuxh1996/Ydocker/cgroups/subsystems"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

// 设置cgroup资源限制
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range(subsystems.SubsystemsIns) {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range (subsystems.SubsystemsIns) {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *CgroupManager) Destory() error {
	for _, subSysIns := range (subsystems.SubsystemsIns) {
		if err := subSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
