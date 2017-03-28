package ast

// EqualExpr A node that represents an equality comparison, such as (a = b).
type EqualExpr struct {
	BinaryExpr
}

// NotEqualExpr An inequality comparison, such as  (a <> b).
type NotEqualExpr struct {
	BinaryExpr
}

// LessThanExpr A "less than" comparison, such as (a < b).
type LessThanExpr struct {
	BinaryExpr
}

// LessThanOrEqualExpr A "less than or equal to" comparison, such as (a <= b).
type LessThanOrEqualExpr struct {
	BinaryExpr
}

// GreaterThanExpr A "greater than" comparison, such as (a > b).
type GreaterThanExpr struct {
	BinaryExpr
}

// GreaterThanOrEqualExpr A "greater than or equal to" comparison, such as (a >= b).
type GreaterThanOrEqualExpr struct {
	BinaryExpr
}

// IsNullExpr A "IS NULL" comparison, such as (null < 3 IS NULL).
type IsNullExpr struct {
	BinaryExpr
}

// IsNotNullExpr A "IS NOT NULL" comparison, such as (a <> 3 IS NOT NULL).
type IsNotNullExpr struct {
	BinaryExpr
}

func (*EqualExpr) exprNode()              {}
func (*NotEqualExpr) exprNode()           {}
func (*LessThanExpr) exprNode()           {}
func (*LessThanOrEqualExpr) exprNode()    {}
func (*GreaterThanExpr) exprNode()        {}
func (*GreaterThanOrEqualExpr) exprNode() {}
