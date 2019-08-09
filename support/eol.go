package support

type Eol struct {
	Sequence []rune
}

func (eol *Eol) IsPartRune(char rune) bool {
	for _, value := range eol.Sequence {
		if value == char {
			return true
		}
	}

	return false
}

func (eol *Eol) RuneSize() int {
	return len(eol.Sequence)
}

func (eol *Eol) Value() []rune {
	return eol.Sequence
}

var LF = &Eol{
	[]rune{'\n'},
}

var CR = &Eol{
	[]rune{'\r'},
}

var CRLF = &Eol{
	[]rune{'\r', '\n'},
}
