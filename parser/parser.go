package parser

import (
	"fmt"
	"monkey-lang-clone/ast"
	"monkey-lang-clone/lexer"
	"monkey-lang-clone/token"
)

// Parser는 구문분석자(글의 짜임을 분석하는 자)로써, 토큰으로 입력으로 받아 AST를 반환하는 것이 목표인 구조체 이다.
type Parser struct {
	// l은 Lexer의 인스턴스 이다.
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

// New 함수는 렉서를 매개변수로 받아 Parser의 인스턴스 주소를 반환하는 함수 입니다.
// 반환되는 *Parser는 parser의 curToken과 peekToken은 설정이 완료된 상태 입니다.
func New(l *lexer.Lexer) *Parser {

	// 새로운 파서를 생성
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// token settings
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram 함수는 모든 토큰들을 구문 분석(글의 짜임 혹은 parsing)을 분석 하여 Statment들로 결과값을 도출하는 함수 입니다.
func (p *Parser) ParseProgram() *ast.Program {
	// 프로그램 생성 (프로그램은 Statment의 집합이다.)
	program := &ast.Program{Statements: []ast.Statement{}}

	/*
		이상적인 토큰들
		index[0] tokenType:"LET", tokenLiteral:"let"
		index[1] tokenType:"IDENT", tokenLiteral:"five"
		index[2] tokenType:"=", tokenLiteral:"="
		index[3] tokenType:"INT", tokenLiteral:"5"
		index[4] tokenType:";", tokenLiteral:";"
		index[5] tokenType:"LET", tokenLiteral:"let"
		index[6] tokenType:"IDENT", tokenLiteral:"ten"
		index[7] tokenType:"=", tokenLiteral:"="
		index[8] tokenType:"INT", tokenLiteral:"10"
		index[9] tokenType:";", tokenLiteral:";"
		index[10] tokenType:"LET", tokenLiteral:"let"
		index[11] tokenType:"IDENT", tokenLiteral:"add"
		index[12] tokenType:"=", tokenLiteral:"="
		index[13] tokenType:"FUNCTION", tokenLiteral:"fn"
		index[14] tokenType:"(", tokenLiteral:"("
		index[15] tokenType:"IDENT", tokenLiteral:"x"
		index[16] tokenType:",", tokenLiteral:","
		index[17] tokenType:"IDENT", tokenLiteral:"y"
		index[18] tokenType:")", tokenLiteral:")"
		index[19] tokenType:"{", tokenLiteral:"{"
		index[20] tokenType:"IDENT", tokenLiteral:"x"
		index[21] tokenType:"+", tokenLiteral:"+"
		index[22] tokenType:"IDENT", tokenLiteral:"y"
		index[23] tokenType:";", tokenLiteral:";"
		index[24] tokenType:"}", tokenLiteral:"}"
		index[25] tokenType:";", tokenLiteral:";"
		index[26] tokenType:"LET", tokenLiteral:"let"
		index[27] tokenType:"IDENT", tokenLiteral:"result"
		index[28] tokenType:"=", tokenLiteral:"="
		index[29] tokenType:"IDENT", tokenLiteral:"add"
		index[30] tokenType:"(", tokenLiteral:"("
		index[31] tokenType:"IDENT", tokenLiteral:"five"
		index[32] tokenType:",", tokenLiteral:","
		index[33] tokenType:"IDENT", tokenLiteral:"ten"
		index[34] tokenType:")", tokenLiteral:")"
		index[35] tokenType:";", tokenLiteral:";"
		index[36] tokenType:"EOF", tokenLiteral:""
	*/

	// 모든 토큰들을 순환하면서
	for p.curToken.Type != token.EOF {
		// statment를 구문 분석 해라!
		stmt := p.parseStatement()

		// 만약 구문 분석한 statment가 nil이라면?
		if stmt != nil {
			// program.Statments에 nil을 추가해라
			program.Statements = append(program.Statements, stmt)
		}

		// 다음 토큰갑으로 설정
		p.nextToken()
	}

	// 구문 문석 후 도출된 Statment의 집합인 Program을 반환
	return program
}

// parseStatment 함수는 현재 토큰을 분석하여 만약 모종의 토큰이면 그 토큰을 구문 분석하여 Statment를 반환하는 함수 이다.
func (p *Parser) parseStatement() ast.Statement {
	// parser의 현재 토큰이 ~
	switch p.curToken.Type {
	// token.LET 이라면
	case token.LET:
		// LetStatment를 구문분석 한 뒤 분석된 후 도출된 Statment를 반환해라
		return p.parseLetStatement()
	default:
		// 없으면 nil을 반환해라
		return nil
	}
}

// parseLetStatement() 함수는 Let 구문을 분석하는 함수 입니다.
// parseStatement()에 의해 호출되며, 현재 토큰의 타입이 token.LET일 때만 호출됩니다.
func (p *Parser) parseLetStatement() *ast.LetStatement {

	// 이상적인 값
	// tokenType:"LET"    tokenLiteral:"let"
	// tokenType:"IDENT"  tokenLiteral:"five"
	// tokenType:"="      tokenLiteral:"="
	// tokenType:"INT"    tokenLiteral:"5"
	// tokenType:";"      tokenLiteral:";"

	// LetStatment 생성
	letStmt := &ast.LetStatement{Token: p.curToken /* token.LET*/}

	// 다음 토큰이 token.IDENT가 아니라면?
	if p.expectedPeek(token.IDENT) == false {
		// nil을 반환한다.
		return nil
	}

	// Identfier 값을 설정 (ident은 token과, Literal을 가진다.)
	letStmt.Name = &ast.Identifier{Token: p.curToken /*token.IDENT*/, Value: p.curToken.Literal /* ex)five */}

	// 다음 토큰이 token.ASSIGN이 아니라면?
	if p.expectedPeek(token.ASSIGN) == false {
		return nil
	}

	// 다음 토큰이 세미콜론이 아니라면?
	for p.curTokenIs(token.SEMICOLON) == false {
		// let five = 5 + 5 + 5; 이런 문장이 올 수 있기 때문에 다음과 같은 처리를 한다.
		// TODO. letStmt.Value 설정해야 함.
		p.nextToken()
	}

	return letStmt
}

// 현재 토큰이 매개변수와 맞는지 bool 타입으로 반환하는 함수
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs함수는 매개변수로 들어온 t가 parser의 다음 토큰의 값과 같다면 참, 아니면 거짓을 반환하는 함수 입니다.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// excpectedPeek 함수는 매개변수로 들어온 t가 parser의 다음 토큰의 값과 같다면,
// 다음 토큰으로 이동하고 참을 반환합니다.
// 그외는 거짓을 반환합니다.
func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

// nextToken 함수는 현재 Parser의 토큰을 다음 토큰값으로 설정해주는 함수 입니다.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
