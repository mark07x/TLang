package repl

import (
	"bufio"
	"fmt"
	"github.com/mark07x/TLang/evaluator"
	"github.com/mark07x/TLang/lexer"
	"github.com/mark07x/TLang/parser"
	"io"
)

const PROMPT = "T> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := evaluator.SharedEnv.NewEnclosedEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Printf("\n")
			scanner = bufio.NewScanner(in)
			continue
		}

		line := scanner.Text() + "\n"
		for len(line) >= 2 && line[len(line) - 2] == '\\' {
			line = line[:len(line) - 2] + "\n"
			fmt.Printf(".. ")
			scanned := scanner.Scan()
			if !scanned {
				fmt.Printf("\n")
				scanner = bufio.NewScanner(in)
				continue
			}
			line = line + scanner.Text() + "\n"
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			evaluator.PrintParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != evaluator.VoidObj {
			_, _ = io.WriteString(out, evaluated.Inspect(2, env))
			_, _ = io.WriteString(out, "\n")
		}
	}
}
