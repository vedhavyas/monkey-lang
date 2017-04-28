package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vedhavyas/monkey-lang/lexer"
	"github.com/vedhavyas/monkey-lang/token"
)

const PROMPT string = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		l := lexer.New(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
