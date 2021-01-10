package ast

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestModify(t *testing.T) {
	one := func() Expression { return &IntegerLiteral{Value: 1} }
	two := func() Expression { return &IntegerLiteral{Value: 2} }

	turnOneIntoTwo := func(node Node) Node {
		integer, ok := node.(*IntegerLiteral)
		if !ok {
			return node
		}

		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	var tests = []struct {
		input    Node
		expected Node
	}{
		{
			input:    one(),
			expected: two(),
		},
		{
			input: &Program{Statements: []Statement{
				&ExpressionStatement{Expression: one()},
			}},
			expected: &Program{Statements: []Statement{
				&ExpressionStatement{Expression: two()},
			}},
		},
		{
			input:    &InfixExpression{Left: one(), Right: two()},
			expected: &InfixExpression{Left: two(), Right: two()},
		},
		{
			input:    &InfixExpression{Left: two(), Right: one()},
			expected: &InfixExpression{Left: two(), Right: two()},
		},
		{
			input:    &PrefixExpression{Right: one()},
			expected: &PrefixExpression{Right: two()},
		},
		{
			input:    &IndexExpression{Left: one(), Index: one()},
			expected: &IndexExpression{Left: two(), Index: two()},
		},
		{
			input: &IfExpression{
				Condition: one(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			expected: &IfExpression{
				Condition: two(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			input:    &ReturnStatement{ReturnValue: one()},
			expected: &ReturnStatement{ReturnValue: two()},
		},
		{
			input:    &LetStatement{Value: one()},
			expected: &LetStatement{Value: two()},
		},
		{
			input: &FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			expected: &FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			input:    &ArrayLiteral{Elements: []Expression{one(), one()}},
			expected: &ArrayLiteral{Elements: []Expression{two(), two()}},
		},
		{
			input: &HashLiteral{Pairs: []HashPair{
				{
					Key:   one(),
					Value: one(),
				},
				{
					Key:   two(),
					Value: one(),
				},
			}},
			expected: &HashLiteral{Pairs: []HashPair{
				{
					Key:   two(),
					Value: two(),
				},
				{
					Key:   two(),
					Value: two(),
				},
			}},
		},
	}

	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)

		if !cmp.Equal(modified, tt.expected) {
			t.Error(cmp.Diff(modified, tt.expected))
		}
	}

}
