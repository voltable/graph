package ast

// ReturnStmt used to return results from the search.
type ReturnStmt struct {
	Maps []*ProjectionMapStmt
	Next Clauses
}

func (*ReturnStmt) exprNode() {}

// NewMapProjectionStmt
func NewReturnStmt(maps ...*ProjectionMapStmt) *ReturnStmt {
	return &ReturnStmt{Maps: maps}
}

func (m *ReturnStmt) GetNext() Clauses {
	return m.Next
}
