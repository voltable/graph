package ast

// MatchStmt used to search for the pattern described in it.
type MatchStmt struct {
	Pattern *Patn
	Next    Clauses
}

func (MatchStmt) patternNode() {}

func (m *MatchStmt) GetPattern() *Patn {
	return m.Pattern
}

func (m *MatchStmt) GetNext() Clauses {
	return m.Next
}
