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

func TestReturnStatements(t *testing.T) {
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
return 5; 
return 10;
return 838383;
`,
			expectedIdentifiers: []string{
				"x",
				"y",
				"foobar",
			},
			wantErr: false,
		},
		// 		{
		// 			name: "invalid_test",
		// 			input: `
		// return return;
		// `,
		// 			wantErr: true,
		// 			errLen:  2,
		// 		},
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
				for _, stmt := range program.Statements {
					returnStmt, ok := stmt.(*ast.ReturnStatement)
					if !ok {
						t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
						continue
					}
					if returnStmt.TokenLiteral() != "return" {
						t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
							returnStmt.TokenLiteral())
					}
				}
			}
		})
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}
