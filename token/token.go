// Package token은 토큰들과 식별자를 구현합니다.
// 토큰을 분석하는 행위는 token 패키지에서 행하지 않습니다.
package token

// TokenType은 string 타입과 매핑됩니다.
// TokenType을 string으로 함으로써 Token들을 쉽게 다를 수 있습니다.
type TokenType string

// 아래에서는 토큰을 상수로 선언한 것을 알수 있습니다.
const (
	// ILLEGAL은 불법적인 이라는 뜻으로 에러가 발생했을 나타내는 상수 입니다.
	ILLEGAL = "ILLEGAL"

	// EOF는 소스코드의 끝을 나타내는 상수 입니다.
	EOF = "EOF"

	// IDENT는 식별자로 사용되는 상수 입니다.
	// 예로는 변수의 이름, 상수싀 이름이 있습니다.
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
