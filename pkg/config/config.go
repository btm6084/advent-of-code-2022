package config

import (
	"log"
	"os"
	"strings"
)

var (
	INPUTFILE = "input.txt"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		INPUTFILE = "example.txt"
	}
}

func RawInput(fn string) []string {
	raw, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(raw), "\n")
}
