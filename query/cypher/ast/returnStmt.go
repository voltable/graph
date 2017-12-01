package ast

// ReturnStmt used to return results from the search.
type ReturnStmt struct {
	Maps []*MapProjectionStmt
}

func (*ReturnStmt) exprNode() {}

// NewMapProjectionStmt
func NewReturnStmt(maps ...*MapProjectionStmt) *ReturnStmt {
	return &ReturnStmt{Maps: maps}
}
