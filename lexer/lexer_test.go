package lexer

import (
	"github.com/care0717/monkey-interpreter/token"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			input: `=+(){},;`,
			expected: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `!-/*5;`,
			expected: []token.Token{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `5 < 10 > 5;`,
			expected: []token.Token{
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `
if (5 < 10) {
  return true;
} else {
  return false;
}`,
			expected: []token.Token{
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			input: `
10 == 10; 
10 != 9;`,
			expected: []token.Token{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`,
			expected: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `
"foobar"
"foo bar"
"hello \"world\""
"hello\n world"
"hello\t\t\tworld"
`,
			expected: []token.Token{
				{Type: token.STRING, Literal: "foobar"},
				{Type: token.STRING, Literal: "foo bar"},
				{Type: token.STRING, Literal: `hello \"world\"`},
				{Type: token.STRING, Literal: `hello\n world`},
				{Type: token.STRING, Literal: `hello\t\t\tworld`},
				{token.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		for i, e := range tt.expected {
			tok := l.NextToken()

			if !cmp.Equal(tok, e) {
				t.Errorf("test[%d] = tokentype wrong. expected=%v, got=%v", i, e, tok)
			}
		}
	}
}
