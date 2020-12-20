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
