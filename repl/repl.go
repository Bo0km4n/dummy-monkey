package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Bo0km4n/dummy-monkey/object"

	"github.com/Bo0km4n/dummy-monkey/evaluator"

	"github.com/Bo0km4n/dummy-monkey/lexer"
	"github.com/Bo0km4n/dummy-monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
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
			printParseErrors(out, p.Errors())
			continue
		}

		evalueated := evaluator.Eval(program, env)
		if evalueated != nil {
			io.WriteString(out, evalueated.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, fmt.Sprintf("could'nt evaluate expression: %s\n", program.String()))
		}
	}
}

func FileExecute(file *os.File, out io.Writer) {
	d, _ := ioutil.ReadAll(file)

	io.WriteString(out, string(d)+"\n"+"(↑ input code)====================================(↓ output)\n")
	env := object.NewEnvironment()
	l := lexer.New(string(d))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
		return
	}

	evalueated := evaluator.Eval(program, env)
	if evalueated != nil {
		io.WriteString(out, evalueated.Inspect())
		io.WriteString(out, "\n")
	} else {
		io.WriteString(out, fmt.Sprintf("could'nt evaluate expression: %s\n", program.String()))
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
