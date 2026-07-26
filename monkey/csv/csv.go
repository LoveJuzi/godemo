package csv

func GetWords(input string) []string {
	var words = []string{}

	var r = reader{input: input, pos: 0}

	if !r.isEnd() && r.isDelimiter() {
		words = append(words, "")
	}

	for {
		if word, ok := r.nextWord(); ok {
			words = append(words, word)
		} else {
			break
		}
	}
	return words
}

const (
	eNORMAL = iota
	eQUOTE
)

type reader struct {
	input string
	pos   int
}

func (r *reader) nextWord() (string, bool) {
	if r.isEnd() {
		return "", false
	}

	r.skipOneDelimeter() // 消耗一个前导符号

	if r.isEnd() {
		return "", true
	}

	return readerFns[r.getType()](r), true
}

func (r *reader) getType() int {
	if r.isQuote() {
		return eQUOTE
	}

	return eNORMAL
}

func (r *reader) isEnd() bool {
	return r.pos >= len(r.input)
}

func (r *reader) isDelimiter() bool {
	if ',' == r.input[r.pos] {
		return true
	}
	return false
}

func (r *reader) isQuote() bool {
	return '"' == r.input[r.pos]
}

func (r *reader) readWord() string {
	var word []byte
	for {
		if r.isEnd() {
			break
		}

		if r.isDelimiter() {
			break
		}

		word = append(word, r.input[r.pos])
		r.pos += 1
	}

	return string(word)
}

func (r *reader) readString() string {
	var word []byte
	r.pos += 1 // skip begin quote character
	for {
		if r.isEnd() {
			break
		}
		if r.isQuote() { // check end quote character
			r.pos += 1
			if r.isEnd() || !r.isQuote() {
				break
			}
			// 转义引号 ""
		}

		word = append(word, r.input[r.pos])
		r.pos += 1
	}

	if r.isEnd() || r.isDelimiter() {
		return string(word)
	}

	panic("quote error")
}

func (r *reader) skipOneDelimeter() {
	if r.isDelimiter() {
		r.pos += 1
	}
}

type readerFn func(r *reader) string

var readerFns []readerFn

func init() {
	readerFns = []readerFn{
		(*reader).readWord,   // eNORMAL
		(*reader).readString, // eQUOTE
	}
}
