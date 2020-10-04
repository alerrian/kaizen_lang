package lexer

import (
	"kaizen_lang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}[],;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LCURL, "{"},
		{token.RCURL, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - type wrong. \nExpected: %q\nGot: %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. \nExpected: %q\nGot: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
