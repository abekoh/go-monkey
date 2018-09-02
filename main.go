package main

import (
	"log"
	"os"

	"github.com/abekoh/monkey/runner"
)

func main() {
	switch len(os.Args) {
	case 1:
		runner.StartRepl(os.Stdin, os.Stdout)
	case 2:
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		runner.StartScript(f, os.Stdout)
	default:
	}

}
