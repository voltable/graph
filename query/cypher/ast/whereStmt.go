package ast

// WhereStmt used to adds constraints to the patterns in a MATCH or OPTIONAL MATCH clause or filters the results of a WITH clause.
type WhereStmt struct {
	Predicate Expr
	Next      Clauses
}

func (*WhereStmt) exprNode() {}

func (m *WhereStmt) GetNext() Clauses {
	return m.Next
}
