package ast

import "github.com/RossMerr/Caudex.Graph/vertices"

// Stmt all statments implement the Stmt interface.
type Stmt interface {
}

// Expr all expression nodes implement the Expr interface.
type Expr interface {
	exprNode()
}

// BlockExpr represents a block that contains a sequence of expressions where variables can be defined.
type BlockExpr struct {
	Expressions []Expr // Gets the expressions in this block.
}

// A BinaryExpr node represents a binary expression.
type BinaryExpr struct {
	Left  Expr // left operand
	Right Expr // right operand
}

func (BinaryExpr) exprNode() {}

// WhereStmt used to adds constraints to the patterns in a MATCH or OPTIONAL MATCH clause or filters the results of a WITH clause.
type WhereStmt struct {
	Predicate Expr
}

func (*WhereStmt) patnNode() {}

// MatchStmt used to search for the pattern described in it.
type MatchStmt struct {
	Pattern Patn
	Next    Stmt
}

// OptionalMatchStmt used to search for the pattern described in it, can match on nil
type OptionalMatchStmt struct {
	Pattern Patn
	Next    Stmt
}

// CreateStmt used to create nodes and relationships.
type CreateStmt struct {
	Pattern Patn
	Next    Stmt
}

// DeleteStmt used to delete graph elements — nodes, relationships or paths.
type DeleteStmt struct {
	Pattern Patn
	Next    Stmt
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

func (DeleteStmt) patternNode()        {}
func (CreateStmt) patternNode()        {}
func (OptionalMatchStmt) patternNode() {}
func (MatchStmt) patternNode()         {}

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
