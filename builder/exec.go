package main

import "os/exec"

func execScript() error {
	cmd := exec.Command("sh", "buildScript.sh")
	return cmd.Run()
}
