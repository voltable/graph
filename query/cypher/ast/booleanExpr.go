package ast

// AndExpr A bitwise or logical AND operation, such as (a And b).
type AndExpr struct {
	BinaryExpr
}

// OrExpr A bitwise or logical OR operation, such as (a Or b).
type OrExpr struct {
	BinaryExpr
}

// NotExpr A bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	BinaryExpr
}

// XorExpr A bitwise or logical XOR operation, such as (a Xor b).
type XorExpr struct {
	BinaryExpr
}

func (*AndExpr) exprNode() {}
func (*OrExpr) exprNode()  {}
func (*NotExpr) exprNode() {}
func (*XorExpr) exprNode() {}
