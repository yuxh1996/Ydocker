package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	p := exec.Command("/usr/bin/yes")
	if err := p.Start(); err != nil {
		panic(err)
	}
	if err := p.Process.Kill(); err != nil {
		panic(err)
	}
	fmt.Println("yes is now a zombie process")
	time.Sleep(10 * time.Second)
	if err := p.Process.Release(); err != nil {
		panic(err)
	}
	fmt.Println("yes remains a zombie process")
	time.Sleep(10 * time.Second)
}
