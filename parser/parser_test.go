package parser

import (
	"testing"

	"github.com/startdusk/tinyscript/ast"
	"github.com/startdusk/tinyscript/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, identifier string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf(`stmt.TokenLiteral not "let" got=%q`, stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf(`stmt not *ast.LetStatement, got=%T`, stmt)
		return false
	}

	if letStmt.Name.Value != identifier {
		t.Errorf(`letStmt.Name.Value not "%s", got=%q`, identifier, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != identifier {
		t.Errorf(`stmt.Name not %q, got=%q`, identifier, letStmt.Name)
		return false
	}

	return true
}
