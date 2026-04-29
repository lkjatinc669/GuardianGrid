package programscanner

import (
	"bytes"
	"os/exec"
)

func runCmd(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}
