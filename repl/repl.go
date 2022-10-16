package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/startdusk/tinyscript/evaluator"
	"github.com/startdusk/tinyscript/lexer"
	"github.com/startdusk/tinyscript/object"
	"github.com/startdusk/tinyscript/parser"
)

const PROMPT = ">> "

const LOGO = ` __,__
_____ _                           _       _   
|_   _(_)_ __  _   _ ___  ___ _ __(_)_ __ | |_ 
  | | | | '_ \| | | / __|/ __| '__| | '_ \| __|
  | | | | | | | |_| \__ \ (__| |  | | |_) | |_ 
  |_| |_|_| |_|\__, |___/\___|_|  |_| .__/ \__|
               |___/                |_|        
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		if scanned := scanner.Scan(); !scanned {
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

func printParserErrors(out io.Writer, errs []string) {
	io.WriteString(out, LOGO)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errs {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
