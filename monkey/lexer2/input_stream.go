package lexer2

type inputStream struct {
	input string
	pos   int
}

func (i *inputStream) getChar() byte {
	if i.pos >= len(i.input) {
		i.pos += 1
		return 0
	}

	defer func() { i.pos += 1 }()
	return i.input[i.pos]
}

func (i *inputStream) ungetChar() {
	i.pos -= 1
}
