package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/startdusk/tinyscript/lexer"
	"github.com/startdusk/tinyscript/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
