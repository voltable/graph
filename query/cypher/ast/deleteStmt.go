package ast

import "github.com/RossMerr/Caudex.Graph/query/cypher/ir"

// DeleteStmt used to delete graph elements — nodes, relationships or paths.
type DeleteStmt struct {
	Pattern ir.Patn
	Next    Clauses
}

func (DeleteStmt) patternNode() {}

func (m *DeleteStmt) GetPattern() ir.Patn {
	return m.Pattern
}

func (m *DeleteStmt) GetNext() Clauses {
	return m.Next
}
