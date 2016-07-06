package stdin

import (
	"log"
	"os"
)

var stdinChars chan byte

func init() {
	stdinChars = make(chan byte)
	go func() {
		disableInputBuffering()
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

func GetCharChan() <-chan byte {
	return stdinChars
}
