package stdin

import "os/exec"

func disableInputBuffering() {
	err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		panic(err)
	}
}
