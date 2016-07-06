package stdin

import "os/exec"

func disableInputBuffering() {
	err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		panic(err)
	}
}
