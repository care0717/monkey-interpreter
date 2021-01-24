package evaluator

import (
	"github.com/care0717/monkey-interpreter/ast"
	"github.com/care0717/monkey-interpreter/lexer"
	"github.com/care0717/monkey-interpreter/object"
	"github.com/care0717/monkey-interpreter/parser"
	"github.com/care0717/monkey-interpreter/token"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDefineMacros(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram *ast.Program
		expectedKey     string
		expectedMacro   *object.Macro
	}{
		{
			input: `
let number = 1;
let function = fn(x, y) { x + y };
let mymacro = macro(x, y) { x + y; };`,
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{
							Type:    token.LET,
							Literal: "let",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "number",
							},
							Value: "number",
						},
						Value: &ast.IntegerLiteral{
							Token: token.Token{
								Type:    token.INT,
								Literal: "1",
							},
							Value: 1,
						},
					},
					&ast.LetStatement{
						Token: token.Token{
							Type:    token.LET,
							Literal: "let",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "function",
							},
							Value: "function",
						},
						Value: &ast.FunctionLiteral{
							Token: token.Token{
								Type:    token.FUNCTION,
								Literal: "fn",
							},
							Parameters: []*ast.Identifier{
								{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "x",
									},
									Value: "x",
								},
								{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "y",
									},
									Value: "y",
								},
							},
							Body: &ast.BlockStatement{
								Token: token.Token{
									Type:    token.LBRACE,
									Literal: "{",
								},
								Statements: []ast.Statement{
									&ast.ExpressionStatement{
										Token: token.Token{
											Type:    token.IDENT,
											Literal: "x",
										},
										Expression: &ast.InfixExpression{
											Token: token.Token{
												Type:    token.PLUS,
												Literal: "+",
											},
											Operator: "+",
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
									},
								},
							},
						},
					},
				},
			},
			expectedKey: "mymacro",
			expectedMacro: &object.Macro{
				Parameters: []*ast.Identifier{
					{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "x",
						},
						Value: "x",
					},
					{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "y",
						},
						Value: "y",
					},
				},
				Body: &ast.BlockStatement{
					Token: token.Token{
						Type:    token.LBRACE,
						Literal: "{",
					},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "x",
							},
							Expression: &ast.InfixExpression{
								Token: token.Token{
									Type:    token.PLUS,
									Literal: "+",
								},
								Operator: "+",
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
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		program := testParseProgram(tt.input)
		env := object.NewEnvironment()
		DefineMacros(program, env)
		if !cmp.Equal(program, tt.expectedProgram) {
			t.Errorf("%T diff %s[-got, +expected]", tt.expectedProgram, cmp.Diff(program, tt.expectedProgram))
		}
		obj, ok := env.Get(tt.expectedKey)
		if !ok {
			t.Errorf("macro not in env")
		}

		macro, ok := obj.(*object.Macro)
		if !ok {
			t.Errorf("object is not Macro. got=%T (%+v)", obj, obj)
		}

		if !cmp.Equal(macro.Parameters, tt.expectedMacro.Parameters) {
			t.Errorf("%T diff %s[-got, +expected]", tt.expectedMacro.Parameters, cmp.Diff(macro.Parameters, tt.expectedMacro.Parameters))
		}
		if !cmp.Equal(macro.Body, tt.expectedMacro.Body) {
			t.Errorf("%T diff %s[-got, +expected]", tt.expectedMacro.Body, cmp.Diff(macro.Body, tt.expectedMacro.Body))
		}
	}
}

func TestExpandMacros(t *testing.T) {
	tests := []struct{
		input string
		expected object.Object
	}{
		{
			input:    `
let infixExpression = macro() { quote(1 + 2); };
infixExpression();
`,
			expected: &object.Integer{Value: 3},
		},
		{
			input:    `
let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); };
reverse(2+2, 10-5)`,
			expected: &object.Integer{Value: 1},
		},
	}

	for _, tt := range tests {


		program := testParseProgram(tt.input)
		env := object.NewEnvironment()
		DefineMacros(program, env)
		expanded := ExpandMacros(program, env)
		got := Eval(expanded, env)

		if !cmp.Equal(got, tt.expected) {
			t.Errorf("%T diff %s[-got, +expected]", tt.expected, cmp.Diff(expanded, tt.expected))
		}
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
