package object

import (
	"bytes"
	"fmt"
	"github.com/care0717/monkey-interpreter/ast"
	"strings"
)

type Type string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

var (
	NULL  = &null{}
	TRUE  = &boolean{Value: true}
	FALSE = &boolean{Value: false}
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type boolean struct {
	Value bool
}

func (b *boolean) Type() Type      { return BOOLEAN_OBJ }
func (b *boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

func (s *String) Type() Type      { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }

type null struct{}

func (n null) Type() Type      { return NULL_OBJ }
func (n null) Inspect() string { return "null" }

type ReturnValue struct {
	Value Object
}

func (r ReturnValue) Type() Type      { return RETURN_VALUE_OBJ }
func (r ReturnValue) Inspect() string { return r.Value.Inspect() }

type Error struct {
	Message string
}

func (e Error) Type() Type      { return ERROR_OBJ }
func (e Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        Environment
}

func (f Function) Type() Type { return FUNCTION_OBJ }

func (f Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b Builtin) Type() Type      { return BUILTIN_OBJ }
func (b Builtin) Inspect() string { return "builtin function" }

type Array struct {
	Elements []Object
}

func (a Array) Type() Type { return ARRAY_OBJ }
func (a Array) Inspect() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
