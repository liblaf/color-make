package shutil

import "os/exec"

func Has(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
