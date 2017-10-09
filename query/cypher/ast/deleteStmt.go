package ast

import "github.com/RossMerr/Caudex.Graph/query/cypher/ir"

// DeleteStmt used to delete graph elements — nodes, relationships or paths.
type DeleteStmt struct {
	Pattern ir.Patn
	Next    Stmt
}

func (DeleteStmt) patternNode() {}

func (m *DeleteStmt) GetPattern() ir.Patn {
	return m.Pattern
}

func (m *DeleteStmt) GetNext() Stmt {
	return m.Next
}
