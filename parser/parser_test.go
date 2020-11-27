package parser

import (
	"fmt"
	"github.com/care0717/monkey-interpreter/ast"
	"github.com/care0717/monkey-interpreter/lexer"
	"github.com/care0717/monkey-interpreter/token"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.LetStatement
	}{
		{
			input: `
let x = 5;
let y = 10;
let foobar = 838383;
`,
			expected: []ast.LetStatement{
				{
					Token: token.Token{
						Type:    token.LET,
						Literal: "let",
					},
					Name: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "x",
						},
						Value: "x",
					},
					Value: nil,
				},
				{
					Token: token.Token{
						Type:    token.LET,
						Literal: "let",
					},
					Name: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "y",
						},
						Value: "y",
					},
					Value: nil,
				},
				{
					Token: token.Token{
						Type:    token.LET,
						Literal: "let",
					},
					Name: &ast.Identifier{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "foobar",
						},
						Value: "foobar",
					},
					Value: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if program == nil {
			t.Errorf("ParseProgram() returned nil")
			continue
		}

		if len(program.Statements) != len(tt.expected) {
			t.Errorf("program.Statements does not contain %d stratements. got=%d", len(tt.expected), len(program.Statements))
			continue
		}
		for i := 0; i < len(program.Statements); i++ {
			s := program.Statements[i]
			if err := testLetStatement(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(s ast.Statement, expect ast.LetStatement) error {
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		return fmt.Errorf("s not *ast.LetStatement, got=%T", s)
	}

	if letStmt.TokenLiteral() != expect.TokenLiteral() {
		return fmt.Errorf("letStmt.TokenLiteral() not %s, got=%s", expect.TokenLiteral(), letStmt.TokenLiteral())
	}

	if !cmp.Equal(letStmt.Name, expect.Name) {
		return fmt.Errorf("letStmt.Name not %s, got=%s", expect.Name, letStmt.Name)
	}

	return nil
}
