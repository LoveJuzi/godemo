package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		var scanned = scanner.Scan()
		if !scanned {
			break
		}

		var line = scanner.Text()
		var l = lexer.New(line)
		var p = parser.New(l)

		var program = p.ParseProgram()
		if 0 != len(p.Errors()) {
			continue
		}

		var evaluated = evaluator.New(program).Eval()
		if nil != evaluated {
			io.WriteString(out, evaluated.Inspect())

			io.WriteString(out, "\n")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(out, "error reading input: %v\n", err)
	}
}
