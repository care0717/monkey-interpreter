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
		return fmt.Errorf("s not *%T, got=%T", expect, s)
	}

	if letStmt.TokenLiteral() != expect.TokenLiteral() {
		return fmt.Errorf("letStmt.TokenLiteral() not %s, got=%s", expect.TokenLiteral(), letStmt.TokenLiteral())
	}

	if !cmp.Equal(letStmt.Name, expect.Name) {
		return fmt.Errorf("letStmt.Name not %s, got=%s", expect.Name, letStmt.Name)
	}

	return nil
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.ReturnStatement
	}{
		{
			input: `
return 5;
return 10;
return 838383;
`,
			expected: []ast.ReturnStatement{
				{
					Token: token.Token{
						Type:    token.RETURN,
						Literal: "return",
					},
				},
				{
					Token: token.Token{
						Type:    token.RETURN,
						Literal: "return",
					},
				},
				{
					Token: token.Token{
						Type:    token.RETURN,
						Literal: "return",
					},
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
			if err := testReturnStatement(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testReturnStatement(s ast.Statement, expect ast.ReturnStatement) error {
	returnStmt, ok := s.(*ast.ReturnStatement)
	if !ok {
		return fmt.Errorf("s not *%T, got=%T", expect, s)
	}

	if returnStmt.TokenLiteral() != expect.TokenLiteral() {
		return fmt.Errorf("returnStmt.TokenLiteral() not %s, got=%s", expect.TokenLiteral(), returnStmt.TokenLiteral())
	}
	return nil
}

func TestIdentifierExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.Identifier
	}{
		{
			input: "foobar;",
			expected: []ast.Identifier{
				{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "foobar",
					},
					Value: "foobar",
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
			if err := testIdentifierExpression(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testIdentifierExpression(s ast.Statement, expect ast.Identifier) error {
	stmt, ok := s.(*ast.ExpressionStatement)
	if !ok {
		return fmt.Errorf("s not *ast.ExpressionStatement, got=%T", s)
	}
	if !cmp.Equal(stmt.Expression, &expect) {
		return fmt.Errorf("%T diff %s", expect, cmp.Diff(&expect, stmt.Expression))
	}
	return nil
}

func TestIntegerLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.IntegerLiteral
	}{
		{
			input: "55;",
			expected: []ast.IntegerLiteral{
				{
					Token: token.Token{
						Type:    token.INT,
						Literal: "55",
					},
					Value: 55,
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
			if err := testIntegerLiteral(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testIntegerLiteral(s ast.Statement, expect ast.IntegerLiteral) error {
	stmt, ok := s.(*ast.ExpressionStatement)
	if !ok {
		return fmt.Errorf("s not *ast.ExpressionStatement, got=%T", s)
	}

	if !cmp.Equal(stmt.Expression, &expect) {
		return fmt.Errorf("%T diff %s", expect, cmp.Diff(&expect, stmt.Expression))
	}
	return nil
}

func TestParsingPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.PrefixExpression
	}{
		{
			input: "!5;",
			expected: []ast.PrefixExpression{
				{
					Token: token.Token{
						Type:    token.BANG,
						Literal: "!",
					},
					Operator: "!",
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
				},
			},
		},
		{
			input: "-15;",
			expected: []ast.PrefixExpression{
				{
					Token: token.Token{
						Type:    token.MINUS,
						Literal: "-",
					},
					Operator: "-",
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "15",
						},
						Value: 15,
					},
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
			if err := testPrefixExpression(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testPrefixExpression(s ast.Statement, expect ast.PrefixExpression) error {
	stmt, ok := s.(*ast.ExpressionStatement)
	if !ok {
		return fmt.Errorf("s not *ast.ExpressionStatement, got=%T", s)
	}

	if !cmp.Equal(stmt.Expression, &expect) {
		return fmt.Errorf("%T diff %s", expect, cmp.Diff(&expect, stmt.Expression))
	}

	return nil
}

func TestParsingInfixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.InfixExpression
	}{
		{
			input: "2 + 5;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.PLUS,
						Literal: "+",
					},
					Operator: "+",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
				},
			},
		},
		{
			input: "5 - 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.MINUS,
						Literal: "-",
					},
					Operator: "-",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "5 - 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.MINUS,
						Literal: "-",
					},
					Operator: "-",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "5 * 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.ASTERISK,
						Literal: "*",
					},
					Operator: "*",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "2 / 5;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.SLASH,
						Literal: "/",
					},
					Operator: "/",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
				},
			},
		},
		{
			input: "5 > 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.GT,
						Literal: ">",
					},
					Operator: ">",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "5 < 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.LT,
						Literal: "<",
					},
					Operator: "<",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "5 == 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.EQ,
						Literal: "==",
					},
					Operator: "==",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "5 != 2;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.NOT_EQ,
						Literal: "!=",
					},
					Operator: "!=",
					Left: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "5",
						},
						Value: 5,
					},
					Right: &ast.IntegerLiteral{
						Token: token.Token{
							Type:    token.INT,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		},
		{
			input: "true == true;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.EQ,
						Literal: "==",
					},
					Operator: "==",
					Left: &ast.Boolean{
						Token: token.Token{
							Type:    token.TRUE,
							Literal: "true",
						},
						Value: true,
					},
					Right: &ast.Boolean{
						Token: token.Token{
							Type:    token.TRUE,
							Literal: "true",
						},
						Value: true,
					},
				},
			},
		},
		{
			input: "true == false;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.EQ,
						Literal: "==",
					},
					Operator: "==",
					Left: &ast.Boolean{
						Token: token.Token{
							Type:    token.TRUE,
							Literal: "true",
						},
						Value: true,
					},
					Right: &ast.Boolean{
						Token: token.Token{
							Type:    token.FALSE,
							Literal: "false",
						},
						Value: false,
					},
				},
			},
		},
		{
			input: "false != false;",
			expected: []ast.InfixExpression{
				{
					Token: token.Token{
						Type:    token.NOT_EQ,
						Literal: "!=",
					},
					Operator: "!=",
					Left: &ast.Boolean{
						Token: token.Token{
							Type:    token.FALSE,
							Literal: "false",
						},
						Value: false,
					},
					Right: &ast.Boolean{
						Token: token.Token{
							Type:    token.FALSE,
							Literal: "false",
						},
						Value: false,
					},
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
			if err := testInfixExpression(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testInfixExpression(statement ast.Statement, expect ast.InfixExpression) interface{} {
	stmt, ok := statement.(*ast.ExpressionStatement)
	if !ok {
		return fmt.Errorf("s not *ast.ExpressionStatement, got=%T", statement)
	}

	if !cmp.Equal(stmt.Expression, &expect) {
		return fmt.Errorf("%T diff %s", expect, cmp.Diff(&expect, stmt.Expression))
	}

	return nil
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"a * -b",
			"(a * (-b))",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{"a + b + c",
			"(a + (b + c))",
		},
		{"a + b - c",
			"(a + (b - c))",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(a + ((b * c) + ((d / e) - f)))",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1+(2+3)+4",
			"(1 + ((2 + 3) + 4))",
		},
		{
			"(5+5)*2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5+1)",
			"(-(5 + 1))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
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

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, but got=%q", tt.expected, actual)
		}
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.IfExpression
	}{
		{
			input: `if (x < y) { x }`,
			expected: []ast.IfExpression{
				{
					Token: token.Token{
						Type:    token.IF,
						Literal: "if",
					},
					Condition: &ast.InfixExpression{
						Token: token.Token{
							Type:    token.LT,
							Literal: "<",
						},
						Operator: "<",
						Left: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "x",
							},
							Value: "x",
						},
						Right: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "y",
							},
							Value: "y",
						},
					},
					Consequence: &ast.BlockStatement{
						Token: token.Token{
							Type:    token.LBRACE,
							Literal: "{",
						},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Token: token.Token{Type: token.IDENT,
									Literal: "x"},
								Expression: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "x",
									},
									Value: "x",
								},
							},
						},
					},
					Alternative: nil,
				},
			},
		},
		{
			input: `if (x > y) { y } else { x }`,
			expected: []ast.IfExpression{
				{
					Token: token.Token{
						Type:    token.IF,
						Literal: "if",
					},
					Condition: &ast.InfixExpression{
						Token: token.Token{
							Type:    token.GT,
							Literal: ">",
						},
						Operator: ">",
						Left: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "x",
							},
							Value: "x",
						},
						Right: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "y",
							},
							Value: "y",
						},
					},
					Consequence: &ast.BlockStatement{
						Token: token.Token{
							Type:    token.LBRACE,
							Literal: "{",
						},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Token: token.Token{Type: token.IDENT,
									Literal: "y"},
								Expression: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "y",
									},
									Value: "y",
								},
							},
						},
					},
					Alternative: &ast.BlockStatement{
						Token: token.Token{
							Type:    token.LBRACE,
							Literal: "{",
						},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Token: token.Token{Type: token.IDENT,
									Literal: "x"},
								Expression: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "x",
									},
									Value: "x",
								},
							},
						},
					},
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
			if err := testIfExpression(s, tt.expected[i]); err != nil {
				t.Error(err)
				break
			}
		}
	}
}

func testIfExpression(statement ast.Statement, expect ast.IfExpression) interface{} {
	stmt, ok := statement.(*ast.ExpressionStatement)
	if !ok {
		return fmt.Errorf("s not *ast.ExpressionStatement, got=%T", statement)
	}

	if !cmp.Equal(stmt.Expression, &expect) {
		return fmt.Errorf("%T diff %s", expect, cmp.Diff(&expect, stmt.Expression))
	}

	return nil
}
