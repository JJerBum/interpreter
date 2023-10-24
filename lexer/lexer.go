package lexer

import (
	"monkey-lang-clone/token"
)

// Lexert는 구조체는 export
// input(사용자가 입력한 monkey-lang의 소스코드)
// position(현재 읽은 index 값)
// readPosition(현재 읽고 있는 문자의 다음 인덱스 깂)
// ch(현재 읽은 문자값)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// 생성자
// 피드백: Lexer 구조체를 반환하는게 아니라, NextToken을 구현하라고 명시해논 인터터페이스를 반환하는 것도 좋다.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar()는
// 문자열 끝에 도달했다면 ch에 eof 할당
// l.ch에 다음 문자값 할당.

// l.position에 l.reapotion 대입
// l.readpotion++

// 패드백: if l.readPosition > len(l.input)  이부분은 현재 인덱스와 배열의 크기를 비교 하고 있다. 이렇게 하면 배열의 크기를 인덱스로 생각을 한번 거처야 하므로 귀찬하 진다. 그러므로 len(l.input) -1 로 표현함으로써 인덱스와 비교하게 코드를 작성하면 좋지 않을까?
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:

		tok.Literal = ""
		tok.Type = token.EOF
	default:

		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// skip Withespace
// 이 함수가 끝나면 l.ch에는 문자로 시작된다.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}

}

// 토큰 생성
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 식별자 읽는 함수
func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // postion ~ l.potions -1
}

// 글자인가?
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 숫자 읽기
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// func (l *Lexer) readNumber() string {
// 	position := l.position

// 	for l.ch >= '0' && l.ch <= '9' {
// 		l.readChar()
// 	}

// }

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
