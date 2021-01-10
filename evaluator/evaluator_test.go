package evaluator

import (
	"fmt"
	"github.com/care0717/monkey-interpreter/ast"
	"github.com/care0717/monkey-interpreter/lexer"
	"github.com/care0717/monkey-interpreter/object"
	"github.com/care0717/monkey-interpreter/parser"
	"github.com/care0717/monkey-interpreter/token"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "5",
			expected: &object.Integer{Value: 5},
		},
		{
			input:    "10;",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "-5",
			expected: &object.Integer{Value: -5},
		},
		{
			input:    "-10;",
			expected: &object.Integer{Value: -10},
		},
		{
			input:    "5 + 5 + 5 + 5 - 10;",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "2 * 2 * 2 * 2 * 2;",
			expected: &object.Integer{Value: 32},
		},
		{
			input:    "-50 + 100 + -50;",
			expected: &object.Integer{Value: 0},
		},
		{
			input:    "5*2+10;",
			expected: &object.Integer{Value: 20},
		},
		{
			input:    "5 + 2 * 10;",
			expected: &object.Integer{Value: 25},
		},
		{
			input:    "20 + 2 * -10;",
			expected: &object.Integer{Value: 0},
		},
		{
			input:    "50 / 2 * 2 + 10;",
			expected: &object.Integer{Value: 60},
		},
		{
			input:    "2*(5+10);",
			expected: &object.Integer{Value: 30},
		},
		{
			input:    "3 * (3 * 3) - 10;",
			expected: &object.Integer{Value: 17},
		},
		{
			input:    "(5 + 10 * 2 + 15 / 3) * 2 + -10;",
			expected: &object.Integer{Value: 50},
		},
	}

	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "true",
			expected: object.TRUE,
		},
		{
			input:    "false;",
			expected: object.FALSE,
		},
		{
			input:    "1 < 2",
			expected: object.TRUE,
		},
		{
			input:    "1 > 2",
			expected: object.FALSE,
		},
		{
			input:    "1 < 1",
			expected: object.FALSE,
		},
		{
			input:    "1 > 1",
			expected: object.FALSE,
		},
		{
			input:    "1 == 1",
			expected: object.TRUE,
		},
		{
			input:    "1 != 1",
			expected: object.FALSE,
		},
		{
			input:    "1 == 2",
			expected: object.FALSE,
		},
		{
			input:    "1 != 2",
			expected: object.TRUE,
		},
		{
			input:    "true == true",
			expected: object.TRUE,
		},
		{
			input:    "false == false",
			expected: object.TRUE,
		},
		{
			input:    "true == false",
			expected: object.FALSE,
		},
		{
			input:    "true != false",
			expected: object.TRUE,
		},
		{
			input:    "(1<2) == true",
			expected: object.TRUE,
		},
		{
			input:    "(1<2) == false",
			expected: object.FALSE,
		},
		{
			input:    "true == (1 > 2)",
			expected: object.FALSE,
		},
		{
			input:    "true == (1 != 2)",
			expected: object.TRUE,
		},
	}

	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    `"hello world!"`,
			expected: &object.String{Value: "hello world!"},
		},
		{
			input:    `"hello" + " " + "world!"`,
			expected: &object.String{Value: "hello world!"},
		},
	}

	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "!true",
			expected: object.FALSE,
		},
		{
			input:    "!false;",
			expected: object.TRUE,
		},
		{
			input:    "!5",
			expected: object.FALSE,
		},
		{
			input:    "!!true;",
			expected: object.TRUE,
		},
		{
			input:    "!!false",
			expected: object.FALSE,
		},
		{
			input:    "!!5;",
			expected: object.TRUE,
		},
	}

	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "if (true) { 10 }",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "if (false) { 10 }",
			expected: object.NULL,
		},
		{
			input:    "if (1) { 10 }",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "if (1<2) { 10 }",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "if (1>2) { 10 }",
			expected: object.NULL,
		},
		{
			input:    "if (1 > 2) { 10 } else {20}",
			expected: &object.Integer{Value: 20},
		},
		{
			input:    "if (1 < 2) { 10 } else { 20 }",
			expected: &object.Integer{Value: 10},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "return 10;",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "return 10;9",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "return 2*5; 9;",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "9; return 2*5; 9;",
			expected: &object.Integer{Value: 10},
		},
		{
			input: `
if (10 > 1) {
  if (10 > 2) { 
    return 10;
  } 
  return 1;
}
`,
			expected: &object.Integer{Value: 10},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "5 + true;",
			expected: &object.Error{Message: "type mismatch: INTEGER + BOOLEAN"},
		},
		{
			input:    "5 * true; 5;",
			expected: &object.Error{Message: "type mismatch: INTEGER * BOOLEAN"},
		},
		{
			input:    "-true;",
			expected: &object.Error{Message: "unknown operator: -BOOLEAN"},
		},
		{
			input:    "true + false;",
			expected: &object.Error{Message: "unknown operator: BOOLEAN + BOOLEAN"},
		},
		{
			input:    "5; true - true; 1;",
			expected: &object.Error{Message: "unknown operator: BOOLEAN - BOOLEAN"},
		},
		{
			input:    "if (10 > 1) { true * true; }",
			expected: &object.Error{Message: "unknown operator: BOOLEAN * BOOLEAN"},
		},
		{
			input: `
if (10 > 1) {
  if (10 > 2) { 
    true / false;
	1;
  } 
  return 1;
}
`,
			expected: &object.Error{Message: "unknown operator: BOOLEAN / BOOLEAN"},
		},
		{
			input:    "foo",
			expected: &object.Error{Message: "identifier not found: foo"},
		},
		{
			input:    `"foo" - "bar"`,
			expected: &object.Error{Message: "unknown operator: STRING - STRING"},
		},
		{
			input:    `{"foo": "bar"}[fn(x) {x}];`,
			expected: &object.Error{Message: "unusable as hash key: FUNCTION"},
		},
	}

	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "let a = 5; a;",
			expected: &object.Integer{Value: 5},
		},
		{
			input:    "let a = 5 * 5 a;",
			expected: &object.Integer{Value: 25},
		},
		{
			input:    "let a = 5; let b = a + 1; b",
			expected: &object.Integer{Value: 6},
		},
		{
			input:    "let a = 5; let b = a; let c = a + b + 5; return c * 10;",
			expected: &object.Integer{Value: 150},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestFunctionObject(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input: "fn(x) { x + 2; };",
			expected: &object.Function{
				Parameters: []*ast.Identifier{
					{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "x",
						},
						Value: "x",
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
				},
				Env: object.NewEnvironment(),
			},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    "let identity = fn(x) { x; }; identity(5);",
			expected: &object.Integer{Value: 5},
		},
		{
			input:    "let identity = fn(x) { let a = x; return a; }; identity(5);",
			expected: &object.Integer{Value: 5},
		},
		{
			input:    "let double = fn(x) { return x * 2; }; double(5);",
			expected: &object.Integer{Value: 10},
		},
		{
			input:    "let add = fn(x, y) { return x + y; }; add(5, 3);",
			expected: &object.Integer{Value: 8},
		},
		{
			input:    "let add = fn(x, y) { return x + y; }; add(5 + 2, add(5, 3));",
			expected: &object.Integer{Value: 15},
		},
		{
			input:    "fn(x){ x; }(5)",
			expected: &object.Integer{Value: 5},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    `len("")`,
			expected: &object.Integer{Value: 0},
		},
		{
			input:    `len("hello world")`,
			expected: &object.Integer{Value: 11},
		},
		{
			input:    `len(1)`,
			expected: &object.Error{Message: "argument to `len` not supported, got INTEGER"},
		},
		{
			input:    `len("one", "two")`,
			expected: &object.Error{Message: "wrong number of arguments. got=2, want=1"},
		},
		{
			input:    `len([1, "two"])`,
			expected: &object.Integer{Value: 2},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input: `[1, 2 * 3, 4 + 5]`,
			expected: &object.Array{
				Elements: []object.Object{
					&object.Integer{Value: 1},
					&object.Integer{Value: 6},
					&object.Integer{Value: 9},
				},
			},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestArrayIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input:    `[1, 2, 3][0]`,
			expected: &object.Integer{Value: 1},
		},
		{
			input:    `[1, 2, 3][1+1]`,
			expected: &object.Integer{Value: 3},
		},
		{
			input:    `let i = 1; [1, 2, 3][i]`,
			expected: &object.Integer{Value: 2},
		},
		{
			input:    `let myArray = [1, 2, 3]; myArray[1] + myArray[2]`,
			expected: &object.Integer{Value: 5},
		},
		{
			input:    `[1, 2, 3][3]`,
			expected: object.NULL,
		},
		{
			input:    `[1, 2, 3][-1]`,
			expected: object.NULL,
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input: `
let t = "two";
{
  "one": 10-9,
  t: 1+1,
  "thr" + "ee": 6/2,
  4: 4,
  true: 5,
  false: 6
}`,
			expected: &object.Hash{
				Pairs: map[object.HashKey]object.HashPair{
					(&object.String{Value: "one"}).HashKey():   {Key: &object.String{Value: "one"}, Value: &object.Integer{Value: 1}},
					(&object.String{Value: "two"}).HashKey():   {Key: &object.String{Value: "two"}, Value: &object.Integer{Value: 2}},
					(&object.String{Value: "three"}).HashKey(): {Key: &object.String{Value: "three"}, Value: &object.Integer{Value: 3}},
					(&object.Integer{Value: 4}).HashKey():      {Key: &object.Integer{Value: 4}, Value: &object.Integer{Value: 4}},
					object.TRUE.HashKey():                      {Key: object.TRUE, Value: &object.Integer{Value: 5}},
					object.FALSE.HashKey():                     {Key: object.FALSE, Value: &object.Integer{Value: 6}},
				},
			},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}
func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		{
			input: `
{"foo": 5}["foo"]
`,
			expected: &object.Integer{Value: 5},
		},
		{
			input: `
{"foo": 5}["bar"]
`,
			expected: object.NULL,
		},
		{
			input: `
let key = "foo"; {"foo": 6}[key]
`,
			expected: &object.Integer{Value: 6},
		},
		{
			input: `
{}["bar"]
`,
			expected: object.NULL,
		},
		{
			input: `
{false: 1}[false]
`,
			expected: &object.Integer{Value: 1},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	} {
		{
			input: `quote(5+8)`,
			expected: &object.Quote{Node: &ast.InfixExpression{
				Token: token.Token{
					Type:    token.PLUS,
					Literal: "+",
				},
				Operator: "+",
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
						Literal: "8",
					},
					Value: 8,
				},
			}},
		},
		{
			input: `quote(foo+bar)`,
			expected: &object.Quote{Node: &ast.InfixExpression{
				Token: token.Token{
					Type:    token.PLUS,
					Literal: "+",
				},
				Operator: "+",
				Left: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "foo",
					},
					Value: "foo",
				},
				Right: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "bar",
					},
					Value: "bar",
				},
			}},
		},
	}
	if errors := testEval(tests); errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func testEval(tests []struct {
	input    string
	expected object.Object
}) []error {
	var errors []error
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluated := Eval(program, env)
		if err := testObject(evaluated, tt.expected); err != nil {
			errors = append(errors, fmt.Errorf("case: %s. err: %w", tt.input, err))
		}
	}
	return errors
}

func testObject(got, expected object.Object) error {
	if !cmp.Equal(got, expected) {
		return fmt.Errorf("%T diff %s[-got, +expected]", expected, cmp.Diff(got, expected))
	}
	return nil
}
