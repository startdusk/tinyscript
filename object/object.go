package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      ObjectType = "INTEGER"
	BOOLEAN_OBJ      ObjectType = "BOOLEAN"
	NULL_OBJ         ObjectType = "NULL"
	RETURN_VALUE_OBJ ObjectType = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// =================================================================================================
// Integer Object
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// =================================================================================================
// Boolean Object
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// =================================================================================================
// Null Object
type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType { return NULL_OBJ }

// =================================================================================================
// Return Value Object
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
