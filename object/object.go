package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/startdusk/tinyscript/ast"
)

type ObjectType string

const (
	INTEGER_OBJ      ObjectType = "INTEGER"
	BOOLEAN_OBJ      ObjectType = "BOOLEAN"
	NULL_OBJ         ObjectType = "NULL"
	RETURN_VALUE_OBJ ObjectType = "RETURN_VALUE"
	ERROR_OBJ        ObjectType = "ERROR"
	FUNCTION_OBJ     ObjectType = "FUNCTION"
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

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// =================================================================================================
// Boolean Object
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// =================================================================================================
// Null Object
type Null struct{}

func (n *Null) Inspect() string { return "null" }

func (n *Null) Type() ObjectType { return NULL_OBJ }

// =================================================================================================
// Return Value Object
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// =================================================================================================
// Error Object
type Error struct {
	Message string
}

func (e *Error) Inspect() string { return "ERROR: " + e.Message }

func (e *Error) Type() ObjectType { return ERROR_OBJ }

// =================================================================================================
// Function Object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

func (f *Function) Inspect() string {
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
