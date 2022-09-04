package parser

import (
	"testing"

	"github.com/startdusk/tinyscript/ast"
	"github.com/startdusk/tinyscript/lexer"
)

func TestLetStatements(t *testing.T) {

	tests := []struct {
		name                string
		input               string
		expectedIdentifiers []string
		wantErr             bool
		errLen              int
	}{
		{
			name: "success_test",
			input: `
let x = 5; 
let y = 10;
let foobar = 838383;
`,
			expectedIdentifiers: []string{
				"x",
				"y",
				"foobar",
			},
			wantErr: false,
		},
		{
			name: "invalid_test",
			input: `
let x 5; 
let = 10;
let 838383;
`,
			wantErr: true,
			errLen:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			program := p.ParseProgram()
			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if tt.wantErr {
				t.Logf("%s want error\n", tt.name)
				if len(p.errors) != tt.errLen {
					t.Fatalf("parser does not contain %d errors, got=%d", tt.errLen, len(p.errors))
				}
			} else {
				if len(program.Statements) != 3 {
					t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
				}
				checkParserErrors(t, p)
				for i := range tt.expectedIdentifiers {
					stmt := program.Statements[i]
					if !testLetStatement(t, stmt, tt.expectedIdentifiers[i]) {
						return
					}
				}
			}
		})
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d error", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}
