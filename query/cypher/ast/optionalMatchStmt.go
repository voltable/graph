package ast

import "github.com/RossMerr/Caudex.Graph/query/cypher/ir"

// OptionalMatchStmt used to search for the pattern described in it, can match on nil
type OptionalMatchStmt struct {
	Pattern ir.Patn
	Next    Clauses
}

func (OptionalMatchStmt) patternNode() {}

func (m *OptionalMatchStmt) GetPattern() ir.Patn {
	return m.Pattern
}

func (m *OptionalMatchStmt) GetNext() Clauses {
	return m.Next
}
