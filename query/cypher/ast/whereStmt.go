package ast

// WhereStmt used to adds constraints to the patterns in a MATCH or OPTIONAL MATCH clause or filters the results of a WITH clause.
type WhereStmt struct {
	Predicate Expr
}

func (*WhereStmt) exprNode() {}
