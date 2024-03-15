package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(test *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, testToken := range tests {
		token := lexer.nextToken()

		if token.Type != testToken.expectedType {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, testToken.expectedType, token.Type)
		}

		if token.Literal != testToken.expectedLiteral {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, testToken.expectedLiteral, token.Literal)
		}
	}
}
