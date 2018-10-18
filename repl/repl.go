package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Bo0km4n/dummy-monkey/lexer"
	"github.com/Bo0km4n/dummy-monkey/parser"
	"github.com/k0kubun/pp"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)

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

		pp.Println(program)
	}
}
