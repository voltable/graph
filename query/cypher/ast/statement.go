package ast

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
	X Expr // left operand
	Y Expr // right operand
}

// PropertyStmt represents a node property.
type PropertyStmt struct {
	Variable string
	Value    string
}

type WhereStmt struct {
	Predicate Expr
}

type Ident struct {
	Data interface{} // denoted object; or nil
}

func (*BinaryExpr) exprNode()   {}
func (*PropertyStmt) exprNode() {}
func (*Ident) exprNode()        {}

type MatchStmt struct {
	Pattern Patn
	Next    Stmt
}

type OptionalMatchStmt struct {
	Pattern Patn
	Next    Stmt
}

type OperatorExpr interface {
	GetX() Expr
	GetY() Expr
	SetX(x Expr)
	SetY(x Expr)
}

func IsOperatorWithFreeXorY(e Expr) bool {
	compar, ok := e.(OperatorExpr)
	if ok {
		if compar.GetX() == nil {
			return true
		} else if compar.GetY() == nil {
			return true
		}
	}
	return false
}
