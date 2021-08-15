package ast

import (
	"testing"

	"github.com/penthious/writing_an_interpreter_in_go/token"
)

func TestString(t *testing.T) {
	want := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != want {
		t.Errorf("program.String() wrong, got = %q, want = %q", program.String(), want)
	}
}
