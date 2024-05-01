package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
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
		fmt.Fprintln(out, program.String())
	}

}
func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprintln(out, "parser errors: ")
		fmt.Fprintf(out, "\t %s \n", msg)
	}
}
