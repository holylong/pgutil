package iocan

import (
	"log"
	"os/exec"
)

func RunCommand(path, name, aten string, arg ...string) (msg string, err error) {
	cmd := exec.Command(name, aten, arg[0])
	cmd.Dir = path
	// err = cmd.Run()
	output, err := cmd.CombinedOutput()
	// log.Println("args:" + cmd.Args)
	if err != nil {
		log.Println("==>> err", err.Error(), "cmd", cmd.Args)
	}
	return string(output), err
}

func RunCommandOne(path string, arg ...string) (msg string, err error) {
	cmd := exec.Command("name", arg...)
	cmd.Dir = path
	err = cmd.Run()
	// log.Println("args:" + cmd.Args)
	if err != nil {
		log.Println("==>> err", err.Error(), "cmd", cmd.Args)
	}
	return
}
