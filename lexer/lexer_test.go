package lexer

import (
	"fmt"
	"monkey-lang-clone/token"

	"testing"
)

func TestNextToken(t *testing.T) {

	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};

	let result= add(five, ten);
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
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
	}

	l := New(input)

	for index, resTok := range tests {

		tok := l.NextToken()
		fmt.Printf("index[%d] tokenType:%q, tokenLiteral:%q\n", index, tok.Type, tok.Literal)
		if tok.Type != resTok.expectedType {

			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, resTok.expectedType, tok.Type)

		}
		if tok.Literal != resTok.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, resTok.expectedLiteral, tok.Literal)
			fmt.Println(index)
		}
	}
}
