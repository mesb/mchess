package lexer

import (
	"testing"
)

type testblock struct {
	word     string
	expected bool
}

var TESTWORDS = []testblock{
	testblock{"move", false},
	testblock{"test", false},
	testblock{"mov", true},
	testblock{"his", false},
	testblock{"last", false},
	testblock{"mver", false},
	testblock{"capture", false},
	testblock{"visit", false},
	testblock{"soc", false},
	testblock{"socrates", false},
}

func TestLex(t *testing.T) {
	for _, cur := range TESTWORDS {
		_, err := Lex(cur.word)

		if err == nil && !cur.expected {
			t.Errorf("%v is not a valid token", cur)
		}
	}
}

func TestLexMov(t *testing.T) {
	word := TESTWORDS[2]

	_, err := Lex(word.word)

	if err != nil && !word.expected {
		t.Errorf("%v is not a valid token", word)
	}
}

func TestLexResult(t *testing.T) {
	str := "mov PN-e4 e5"

	toks, err := Lex(str)

	if err == nil {
		t.Errorf("%v", toks)
	}
}
