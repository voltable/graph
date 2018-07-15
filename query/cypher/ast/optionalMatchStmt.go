package ast

// OptionalMatchStmt used to search for the pattern described in it, can match on nil
type OptionalMatchStmt struct {
	Pattern *Patn
	Next    Clauses
}

func (OptionalMatchStmt) patternNode() {}

func (m *OptionalMatchStmt) GetPattern() *Patn {
	return m.Pattern
}

func (m *OptionalMatchStmt) GetNext() Clauses {
	return m.Next
}
