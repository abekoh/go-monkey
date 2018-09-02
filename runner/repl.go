package runner

import (
	"bufio"
	"fmt"
	"io"

	"github.com/abekoh/monkey/evaluator"
	"github.com/abekoh/monkey/lexer"
	"github.com/abekoh/monkey/object"
	"github.com/abekoh/monkey/parser"
)

const PROMPT = ">> "

func StartRepl(in io.Reader, out io.Writer) {
	io.WriteString(out, "Hello abekoh! This is the Monkey programming language!\n")
	io.WriteString(out, "Feel free to type in commands\n")
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
