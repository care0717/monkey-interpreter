package token

type Type string

var keywords = map[string]Type {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL Type = "ILLEGAL"
	EOF          = "EOF"

	// 識別子・リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    Type
	Literal string
}

func NewToken(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
