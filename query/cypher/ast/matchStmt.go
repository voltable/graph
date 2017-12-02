package ast

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
)

// MatchStmt used to search for the pattern described in it.
type MatchStmt struct {
	Pattern ir.Patn
	Next    Clauses
}

func (MatchStmt) patternNode() {}

func (m *MatchStmt) GetPattern() ir.Patn {
	return m.Pattern
}

func (m *MatchStmt) GetNext() Clauses {
	return m.Next
}
