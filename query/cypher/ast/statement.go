package ast

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// Stmt all statments implement the Stmt interface.
type Stmt interface {
}

// Expr all expression nodes implement the Expr interface.
type Expr interface {
	exprNode()
}

// Clauses in the Cypher query language.
type Clauses interface {
	GetPattern() ir.Patn
	GetNext() Stmt
}

// NonTerminalExpr is a NonTerminal symbol which can still be broken down e.g. a BooleanExpr
type NonTerminalExpr interface {
	exprNode()
	interpretNode()

	GetLeft() InterpretExpr
	GetRight() InterpretExpr
	SetLeft(x InterpretExpr)
	SetRight(x InterpretExpr)
}

// SingleNonTerminalExpr is a NonTerminal symbol which only can be broken down once e.g. a NotExpr
type SingleNonTerminalExpr interface {
	exprNode()
	interpretNode()

	GetValue() InterpretExpr
	SetValue(x InterpretExpr)
}

// TerminalExpr is a Terminal symbol which cannot be broken down further e.g. a Ident
type TerminalExpr interface {
	exprNode()
	interpretNode()

	GetValue() interface{}
	SetValue(x interface{})
}

// InterpretExpr is the base interface for the NonTerminalExpr and TerminalExpr
type InterpretExpr interface {
	exprNode()
	interpretNode()
	Interpret(variable string, vertex *vertices.Vertex) interface{}
}

type PatternStmt interface {
	patternNode()
}

// Precedence sorts Expr by their precedence
func Precedence(item Expr) int {
	if b, ok := item.(*BooleanExpr); ok {
		return BooleanPrecedence(*b)
	} else if b, ok := item.(*NotExpr); ok {
		return NotPrecedence(*b)
	} else if b, ok := item.(*ComparisonExpr); ok {
		return ComparisonPrecedence(*b)
	} else if b, ok := item.(MathematicalExpr); ok {
		return MathPrecedence(b)
	} else if b, ok := item.(ParenthesesExpr); ok {
		return ParenthesesPrecedence(b)
	}
	return 20
}
