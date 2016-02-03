package stdin

import (
	"log"
	"os"
	"os/exec"
)

var stdinChars chan byte

func init() {
	stdinChars = make(chan byte)
	go func() {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		var b []byte = make([]byte, 1)
		for {
			_, err := os.Stdin.Read(b)
			if err != nil {
				log.Fatal(err)
			}
			stdinChars <- b[0]
		}
	}()
}
