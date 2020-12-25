package evaluator

import (
	"fmt"
	"github.com/care0717/monkey-interpreter/lexer"
	"github.com/care0717/monkey-interpreter/object"
	"github.com/care0717/monkey-interpreter/parser"
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

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if err := testObject(evaluated, tt.expected); err != nil {
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

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if err := testObject(evaluated, tt.expected); err != nil {
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

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if err := testObject(evaluated, tt.expected); err != nil {
			t.Error(err)
		}
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct{
		input    string
		expected object.Object
	} {
		{
			input:    "if (true) { 10 }",
			expected: &object.Integer{Value:10},
		},
		{
			input:    "if (false) { 10 }",
			expected: object.NULL,
		},
		{
			input:    "if (1) { 10 }",
			expected: &object.Integer{Value:10},
		},
		{
			input:    "if (1<2) { 10 }",
			expected: &object.Integer{Value:10},
		},
		{
			input:    "if (1>2) { 10 }",
			expected: object.NULL,
		},
		{
			input:    "if (1 > 2) { 10 } else {20}",
			expected: &object.Integer{Value:20},
		},
		{
			input:    "if (1 < 2) { 10 } else { 20 }",
			expected: &object.Integer{Value:10},
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if err := testObject(evaluated, tt.expected); err != nil {
			t.Error(err)
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	return Eval(program)
}

func testObject(got, expected object.Object) error {
	if !cmp.Equal(got, expected) {
		return fmt.Errorf("%T diff %s", expected, cmp.Diff(got, expected))
	}
	return nil
}
