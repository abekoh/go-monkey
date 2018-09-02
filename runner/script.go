package runner

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/abekoh/monkey/evaluator"
	"github.com/abekoh/monkey/lexer"
	"github.com/abekoh/monkey/object"
	"github.com/abekoh/monkey/parser"
)

func StartScript(in io.Reader, out io.Writer) {
	bs, err := ioutil.ReadAll(in)
	env := object.NewEnvironment()

	if err != nil {
		log.Fatal(err)
	}
	l := lexer.New(string(bs))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
