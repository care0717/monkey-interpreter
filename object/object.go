package object

import "fmt"

type Type string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
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
