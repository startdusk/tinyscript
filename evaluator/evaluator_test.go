package evaluator

import (
	"testing"

	"github.com/startdusk/tinyscript/lexer"
	"github.com/startdusk/tinyscript/object"
	"github.com/startdusk/tinyscript/parser"
)

func TestEvalIntergerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		// {"-10", -10}, not supported yet
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	res, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not integer. got=%T (%+v)", obj, obj)
		return false
	}

	if res.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", res.Value, expected)
		return false
	}

	return true
}
