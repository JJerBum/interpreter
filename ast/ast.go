package ast

import "monkey-lang-clone/token"

type Node interface {
	TokenLiteral() string // Used for debugging and testing purposes
}

// Statment does not make a value
type Statement interface {
	Node
	statementNode() // dummy function -> To prevent potential implementation
}

// Expression makes value
type Expression interface {
	Node
	expressionNode() // dummy functino -> To prevent potential implementation
}

// A collection of ast tree statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Implemetns Statement
type LetStatement struct {
	Token token.Token
	Name  *Identifier // ex) foo, bar
	Value Expression  // ex) add(2, 2) * 5 / 10, 5, 5 * 5
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}

// Implements Expression
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}
