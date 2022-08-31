package lexer

import (
	"testing"

	"github.com/startdusk/tinyscript/token"
)

func TestNextToken(t *testing.T) {
	type expect struct {
		expectedType    token.TokenType
		expectedLiteral string
	}
	tests := []struct {
		input   string
		expects []expect
	}{
		{
			input: "=+(){},;",
			expects: []expect{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: "{},;",
			expects: []expect{
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `
			let five = 5;
			let ten = 10;
			let add = fn(x, y) {
				x + y;
			};
			let result = add(five, ten);
			`,
			expects: []expect{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for i, tt := range tests {
		l := New(tt.input)
		t.Run(tt.input, func(t *testing.T) {
			for _, expect := range tt.expects {
				tok := l.NextToken()
				if tok.Type != expect.expectedType {
					t.Fatalf("tests[%s]-[%d] - tokentype wrong. expected=%q, got=%q",
						tt.input, i, expect.expectedType, tok.Type)
				}
				if tok.Literal != expect.expectedLiteral {
					t.Fatalf("tests[%s]-[%d] - literal wrong. expected=%q, got=%q",
						tt.input, i, expect.expectedLiteral, tok.Literal)
				}
			}
		})
	}
}
