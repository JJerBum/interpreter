package ast

import "monkey-lang-clone/token"

// Abstract Systex Tree 이므로, Node가 존재
type Node interface {
	TokenLiteral() string // Used for debugging and testing purposes
}

// Statment does not make a value
// Statment는 실행가능한 최소의 독립적인 코드 조각을 일컫는다.
type Statement interface {
	Node
	statementNode() // dummy function -> To prevent potential implementation
}

// Expression makes value
// Expression은 하나 이상의 값으로 표현될 수 있는 코드 조각이다.
type Expression interface {
	Node
	expressionNode() // dummy functino -> To prevent potential implementation
}

// A collection of ast tree statements
// Proram은 실행 가능한 최소의 독립적인 코드의 잡합이다.
type Program struct {
	Statements []Statement
}

// TokenLiteral 메서드는 Node interface implement 하며, p.Statments의 0번쨰 토큰을 반환하는 함수 입니다.
// 재귀식으로 동작되며, 방금 선언한 함수를 방금 이용했습니다. (신기하네 이놈)
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Implemetns Statement
type LetStatement struct {
	// token.LET
	Token token.Token

	// ex) foo, bar (Statment는 Expression을 포함하는 관계를 가지고 있다.)
	Name *Identifier

	// ex) add(2, 2) * 5 / 10, 5, 5 * 5
	Value Expression
}

// TokenLiteral 함수는 Node interface implement 합니다.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// statmentNode 함수는 Statment interface implement 합니다.
func (ls *LetStatement) statementNode() {}

// Implements Expression
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral 함수는 Node interface implement 합니다.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// expressionNode 함수는 Expression interface implement 합니다.
func (i *Identifier) expressionNode() {}
