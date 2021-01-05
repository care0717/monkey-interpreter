package ast

import (
	"github.com/care0717/monkey-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    *Program
		expected string
	}{
		{
			input: &Program{
				Statements: []Statement{
					&LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name: Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "myVar"},
							Value: "myVar",
						},
						Value: Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
							Value: "anotherVar",
						},
					},
				},
			},
			expected: "let myVar = anotherVar;",
		},
	}
	for _, tt := range tests {
		if got := tt.input.String(); got != tt.expected {
			t.Errorf("program.String() wrong. got=%q", got)
		}
	}
}
