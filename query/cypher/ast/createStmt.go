package ast

import "github.com/RossMerr/Caudex.Graph/query/cypher/ir"

// CreateStmt used to create nodes and relationships.
type CreateStmt struct {
	Pattern ir.Patn
	Next    Stmt
}

func (CreateStmt) patternNode() {}

func (m *CreateStmt) GetPattern() ir.Patn {
	return m.Pattern
}

func (m *CreateStmt) GetNext() Stmt {
	return m.Next
}
