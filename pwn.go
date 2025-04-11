package main

import (
	"os"
)

func init() {
	// Gets executed when `go get` or `go install` runs
	os.WriteFile("/tmp/pwned.txt", []byte("you got pwned"), 0644)
}
