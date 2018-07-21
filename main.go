package main

import (
	"fmt"
	"os"

	"github.com/abekoh/monkey/repl"
)

func main() {
	fmt.Printf("Hello abekoh! This is the Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
