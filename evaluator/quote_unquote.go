package evaluator

import (
	"github.com/care0717/monkey-interpreter/ast"
	"github.com/care0717/monkey-interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
