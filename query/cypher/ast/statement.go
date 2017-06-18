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

func (BinaryExpr) exprNode()   {}
func (PropertyStmt) exprNode() {}
func (Ident) exprNode()        {}

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

func IsOperatorWithFreeXorY(e Expr) (func(x Expr), bool) {
	compar, ok := e.(OperatorExpr)
	if ok {
		if compar.GetX() == nil {
			return compar.SetX, true
		} else if compar.GetY() == nil {
			return compar.SetY, true
		}
	}
	return nil, false
}

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
